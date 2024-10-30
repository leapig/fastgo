package service

import (
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/dal/entity"
	"github.com/leapig/fastgo/app/dal/model"
	"github.com/leapig/fastgo/app/library/aliyun"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"math/rand"
	"strconv"
	"time"
)

type Files interface {
	SaveFile(string, model.Oss) (model.Oss, error)
	RemoveFile(string, model.Oss) error
}

// files 接口规范实现类
type files struct {
	dao dao.Dao
}

// NewFiles 实例化接口规范实现类
func NewFiles(dao dao.Dao) Files {
	return &files{
		dao: dao,
	}
}

// SaveFile 保存文件业务层实现
func (f *files) SaveFile(tenant string, object model.Oss) (model.Oss, error) {
	var obj model.Oss
	var err error
	object.Name = tenant + "/" + time.Now().Format("20060102") + "/" + time.Now().Format("150405") + strconv.Itoa(rand.Intn(10000)) + object.Name
	obj, err = aliyun.Oss().PutObject(object)
	if err != nil {
		return obj, err
	}
	_, err = f.dao.Files().Create(&entity.Files{
		Pk:           helper.GetRid(helper.FilesT),
		EnterprisePk: utils.StringToInt64(tenant),
		Name:         obj.Name,
		Size:         obj.Size,
		Extension:    obj.Extension,
		Suffix:       obj.Suffix,
	})
	return obj, err
}

// RemoveFile 移除文件业务层实现
func (f *files) RemoveFile(tenant string, object model.Oss) error {
	if err := f.dao.Files().Delete(&entity.Files{
		EnterprisePk: utils.StringToInt64(tenant),
		Name:         object.Name,
	}); err != nil {
		return err
	}
	return aliyun.Oss().RemoveObject(object)
}
