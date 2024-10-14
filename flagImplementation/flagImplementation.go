package flagImplementation

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func wordsCounter(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanWords)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	if err := fileScanner.Err(); err != nil {
		return count, err
	}

	return count, nil
}

func charsCounter(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanRunes)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	if err := fileScanner.Err(); err != nil {
		return count, err
	}

	return count, nil
}

func GetFileState(filePath string, byteSizeFlag bool, linesFlag bool, wordsFlag bool, charsFlag bool) ([]string, error) {
	var fileOutput = make([]string, 0)

	if byteSizeFlag {
		fileStat, err := os.Stat(filePath)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(int(fileStat.Size())))
	}

	if linesFlag {
		file, err := os.Open(filePath)
		if err != nil {
			return fileOutput, err
		}
		defer file.Close()

		numberOfLines, err := lineCounter(file)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfLines))
	}

	if wordsFlag {
		file, err := os.Open(filePath)
		if err != nil {
			return fileOutput, err
		}
		defer file.Close()

		numberOfWords, err := wordsCounter(file)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfWords))
	}

	if charsFlag {
		file, err := os.Open(filePath)
		if err != nil {
			return fileOutput, err
		}
		defer file.Close()

		numberOfChars, err := charsCounter(file)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfChars))
	}

	fileOutput = append(fileOutput, filePath)
	return fileOutput, nil
}

func PrintFileOutput(fileOutput []string) {
	outputFormatted := make([]string, 0)
	for i, c := range fileOutput {
		if i == len(fileOutput)-1 {
			outputFormatted = append(outputFormatted, fmt.Sprintf("%s", c))
		} else {
			outputFormatted = append(outputFormatted, fmt.Sprintf("%12s", c))
		}
	}
	fmt.Println(strings.Join(outputFormatted, " "))
}
