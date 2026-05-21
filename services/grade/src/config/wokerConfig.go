package config

import (
	"grade/logic"
	"grade/models"
)

var jobs chan models.Job;

func SummonWorkers(n int) {
	for i := 0; i < n; i++ {
		go worker(jobs);
	}
}

func CallWorker(job models.Job) {
	jobs <- job;
}

func worker(jobs <-chan models.Job) {
	for job := range jobs {
		logic.Grade(job);
	}
}