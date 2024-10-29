package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var PrivateKey = []byte("78e6c5dd07e531509f99fc9c9d52743af74b766a")

type UserAccessToken struct {
	UserAccessToken          string `json:"UserAccessToken"`          //用户鉴权
	UserAccessTokenExpireIn  int32  `json:"ExpireIn"`                 //过期时间
	UserRefreshToken         string `json:"UserRefreshToken"`         //用户刷新鉴权
	UserRefreshTokenExpireIn int32  `json:"UserRefreshTokenExpireIn"` //过期时间
}

type TenantAccessToken struct {
	TenantAccessToken          string `json:"TenantAccessToken"`          //用户鉴权
	TenantAccessTokenExpireIn  int32  `json:"TenantAccessTokenExpireIn"`  //过期时间
	TenantRefreshToken         string `json:"TenantRefreshToken"`         //用户刷新鉴权
	TenantRefreshTokenExpireIn int32  `json:"TenantRefreshTokenExpireIn"` //过期时间
}
type UserClaim struct {
	jwt.RegisteredClaims
	UserPk string `json:"user_pk"`
}
type TenantClaim struct {
	jwt.RegisteredClaims
	TenantPk     string `json:"tenant_pk"`
	UserPk       string `json:"user_pk"`
	TenantUserPk string `json:"tenant_user_pk"`
}

type AuthMessage struct {
	TenantPk     string `json:"tenant_pk"`
	UserPk       string `json:"user_pk"`
	TenantUserPk string `json:"tenant_user_pk"`
}

// GenerateUserToken 生成一个用户访问令牌和一个刷新令牌
func GenerateUserToken(userPk string) UserAccessToken {
	// 创建一个新的 JWT 对象，并使用 HS256 签名方法
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		// 注册声明
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发者
			Issuer: "user_access_token",
			// 过期时间：当前时间 + 2 小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			// 签发时间：当前时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间：当前时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		// 用户主键
		UserPk: userPk,
	})
	// 将 JWT 对象签名字符串化
	tokenString, _ := accessToken.SignedString(PrivateKey)

	// 创建一个新的 JWT 对象，并使用 HS256 签名方法
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		// 注册声明
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发者
			Issuer: "refresh_user_access_token",
			// 过期时间：当前时间 + 12 小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			// 签发时间：当前时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间：当前时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		// 用户主键
		UserPk: userPk,
	})
	// 将 JWT 对象签名字符串化
	refreshTokenString, _ := refreshToken.SignedString(PrivateKey)

	// 返回用户访问令牌和刷新令牌
	return UserAccessToken{UserAccessToken: tokenString, UserRefreshToken: refreshTokenString, UserAccessTokenExpireIn: 2 * 60 * 60, UserRefreshTokenExpireIn: 12 * 60 * 60}
}

// ParseUserToken 解析用户访问令牌并验证其有效性
func ParseUserToken(userAccessToken string) (*AuthMessage, error) {
	// 解析 JWT 字符串
	token, err := jwt.Parse(userAccessToken, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否为 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		} else {
			return PrivateKey, nil
		}
	})
	// 如果解析过程中没有错误
	if err == nil {
		// 尝试将声明转换为 MapClaims 类型
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 获取签发者，如果获取成功且签发者为 "user_access_token"
			if issuer, err := claims.GetIssuer(); err == nil && issuer == "user_access_token" {
				// 返回包含用户主键的 AuthMessage 对象
				return &AuthMessage{
					UserPk: claims["user_pk"].(string),
				}, nil
			}
		}
	}
	// 如果解析失败或验证未通过，返回一个空的 AuthMessage 对象和一个错误信息
	return &AuthMessage{}, errors.New("无权访问")
}

// RefreshUserToken 刷新用户访问令牌
func RefreshUserToken(refreshToken string) UserAccessToken {
	// 解析 refreshToken
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		} else {
			return PrivateKey, nil
		}
	})
	// 若解析失败，返回空的 UserAccessToken
	if err != nil {
		return UserAccessToken{}
	}
	// 若解析成功，验证 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 生成新的用户访问令牌
		return GenerateUserToken(claims["user_pk"].(string))
	}
	// 若 token 无效，返回空的 UserAccessToken
	return UserAccessToken{}
}

// GenerateTenantToken 生成租户访问令牌和刷新令牌
func GenerateTenantToken(tenantPk, userPk, tenantUserPk string) TenantAccessToken {
	// 创建一个新的 HS256 签名的 JWT
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, TenantClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发者
			Issuer: "tenant_access_token",
			// 过期时间：当前时间 + 2 小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			// 签发时间：当前时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间：当前时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		TenantPk:     tenantPk,
		UserPk:       userPk,
		TenantUserPk: tenantUserPk,
	})
	// 将 accessToken 转换为字符串
	tokenString, _ := accessToken.SignedString(PrivateKey)

	// 创建一个新的 HS256 签名的 JWT 作为刷新令牌
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, TenantClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发者
			Issuer: "refresh_tenant_refresh_token",
			// 过期时间：当前时间 + 12 小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			// 签发时间：当前时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间：当前时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		TenantPk:     tenantPk,
		UserPk:       userPk,
		TenantUserPk: tenantUserPk,
	})
	// 将 refreshToken 转换为字符串
	refreshTokenString, _ := refreshToken.SignedString(PrivateKey)

	// 返回 TenantAccessToken 结构体，包含 accessToken、refreshToken 及其过期时间
	return TenantAccessToken{TenantAccessToken: tokenString, TenantRefreshToken: refreshTokenString, TenantAccessTokenExpireIn: 2 * 60 * 60, TenantRefreshTokenExpireIn: 12 * 60 * 60}
}

// ParseTenantToken 解析租户访问令牌
func ParseTenantToken(tenantAccessToken string) (*AuthMessage, error) {
	// 解析 tenantAccessToken
	token, err := jwt.Parse(tenantAccessToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		} else {
			return PrivateKey, nil
		}
	})
	// 若解析失败，返回错误
	if err == nil {
		// 若解析成功，验证 token 是否有效
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 返回 AuthMessage 结构体，包含 tenantPk、userPk、tenantUserPk
			return &AuthMessage{
				UserPk:       claims["user_pk"].(string),
				TenantPk:     claims["tenant_pk"].(string),
				TenantUserPk: claims["tenant_user_pk"].(string),
			}, nil
		}
	}
	// 若 token 无效，返回错误
	return nil, errors.New("无权访问")
}

// RefreshTenantToken 刷新租户访问令牌
func RefreshTenantToken(refreshToken string) TenantAccessToken {
	// 解析 refreshToken
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		} else {
			return PrivateKey, nil
		}
	})
	// 若解析失败，返回空的 TenantAccessToken
	if err != nil {
		return TenantAccessToken{}
	}
	// 若解析成功，验证 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 生成新的租户访问令牌
		return GenerateTenantToken(claims["tenant_pk"].(string), claims["user_pk"].(string), claims["tenant_user_pk"].(string))
	}
	// 若 token 无效，返回空的 TenantAccessToken
	return TenantAccessToken{}
}
