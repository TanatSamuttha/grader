package logic

import (
	"encoding/json"
	"grade/models"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var SocketMap map[string]*websocket.Conn;

var SocketMutex sync.RWMutex

var GradeResBuffer chan models.GradeResDTO = make(chan models.GradeResDTO);

func SendResult(GradeResBuffer <- chan models.GradeResDTO) {
	for res := range GradeResBuffer {
		jobID := res.JobID;

		SocketMutex.Lock();
		conn := SocketMap[jobID];
		SocketMutex.Unlock();

		res.JobID = "";

		resJson, err := json.Marshal(res);
		if err != nil {
			log.Println("Error json marshal -> " + err.Error());
		}

		conn.WriteMessage(websocket.TextMessage, resJson);
	}
}