package sd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"

	"qtask/pkg/constant"
)

func ConsulRegister(
	consulAddress string,
	consulPort string,
	advertiseAddress string,
	advertisePort string,
	serviceName string,
	logger log.Logger) sd.Registrar {

	rand.Seed(time.Now().UTC().UnixNano())

	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = consulAddress + ":" + consulPort
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	check := api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%v:%v/%v", advertiseAddress, advertisePort, serviceName),
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	num := rand.Intn(100) // to make service ID unique
	asr := api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", constant.QTaskSrvPrefix, serviceName, strconv.Itoa(num)), //unique service ID
		Name:    fmt.Sprintf("%v-%v", constant.QTaskSrvPrefix, serviceName),
		Address: advertiseAddress,
		Port:    port,
		Tags:    []string{serviceName},
		Check:   &check,
	}
	return consulsd.NewRegistrar(client, &asr, logger)
}
