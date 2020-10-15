package pbxproj

import (
	"sort"
	"strings"
)

func getPbxMapSortedKeys(m pbxMap) []string {
	keys := make([]string, 0, len(m))
	isaExist := false
	for key := range m {
		if strings.EqualFold(key, "isa") {
			isaExist = true
			continue
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if isaExist {
		keys = append([]string{"isa"}, keys...)
	}

	return keys
}
