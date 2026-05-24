package logic

import (
	"grade/models"
	"log"

	"github.com/docker/docker/api/types/container"
)

func Grade(job models.Job, container *container.CreateResponse) {
	log.Println(job);
	log.Println(container);
}