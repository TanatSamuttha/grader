package handler

import (
	"grade/logic"

	"github.com/gofiber/contrib/websocket"
)

func ReturnResult(conn *websocket.Conn) {
	jobID := conn.Cookies("job_id");
	logic.SocketMutex.Lock();
	logic.SocketMap[jobID] = conn;
	logic.SocketMutex.Unlock();
}