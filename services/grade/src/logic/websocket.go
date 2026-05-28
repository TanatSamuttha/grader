package logic

import (
	"encoding/json"
	"grade/models"
	"log"
	"sync"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
)

var SocketMap map[string]*websocket.Conn;

var SocketMutex sync.RWMutex

var GradeResBuffer chan models.GradeResDTO = make(chan models.GradeResDTO);

func SendResult(GradeResBuffer <- chan models.GradeResDTO) {
	for res := range GradeResBuffer {
		jobID := res.JobID;

		var conn *websocket.Conn

		for {
			SocketMutex.RLock()
			conn = SocketMap[jobID]
			SocketMutex.RUnlock()

			if conn != nil {
				break
			}

			time.Sleep(200 * time.Millisecond);
		}

		res.JobID = "";

		resJson, err := json.Marshal(res);
		if err != nil {
			log.Println("Error json marshal -> " + err.Error());
		}

		conn.WriteMessage(websocket.TextMessage, resJson);
	}
}