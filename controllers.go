package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Agilen/Client/commands"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	BaseURL  = "http://localhost:10000"
	JSONType = "application/json"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

type (
	UserInfo struct {
		Nickname string
		Login    string
	}

	RegistrationRequest struct {
		Login    string
		Mail     string
		Password string
	}

	RegistrationResponce struct {
		Nickname string
		Login    string
	}

	LoginRequest struct {
		Login    string
		Password string
	}

	LoginResponce struct {
		Nickname string
		Login    string
	}

	FindUserRequest struct {
		Parm string
	}

	FindUserResponce struct {
		Users []UserInfo
	}

	SendMessageRequest struct {
		To      string //Login
		Message string
	}
)

func (s *Server) CreateUserController(c echo.Context) error {
	req := new(RegistrationRequest)
	if err := c.Bind(req); err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	resp, err := commands.POST(Join(BaseURL, c.Request().URL.String()), JSONType, body)
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	var r RegistrationResponce
	if err := json.Unmarshal(resp, &r); err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, r)
}

func (s *Server) LoginController(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}
	fmt.Println("1")
	body, err := json.Marshal(req)
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}
	fmt.Println("1")
	fmt.Println(string(body))
	resp, err := commands.POST(Join(BaseURL, c.Request().URL.String()), JSONType, body)
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}
	fmt.Println("1")
	var r LoginResponce
	if err := json.Unmarshal(resp, &r); err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}
	fmt.Println("1")
	return c.JSON(http.StatusOK, r)
}

func (s *Server) FindUserController(c echo.Context) error {
	resp, err := commands.GET(c.Request().URL.String())
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusBadRequest)
	}

	var r FindUserResponce
	if err := json.Unmarshal(resp, &r); err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, r)
}

func (s *Server) SendMessageController(c echo.Context) error {
	req := new(SendMessageRequest)
	if err := c.Bind(req); err != nil {
		return s.HttpErrorHandler(c, err, http.StatusBadRequest)
	}

	body, err := json.Marshal(&req)
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	_, err = commands.POST(Join(BaseURL), JSONType, (body))
	if err != nil {
		return s.HttpErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, nil)
}
