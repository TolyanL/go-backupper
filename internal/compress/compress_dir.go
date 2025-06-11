package compress

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zstd"
)

func CompressDir(logger *log.Logger, paths []string, dstPath string) (int64, error) {
	tarWriter, zstdWriter, dstFile, err := createTarZstdWriter(dstPath)
	if err != nil {
		return 0, err
	}

	var totalWritten int64 = 0

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			logger.Printf("Ошибка при получении статистики для файла %s: %v", path, err)
			continue
		}

		if info.IsDir() {
			written, err := addDirToTar(path, tarWriter)
			if err != nil {
				logger.Printf("Ошибка при добавлении директории %s: %v", path, err)
			}
			totalWritten += written
		} else {
			written, err := addFileToTar(path, tarWriter)
			if err != nil {
				logger.Printf("Ошибка при добавлении файла %s: %v", path, err)
			}
			totalWritten += written
		}
	}

	if err := tarWriter.Close(); err != nil {
		return totalWritten, err
	}
	if err := zstdWriter.Close(); err != nil {
		return totalWritten, err
	}
	if err := dstFile.Close(); err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func createTarZstdWriter(dstPath string) (*tar.Writer, *zstd.Encoder, *os.File, error) {
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return nil, nil, nil, err
	}

	zstdWriter, err := zstd.NewWriter(dstFile)
	if err != nil {
		dstFile.Close()
		return nil, nil, nil, err
	}

	tarWriter := tar.NewWriter(zstdWriter)
	return tarWriter, zstdWriter, dstFile, nil
}

func addFileToTar(filePath string, tarWriter *tar.Writer) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return 0, err
	}
	header.Name = filepath.Base(filePath)

	if err := tarWriter.WriteHeader(header); err != nil {
		return 0, err
	}

	written, err := io.Copy(tarWriter, file)
	return written, err
}

func addDirToTar(dirPath string, tarWriter *tar.Writer) (int64, error) {
	var totalWritten int64 = 0

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == dirPath && info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(filepath.Dir(dirPath), path)
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			written, err := io.Copy(tarWriter, file)
			totalWritten += written
			return err
		}

		return nil
	})

	return totalWritten, err
}
