package jobs

import (
	"fmt"
	"path"
	"time"
)

func BackupMediaCmd(containerName, mediaPath, workDirPath string) (string, []string) {
	date := time.Now().Format("2006-01-02_15-04")
	archiveFilename := fmt.Sprintf("media_%s_%s.tar", containerName, date)

	archiveMediaCmd := fmt.Sprintf(
		"docker exec -i %s tar -cf %s %s",
		containerName,
		archiveFilename,
		mediaPath,
	)
	cpArchiveCmd := fmt.Sprintf(
		"docker cp -q %s:./code/%s %s",
		containerName,
		archiveFilename,
		workDirPath,
	)
	rmArchiveCmd := fmt.Sprintf(
		"docker exec -i %s rm %s",
		containerName,
		archiveFilename,
	)

	archiveFilePath := path.Join(workDirPath, archiveFilename)

	return archiveFilePath, []string{archiveMediaCmd, cpArchiveCmd, rmArchiveCmd}
}
