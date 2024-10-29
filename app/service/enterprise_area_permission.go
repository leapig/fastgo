package service

import (
	"errors"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
)

type EnterpriseAreaPermission interface {
	CreateEnterpriseAreaPermission(areaPermission *entity.EnterpriseAreaPermission) (*entity.EnterpriseAreaPermission, error)
	SelectEnterpriseAreaPermission(areaPermission *entity.EnterpriseAreaPermission, page *entity.Pagination) ([]*entity.EnterpriseAreaPermission, int32, error)
	DeleteEnterpriseAreaPermission(areaPermission *entity.EnterpriseAreaPermission) error
	FindList(areaPermission *entity.EnterpriseAreaPermission) ([]*entity.EnterpriseAreaPermission, error)
}

type enterpriseAreaPermission struct {
	dao dao.Dao
}

func NewEnterpriseAreaPermission(dao dao.Dao) EnterpriseAreaPermission {
	return &enterpriseAreaPermission{dao: dao}
}

func (e *enterpriseAreaPermission) CreateEnterpriseAreaPermission(areaPermission *entity.EnterpriseAreaPermission) (*entity.EnterpriseAreaPermission, error) {
	if bol, bolErr := e.dao.EnterpriseAreaPermission().CheckRepetition(areaPermission); bolErr == nil {
		if bol {
			return e.dao.EnterpriseAreaPermission().Create(areaPermission)
		} else {
			return nil, errors.New("重复添加")
		}
	} else {
		return nil, bolErr
	}
}

func (e *enterpriseAreaPermission) SelectEnterpriseAreaPermission(areaPermission *entity.EnterpriseAreaPermission, page *entity.Pagination) ([]*entity.EnterpriseAreaPermission, int32, error) {
	if rows, err := e.dao.EnterpriseAreaPermission().Select(areaPermission, page); err == nil {
		count, _ := e.dao.EnterpriseAreaPermission().Count(areaPermission)
		return rows, count, err
	} else {
		return nil, 0, err
	}
}

func (e *enterpriseAreaPermission) DeleteEnterpriseAreaPermission(areaPermission *entity.EnterpriseAreaPermission) error {
	return e.dao.EnterpriseAreaPermission().Delete(areaPermission)
}

func (e *enterpriseAreaPermission) FindList(areaPermission *entity.EnterpriseAreaPermission) ([]*entity.EnterpriseAreaPermission, error) {
	return e.dao.EnterpriseAreaPermission().FindList(areaPermission)
}
