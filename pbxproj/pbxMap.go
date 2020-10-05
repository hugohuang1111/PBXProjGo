package pbxproj

import "strings"

type pbxMap []pbxMapSection

type pbxMapSearchCallback func(key string, val interface{}) bool

func (m pbxMap) getValue(key string) interface{} {
	for _, section := range m {
		v, ok := section.value[key]
		if ok {
			return v
		}
	}

	return nil
}

func (m pbxMap) search(cb pbxMapSearchCallback) (string, interface{}) {
	for _, section := range m {
		for k, v := range section.value {
			if cb(k, v) {
				return k, v
			}
		}
	}

	return "", nil
}

func (m pbxMap) searchWithSection(sectionName string, cb pbxMapSearchCallback) (string, interface{}) {
	for _, section := range m {
		if 0 != strings.Compare(section.name, sectionName) {
			continue
		}
		for k, v := range section.value {
			if cb(k, v) {
				return k, v
			}
		}
	}

	return "", nil
}

func (m pbxMap) addValue(sectionName string, key string, value interface{}) {
	for _, section := range m {
		if 0 == strings.Compare(section.name, sectionName) {
			if vi, ok := section.value[key]; ok {
				arr := vi.([]interface{})
				section.value[key] = append(arr, value)
			} else {
				section.value[key] = [...]interface{}{value}
			}
			break
		}
	}
}

func (m pbxMap) setValue(sectionName string, key string, value interface{}) {
	for _, section := range m {
		if 0 == strings.Compare(section.name, sectionName) {
			section.value[key] = value
			break
		}
	}
}
