package pbxproj

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var uuidMap = map[string]string{}

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

func transUUID2String(uuid string) string {
	if cmt, ok := uuidMap[uuid]; ok {
		return fmt.Sprintf("%s %s", uuid, cmt)
	}

	return uuid
}
