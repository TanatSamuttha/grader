package handler

import (
	"grade/logic"
	"log"

	"github.com/gofiber/contrib/v3/websocket"
)

func ReturnResult(conn *websocket.Conn) {
	jobID := conn.Cookies("job_id");
	log.Println("WebSocket connected " + jobID + conn.IP());
	logic.SocketMutex.Lock();
	logic.SocketMap[jobID] = conn;
	logic.SocketMutex.Unlock();
	
	defer func() {
		log.Println("WebSocket disconnected ->", jobID)

		logic.SocketMutex.Lock()
		delete(logic.SocketMap, jobID)
		logic.SocketMutex.Unlock()
		
		conn.Close()
	}()

	logic.WorkingJobsMutex.Lock();
	jobExist := logic.WorkingJobs[jobID];
	logic.WorkingJobsMutex.Unlock();
	if !jobExist {
		return;
	}

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error ->", err.Error())
			break
		}
	}
}