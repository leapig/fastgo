package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Sms interface {
	SendCode(phone, tenantRid string) error
	CheckCode(phone, code string) error
}

// Sms 短信息
type sms struct {
	dao dao.Dao
}

// NewSms 实例化接口规范实现类
func NewSms(dao dao.Dao) Sms {
	return &sms{dao: dao}
}

// SendCode 发送验证码
func (o *sms) SendCode(phone, tenantRid string) error {
	key, err := totp.Generate(totp.GenerateOpts{Issuer: "otp", AccountName: "checkCode", SecretSize: 64})
	if err != nil {
		return err
	}
	code, te := totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{Period: 300, Digits: otp.DigitsSix, Algorithm: otp.AlgorithmSHA1})
	if te != nil {
		return te
	}
	content := "【AI门岗】您的短信验证码：" + code + "，5分钟内有效，请勿泄漏。"
	if sendMsgErr := o.send(content, phone); sendMsgErr != nil {
		return sendMsgErr
	} else {
		return o.dao.Redis().SetEx(phone, 300, code)
	}
}

// CheckCode 验证验证码
func (o *sms) CheckCode(phone, code string) error {
	if sCode, err := o.dao.Redis().Get(phone); err != nil || sCode == nil || string(sCode.([]byte)) != code {
		return errors.New("验证码错误")
	} else {
		_, _ = o.dao.Redis().Del(phone)
		return nil
	}
}

// send 发送短信
func (o *sms) send(content, phone string) error {
	//请求参数s
	type Data struct {
		UserId    string `json:"userid"` //用户账号
		Pwd       string `json:"pwd"`    //用户密码 :md5加密（userid值大写+固定字符串：00000000+明文pwd+timestamp）
		Mobile    string `json:"mobile"`
		Content   string `json:"content"`
		TimeStamp string `json:"timestamp"`
	}
	timestamp := time.Now().Format("0102150405")
	pwd := utils.MD5(strings.ToUpper(os.Getenv("SMS_USERID")) + "00000000" + os.Getenv("SMS_PWD") + timestamp)
	//请求参数的封装
	q := Data{
		UserId:    os.Getenv("SMS_USERID"),
		Pwd:       pwd,
		Mobile:    phone,
		Content:   content,
		TimeStamp: timestamp,
	}
	payload, _ := json.Marshal(q)
	req, _ := http.NewRequest(http.MethodPost, os.Getenv("SMS_DOMAIN")+"/sms/v2/std/send_single", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, _ := client.Do(req)
	resp, _ := io.ReadAll(response.Body)
	//应答公共参数
	type Res struct {
		Result int    `json:"result"`
		Desc   string `json:"desc"`
		Msgid  string `json:"msgid"`
		Custid string `json:"custid"`
	}
	res := &Res{}
	_ = json.Unmarshal(resp, res)
	var status int8 = 1
	reason := "发送成功"
	if res.Result != 0 {
		status = 2
		reason = "发送失败：" + res.Desc
	}
	_, err := o.dao.Sms().Create(&entity.Sms{
		Pk:      utils.StringToInt64(helper.Rid(helper.Sms)),
		Phone:   phone,
		Content: content,
		Status:  status,
		Reason:  reason,
		Type:    1,
	})
	return err
}
