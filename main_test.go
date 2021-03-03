package notifyxf_test

import (
	"notifyxf"
	"os"
	"testing"
)

func TestNotify(t *testing.T) {
	if err := notifyxf.Notify(os.Getenv("NOTIFYXF_TOKEN"), "my message"); err != nil {
		t.Fatal(err.Error())
	}
}
