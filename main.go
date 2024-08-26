package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	content := readFromArgs()
	if content == "" {
		content = readFromStdin()
	}
	if content == "" {
		fmt.Printf("Usage: uts <unix timestamp>\n")
		os.Exit(1)
	}

	timestamp, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	var unixTimestamp time.Time

	// Dec 31 2100 23:59:59
	if timestamp > math.MaxInt32 {
		unixTimestamp = time.Unix(timestamp/1000000000, timestamp%1000000000)
	} else {
		unixTimestamp = time.Unix(timestamp, 0)
	}

	fmt.Printf("%s\n", unixTimestamp.Format(time.RFC1123))
}

func readFromArgs() string {
	args := os.Args
	if len(args) < 2 {
		return ""
	}

	return args[1]
}

func readFromStdin() string {
	reader := os.Stdin

	buffer, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}

	content := strings.Trim(string(buffer), " \n")

	return content
}
