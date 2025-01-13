package config

func Init() {
	// 初始化数据库
	ConnectToDB("root", "root", "127.0.0.1", 3306, "go_db")
	// 初始化路由
	InitRouters()
}
