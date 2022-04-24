package notifyxf

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	url = "https://notifyxf.com/api/message"
)

type Notifier struct {
	token     string
	handle    string
	parseMode string
	redact    bool
	preview   bool
}

// NewNotifier will create a new Notifier service
func NewNotifier(token string, opts ...Option) (*Notifier, error) {
	if token == "" {
		return nil, errors.New("`token` is empty")
	}

	n := &Notifier{token: token}
	for _, opt := range opts {
		opt(n)
	}

	return n, nil
}

// Notify will notify with given `message`
func (n *Notifier) Notify(message string) error {
	if message == "" {
		log.Fatalln(errors.New("`msg` is empty"))
	}

	dat := struct {
		Token     string `json:"access_token"`
		Message   string `json:"message"`
		Handle    string `json:"handle,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		Redact    bool   `json:"redact,omitempty"`
		Preview   bool   `json:"disable_web_page_preview,omitempty"`
	}{
		Token:     n.token,
		Message:   message,
		Handle:    n.handle,
		ParseMode: n.parseMode,
		Redact:    n.redact,
		Preview:   n.preview,
	}

	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(dat)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}
