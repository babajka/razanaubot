package recurrent

import "razanaubot/pkg/data"

type (
	telegramService interface {
		SendMessage(string) error
	}
)

type Service struct{}

func BuildService(dt data.Data, tg telegramService) *Service {
	return &Service{}
}

func (s *Service) StartWait() error {
	return nil
}
