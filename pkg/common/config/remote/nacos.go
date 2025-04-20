package remote

// import (
// 	"encoding/json"

// 	"github.com/nacos-group/nacos-sdk-go/v2/clients"
// 	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
// 	"github.com/nacos-group/nacos-sdk-go/v2/vo"
// )

// // var

// func GetConfig(dataId, group string, conf any) error {
// 	content, err := GetConfigContent(dataId, group)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal([]byte(content), conf)
// }

// func GetConfigContent(dataId, group string) (string, error) {
// 	sc := constant.ServerConfig{
// 		IpAddr: nacosConfig.Host,
// 		Port:   nacosConfig.Port,
// 	}

// 	cc := constant.ClientConfig{
// 		NamespaceId: nacosConfig.Namespace,
// 		OpenKMS:     true,
// 		TimeoutMs:   5000,
// 		LogLevel:    "debug",
// 	}

// 	client, err := clients.NewConfigClient(
// 		vo.NacosClientParam{
// 			ClientConfig:  &cc,
// 			ServerConfigs: []constant.ServerConfig{sc},
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return client.GetConfig(vo.ConfigParam{
// 		DataId: dataId,
// 		Group:  group,
// 	})
// }
