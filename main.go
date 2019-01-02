package main

import (
	"os"
	"flag"
	"fmt"
	"bufio"
	"io"
	"strings"
)

func main() {
	os.Exit(mainFunc())
}

func mainFunc() int {
	flag.Usage = Usage
	
	flag.Parse()

	file, section, key := "-", "", ""
	
	nargs := flag.NArg()

	if(nargs > 3 || nargs == 0) {
		Usage()
	}
	if(nargs == 3) {
		file = flag.Arg(nargs - 3)	
	}
	if(nargs >= 2) {
		section = strings.ToLower(flag.Arg(nargs - 2))	
	}
	if(nargs >= 1) {
		key = strings.ToLower(flag.Arg(nargs - 1))
	}

	reader := os.Stdin
	var err error
	if(file != "-") {
		reader, err = os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open %s: %s\n", file, err)
			return 2
		}
		defer reader.Close()
	}
	bufreader := bufio.NewReader(reader)

	cursection := ""
	err = nil
	var lineBytes []byte
	for {
		if err == io.EOF {
			break
		}

		lineBytes, _, err = bufreader.ReadLine()
		if err != nil && err != io.EOF {			
			fmt.Fprintf(os.Stderr, "Error reading %s: %s\n", file, err)
			return 2
		}
		line := string(lineBytes)
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == ';' {
			continue
		}
		if line[0] == '[' && line[len(line) - 1] == ']' {
			cursection = strings.ToLower(line[1:len(line) - 1])
			continue
		}
		idx := strings.IndexByte(line, '=')
		if idx != -1 {
			curkey := strings.ToLower(strings.TrimSpace(line[0:idx]))
			curval := strings.TrimSpace(line[idx + 1:len(line)])
			if curval[0] == '"' && curval[len(curval) -1] == '"' {
				curval = curval[1:len(curval) - 1]
			}
			if curkey == key && cursection == section {
				fmt.Printf("%s\n", curval)
				return 0
			}
		}
		
	}
	
	return 1
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
        fmt.Fprintf(os.Stderr, "  %s [FILE] [SECTION] VARIABLE\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  -h, --help: Print this help\n")
	os.Exit(0)
}
