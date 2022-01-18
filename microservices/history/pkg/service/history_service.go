package service

type HistoryService struct {
}

func NewHistoryService() *HistoryService {
	return &HistoryService{}
}

func (hs *HistoryService) Status() (string, error) {
	return "history service is alive", nil
}
