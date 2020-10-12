package pbxproj

import (
	"sort"
	"strings"
)

type pbxMapKeyArr []string

func (s pbxMapKeyArr) Len() int { return len(s) }

func (s pbxMapKeyArr) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s pbxMapKeyArr) Less(i, j int) bool {
	a := s[i]
	if strings.HasPrefix(a, "\"") {
		a = a[1:]
	}
	b := s[j]
	if strings.HasPrefix(b, "\"") {
		b = b[1:]
	}
	if strings.Compare(a, b) < 0 {
		return true
	}
	return false
}

func getPbxMapSortedKeys(m pbxMap) pbxMapKeyArr {
	keys := make(pbxMapKeyArr, 0, len(m))
	isaExist := false
	for key := range m {
		if strings.EqualFold(key, "isa") {
			isaExist = true
			continue
		}
		keys = append(keys, key)
	}

	sort.Sort(keys)
	if isaExist {
		keys = append([]string{"isa"}, keys...)
	}

	return keys
}
