package utils

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	"github.com/golkhandani/taskmate/exceptions"
	"github.com/golkhandani/taskmate/models"
)

func ReadDataFile() (*os.File, []byte) {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	path := filepath.Join(dirname, "..", "tasks.json")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	exceptions.HandleErr(err)
	stats, err := file.Stat()
	exceptions.HandleErr(err)
	fileContent := make([]byte, stats.Size())
	file.Read(fileContent)
	return file, fileContent
}

func SaveDataFile(file *os.File, tasks *[]models.Task) {
	file.Seek(0, 0)
	file.Truncate(0)

	updatedContent, err := json.MarshalIndent(*tasks, "", " ")
	exceptions.HandleErr(err)

	writer := bufio.NewWriter(file)

	_, err = writer.Write(updatedContent)
	exceptions.HandleErr(err)
	writer.Flush()
}
