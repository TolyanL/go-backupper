package compress

import (
	"backupper/internal/errors"
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
)

func Compress(srcPath, dstPath string) (int64, error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return 0, errors.ErrFileNotFound
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	encoder, err := zstd.NewWriter(dstFile)
	if err != nil {
		return 0, err
	}

	written, err := io.Copy(encoder, srcFile)
	if err != nil {
		return 0, err
	}

	err = encoder.Close()
	if err != nil {
		return written, err
	}

	return written, nil
}
