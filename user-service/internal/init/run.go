package init

func Run() {
	ConfigLoader()

	r := InitRouter()
	r.Run()

}
