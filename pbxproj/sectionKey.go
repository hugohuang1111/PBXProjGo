package pbxproj

import "strings"

type sectionKeyArr []string

func (s sectionKeyArr) Len() int { return len(s) }

func (s sectionKeyArr) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s sectionKeyArr) Less(i, j int) bool {
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
