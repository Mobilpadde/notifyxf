package notifyxf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	url = "https://notifyxf.com/api/message"
)

// Message is to contain message Data for NotifyXF
type Message struct {
	Token   string `json:"access_token"`
	Message string `json:"message"`
}

// Notify will notify with message `msg`
func Notify(token, msg string) error {
	if token == "" {
		log.Fatalln(errors.New("no token specified"))
	}

	if msg == "" {
		log.Fatalln(errors.New("no msg is empty"))
	}

	dat := Message{
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
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return nil
}

// Middleware to use as a middleware for gin to notify errors with `Notify`
func Middleware(token string) gin.HandlerFunc {
	if token == "" {
		log.Fatalln(errors.New("no token specified"))
	}

	return func(c *gin.Context) {
		c.Next()
		ginErr := c.Errors.Last()
		if ginErr == nil {
			return
		}

		err := Notify(token, ginErr.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
}
