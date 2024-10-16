package wcImplementation

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func counter(r io.Reader, splitterType func(data []byte, atEOF bool) (advance int, token []byte, err error)) (int, error) {
	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(splitterType)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	if err := fileScanner.Err(); err != nil {
		return count, err
	}

	return count, nil
}

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

func GetFileState(file OSFile, byteSizeFlag bool, linesFlag bool, wordsFlag bool, charsFlag bool) ([]string, error) {
	var fileOutput = make([]string, 0)

	if byteSizeFlag {
		numberOfBytes, err := counter(file, bufio.ScanBytes)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfBytes))
		file.Seek(0, io.SeekStart)
	}

	if linesFlag {
		numberOfLines, err := lineCounter(file)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfLines))
		file.Seek(0, io.SeekStart)
	}

	if wordsFlag {
		numberOfWords, err := counter(file, bufio.ScanWords)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfWords))
		file.Seek(0, io.SeekStart)
	}

	if charsFlag {
		numberOfChars, err := counter(file, bufio.ScanRunes)
		if err != nil {
			return fileOutput, err
		}

		fileOutput = append(fileOutput, strconv.Itoa(numberOfChars))
		file.Seek(0, io.SeekStart)
	}

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
