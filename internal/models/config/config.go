package config

import (
	"backupper/internal/models/logger"
	"backupper/internal/models/task"
)

type Config struct {
	StoreDir string         `yaml:"store_dir"`
	Logger   *logger.Logger `yaml:"logger"`
	Tasks    []task.Task    `yaml:"tasks"`
}
