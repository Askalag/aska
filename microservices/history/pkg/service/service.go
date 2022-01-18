package service

type Service struct {
	HSService History
}

type History interface {
	Status() (string, error)
}

func NewService() *Service {
	return &Service{HSService: NewHistoryService()}
}
