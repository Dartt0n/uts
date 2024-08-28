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
	// todo: parse flags like -h, --help, -v, --version
	userValue := readFromArgs()
	if userValue == "" {
		userValue = readFromStdin()
	}
	if userValue == "" {
		fmt.Printf(`Usage: uts <unix timestamp>

Examples:
# seconds precision
$ uts 1724692825
> Mon, 26 Aug 2024 20:20:25 UTC

# nanoseconds precision
$ uts 1723140436809000000
> Thu, 08 Aug 2024 21:07:16 UTC

# pipe from stdin
$ echo 1724692825 | uts
> Mon, 26 Aug 2024 20:20:25 UTC
`)
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

	if stats.Size() == 0 {
		return ""
	}

	buffer, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}

	content := strings.Trim(string(buffer), " \n")

	return content
}
