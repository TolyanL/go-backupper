package jobs

import (
	"backupper/internal/errors"
	"backupper/internal/models/task"
	"backupper/internal/ssh"

	"github.com/sirupsen/logrus"
)

func Execute(task *task.Task, logger *logrus.Logger) (string, error) {
	connection, err := ssh.NewSSHConfig(task.User, task.Address)
	if err != nil {
		return "", err
	}

	workDir := CreateWorkDir(task.Job.WorkDirPath, task.Name)
	ssh.RunCommand(connection, task.Address, "mkdir -p "+workDir)

	if task.Job.UseCommand != "" {
		_, err := ssh.RunCommand(connection, task.Address, task.Job.UseCommand)
		if err != nil {
			logger.Error("Error when executing user command: ", err)
		}
	}

	if task.Postgres.Database != "" {
		filePath, backupDbCmd := BackupDatabaseCmd(
			workDir,
			task.Job.ContainerName,
			task.Postgres.Database,
			task.Postgres.User,
		)

		_, err := ssh.RunCommand(connection, task.Address, backupDbCmd)
		if err != nil {
			return "", errors.ErrDataBaseBak
		}
		return filePath, nil
	}

	if task.Job.MediaPath != "" {
		filePath, backupMediaCmd := BackupMediaCmd(
			task.Job.ContainerName,
			task.Job.MediaPath,
			workDir,
		)

		for _, cmd := range backupMediaCmd {
			_, err := ssh.RunCommand(connection, task.Address, cmd)
			if err != nil {
				return "", errors.ErrMediaBak
			}
		}
		return filePath, nil
	}

	return "", errors.ErrWrongTaskBody
}
