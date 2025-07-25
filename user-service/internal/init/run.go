package init

func Run() {
	ConfigLoader()

	InitPostgres()
	InitRedis()
	InitNats()

	r := InitRouter()
	r.Run()

}
