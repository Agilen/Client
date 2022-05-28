package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	router  *echo.Echo
	logger  *logrus.Logger
	Sockets *Sockets
}

func NewServer() *Server {
	s := &Server{
		router: echo.New(),
		logger: logrus.New(),
	}

	s.configureRouter()

	return s
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

	s := NewServer()

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
