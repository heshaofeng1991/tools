package registry

import (
	"core/config"
)

type Empty struct {
}

func (d *Empty) Register(service *config.Service) error {
	return nil
}

func (d *Empty) Deregister(service *config.Service) error {
	return nil
}

func (d *Empty) GetService(service config.Service) ([]*config.Service, error) {
	return nil, nil
}

func (d *Empty) ListServices() ([]*config.Service, error) {
	return nil, nil
}
