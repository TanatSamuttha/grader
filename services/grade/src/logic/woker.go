package logic

import (
	"context"
	"grade/config"
	"grade/models"
	"log"
	"sync"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

var jobs chan models.Job;

var WorkingJobs map[string]bool = make(map[string]bool);
var WorkingJobsMutex sync.RWMutex;

func SummonWorkers(n int) {
	jobs = make(chan models.Job, n);
	for i := 0; i < n; i++ {
		go worker();
	}
}

func CallWorker(job models.Job) {
	jobs <- job;
}

func worker() {
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
				HostConfig: &container.HostConfig{
					Resources: container.Resources{
						Memory:     512 * 1024 * 1024,
						MemorySwap: 512 * 1024 * 1024,
						NanoCPUs:   1_000_000_000,
					},
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

		WorkingJobsMutex.Lock();
		delete(WorkingJobs, job.ID);
		WorkingJobsMutex.Unlock();

		if conn != nil {
			conn.Close();
			log.Println("WebSocket connection closed");
		}
	}
}