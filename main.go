package notifyxf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	errs "github.com/go-errors/errors"
)

const (
	url = "https://notifyxf.com/api/message"
)

type message struct {
	Token   string `json:"access_token"`
	Message string `json:"message"`
}

// Notify will notify with message `msg`
func Notify(token, msg string) error {
	if token == "" {
		log.Fatalln(errors.New("`token` is empty"))
	}

	if msg == "" {
		log.Fatalln(errors.New("`msg` is empty"))
	}

	dat := message{
		Token:   token,
		Message: msg,
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

// Recover to use as a middleware for gin to notify panics with `Notify`
func Recover(token, title string) gin.HandlerFunc {
	if token == "" {
		panic(errors.New("no token specified"))
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				pc, fn, line, _ := runtime.Caller(2)
				err := Notify(token, fmt.Sprintf("Error in `%s` (%s[%s:%d]): `%v`\n\n%s", title, runtime.FuncForPC(pc).Name(), fn, line, err, errs.Wrap(err, 2).ErrorStack()))
				if err != nil {
					log.Println(err)
				}
			}
		}()

		c.Next()
	}
}
