package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type UserRealNameAuthenticationLog interface {
	VerifyUser(userPk, name, idCard, face, url string) error
	CheckRealNameAuthenticationLog(userPk string) error
}

type userRealNameAuthenticationLog struct {
	dao dao.Dao
}

// NewUserRealNameAuthenticationLog 实例化接口规范实现类
func NewUserRealNameAuthenticationLog(dao dao.Dao) UserRealNameAuthenticationLog {
	return &userRealNameAuthenticationLog{dao: dao}
}

func (o *userRealNameAuthenticationLog) VerifyUser(userPk, name, idCard, face, url string) error {
	_, _, code, _, err := verifyOnline(face, idCard, name)
	if err != nil {
		logger.Error("人证核验有问题！！！！！！！！！！！！！！！！！！！！！")
		return err
	}
	if code == 0 {
		pk := helper.Rid(helper.UserRealNameAuthenticationLog)
		_, err1 := o.dao.UserRealNameAuthenticationLog().Create(&entity.UserRealNameAuthenticationLog{
			Pk:     utils.StringToInt64(pk),
			UserPk: utils.StringToInt64(userPk),
			Name:   name,
			IdCard: idCard,
			Face:   url,
		})
		if err1 != nil {
			logger.Error(err1)
		}
	} else {
		return errors.New("比对错误！！")
	}
	return nil
}

func (o *userRealNameAuthenticationLog) CheckRealNameAuthenticationLog(userPk string) error {
	if count, err := o.dao.UserRealNameAuthenticationLog().Count(&entity.UserRealNameAuthenticationLog{UserPk: utils.StringToInt64(userPk)}); err != nil {
		return err
	} else {
		if count > 0 {
			return nil
		} else {
			return errors.New("该人员未实名")
		}
	}
}

func verifyOnline(face, idCard, name string) (string, string, int, bool, error) {
	token, err := getDaBaiAccessToken()
	if err != nil {
		return "", "", -1, false, err
	}
	type Req struct {
		AccessToken string `json:"accessToken"`
		AuthData    struct {
			Mode   string `json:"mode"`
			IdInfo struct {
				IdNum    string `json:"idNum"`
				FullName string `json:"fullName"`
			} `json:"idInfo"`
			Portrait string `json:"portrait"`
		} `json:"authData"`
	}
	data := Req{}
	if face == "" {
		// 无照片模式
		data.AuthData.Mode = "64"
	} else {
		faceSub, err := subBase64(face)
		if err != nil {
			return "", "", -1, false, err
		}
		data.AuthData.Mode = "66"
		data.AuthData.Portrait = strings.Replace(faceSub, "\u0000", "", -1)
	}
	data.AccessToken = token
	data.AuthData.IdInfo.IdNum = idCard
	data.AuthData.IdInfo.FullName = name
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, os.Getenv("DABAI_DOMAIN")+"/v2/api/simpauth", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("getAuthIdentity:", err)
		return "", "", -1, false, errors.New("网络错误")
	}
	logger.Info(err)
	resp, _ := io.ReadAll(response.Body)
	type Res struct {
		Code int    `json:"retCode"`
		Msg  string `json:"retMessage"`
	}
	res := &Res{}
	_ = json.Unmarshal(resp, res)
	logger.Info(string(resp))
	if res.Code != 0 {
		return string(payload), string(resp), res.Code, true, errors.New(res.Msg)
	}
	return string(payload), string(resp), res.Code, true, nil

}

func getDaBaiAccessToken() (string, error) {
	clientId := os.Getenv("DABAI_CLIENTID")
	clientSecret := os.Getenv("DABAI_CLIENTSECRET")
	p := url.Values{}
	p.Add("clientId", clientId)
	p.Add("clientSecret", clientSecret)
	logger.Info("params:", p)
	req, _ := http.NewRequest(http.MethodGet, os.Getenv("DABAI_DOMAIN")+"/v2/api/getaccesstoken?"+p.Encode(), nil)
	response, err := http.DefaultClient.Do(req)
	logger.Info(err)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	resp, _ := io.ReadAll(response.Body)
	type Res struct {
		Code        int    `json:"retCode"`
		Msg         string `json:"retMessage"`
		AccessToken string `json:"accessToken"`
		ExpiresIn   int    `json:"expireSeconds"`
	}
	res := &Res{}
	_ = json.Unmarshal(resp, res)
	return res.AccessToken, nil
}

func subBase64(face string) (string, error) {
	ddd, _ := base64.StdEncoding.DecodeString(face) //成图片文件并把文件写入到buffer
	bbb := bytes.NewBuffer(ddd)
	m, _, err := image.Decode(bbb)
	if err != nil {
		return "", err
	}
	subImg := resize.Thumbnail(800, 800, m, resize.NearestNeighbor)
	emptyBuff := bytes.NewBuffer(nil)       //开辟一个新的空buff
	_ = jpeg.Encode(emptyBuff, subImg, nil) //img写入到buff
	dist := make([]byte, 1024*1024)         //开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes())
	return string(dist), nil
}
