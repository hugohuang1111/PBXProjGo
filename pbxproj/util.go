package pbxproj

import (
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

	return s[0:24]
}

func findSectionFromMap(m pbxMap, sectionName string) pbxMapSection {
	for idx, section := range m {
		if 0 == len(sectionName) {
			if 0 == len(section.name) {
				return section
			}
		} else if 0 == strings.Compare(section.name, sectionName) {
			return section
		}
	}
	return pbxMapSection{}
}

func findValueFromSection(section pbxMapSection, name string) interface{} {
	if val, ok := section.value[name]; ok {
		return val
	}

	return nil
}

func findValueFromMap(m pbxMap, sectionName string, name string) interface{} {
	sec := findSectionFromMap(m, sectionName)
	return findValueFromSection(sec, name)
}

func findObjs(pbx pbxMap) pbxMap {
	objMap := pbx[0].value["objects"].(pbxMap)

	return objMap
}

func findSection(objMap pbxMap, name string) map[string]interface{} {
	for idx, section := range objMap {
		if 0 == strings.Compare(section.name, "PBXGroup") {
			return section.value
		}
	}
	return nil
}
