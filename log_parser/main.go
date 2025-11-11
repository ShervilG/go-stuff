package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg *sync.WaitGroup

func parseInfoLog(line chan string, file *os.File) {
	for line := range line {
		file.WriteString(line + "\n")
	}
	wg.Done()
}

func parseWarnLog(line chan string, file *os.File) {
	for line := range line {
		file.WriteString(line + "\n")
	}
	wg.Done()
}

func parseErrorLog(line chan string, file *os.File) {
	for line := range line {
		file.WriteString(line + "\n")
	}
	wg.Done()
}

func main() {
	wg = &sync.WaitGroup{}
	infoLineChan := make(chan string, 1)
	warnLineChan := make(chan string, 1)
	errorLineChan := make(chan string, 1)

	infoFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer infoFile.Close()

	warnFile, err := os.OpenFile("warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer warnFile.Close()

	errorFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	wg.Add(3)
	go parseInfoLog(infoLineChan, infoFile)
	go parseWarnLog(warnLineChan, warnFile)
	go parseErrorLog(errorLineChan, errorFile)

	readFile, err := os.OpenFile("/home/shervil/Documents/repos/go-stuff/log_parser/sample_app.log", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	buffReader := bufio.NewReader(readFile)
	s, e := buffReader.ReadString('\n')
	for e == nil {
		if len(s) >= 5 {
			if strings.Contains(s, "INFO") {
				infoLineChan <- s
			} else if strings.Contains(s, "WARN") {
				warnLineChan <- s
			} else if strings.Contains(s, "ERROR") {
				errorLineChan <- s
			}
		}
		s, e = buffReader.ReadString('\n')
	}

	close(infoLineChan)
	close(warnLineChan)
	close(errorLineChan)

	wg.Wait()
	fmt.Println("Log parsing completed.")
}
