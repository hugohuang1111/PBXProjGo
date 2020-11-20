package pbxproj

import (
	"sort"
	"strings"
)

type pbxMap map[string]interface{}

func (m pbxMap) getValue(keys ...string) interface{} {
	var val interface{} = m
	var valMap pbxMap
	var ok bool
	for _, key := range keys {
		if valMap, ok = val.(pbxMap); !ok {
			return nil
		}
		if val, ok = valMap[key]; !ok {
			return nil
		}
	}

	return val
}

func (m pbxMap) getValueString(keys ...string) string {
	if nil == m {
		return ""
	}
	v := m.getValue(keys...)
	if nil == v {
		return ""
	}
	return v.(string)
}

func (m pbxMap) getName() string {
	val, ok := m["name"]
	sName := ""
	if ok {
		sName = val.(string)
	} else {
		val, ok = m["path"]
		if ok {
			sName = val.(string)
		}
	}

	return sName
}

func (m pbxMap) getValueMap(keys ...string) pbxMap {
	if nil == m {
		return nil
	}
	v := m.getValue(keys...)
	if nil == v {
		return nil
	}
	return v.(pbxMap)
}

func (m pbxMap) setPath(path string, sourcetree string) {
	if nil == m {
		return
	}
	m["sourceTree"] = sourcetree
	m["path"] = path
}

func (m pbxMap) appendChild(key string, val interface{}) {
	valItem, ok := m[key]
	var children []interface{} = nil

	if ok {
		children = makeSureIsStringArray(valItem)
	} else {
		children = make([]interface{}, 0, 2)
	}

	if exist, _ := inArray(val, children); !exist {
		children = append(children, val)
		m[key] = children
	}
}

/*
 * Map Section related
 */

func getSectionNames(m pbxMap) []string {
	sections := make([]string, 0)

	for _, val := range m {
		m := val.(pbxMap)
		if valISA, ok := m["isa"]; ok {
			valStr := valISA.(string)
			if !isContain(sections, valStr) {
				sections = append(sections, valStr)
			}
		}
	}

	sort.Strings(sections)

	return sections
}

func pbxSourceTreeMatch(m pbxMap, st string) bool {
	if v, ok := m["sourceTree"]; ok {
		if 0 == strings.Compare(st, v.(string)) {
			return true
		}
		return false
	}

	if 0 == strings.Compare(st, DEFAULT_SOURCETREE) {
		return true
	}
	return false
}

/*
 * PBXGroup related
 */

func pbxGroupPath(m pbxMap) string {
	if v, ok := m["path"]; ok {
		return v.(string)
	}

	return ""
}

func pbxGroupName(m pbxMap) string {
	if v, ok := m["name"]; ok {
		return v.(string)
	} else if v, ok := m["path"]; ok {
		return v.(string)
	}

	return ""
}

func pbxGroupIncludeChild(m pbxMap, gid string) bool {
	if nil == m {
		return false
	}
	val, ok := m["children"]
	if !ok {
		return false
	}
	valChildren := val.([]interface{})
	for _, valChild := range valChildren {
		if 0 == strings.Compare(valChild.(string), gid) {
			return true
		}
	}
	return false
}

func makeSureIsStringArray(val interface{}) []interface{} {
	sa := make([]interface{}, 0, 2)
	switch val.(type) {
	case string:
		s := val.(string)
		sa = append(sa, s)
	case []interface{}:
		sa = val.([]interface{})
	}

	return sa
}
