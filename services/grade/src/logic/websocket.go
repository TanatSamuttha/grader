package logic

import (
	"encoding/json"
	"grade/models"
	"log"
	"sync"

	"github.com/gofiber/contrib/v3/websocket"
)

var SocketMap map[string]*websocket.Conn;

var SocketMutex sync.RWMutex;

var GradeResBuffer chan models.GradeResJob = make(chan models.GradeResJob);

func SendResult() {
	SocketMap = make(map[string]*websocket.Conn);
	for res := range GradeResBuffer {
		jobID := res.JobID;
		log.Println("Assign websocket job" + jobID);

		resDTO := models.GradeResDTO{
			Task: res.Task,
			Result: res.Result,
			Error: res.Error,
		}

		resJson, err := json.Marshal(resDTO);
		log.Println("send -> " + string(resJson));
		if err != nil {
			log.Println("Error json marshal -> " + err.Error());
		}

		conn := res.Conn;
		err = conn.WriteMessage(websocket.TextMessage, resJson)
		if err != nil {
			log.Println("write websocket error ->", err.Error())
		}
	}
}