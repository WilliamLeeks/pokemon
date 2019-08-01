package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadLog ...
func ReadLog(path string) ([]int, error) {
	var ids []int

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return ids, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		id, err := strconv.Atoi(s)
		if err == nil {
			ids = append(ids, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return ids, err
	}

	return ids, nil
}

// WriteLog ...
func WriteLog(path, data string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(data)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

// Exists checks if a file exists
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Remove deletes the file or directory at path
func Remove(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Println(err)
	}
}
