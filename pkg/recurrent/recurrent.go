package recurrent

import "razanaubot/pkg/data"

type (
	telegramWithChannel interface {
		SendMessage(string) error
	}
)

type Service struct {
	dt data.Data

	tgService telegramWithChannel
}

func BuildService(dt data.Data, tgService telegramWithChannel) *Service {
	return &Service{
		dt:        dt,
		tgService: tgService,
	}
}

func (s *Service) StartWait() error {

	return nil
}
