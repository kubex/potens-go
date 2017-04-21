package potens_test

import (
	"testing"
	"github.com/kubex/potens-go"
)

func TestValidateGlobalAppID(t *testing.T) {
	if potens.ValidateGlobalAppID("/asd") == nil {
		t.Fail()
	}

	if potens.ValidateGlobalAppID("asd/") == nil {
		t.Fail()
	}
	if potens.ValidateGlobalAppID("/a") == nil {
		t.Fail()
	}
	if potens.ValidateGlobalAppID("d") == nil {
		t.Fail()
	}
	if potens.ValidateGlobalAppID("dwfwe") == nil {
		t.Fail()
	}
	if potens.ValidateGlobalAppID("d/a") != nil {
		t.Fail()
	}
}

func TestSplitGaID(t *testing.T) {
	vendor, app, err := potens.SplitGaID("cubex/app")
	if err != nil {
		t.Error(err.Error())
	}
	if vendor != "cubex" {
		t.Error("Incorrect vendor ID")
	}
	if app != "app" {
		t.Error("Incorrect app ID")
	}

	vendor, app, err = potens.SplitGaID("cubex")
	if err == nil {
		t.Error("Failed to error on invalid GAID")
	}
	vendor, app, err = potens.SplitGaID("cubex/klhfw/fwejhfew")
	if err == nil {
		t.Error("Failed to error on invalid GAID")
	}
}