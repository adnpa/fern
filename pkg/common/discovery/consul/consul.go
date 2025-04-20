package discovery

// import (
// 	"fmt"
// 	"log"

// 	"github.com/google/uuid"
// 	"github.com/hashicorp/consul/api"
// )

// type Discovery interface{}

// type ConsulConfig struct {
// 	Host string `json:"host,omitempty"`
// 	Port int    `json:"port,omitempty"`
// }

// var (
// 	consulConfig *ConsulConfig
// 	cli          *api.Client
// )

// func init() {
// 	err := config.GetConfig("consul", "dev", &consulConfig)
// 	if err != nil {
// 		panic(err)
// 	}

// 	cfg := api.DefaultConfig()
// 	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)
// 	cli, err = api.NewClient(cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func Register(addr string, port int, name string) error {
// 	log.Println(addr, port, name)
// 	check := api.AgentServiceCheck{
// 		GRPC:                           fmt.Sprintf("%s:%d", addr, port),
// 		Timeout:                        "3s",
// 		Interval:                       "10s",
// 		DeregisterCriticalServiceAfter: "10s",
// 	}
// 	serverId := uuid.NewString()
// 	registeration := api.AgentServiceRegistration{
// 		ID:      serverId,
// 		Address: addr,
// 		Port:    port,
// 		Name:    name,
// 		// Tags:    tags,
// 		Check: &check,
// 	}
// 	err := cli.Agent().ServiceRegister(&registeration)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func AllService() map[string]*api.AgentService {
// 	data, err := cli.Agent().Services()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return data
// }

// func FilterService(name string) {

// 	cfg := api.DefaultConfig()
// 	cfg.Address = "localhost:8500"
// 	cli, err := api.NewClient(cfg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	data, err := cli.Agent().ServicesWithFilter(`Service == ""`)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return data
// }

// func UnRegister() {
// 	Register("loaclhost", 11000, "user-web", []string{"maxsho", "bobby"}, "user-web")
// }

// 负载均衡和服务发现
