package init

import (
	"fmt"
	v "github.com/spf13/viper"
	"music-streaming-microservices/user-service/global"
)

func ConfigLoader() {
	viper := v.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Server Port:: ", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Configs); err != nil {
		panic(fmt.Errorf("Fatal error unmarshal config file: %s \n", err))
	}
}
