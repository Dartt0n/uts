package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type sFmt struct{}

func (u sFmt) String() string { return "unix-seconds" }

func (u sFmt) Match(value string) bool {
	return regexp.MustCompile(`^\d{1,11}$`).Match([]byte(value))
}

func (u sFmt) Parse(value string) (time.Time, error) {
	if value, err := strconv.ParseInt(value, 10, 64); err != nil {
		return time.Time{}, err
	} else {
		return time.Unix(value, 0), nil
	}
}

type msFmt struct{}

func (u msFmt) String() string { return "unix-milliseconds" }

func (u msFmt) Match(value string) bool {
	return regexp.MustCompile(`^\d{12,14}$`).Match([]byte(value))
}

func (u msFmt) Parse(value string) (time.Time, error) {
	if value, err := strconv.ParseInt(value, 10, 64); err != nil {
		return time.Time{}, err
	} else {
		return time.Unix(value/1000, value%1000*1000000), nil
	}
}

type usFmt struct{}

func (u usFmt) String() string { return "unix-microseconds" }

func (u usFmt) Match(value string) bool {
	return regexp.MustCompile(`^\d{15,16}$`).Match([]byte(value))
}

func (u usFmt) Parse(value string) (time.Time, error) {
	if value, err := strconv.ParseInt(value, 10, 64); err != nil {
		return time.Time{}, err
	} else {
		return time.Unix(value/1000000, value%1000000*1000), nil
	}
}

type nsFmt struct{}

func (u nsFmt) String() string { return "unix-nanoseconds" }

func (u nsFmt) Match(value string) bool {
	return regexp.MustCompile(`^\d{16,}$`).Match([]byte(value))
}

func (u nsFmt) Parse(value string) (time.Time, error) {
	if value, err := strconv.ParseInt(value, 10, 64); err != nil {
		return time.Time{}, err
	} else {
		return time.Unix(value/1000000000, value%1000000000), nil
	}
}

type fsFmt struct{}

func (u fsFmt) String() string { return "unix-float-seconds" }

func (u fsFmt) Match(value string) bool {
	return regexp.MustCompile(`^\d{1,11}\.\d{1,9}$`).Match([]byte(value))
}

func (u fsFmt) Parse(value string) (time.Time, error) {
	parts := strings.Split(value, ".")
	if len(parts) != 2 {
		return time.Time{}, errors.New("invalid float seconds format")
	}

	for len(parts[1]) < 9 {
		parts[1] += "0"
	}

	sec, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	nsecs, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(sec, nsecs), nil
}
