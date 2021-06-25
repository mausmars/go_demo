package fsm_service

import (
	"testing"
)

func TestMatchService(t *testing.T) {
	service := NewFSMService()
	service.Startup()
}
