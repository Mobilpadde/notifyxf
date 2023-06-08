package notifyxf_test

import (
	"os"
	"testing"

	"github.com/Mobilpadde/notifyxf/v2"
)

func TestHelloWorld(t *testing.T) {
	n, err := notifyxf.NewNotifier(os.Getenv("NOTIFYXF_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}

	n.Notify("Hello, world!")
	t.Log("Notified!")
}

func TestWithHandle(t *testing.T) {
	n, err := notifyxf.NewNotifier(os.Getenv("NOTIFYXF_TOKEN"), notifyxf.WithHandle(os.Getenv("NOTIFYXF_HANDLE")))
	if err != nil {
		t.Fatal(err)
	}

	n.Notify("Hello, world!")
	t.Log("Notified!")
}
