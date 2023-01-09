package config_test

import (
	"github.com/emilebui/demoX-rk_worker/pkg/config"
	"testing"
)

func TestGet(t *testing.T) {
	check := config.Get("../../config.yaml")
	if check == nil {
		t.Errorf("get config failed")
	}
}
