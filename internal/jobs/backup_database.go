package jobs

import (
	"fmt"
	"path"
	"strings"
	"time"
)

func BackupDatabaseCmd(workDirPath, containerName, dbName, dbUser string) (string, string) {
	dateNow := time.Now().Format("2006-01-02_15-04")

	archiveFilename := fmt.Sprintf("db_%s_%s.tar", dbName, dateNow)
	dbDumpFilename := fmt.Sprintf("%s_dump.sql", dbName)

	dumpFilePath := path.Join(workDirPath, dbDumpFilename)
	arichiveFilePath := path.Join(workDirPath, archiveFilename)

	dbDumpCmd := fmt.Sprintf(
		"docker exec -i %s pg_dump -U %s -d %s > %s",
		containerName,
		dbUser,
		dbName,
		dumpFilePath,
	)
	createArchiveCmd := fmt.Sprintf(
		"tar -cf %s -C %s %s",
		arichiveFilePath,
		workDirPath,
		dbDumpFilename,
	)
	removeDbDumpCmd := fmt.Sprintf("rm %s", path.Join(workDirPath, dbDumpFilename))

	commands := []string{
		dbDumpCmd,
		createArchiveCmd,
		removeDbDumpCmd,
	}

	cmd := strings.Join(commands, " && ")
	return arichiveFilePath, cmd
}
