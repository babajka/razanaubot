package telegram

type Service struct{}

func BuildService() *Service {
	return &Service{}
}

func (s *Service) SendMessage(msg string) error {
	return nil
}
