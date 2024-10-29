package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
)

type UserClient interface {
	DeleteByUserPkAndClientType(client *entity.UserClient) error
}

// userClient 接口规范实现类
type userClient struct {
	dao dao.Dao
}

// NewUserClient 实例化接口规范实现类
func NewUserClient(dao dao.Dao) UserClient {
	return &userClient{dao: dao}
}

func (o *userClient) DeleteByUserPkAndClientType(client *entity.UserClient) error {
	return o.dao.UserClient().DeleteByUserPkAndClientType(client)
}
