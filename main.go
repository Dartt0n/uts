package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const VERSION = `v0.1.1`

const HELP_MESSAGE = `Convert unix timestamp to human readable format

Usage: uts [options] <unix timestamp>

Options:
  -h, --help      Show this help message
  -v, --version   Show version

Examples:
# seconds precision
$ uts 1724692825
> Mon, 26 Aug 2024 20:20:25 UTC

# nanoseconds precision
$ uts 1723140436809000000
> Thu, 08 Aug 2024 21:07:16 UTC

# pipe from stdin
$ echo 1724692825 | uts
> Mon, 26 Aug 2024 20:20:25 UTC`

func main() {
	var (
		help, version bool
	)

	flag.BoolVar(&help, "h", false, "Show this help message")
	flag.BoolVar(&help, "help", false, "Show this help message")
	flag.BoolVar(&version, "v", false, "Show version")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.Parse()

	if help {
		fmt.Println(HELP_MESSAGE)
		os.Exit(0)
	}

	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	userValue := readFromArgs()
	if userValue == "" {
		userValue = readFromStdin()
	}
	if userValue == "" {
		fmt.Printf(HELP_MESSAGE)
		os.Exit(1)
	}

	unixtime, err := strconv.ParseInt(userValue, 10, 64)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	var timestamp time.Time
	// todo: add support for milliseconds precision, floating point unixtime and different formats
	if unixtime > math.MaxInt32 { // 19 Jan 2038 06:14:07
		timestamp = time.Unix(unixtime/1000000000, unixtime%1000000000)
	} else {
		timestamp = time.Unix(unixtime, 0)
	}

	fmt.Printf("%s\n", timestamp.Format(time.RFC1123))
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

	stats, err := reader.Stat()
	if err != nil {
		return ""
	}

	if stats.Mode()&os.ModeNamedPipe != 0 {
		return ""
	}

	buffer, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}

	content := strings.Trim(string(buffer), " \n")

	return content
}
