package publisher

import (
	natsPublisher "todo_list/src/infra/broker/nats/publisher"

	"github.com/stretchr/testify/mock"
)

type MockPublisher struct {
	mock.Mock
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{}
}

var _ natsPublisher.PublisherInterface = &MockPublisher{}

func (o *MockPublisher) Nats(data []byte, subject string) error {
	args := o.Called(data, subject)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
