package consul

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"net/http"
	"strconv"
)

func Reg(name, id, host, port string, tags []string) error {
	defaultConfig := api.DefaultConfig()
	h := "127.0.0.1"
	p := "8500"
	defaultConfig.Address = fmt.Sprintf("%s:%s", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = defaultConfig.Address
	agentServiceRegistration.Name = name
	agentServiceRegistration.ID = id
	intPort, _ := strconv.Atoi(port)
	agentServiceRegistration.Port = intPort
	agentServiceRegistration.Tags = tags
	serverAddr := fmt.Sprintf("http://%s:%s/health", host, port)
	check := api.AgentServiceCheck{
		// GRPC: serverAddr,
		HTTP:                           serverAddr,
		Timeout:                        "3s",
		Interval:                       "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	agentServiceRegistration.Check = &check
	return client.Agent().ServiceRegister(agentServiceRegistration)
}

func GetServiceList() error {
	defaultConfig := api.DefaultConfig()
	h := "172.26.160.1"
	p := "8500"
	defaultConfig.Address = fmt.Sprintf("%s:%s", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	serviceList, err := client.Agent().Services()
	if err != nil {
		return err
	}
	for k, v := range serviceList {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--------------------")
	}
	return nil
}

func FilterService() error {
	defaultConfig := api.DefaultConfig()
	h := "127.0.0.1"
	p := "8500"
	defaultConfig.Address = fmt.Sprintf("%s:%s", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	serviceList, err := client.Agent().ServicesWithFilter("Service==go-maxms")
	if err != nil {
		return err
	}
	for k, v := range serviceList {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--------------------")
	}
	return nil
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
