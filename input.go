package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func read(reader io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(reader)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []byte{}, err
	}

	content := bytes.NewBuffer(make([]byte, 0, 1024))

	for _, line := range lines {
		content.WriteString(line)

		content.WriteString("\n")
	}

	return content.Bytes(), nil
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return []byte{}, err
	}

	defer file.Close()

	return read(file)
}

func readPipe() ([]byte, error) {
	return read(os.Stdin)
}
