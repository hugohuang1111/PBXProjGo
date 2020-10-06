package pbxproj

import (
	"os"
	"strings"

	"github.com/google/uuid"
)

func isUUID(s string) bool {
	if 24 != len(s) {
		return false
	}
	match := false
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			match = true
		} else {
			match = false
		}

		if !match {
			return false
		}
	}

	return true
}

func genUUID() string {
	u1 := uuid.New()
	s := u1.String()

	s = strings.Replace(s, "-", "", -1)

	return strings.ToUpper(s[0:24])
}

func fileIsExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
