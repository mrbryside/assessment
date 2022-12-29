package util

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	serverPort = ":2565"
)

type testHelper struct{}

func newTestHelper() testHelper {
	return testHelper{}
}

func (t testHelper) Uri(paths ...string) string {
	host := fmt.Sprintf("http://localhost%v", serverPort)
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func (t testHelper) Request(method, url string, body io.Reader) *httpResponse {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &httpResponse{res, err}
}

func (t testHelper) InitItEcho(eh *echo.Echo, setRoute func()) *echo.Echo {
	go func(e *echo.Echo) {
		e.Validator = Validator(validator.New())
		setRoute()
		eh.Start(serverPort)
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost%v", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
	return eh
}
func (t testHelper) Seeder(model interface{}, payload string, args ...string) error {
	body := bytes.NewBufferString(payload)
	err := t.Request(http.MethodPost, t.Uri(args...), body).Decode(&model)
	if err != nil {
		return err
	}
	return nil
}
