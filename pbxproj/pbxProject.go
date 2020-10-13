package pbxproj

import (
	"fmt"
	"path"
	"strings"
)

func (pbx PBXProject) addFile(target string, group string, file string) {
	frMap := newPBXFileRef(file)
	dGroup := detectGroup(frMap)
	if 0 == len(group) {
		group = dGroup
	}
	frid := pbx.addFileReference(frMap)
	pbx.addToGroup(group, frid)

	var bpMap pbxMap
	switch dGroup {
	case "Sources":
		{
			_, bpMap = pbx.getSourcesBuildPhase(target)
		}
	case "Frameworks", "Embed Frameworks":
		{
			_, bpMap = pbx.getFrameworksBuildPhase(target)
		}
	case "Resources":
		{
			_, bpMap = pbx.getResourcesBuildPhase(target)
		}
	}

	if nil != bpMap {
		var val []interface{}
		if v, ok := bpMap["files"]; ok {
			val = v.([]interface{})
		} else {
			val = make([]interface{}, 0)
		}
		val = append(val, pbx.addBuildFile(frid))
		bpMap["files"] = val
	}
}

func (pbx PBXProject) getProject() pbxMap {
	objMap := pbx.project["objects"].(pbxMap)
	for _, value := range objMap {
		valMap := value.(pbxMap)
		if 0 == strings.Compare(valMap["isa"].(string), "PBXProject") {
			return valMap
		}
	}

	return nil
}

func (pbx PBXProject) getNativeTarget(target string) (string, pbxMap) {
	objMap := pbx.project["objects"].(pbxMap)
	if 0 == len(target) {
		return "", nil
	}
	if isUUID(target) {
		if v, ok := objMap[target]; ok {
			return target, v.(pbxMap)
		}
		return "", nil
	}

	for key, value := range objMap {
		valMap := value.(pbxMap)
		if 0 == strings.Compare(valMap["isa"].(string), "PBXNativeTarget") {
			if 0 == strings.Compare(valMap["name"].(string), target) {
				return key, valMap
			}
		}
	}

	return "", nil
}

func (pbx PBXProject) objectIsa(uuid string, isa string) bool {
	objMap := pbx.project["objects"].(pbxMap)
	if val, ok := objMap[uuid]; ok {
		valMap := val.(pbxMap)
		str := valMap["isa"].(string)
		if 0 == strings.Compare(str, isa) {
			return true
		}
	}

	return false
}

func (pbx PBXProject) getSourcesBuildPhase(target string) (string, pbxMap) {
	_, targetMap := pbx.getNativeTarget(target)
	if nil == targetMap {
		return "", nil
	}
	objMap := pbx.project["objects"].(pbxMap)
	for _, val := range targetMap["buildPhases"].([]interface{}) {
		bpid := val.(string)
		if pbx.objectIsa(bpid, "PBXSourcesBuildPhase") {
			return bpid, objMap[bpid].(pbxMap)
		}
	}

	sbp := make(pbxMap)
	sbp["isa"] = "PBXSourcesBuildPhase"
	sbp["buildActionMask"] = "2147483647"
	sbp["files"] = make([]interface{}, 0)
	sbp["runOnlyForDeploymentPostprocessing"] = "0"
	key := genUUID()
	objMap[key] = sbp

	return key, sbp
}

func (pbx PBXProject) getResourcesBuildPhase(target string) (string, pbxMap) {
	_, targetMap := pbx.getNativeTarget(target)
	if nil == targetMap {
		return "", nil
	}
	objMap := pbx.project["objects"].(pbxMap)
	for _, val := range targetMap["buildPhases"].([]interface{}) {
		bpid := val.(string)
		if pbx.objectIsa(bpid, "PBXResourcesBuildPhase") {
			return bpid, objMap[bpid].(pbxMap)
		}
	}

	return "", nil
}

func (pbx PBXProject) getFrameworksBuildPhase(target string) (string, pbxMap) {
	_, targetMap := pbx.getNativeTarget(target)
	if nil == targetMap {
		return "", nil
	}
	objMap := pbx.project["objects"].(pbxMap)
	for _, val := range targetMap["buildPhases"].([]interface{}) {
		bpid := val.(string)
		if pbx.objectIsa(bpid, "PBXFrameworksBuildPhase") {
			return bpid, objMap[bpid].(pbxMap)
		}
	}

	return "", nil
}

func (pbx PBXProject) addFileReference(frMap pbxMap) string {
	key := genUUID()
	objMap := pbx.project["objects"].(pbxMap)
	objMap[key] = frMap

	pbx.uuidMap[key] = fmt.Sprintf("/* %v */", frMap.getValue("name"))

	return key
}

func (pbx PBXProject) addToGroup(group string, frid string) string {
	gid, grpMap := pbx.getOrCreateGroup(group)
	valChildren := grpMap["children"]
	var children []interface{}
	if nil == valChildren {
		children = make([]interface{}, 0)
	} else {
		children = valChildren.([]interface{})
	}
	children = append(children, frid)
	grpMap["children"] = children

	return gid
}

func (pbx PBXProject) getOrCreateGroup(group string) (string, pbxMap) {
	objMap := pbx.project["objects"].(pbxMap)

	// try group as key
	if val, ok := objMap[group]; ok {
		return group, val.(pbxMap)
	}

	// try group as name
	for key, value := range objMap {
		valMap := value.(pbxMap)
		if 0 == strings.Compare(valMap["isa"].(string), "PBXGroup") &&
			0 == strings.Compare(valMap["name"].(string), group) {
			return key, valMap
		}
	}

	// create group
	key := genUUID()
	m := make(pbxMap)
	m = make(map[string]interface{})
	m["isa"] = "PBXGroup"
	m["name"] = path.Base(group)
	if fileIsExist(group) {
		m["path"] = group
	}
	m["sourceTree"] = DEFAULT_SOURCETREE

	objMap[key] = m

	pbx.uuidMap[key] = fmt.Sprintf("/* %v */", m["name"])

	return key, m
}

func (pbx PBXProject) addBuildFile(frid string) string {
	objMap := pbx.project["objects"].(pbxMap)

	key := genUUID()

	m := make(pbxMap)
	m = make(map[string]interface{})
	m["isa"] = "PBXBuildFile"
	m["fileRef"] = frid

	objMap[key] = m

	return key
}

func (pbx PBXProject) getBuildSetting(target string, mode string) pbxMap {
	objMap := pbx.project["objects"].(pbxMap)

	_, targetMap := pbx.getNativeTarget(target)
	bcl := targetMap["buildConfigurationList"].(string)
	bclMap := objMap[bcl].(pbxMap)
	for _, v := range bclMap["buildConfigurations"].([]interface{}) {
		bcid := v.(string)
		if valBC, ok := objMap[bcid]; ok {
			bcMap := valBC.(pbxMap)
			if 0 == strings.Compare(bcMap["name"].(string), mode) {
				return bcMap["buildSettings"].(pbxMap)
			}
		}
	}

	return nil
}
