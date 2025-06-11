package jobs

import (
	"path"
	"time"
)

func CreateWorkDir(workDirPath string, taskName string) string {
	bakDir := path.Join(workDirPath, ".backupper")
	dateNow := time.Now().Format("15-04_02-01-2006")
	return path.Join(bakDir, dateNow, taskName)
}
