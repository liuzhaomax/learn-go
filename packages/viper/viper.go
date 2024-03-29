package main

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

type DBConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	NilTest  string `mapstructure:"nil_test"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string `mapstructure:"serverName"`
	RedisConfig
}

func GetEnv(s string) string {
	viper.AutomaticEnv()
	return viper.GetString(s)
}

func main() {
	v := viper.New()
	v.SetConfigFile("./packages/viper/config.yaml")

	// 读取未定义schema的配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(v.Get("name"))

	// 读取已定义schema的配置文件
	var dbConfig DBConfig
	err = v.Unmarshal(&dbConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbConfig.Username)
	fmt.Println(dbConfig.Password)
	fmt.Println(dbConfig.NilTest)
	fmt.Println(reflect.TypeOf(dbConfig.NilTest))

	// 读取环境变量 Mac和linux可以使用 export ENV=dev 直接设置环境变量，Windows要配环境变量并重启IDEA
	env := GetEnv("ENV")
	configPath := "./packages/viper/" + env + ".yaml"
	v.SetConfigFile(configPath)
	v.ReadInConfig()
	var serverConfig ServerConfig
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("env is %s", env))
	fmt.Println(serverConfig.ServerName)
	fmt.Println(serverConfig.RedisConfig.Host)
	fmt.Println(serverConfig.RedisConfig.Port)
}
