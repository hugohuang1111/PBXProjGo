package pbxproj

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
