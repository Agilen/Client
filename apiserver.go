package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Agilen/Client/clientcrypto"
	"github.com/Agilen/Client/commands"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"honnef.co/go/tools/config"
)

type HttpError struct {
	Controller string
	Method     string
	Body       string
	Error      string
}

type Server struct {
	router    *echo.Echo
	logger    *logrus.Logger
	sockets   *Sockets
	cc        *clientcrypto.CryptoContext
	serverKey []byte
}

func NewServer() (*Server, error) {
	s := &Server{
		router: echo.New(),
		logger: logrus.New(),
	}

	var err error
	s.cc, err = clientcrypto.NewCryptoContext()
	if err != nil {
		return nil, err
	}

	s.serverKey, err = commands.GET(Join(BaseURL, "/public"))
	if err != nil {
		return nil, err
	}
	fmt.Println(s.serverKey)

	s.configureRouter()

	return s, nil
}

func (s *Server) configureRouter() {
	app := s.router
	s.router.POST("/reg", s.CreateUserController)
	app.POST("/login", s.LoginController)

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func Start(config *config.Config) error {

	s, err := NewServer()
	if err != nil {
		return err
	}

	return http.ListenAndServe(":10001", s)
}

func (s *Server) HttpErrorHandler(c echo.Context, err error, httpErr int) error {
	url := c.Request().URL.String()
	method := c.Request().Method
	body, _ := ioutil.ReadAll(c.Request().Body)
	k, _ := json.MarshalIndent(HttpError{Controller: url, Method: method, Body: string(body), Error: err.Error()}, "", "\t")
	s.logger.Error(string(k))

	return c.JSON(httpErr, err.Error())
}
