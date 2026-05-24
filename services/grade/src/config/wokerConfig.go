package config

import (
	"context"
	"grade/logic"
	"grade/models"
	"log"

	"github.com/docker/docker/api/types/container"
)

var jobs chan models.Job;

func SummonWorkers(n int) {
	jobs = make(chan models.Job, n);
	for i := 0; i < n; i++ {
		go worker(jobs);
	}
}

func CallWorker(job models.Job) {
	jobs <- job;
}

func worker(jobs <-chan models.Job) {
	ctx := context.Background();
	for job := range jobs {
		container, err := DockerClient.ContainerCreate(
			ctx,
			&container.Config{
				Image: "docker.io/library/gcc:latest",
				Tty: false,
			},
			nil,
			nil,
			nil,
			"",
		);
		if err != nil {
			log.Println("Error create container -> " + err.Error());
		}
		log.Println("Created new container -> " + container.ID);
		logic.Grade(job, &container);
	}
}