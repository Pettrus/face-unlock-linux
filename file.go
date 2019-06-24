package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Kagami/go-face"
)

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return linesFromReader(f)
}

func SaveFaceDescriptions(face face.Descriptor) {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Create("/lib/security/go-face-unlock/faces/" + fmt.Sprintf("%d.txt", rand.Int()))
	if err != nil {
		log.Fatalf("Error saving descriptions")
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range face {
		fmt.Fprintln(w, line)
	}

	w.Flush()
}

/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func RemoveStringFromFile(path, str string) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		if line != str {
			fileContent += line
			fileContent += "\n"
		}
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func ReturnFilesOnFolder(directory string) []os.FileInfo {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	return files
}
