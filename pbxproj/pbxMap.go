package pbxproj

import "strings"

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

/*
 * PBXGroup related
 */
func pbxGroupSourceTreeMatch(m pbxMap, st string) bool {
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
