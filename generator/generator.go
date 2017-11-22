package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"bufio"
	"io"
	"strings"
)

func readFile(w io.Writer, path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	isFirst := true
	for scanner.Scan() {
		if isFirst == false {
			p := strings.Split(scanner.Text(), ",")
			fmt.Fprint(w, "\t\"" + p[0] + "\":" + p[1] + ",\n")
		} else {
			isFirst = false
		}
	}
}
func main() {
	buffer := bytes.NewBuffer(nil)
	fmt.Fprint(buffer, `package main
var names = map[string]int{
`)
	files, err := filepath.Glob(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	for _, f := range files {
		readFile(buffer, f)
	}

	fmt.Fprint(buffer, `}
func getNames(n string) (int, bool) {
	m, ok := names[n]
	return m, ok
}`)
	if err := ioutil.WriteFile("names.go", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}
