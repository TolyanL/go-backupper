package task

import "backupper/internal/models/job"

type Task struct {
	Name     string         `yaml:"name"`
	User     string         `yaml:"user"`
	Address  string         `yaml:"address"`
	Job      job.Job        `yaml:"job"`
	Postgres job.Postgresql `yaml:"postgresql"`
}
