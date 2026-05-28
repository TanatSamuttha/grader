package logic

import (
	"context"
	"grade/config"
	"grade/models"
	"log"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
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
			client.ContainerCreateOptions{
				Config: &container.Config{
					Image: "docker.io/library/gcc:latest",
					WorkingDir: "/workspace",
					Cmd: []string{
						"sleep",
						"infinity",
					},
					Tty: false,
				},	
			},
		);
		if err != nil {
			log.Println("Error create container -> " + err.Error());
			config.DockerClient.ContainerRemove(
				ctx,
				resp.ID,
				client.ContainerRemoveOptions{
					Force: true,
				},
			);
			continue;
		}

		_, err = config.DockerClient.ContainerStart(
			ctx,
			resp.ID,
			client.ContainerStartOptions{},
		);

		if err != nil {
			log.Println("Error start container -> " + err.Error());
			config.DockerClient.ContainerRemove(
				ctx,
				resp.ID,
				client.ContainerRemoveOptions{
					Force: true,
				},
			);
			continue;
		}
		log.Println("Created new container -> " + resp.ID);

		err = Grade(job, &resp, ctx);
		if err != nil {
			log.Println("Error grade -> " + err.Error());
			config.DockerClient.ContainerRemove(
				ctx,
				resp.ID,
				client.ContainerRemoveOptions{
					Force: true,
				},
			);
			continue;
		}

		config.DockerClient.ContainerRemove(
			ctx,
			resp.ID,
			client.ContainerRemoveOptions{
				Force: true,
			},
		);

		SocketMutex.Lock();
		conn := SocketMap[job.ID];
		delete(SocketMap, job.ID);
		SocketMutex.Unlock();

		conn.Close();
	}
}