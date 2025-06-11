package backupper

import (
	"backupper/internal/models/config"
	"backupper/internal/models/task"

	"github.com/sirupsen/logrus"
)

type Backupper struct {
	Config *config.Config
	Logger *logrus.Logger
}

func NewBackuper(config *config.Config, logger *logrus.Logger, tasks []task.Task) *Backupper {
	return &Backupper{
		Config: config,
		Logger: logger,
	}
}
