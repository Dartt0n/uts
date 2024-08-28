package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const VERSION = `v0.2.0`

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

type userFormat interface {
	fmt.Stringer
	Match(string) bool
	Parse(string) (time.Time, error)
}

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

	userFormats := []userFormat{sFmt{}, msFmt{}, usFmt{}, nsFmt{}, fsFmt{}}

	anyFormatUsed := false
	for _, format := range userFormats {
		if !format.Match(userValue) {
			continue
		}

		timestamp, err := format.Parse(userValue)
		if err != nil {
			fmt.Printf("failed to parse \"%s\" value using \"%s\" format with error: %s\n", userValue, format, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", timestamp.Format(time.RFC1123))
		anyFormatUsed = true
		break
	}

	if !anyFormatUsed {
		fmt.Printf("failed to parse \"%s\" value using any of the supported formats\n", userValue)
		os.Exit(1)
	}
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

	if stats.Mode()&os.ModeNamedPipe == 0 {
		return ""
	}

	buffer, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}

	content := strings.Trim(string(buffer), " \n")

	return content
}
