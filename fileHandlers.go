package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type File struct {
	Name         string `bson:"name" json:"name"`
	OriginalName string `bson:"originalName" json:"originalName"`
	Path         string `bson:"path" json:"path"`
}

func HandleFile(file *multipart.FileHeader, filePrefix string, serverPath string, directoryName string) (File, error) {
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		err := os.MkdirAll(directoryName, 0777)
		if err != nil {
			return File{}, err
		}
	}

	var fileToInsert File

	originalFileName := file.Filename
	fileToInsert.OriginalName = originalFileName
	newFileName := filePrefix + "-" + originalFileName
	fileToInsert.Name = newFileName
	out, err := os.Create(directoryName + newFileName)
	if err != nil {
		return fileToInsert, err
	}

	defer out.Close()

	openFile, errOpen := file.Open()
	if errOpen != nil {
		return fileToInsert, err
	}

	_, errCopy := io.Copy(out, openFile)
	if errCopy != nil {
		return fileToInsert, err
	}

	filepath := serverPath + newFileName
	fileToInsert.Path = filepath

	return fileToInsert, nil
}

func DeleteFile(fileName string, directoryName string) error {
	//open directory
	dirRead, _ := os.Open(directoryName)
	dirFiles, _ := dirRead.Readdir(0)

	for index := range dirFiles {
		currentFile := dirFiles[index]
		// Get name of file and its full path.
		currentFileName := currentFile.Name()
		fullPath := directoryName + currentFileName
		if fullPath == directoryName+fileName {
			_ = os.Remove(fullPath)
			fmt.Println("Removed file:", fullPath)
		}
	}

	return nil
}
