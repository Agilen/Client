package client

import (
	"errors"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type Sockets struct {
	ClientConn *websocket.Conn
	ServerConn *websocket.Conn
}

func NewSockets() *Sockets {
	return &Sockets{}
}

func (s *Sockets) ClientSocket(path string) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	s.ClientConn = conn

	for {
		mType, mess, err := conn.ReadMessage()
		if mType == websocket.CloseMessage {
			return errors.New("conn closed")
		}
		if err != nil {
			return err
		}

		err = s.ServerConn.WriteMessage(mType, mess)
		if err != nil {
			return err
		}

	}

}

func (s *Sockets) ServerSocket(c echo.Context) error {
	u := websocket.Upgrader{}

	conn, err := u.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return err
	}

	s.ServerConn = conn

	return nil
}
