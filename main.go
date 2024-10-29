package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/dal/dao"
	"github.com/leapig/fastgo/app/library/helper"
	utils "github.com/leapig/fastgo/app/library/util"
	"github.com/leapig/fastgo/app/router"
	"github.com/leapig/fastgo/app/service"
	_ "github.com/leapig/fastgo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// 读取配置
	helper.Config()
}

func main() {
	// 业务服务
	db := helper.RegisterGormMysql()
	rs := helper.RegisterRedis()
	service.InitSvc(dao.NewDao(db.Db, rs))
	defer rs.CloseRedis()
	// 接口服务
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(utils.TraceId())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Load(r.Group("/open-apis"))
	_ = r.Run(":80")
}
