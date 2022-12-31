package service

import (
	"context"
	"errors"
	"log"

	"combo/gen/healthy"
)

// healthy service example implementation.
// The example methods log the requests and return zero values.
type healthysrvc struct {
	logger *log.Logger
}

// NewHealthy returns the healthy service implementation.
func NewHealthy(logger *log.Logger) healthy.Service {
	return &healthysrvc{logger}
}

// Get implements get.
func (s *healthysrvc) Get(ctx context.Context) (err error) {
	s.logger.Print("healthy.get")

	return errors.Unwrap(err)
}
