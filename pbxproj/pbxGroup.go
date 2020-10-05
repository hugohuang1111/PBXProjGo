package pbxproj

import "path"

func newPBXGroup(filePath string) pbxMap {
	m := make(pbxMap, 0)
	s := pbxMapSection{}
	s.value = make(map[string]interface{})
	s.value["isa"] = "PBXGroup"
	s.value["name"] = path.Base(filePath)
	if fileIsExist(filePath) {
		s.value["path"] = filePath
	}
	s.value["sourceTree"] = DEFAULT_SOURCETREE

	m = append(m, s)

	return m
}
