package main

import (
	"backupper/internal/compress"
	"backupper/internal/errors"
	"backupper/internal/jobs"
	"backupper/internal/log"
	"backupper/internal/models/backupper"
	"backupper/internal/models/config"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "./config.yaml"

var configPathFlag = flag.String("config", CONFIG_PATH, "Path to config file")

func main() {
	flag.Parse()

	config, err := readConfig(*configPathFlag)
	if err != nil {
		fmt.Println("Failed to read config file: ", err)
		os.Exit(1)
	}

	logger, err := log.NewLogger(config.Logger.OutputDir, config.Logger.Format, config.Logger.Level)
	if err != nil {
		fmt.Println("Failed to create logger: ", err)
		os.Exit(1)
	}
	bu := backupper.NewBackuper(config, logger, nil)

	bu.Logger.Info("👋 Hello from backupper's logger")

	for _, task := range bu.Config.Tasks {
		startTime := time.Now()

		filePath, err := jobs.Execute(&task, bu.Logger)
		if err != nil {
			logger.Errorf("Got error when executing task %s: %s", task.Name, err)
			continue
		}

		localFilePath, err := jobs.CopyToLocal(task.User, task.Address, filePath, bu.Config.StoreDir)
		if err != nil {
			logger.Errorf("Got error when copying file to local: %s", err)
			continue
		}

		newFilename := path.Base(localFilePath) + ".zst"
		newFilePath := path.Join(bu.Config.StoreDir, newFilename)

		size, err := compress.Compress(
			localFilePath,
			newFilePath,
		)
		if err != nil {
			logger.Errorf("Got error when compressing file %s: %s", localFilePath, err)
			continue
		}

		newFile, err := os.Stat(newFilePath)
		if err != nil {
			logger.Errorf("Got error when getting file info: %s", err)
			continue
		}

		os.Remove(localFilePath)

		endTime := time.Now()
		jobDuration := formatDuration(endTime.Sub(startTime))

		oldFileSize := formatBytes(size)
		newFileSize := formatBytes(newFile.Size())

		bu.Logger.Infof(
			"✅ in %s | Completed task '%s': 📦 archive name '%s' size: %s ➡️  %s",
			jobDuration,
			task.Name,
			newFilename,
			oldFileSize,
			newFileSize,
		)
	}
}

func readConfig(path string) (*config.Config, error) {
	var config map[string]config.Config

	cfgFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read config file")
		return nil, err
	}

	err = yaml.Unmarshal(cfgFile, &config)
	if err != nil {
		return nil, err
	}

	if cfg, ok := config["backupper"]; ok {
		cfg.StoreDir, err = filepath.Abs(cfg.StoreDir)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
	}
	return nil, errors.ErrConfigNotFound
}

func formatBytes(bytes int64) string {
	kb := bytes / 1024
	mb := kb / 1024

	if mb > 0 {
		return fmt.Sprintf("%.1f MB", float64(bytes)/float64(1024*1024))
	}
	return fmt.Sprintf("%.1f KB", float64(kb))
}

func formatDuration(duration time.Duration) string {
	return fmt.Sprintf("%.2fs", duration.Seconds())
}
