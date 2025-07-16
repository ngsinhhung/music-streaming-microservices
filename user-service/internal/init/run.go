package init

func Run() {
	ConfigLoader()

	InitPostgres()
	InitRedis()

	r := InitRouter()
	r.Run()

}
