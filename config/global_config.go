package config

func Init() {
	// 初始化数据库
	ConnectToSQLiteDB("SQLite.db", true)
	// 初始化路由
	InitRouters()
}
