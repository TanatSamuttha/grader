package logic

import (
	"context"
	"grade/config"
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
		resp, err := config.DockerClient.ContainerCreate(
			ctx,
			&container.Config{
				Image: "docker.io/library/gcc:latest",
				WorkingDir: "/workspace",
				Cmd: []string{
					"sleep",
					"60",
				},
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

		err = config.DockerClient.ContainerStart(
			ctx,
			resp.ID,
			container.StartOptions{},
		)

		if err != nil {
			log.Println("Error start container -> " + err.Error())
			continue
		}
		log.Println("Created new container -> " + resp.ID);

		err = Grade(job, &resp, ctx);
		if err != nil {
			log.Println("Error grade -> " + err.Error())
			continue
		}

		// config.DockerClient.ContainerRemove(
		// 	ctx,
		// 	resp.ID,
		// 	container.RemoveOptions{
		// 		Force: true,
		// 	},
		// );
	}
}