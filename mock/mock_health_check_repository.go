package mock

import (
	"github.com/frencius/go-skeleton/model"
)

type MockHealthCheckRepository struct {
	PingFunc func() (*model.Ping, error)
}

func (m *MockHealthCheckRepository) Ping() (*model.Ping, error) {
	return m.PingFunc()
}
