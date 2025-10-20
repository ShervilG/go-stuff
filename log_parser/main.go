package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg *sync.WaitGroup

func parseInfoLog(line string, file *os.File) {
	file.WriteString(line + "\n")
	wg.Done()
}

func parseWarnLog(line string, file *os.File) {
	file.WriteString(line + "\n")
	wg.Done()
}

func parseErrorLog(line string, file *os.File) {
	file.WriteString(line + "\n")
	wg.Done()
}

func main() {
	wg = &sync.WaitGroup{}

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
				wg.Add(1)
				go parseInfoLog(s, infoFile)
			} else if strings.Contains(s, "WARN") {
				wg.Add(1)
				go parseWarnLog(s, warnFile)
			} else if strings.Contains(s, "ERROR") {
				wg.Add(1)
				go parseErrorLog(s, errorFile)
			}
		}
		s, e = buffReader.ReadString('\n')
	}

	wg.Wait()
	fmt.Println("Log parsing completed.")
}
