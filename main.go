package main

import (
	"thor-backend/internal/controller"
	"thor-backend/internal/dao"
	"thor-backend/internal/logic"
	"thor-backend/internal/setting"
)


func main() {
	// 加载配置文件
	setting.Init()
	// 加载数据库
	d := dao.Init()
	// 加载应用
	l := logic.Init(d)
	// 启动服务
	s := controller.Init(l)
	// 平滑关闭
	controller.GracefulShutdown(s)
}
