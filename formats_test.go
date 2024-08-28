package main

import (
	"testing"
	"time"
)

func TestSFmtString(t *testing.T) {
	t.Parallel()

	s := sFmt{}
	if s.String() != "unix-seconds" {
		t.Errorf("Expected 'unix-seconds', but got '%s'", s.String())
	}
}

func TestSFmtMatch(t *testing.T) {
	t.Parallel()

	s := sFmt{}
	testCases := []struct {
		value string
		match bool
	}{
		{"1234567890", true},
		{"12345678901", true},
		{"123456789012", false},
		{"abcdef", false},
	}

	for _, tc := range testCases {
		if s.Match(tc.value) != tc.match {
			t.Errorf("Expected Match(%s) to be %v, but got %v", tc.value, tc.match, s.Match(tc.value))
		}
	}
}

func TestSFmtParse(t *testing.T) {
	t.Parallel()

	s := sFmt{}
	testCases := []struct {
		value    string
		expected time.Time
		err      bool
	}{
		{"1234567890", time.Unix(1234567890, 0), false},
		{"abcdef", time.Time{}, true},
	}

	for _, tc := range testCases {
		result, err := s.Parse(tc.value)
		if (err != nil) != tc.err {
			t.Errorf("Expected error to be %v, but got %v", tc.err, err != nil)
		} else if !tc.err && result != tc.expected {
			t.Errorf("Expected Parse(%s) to be %v, but got %v", tc.value, tc.expected, result)
		}
	}
}

func TestMSFmtString(t *testing.T) {
	t.Parallel()

	ms := msFmt{}
	if ms.String() != "unix-milliseconds" {
		t.Errorf("Expected 'unix-milliseconds', but got '%s'", ms.String())
	}
}

func TestMSFmtMatch(t *testing.T) {
	t.Parallel()

	ms := msFmt{}
	testCases := []struct {
		value string
		match bool
	}{
		{"1234567890123", true},
		{"123456789012", true},
		{"1234567890", false},
		{"abcdef", false},
	}

	for _, tc := range testCases {
		if ms.Match(tc.value) != tc.match {
			t.Errorf("Expected Match(%s) to be %v, but got %v", tc.value, tc.match, ms.Match(tc.value))
		}
	}
}

func TestMSFmtParse(t *testing.T) {
	t.Parallel()

	ms := msFmt{}
	testCases := []struct {
		value    string
		expected time.Time
		err      bool
	}{
		{"1234567890123", time.Unix(1234567890, 123000000), false},
		{"abcdef", time.Time{}, true},
	}

	for _, tc := range testCases {
		result, err := ms.Parse(tc.value)
		if (err != nil) != tc.err {
			t.Errorf("Expected error to be %v, but got %v", tc.err, err != nil)
		} else if !tc.err && result != tc.expected {
			t.Errorf("Expected Parse(%s) to be %v, but got %v", tc.value, tc.expected, result)
		}
	}
}

func TestUSFmtString(t *testing.T) {
	t.Parallel()

	us := usFmt{}
	if us.String() != "unix-microseconds" {
		t.Errorf("Expected 'unix-microseconds', but got '%s'", us.String())
	}
}

func TestUSFmtMatch(t *testing.T) {
	t.Parallel()

	us := usFmt{}
	testCases := []struct {
		value string
		match bool
	}{
		{"1234567890123456", true},
		{"123456789012345", true},
		{"1234567890123", false},
		{"abcdef", false},
	}

	for _, tc := range testCases {
		if us.Match(tc.value) != tc.match {
			t.Errorf("Expected Match(%s) to be %v, but got %v", tc.value, tc.match, us.Match(tc.value))
		}
	}
}

func TestUSFmtParse(t *testing.T) {
	t.Parallel()

	us := usFmt{}
	testCases := []struct {
		value    string
		expected time.Time
		err      bool
	}{
		{"1234567890123456", time.Unix(1234567890, 123456000), false},
		{"abcdef", time.Time{}, true},
	}

	for _, tc := range testCases {
		result, err := us.Parse(tc.value)
		if (err != nil) != tc.err {
			t.Errorf("Expected error to be %v, but got %v", tc.err, err != nil)
		} else if !tc.err && result != tc.expected {
			t.Errorf("Expected Parse(%s) to be %v, but got %v", tc.value, tc.expected, result)
		}
	}
}

func TestNSFmtString(t *testing.T) {
	t.Parallel()

	ns := nsFmt{}
	if ns.String() != "unix-nanoseconds" {
		t.Errorf("Expected 'unix-nanoseconds', but got '%s'", ns.String())
	}
}

func TestNSFmtMatch(t *testing.T) {
	t.Parallel()

	ns := nsFmt{}
	testCases := []struct {
		value string
		match bool
	}{
		{"1234567890123456789", true},
		{"1234567890123456", true},
		{"123456789012345", false},
		{"abcdef", false},
	}

	for _, tc := range testCases {
		if ns.Match(tc.value) != tc.match {
			t.Errorf("Expected Match(%s) to be %v, but got %v", tc.value, tc.match, ns.Match(tc.value))
		}
	}
}

func TestNSFmtParse(t *testing.T) {
	t.Parallel()

	ns := nsFmt{}
	testCases := []struct {
		value    string
		expected time.Time
		err      bool
	}{
		{"1234567890123456789", time.Unix(1234567890, 123456789), false},
		{"abcdef", time.Time{}, true},
	}

	for _, tc := range testCases {
		result, err := ns.Parse(tc.value)
		if (err != nil) != tc.err {
			t.Errorf("Expected error to be %v, but got %v", tc.err, err != nil)
		} else if !tc.err && result != tc.expected {
			t.Errorf("Expected Parse(%s) to be %v, but got %v", tc.value, tc.expected, result)
		}
	}
}

func TestFSFmtString(t *testing.T) {
	t.Parallel()

	fs := fsFmt{}
	if fs.String() != "unix-float-seconds" {
		t.Errorf("Expected 'unix-float-seconds', but got '%s'", fs.String())
	}
}

func TestFSFmtMatch(t *testing.T) {
	t.Parallel()

	fs := fsFmt{}
	testCases := []struct {
		value string
		match bool
	}{
		{"1234567890.123", true},
		{"1234567890.1", true},
		{"1234567890", false},
		{"abcdef", false},
	}

	for _, tc := range testCases {
		if fs.Match(tc.value) != tc.match {
			t.Errorf("Expected Match(%s) to be %v, but got %v", tc.value, tc.match, fs.Match(tc.value))
		}
	}
}

func TestFSFmtParse(t *testing.T) {
	t.Parallel()

	fs := fsFmt{}
	testCases := []struct {
		value    string
		expected time.Time
		err      bool
	}{
		{"1234567890.123", time.Unix(1234567890, 123000000), false},
		{"1564670787.9459848", time.Unix(1564670787, 945984800), false},
		{"abcdef", time.Time{}, true},
	}

	for _, tc := range testCases {
		result, err := fs.Parse(tc.value)
		if (err != nil) != tc.err {
			t.Errorf("Expected error to be %v, but got %v", tc.err, err != nil)
		} else if !tc.err && result != tc.expected {
			t.Errorf("Expected Parse(%s) to be %v, but got %v", tc.value, tc.expected, result)
		}
	}
}
