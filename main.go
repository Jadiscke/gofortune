package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var files []string

func visit (path string, f os.FileInfo, err error) error {
	if strings.Contains(path, "/off/") {
		return nil
	}

	if filepath.Ext(path) == ".dat" {
		return nil
	}

	if f.IsDir() {
		return nil
	}

	files = append(files, path)
	return nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func printQuote(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	quotes := string(b)

	quotesSlice := strings.Split(quotes, "%")
	j := randomInt(1, len(quotesSlice))

	fmt.Println(quotesSlice[j])
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fortuneCommand := exec.Command("fortune", "-f")
	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}
	fortuneCommand.Start()
	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()
	line := outputStream.Text()
	root := line[strings.Index(line,"/"):]

	err = filepath.Walk(root,visit)

	if err != nil {
		panic(err)
	}

	i := randomInt(1, len(files))
	randomFile := files[i]
	err = printQuote(randomFile)

	if err != nil {
		panic(err)
	}
}