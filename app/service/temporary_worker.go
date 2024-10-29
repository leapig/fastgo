package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
)

type TemporaryWorker interface {
	SelectTemporaryWorker(en *model.TemporaryWorker, pg *entity.Pagination) ([]*model.TemporaryWorker, int32, error)
	CreateTemporaryWorker(en *entity.TemporaryWorker) (*entity.TemporaryWorker, error)
}

type temporaryWorker struct {
	dao dao.Dao
}

func NewTemporaryWorker(dao dao.Dao) TemporaryWorker {
	return &temporaryWorker{dao: dao}
}

func (t *temporaryWorker) SelectTemporaryWorker(en *model.TemporaryWorker, pg *entity.Pagination) ([]*model.TemporaryWorker, int32, error) {
	if res, err := t.dao.TemporaryWorker().SelectDetail(en, pg); err == nil {
		count, _ := t.dao.TemporaryWorker().Count(en)
		return res, count, err
	} else {
		return nil, 0, err
	}
}

func (t *temporaryWorker) CreateTemporaryWorker(en *entity.TemporaryWorker) (*entity.TemporaryWorker, error) {
	return t.dao.TemporaryWorker().Create(en)
}
