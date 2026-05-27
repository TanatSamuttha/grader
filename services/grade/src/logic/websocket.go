package logic

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var SocketMap map[string]*websocket.Conn;

var SocketMutex sync.RWMutex