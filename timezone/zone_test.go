package timezone

import (
	"testing"
	"time"
)

func TestSetGlobalTimeLocalToChina(t *testing.T) {
	loc, _ := time.LoadLocation("America/Atka")
	time.Local = loc
	t.Log(time.Local, time.Now())
	if time.Local.String() != "America/Atka" {
		t.Fatal("error time zone")
	}

	SetGlobalTimeLocalToChina()
	t.Log(time.Local, time.Now())
	if time.Local.String() != "Asia/Shanghai" {
		t.Fatal("error time zone")
	}
}
