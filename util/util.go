package util

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

func StringSplitByParagraph(str string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(strings.TrimSpace(str), "\n")
	return regexp.MustCompile(`(\n+)|(\n+\s+\n+)`).Split(strNormalized, -1)
}

func StringSplitBySpace(str string) []string {
	return regexp.MustCompile(`(\s+)|(\n+\s+\n+)`).Split(strings.TrimSpace(str), -1)
}

func StringDeleteVoidAndFormat(str string) string {
	return regexp.MustCompile(`(\s+)|(\n+)|(\r+)|(\f+)|(\t)|(\v)`).ReplaceAllString(str, "")
}

func JsonWithFalseEscapeHTML(str []string) string {
	Buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(Buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(str)
	if err != nil {
		return "[]"
	}
	return strings.TrimSpace(Buffer.String())
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func IsNum(n string) bool {
	_, err := strconv.Atoi(n)
	if err != nil {
		return false
	}
	return true
}