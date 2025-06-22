package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileService interface {
	Save(file *multipart.FileHeader) (string, error)
	Get(filename string) (string, error)
	Delete(filename string) error
}

type fileService struct {
	basePath string
}

func NewFileService(path string) (FileService, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return &fileService{basePath: path}, nil
}

func (f fileService) Save(file *multipart.FileHeader) (string, error) {
	dst := filepath.Join(f.basePath, file.Filename)
	if err := SaveUploadedFile(file, dst); err != nil {
		return "", err
	}
	return dst, nil
}
func (f fileService) Get(filename string) (string, error) {
	path := filepath.Join(f.basePath, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	return path, nil
}
func (f fileService) Delete(filename string) error {
	path := filepath.Join(f.basePath, filename)
	return os.Remove(path)
}

func SaveUploadedFile(fh *multipart.FileHeader, dst string) error {
	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, file)
	return err
}
