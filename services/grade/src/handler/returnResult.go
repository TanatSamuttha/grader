package handler

import (
	"grade/logic"

	"github.com/gofiber/contrib/v3/websocket"
)

func ReturnResult(conn *websocket.Conn) {
	jobID := conn.Cookies("job_id");
	logic.SocketMutex.Lock();
	logic.SocketMap[jobID] = conn;
	logic.SocketMutex.Unlock();
}