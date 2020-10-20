package pbxproj

import (
	"os"
	"path/filepath"
	"reflect"
)

func isFileExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func isStringNeedQuote(s string) bool {
	if 0 == len(s) {
		return true
	}
	for _, c := range s {
		if c == ' ' || c == '-' || c == '?' || c == '+' || c == '@' || c == '$' || c == ',' || c == '=' ||
			c == '[' || c == ']' || c == '{' || c == '}' || c == '<' || c == '>' || c == '(' || c == ')' {
			return true
		}
	}

	return false
}

func isContain(arr []string, item string) bool {
	for _, eachItem := range arr {
		if eachItem == item {
			return true
		}
	}

	return false
}

func getExecutePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if nil != err {
		return ""
	}
	return dir
}

func changePathToAbs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	dir := getExecutePath()
	return filepath.Join(dir, path)
}

func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
