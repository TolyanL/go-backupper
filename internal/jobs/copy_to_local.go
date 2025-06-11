package jobs

import (
	"fmt"
	"os/exec"
	"path"
)

func CopyToLocal(user, host, srcPath, dstPath string) (string, error) {
	remotePath := fmt.Sprintf("%s@%s:%s", user, host, srcPath)
	cmd := exec.Command("scp", remotePath, dstPath)

	err := cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	localPath := path.Join(dstPath, path.Base(srcPath))

	return localPath, nil
}
