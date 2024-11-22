package env

import (
	"bytes"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/novando/go-ska/pkg/logger"
	"github.com/spf13/viper"
)

type Consul struct {
	client *consul.Client
	logger *logger.Logger
}

func InitConsul(host string, port int, schema string, l ...*logger.Logger) *Consul {
	log := logger.Call()
	if len(l) > 0 {
		log = l[0]
	}
	config := consul.Config{
		Address: fmt.Sprintf("%s:%d", host, port),
		Scheme:  schema,
	}
	client, err := consul.NewClient(&config)
	if log != nil {
		if err != nil {
			logger.Call().Fatalf("Error initializing Consul client: %v\n", err)
			return nil
		}
		logger.Call().Infof("Consul client initialized successfully")
	}

	return &Consul{
		client: client,
		logger: log,
	}
}

func (c Consul) RetrieveConfiguration(key string, configType string) {
	c.logger.Infof("key: %s, configType: %s", key, configType)

	pair, _, err := c.client.KV().Get(key, nil)
	if err != nil {
		c.logger.Fatalf("Error retrieving key %s: %s", key, err)
	}

	if pair != nil {
		envData := pair.Value
		viper.SetConfigType(configType)
		if err := viper.ReadConfig(bytes.NewBuffer(envData)); err != nil {
			c.logger.Fatalf("Error reading configuration for key %s: %s", key, err)
		}

		c.logger.Infof("Configuration registered for key: %s", key)
	}
}

func (c Consul) RegisterService(id, name, host string, port int) {

	reg := &consul.AgentServiceRegistration{
		ID:   id,
		Name: name,
		Port: port,
		Check: &consul.AgentServiceCheck{
			TCP:      fmt.Sprintf("%s:%d", host, port),
			Interval: "10s",
			Timeout:  "2s",
		},
	}

	c.logger.Infof("Registering service with ID: %s, Name: %s", id, name)

	err := c.client.Agent().ServiceRegister(reg)
	if err != nil {
		c.logger.Fatalf("Failed to register service: %v", err)
	}
}

func (c Consul) DeregisterService(id string) {
	err := c.client.Agent().ServiceDeregister(id)
	if err != nil {
		c.logger.Infof("Error deregistering service: %v", err)
	}
}

func (c Consul) RetrieveServiceUrl(id string) (string, error) {
	c.logger.Infof("Retrieving service URL for id: %s", id)

	service, _, err := c.client.Agent().Service(id, nil)
	if err != nil {
		c.logger.Infof("Error retrieving service: %v", err)
		return "", err
	}

	url := fmt.Sprintf("http://%s:%v", service.Address, service.Port)
	c.logger.Infof("Service URL: %s", url)

	return url, nil
}
