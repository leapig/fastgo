package service

import (
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
)

type TurnoverRecord interface {
	SelectTurnoverRecord(record *entity.TurnoverRecord, pg *entity.Pagination) ([]*model.TurnoverRecord, int32, error)
	ApplicationForResignation(record *entity.TurnoverRecord) (*entity.TurnoverRecord, error)
	ApprovalApplication(record *entity.TurnoverRecord, userPk int64) (*entity.TurnoverRecord, error)
}

type turnoverRecord struct {
	dao dao.Dao
}

func NewTurnoverRecord(dao dao.Dao) TurnoverRecord {
	return &turnoverRecord{dao: dao}
}

func (o *turnoverRecord) SelectTurnoverRecord(record *entity.TurnoverRecord, pg *entity.Pagination) ([]*model.TurnoverRecord, int32, error) {
	en := &model.TurnoverRecord{
		EnterprisePk: record.EnterprisePk,
	}
	res, err := o.dao.TurnoverRecord().SelectDetail(en, pg)
	if err != nil {
		return nil, 0, err
	} else {
		count, _ := o.dao.TurnoverRecord().Count(en)
		return res, count, err
	}

}

func (o *turnoverRecord) ApplicationForResignation(record *entity.TurnoverRecord) (*entity.TurnoverRecord, error) {
	if count, countErr := o.dao.TurnoverRecord().Count(&model.TurnoverRecord{
		EnterprisePk:     record.EnterprisePk,
		EnterpriseUserPk: record.EnterpriseUserPk,
		ApplyStatus:      1,
	}); countErr != nil {
		return nil, countErr
	} else {
		if count > 0 {
			return nil, errors.New("该人员已申请！")
		}
		return o.dao.TurnoverRecord().Create(&entity.TurnoverRecord{
			EnterprisePk:     record.EnterprisePk,
			EnterpriseUserPk: record.EnterpriseUserPk,
		})
	}
}

func (o *turnoverRecord) ApprovalApplication(record *entity.TurnoverRecord, userPk int64) (*entity.TurnoverRecord, error) {
	tr, _ := o.dao.TurnoverRecord().Find(&entity.TurnoverRecord{Pk: record.Pk})
	res, err := o.dao.TurnoverRecord().Update(record)
	if err != nil {
		return nil, err
	} else {
		if record.ApplyStatus == 1 {
			//离职审批通过
			o.dao.EnterpriseUser().Update(&entity.EnterpriseUser{Pk: tr.EnterprisePk, Status: 2})
			//删除临时人员记录
			o.dao.TemporaryWorker().DeleteByEnterpriseUserPk(&entity.TemporaryWorker{EnterpriseUserPk: tr.EnterpriseUserPk})
			//删除关联关系
			o.dao.Member().DeleteByEnterprisePkAndUserPk(&entity.Member{UserPk: tr.EnterpriseUserPk, EnterprisePk: tr.EnterprisePk})
			delErr := o.dao.UserPermission().DeleteByUserPkAndEnterprisePk(&entity.UserPermission{UserPk: userPk, EnterprisePk: res.EnterprisePk})
			if delErr != nil {
				logger.Error(delErr)
			}
		}
	}
	return res, err
}
