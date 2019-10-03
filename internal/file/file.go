package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadLog ...
func ReadLog(path string) ([]int, error) {
	var ids []int

	//f, err := os.Open(path)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		return ids, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(line[1])
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ids, nil
}

// WriteLog ...
func WriteLog(path, line string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(line)
	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()
	if err != nil {
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
