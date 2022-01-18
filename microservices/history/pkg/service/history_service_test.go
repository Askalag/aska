package service_test

import (
	"github.com/Askalag/aska/microservices/history/pkg/service"
	"testing"
)

func TestHistoryService_Status(t *testing.T) {
	hs := service.NewHistoryService()
	eq := "history service is alive"

	st, err := hs.Status()
	if err != nil {
		t.Errorf("expected nil err, but: '%s'", err.Error())
	}

	if st != eq {
		t.Errorf("expected: '%s', but: '%s'", eq, st)
	}
}
