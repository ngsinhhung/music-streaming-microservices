package init

import (
	"fmt"
	v "github.com/spf13/viper"
	"log"
	"music-streaming-microservices/email-service/global"
)

func ConfigLoader() {
	viper := v.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file: %s \n", err))
	}

	if err := viper.Unmarshal(&global.Configs); err != nil {
		panic(fmt.Errorf("Error unmarshalling config: %s \n", err))
	}

	log.Println("Config loaded successfully")
	log.Println("Server Port:: ", viper.GetInt("server.port"))
}
