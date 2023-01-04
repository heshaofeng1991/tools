package registry

import (
	"core/config"
	"core/utils"
	"fmt"
	"strings"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	cli  *api.Client
	conf *config.Consul
}

func NewConsul(cli *api.Client, conf *config.Consul) *Consul {
	return &Consul{
		cli:  cli,
		conf: conf,
	}
}

func (c *Consul) Register(service *config.Service) error {
	if len(strings.TrimSpace(service.Host)) == 0 {
		service.Host = utils.GetOutboundIP()
	}
	c.conf.Id = utils.MD5(fmt.Sprintf("%s%d%s", service.Host, service.Port, service.Name))
	param := &api.AgentServiceRegistration{
		ID:      c.conf.Id,
		Name:    service.Name,
		Tags:    []string{service.Cluster},
		Address: service.Host,
		Port:    int(service.Port),
	}

	param.Check = &api.AgentServiceCheck{
		Interval: "5s",
		Timeout:  "5s",
	}

	if service.CheckAddr == "" {
		param.Check.GRPC = fmt.Sprintf("%s:%d", service.Host, service.Port)
	} else {
		param.Check.HTTP = fmt.Sprintf("http://%s:%d%s", service.Host, service.Port, service.CheckAddr)
	}

	return c.cli.Agent().ServiceRegister(param)
}

func (c *Consul) Deregister(service *config.Service) error {
	return c.cli.Agent().ServiceDeregister(c.conf.Id)
}

func (c *Consul) GetService(service config.Service) ([]*config.Service, error) {
	options := &api.QueryOptions{
		Namespace: service.Namespace,
	}

	list := make([]*config.Service, 0)
	opts, err := c.cli.Agent().ServicesWithFilterOpts(fmt.Sprintf("Service == \"%s\"", service.Name), options)
	if err != nil {
		return nil, err
	}
	for _, v := range opts {
		list = append(list, &config.Service{
			Name:      v.Service,
			Metadata:  nil,
			Host:      v.Address,
			Port:      uint64(v.Port),
			Namespace: v.Namespace,
		})
	}

	return list, nil
}

func (c *Consul) ListServices() ([]*config.Service, error) {
	list := make([]*config.Service, 0)
	opts, err := c.cli.Agent().Services()
	if err != nil {
		return nil, err
	}
	for _, v := range opts {
		list = append(list, &config.Service{
			Name:      v.Service,
			Metadata:  nil,
			Host:      v.Address,
			Port:      uint64(v.Port),
			Namespace: v.Namespace,
		})
	}
	return list, nil
}
