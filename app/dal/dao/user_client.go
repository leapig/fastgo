package dao

import (
	"errors"
	"github.com/fatih/structs"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/library/helper"
	"gorm.io/gorm"
)

type UserClient interface {
	FindList(*entity.UserClient) ([]*entity.UserClient, error)
	DeleteByUserPkAndClientType(*entity.UserClient) error
	SaveUserClientSubscribe(*entity.UserClient)
	CreateUserClient(*entity.UserClient) error
	UpdateUserClientSubscribe(*entity.UserClient) error
	Find(*entity.UserClient) (*entity.UserClient, error)
}

type userClient struct {
	db *gorm.DB
	rs *helper.Redis
}

func NewUserClient(db *gorm.DB, rs *helper.Redis) UserClient {
	return &userClient{
		db: db,
		rs: rs,
	}
}

func (o *userClient) FindList(p *entity.UserClient) ([]*entity.UserClient, error) {
	sql := o.db.Model(&entity.UserCredentials{})
	if p.UserPk != 0 {
		sql.Where("user_pk = ?", p.UserPk)
	}
	var rows []*entity.UserClient
	tx := sql.Find(&rows)
	return rows, tx.Error
}

func (o *userClient) DeleteByUserPkAndClientType(client *entity.UserClient) error {
	sql := o.db.Model(&entity.UserClient{})
	sql.Where("user_pk = ? and client_type =?", client.UserPk, client.ClientType)
	tx := sql.Unscoped().Delete(&client)
	return tx.Error
}

// SaveUserClientSubscribe 保存公众号关注状态
func (o *userClient) SaveUserClientSubscribe(uu *entity.UserClient) {
	if u, err := o.Find(uu); err == nil && u.Pk != 0 {
		if u.WxUnionid != "" {
			uu.WxUnionid = u.WxUnionid
		}
		_ = o.UpdateUserClientSubscribe(uu)
	} else {
		_ = o.CreateUserClient(uu)
	}
}

// Find 获取用户绑定关系数据层具体实现
func (o *userClient) Find(uc *entity.UserClient) (*entity.UserClient, error) {
	var row *entity.UserClient
	sql := o.db.Model(&entity.UserClient{}).Where("client_id = ? and client_type = ?", uc.ClientId, uc.ClientType)
	if uc.OpenId != "" {
		sql = sql.Where("open_id = ?", uc.OpenId)
	}
	if uc.UserPk != 0 {
		sql = sql.Where("user_pk = ?", uc.UserPk)
	}
	if uc.WxUnionid != "" {
		sql = sql.Where("wx_unionid = ?", uc.WxUnionid)

	}
	tx := sql.Find(&row)
	return row, tx.Error
}

func (o *userClient) UpdateUserClientSubscribe(uc *entity.UserClient) error {
	if tx := o.db.Model(&entity.UserClient{}).Where("open_id=? and client_id=? and client_type=?", uc.OpenId, uc.ClientId, uc.ClientType).Updates(structs.Map(struct {
		WxUnionid   string `json:"wx_unionid" db:"" gorm:"column:;comment:wx_unionid"`
		WxSubscribe *int8  `json:"wx_subscribe" db:"wx_subscribe" gorm:"column:wx_subscribe;comment:关注状态"   structs:"wx_subscribe"`
	}{WxSubscribe: uc.WxSubscribe, WxUnionid: uc.WxUnionid})); tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("数据操作失败")
	} else {
		return nil
	}
}

func (o *userClient) CreateUserClient(uc *entity.UserClient) error {
	uc.Pk = helper.GetRid(helper.UserClient)
	if tx := o.db.Create(&uc); tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("数据操作失败")
	} else {
		return nil
	}
}
