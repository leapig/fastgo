package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
)

type EnterpriseUserAttachment interface {
	Create(attachment *entity.EnterpriseUserAttachment) (*entity.EnterpriseUserAttachment, error)
	Delete(attachment *entity.EnterpriseUserAttachment) error
	List(attachment *entity.EnterpriseUserAttachment) ([]*entity.EnterpriseUserAttachment, error)
}

type enterpriseUserAttachment struct {
	dao dao.Dao
}

func NewEnterpriseUserAttachment(dao dao.Dao) EnterpriseUserAttachment {
	return &enterpriseUserAttachment{dao: dao}
}

func (o *enterpriseUserAttachment) Create(attachment *entity.EnterpriseUserAttachment) (*entity.EnterpriseUserAttachment, error) {
	return o.dao.EnterpriseUserAttachment().Create(attachment)
}

func (o *enterpriseUserAttachment) Delete(attachment *entity.EnterpriseUserAttachment) error {
	return o.dao.EnterpriseUserAttachment().Delete(attachment)
}

func (o *enterpriseUserAttachment) List(attachment *entity.EnterpriseUserAttachment) ([]*entity.EnterpriseUserAttachment, error) {
	return o.dao.EnterpriseUserAttachment().List(attachment)
}
