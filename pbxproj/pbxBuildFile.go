package pbxproj

func newPBXBuildFile(fruuid string) pbxMap {
	m := make(pbxMap, 0, 1)
	s := pbxMapSection{}
	s.value = make(map[string]interface{})
	s.value["isa"] = "PBXBuildFile"
	s.value["fileRef"] = fruuid

	m = append(m, s)

	return m
}
