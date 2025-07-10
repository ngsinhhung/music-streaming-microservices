package init

func Run() {
	ConfigLoader()

	InitPostgres()

	r := InitRouter()
	r.Run()

}
