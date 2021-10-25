package xstrings_test

import (
	"ksitigarbha/xstrings"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	if xstrings.IsValidEmail("zzz") {
		t.Error("Expect False but True")
	}
	if !xstrings.IsValidEmail("zhouteng@126.com") {
		t.Error("Expect True but False")
	}
	if !xstrings.IsValidEmail("zhoutd.eng@126.com") {
		t.Error("Expect True but False")
	}
}

func TestCheckPhone(t *testing.T) {
	if xstrings.IsValidPhone("124") {
		t.Error("Expect False but True")
	}
	if !xstrings.IsValidPhone("+8618911449894") {
		t.Error("Expect True but False")
	}
	if !xstrings.IsValidPhone("18911449894") {
		t.Error("Expect True but False")
	}
}
