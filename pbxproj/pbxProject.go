package pbxproj

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (pbx PBXProject) addFile(target string, group string, file string) {
	frid, frMap := pbx.createOrFindFileReferencesByPath(file)

	dGroup := detectGroup(frMap)

	if len(group) > 0 {
		pbx.addToGroup(group, frid)
	} else {
		pbx.addToGroup(dGroup, frid)
	}

	var bpMap pbxMap
	fileid, _ := pbx.addBuildFileIf(target, frid)
	if 0 == len(fileid) {
		fileid = frid
	}
	switch dGroup {
	case "Sources":
		_, bpMap = pbx.getSourcesBuildPhase(target)
	case "Frameworks", "Embed Frameworks":
		_, bpMap = pbx.getFrameworksBuildPhase(target)
	case "Resources":
		_, bpMap = pbx.getResourcesBuildPhase(target)
		fileid = frid
	}
	if nil != bpMap {
		bpMap.appendChild("files", fileid)
	}
}

func (pbx PBXProject) createOrFindFileReferencesByPath(path string) (string, pbxMap) {
	frid, frMap := pbx.findFileReferencesByPath(path)
	if nil == frMap {
		frMap = pbx.newPBXFileRef(path)
		frid = pbx.addFileReference(frMap)
	}

	return frid, frMap
}

func (pbx PBXProject) findFileReferencesByPath(path string) (string, pbxMap) {
	objMap := pbx.project["objects"].(pbxMap)
	baseName := filepath.Base(path)
	for uuid, value := range objMap {
		valMap := value.(pbxMap)
		if valMap.getValueString("isa") == "PBXFileReference" &&
			valMap.getValueString("name") == baseName {
			if pbxSourceTreeMatch(valMap, "<group>") {
				if pbx.getAbsPathByMap(uuid, valMap) == path {
					return uuid, valMap
				}
			} else {
				return uuid, valMap
			}
		}
	}

	return "", nil
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

func (pbx PBXProject) addToGroup(group string, frid string) (string, pbxMap) {
	gid, grpMap := pbx.getOrCreateGroup(group)
	grpMap.appendChild("children", frid)

	return gid, grpMap
}

func (pbx PBXProject) getOrCreateGroup(group string) (string, pbxMap) {
	objMap := pbx.project["objects"].(pbxMap)

	// try group as key
	if val, ok := objMap[group]; ok {
		return group, val.(pbxMap)
	}

	projMap := pbx.getProject()
	rootGroupID := projMap.getValueString("mainGroup")
	gs := strings.Split(group, "/")
	rootGroup := objMap[rootGroupID].(pbxMap)

	groupMap := rootGroup
	curDir := pbx.projectDir
	var curGroupID string
	for _, gname := range gs {
		find := false
		if path, ok := groupMap["path"]; ok {
			curDir = filepath.Join(curDir, path.(string))
		}
		for _, val := range groupMap["children"].([]interface{}) {
			gid := val.(string)
			childMap := objMap[gid].(pbxMap)
			if childMap.getValueString("name") == gname {
				find = true
				curGroupID = gid
				groupMap = childMap
				break
			}
		}
		if !find {
			curGroupID = genUUID()
			m := make(pbxMap)
			m = make(map[string]interface{})
			m["isa"] = "PBXGroup"
			m["name"] = path.Base(gname)
			m["sourceTree"] = DEFAULT_SOURCETREE
			if isFileExist(filepath.Join(curDir, gname)) {
				m["path"] = gname
			}

			objMap[curGroupID] = m
			groupMap.appendChild("children", curGroupID)

			pbx.uuidMap[curGroupID] = fmt.Sprintf("/* %v */", m["name"])

			groupMap = m
		}
	}

	return curGroupID, groupMap
}

// addBuildFileIf add build file if needed or if not exist
func (pbx PBXProject) addBuildFileIf(target string, frid string) (retBFID string, exists bool) {
	objMap := pbx.project["objects"].(pbxMap)

	bfids := make([]string, 0)
	for uuid, value := range objMap {
		valMap := value.(pbxMap)
		if valMap.getValueString("isa") == "PBXBuildFile" {
			if valMap.getValueString("fileRef") == frid {
				bfids = append(bfids, uuid)
			}
		}
	}

	frMap := objMap[frid].(pbxMap)
	var bpMap pbxMap
	switch detectBuildPhase(frMap) {
	case "Frameworks":
		_, bpMap = pbx.getFrameworksBuildPhase(target)
	case "Sources":
		_, bpMap = pbx.getSourcesBuildPhase(target)
	case "Resources":
		_, bpMap = pbx.getResourcesBuildPhase(target)
	}

	if nil == bpMap {
		return "", false
	}

	bpChildren := bpMap.getValue("files")
	for _, bfid := range bfids {
		if ok, _ := inArray(bfid, bpChildren); ok {
			retBFID = bfid
			exists = true
			break
		}
	}
	if exists {
		return
	}

	retBFID = genUUID()

	m := make(pbxMap)
	m = make(map[string]interface{})
	m["isa"] = "PBXBuildFile"
	m["fileRef"] = frid

	objMap[retBFID] = m

	return
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
	if nil == targetMap {
		return nil
	}
	bcl := targetMap.getValueString("buildConfigurationList")
	bclMap := objMap.getValueMap(bcl)
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

func (pbx PBXProject) getGroupAbsPath(gid string) string {
	objMap := pbx.project["objects"].(pbxMap)

	var curGid string = gid
	parts := make([]string, 0)
	groupMap := objMap[curGid].(pbxMap)
	parts = append(parts, pbxGroupPath(groupMap))
	if !pbxSourceTreeMatch(groupMap, "<absolute>") {
		for key, value := range objMap {
			valMap := value.(pbxMap)
			if 0 == strings.Compare(valMap["isa"].(string), "PBXGroup") {
				if pbxGroupIncludeChild(valMap, gid) {
					curGid = key
					groupMap = objMap[curGid].(pbxMap)
					parts = append(parts, pbxGroupPath(groupMap))
					if pbxSourceTreeMatch(groupMap, "<absolute>") {
						break
					}
				}
			}
		}
	}

	spath := ""
	for i := len(parts) - 1; i >= 0; i-- {
		if len(parts[i]) > 0 {
			spath = filepath.Join(spath, parts[i])
		}
	}

	return filepath.Join(pbx.projectDir, spath)
}

func (pbx PBXProject) getAbsPathByID(uid string) string {
	objMap := pbx.project["objects"].(pbxMap)

	valMap, ok := objMap[uid]
	if !ok {
		return ""
	}
	frMap := valMap.(pbxMap)

	return pbx.getAbsPathByMap(uid, frMap)
}

func (pbx PBXProject) getAbsPathByMap(uuid string, valMap pbxMap) string {
	objMap := pbx.project["objects"].(pbxMap)
	pathParts := make([]string, 0)
	pathParts = append(pathParts, valMap.getValueString("path"))

	curUUID := uuid
	var isGroupSourcetree bool
	run := true

	isGroupSourcetree = pbxSourceTreeMatch(valMap, "<group>")
	run = isGroupSourcetree
	for run {
		run = false
		for uuid, value := range objMap {
			valMap := value.(pbxMap)
			if valMap.getValueString("isa") == "PBXGroup" &&
				pbxGroupIncludeChild(valMap, curUUID) {
				curUUID = uuid
				pathParts = append(pathParts, valMap.getValueString("path"))
				isGroupSourcetree = pbxSourceTreeMatch(valMap, "<group>")
				run = isGroupSourcetree
				break
			}
		}
		if !run {
			break
		}
	}

	spath := ""
	for i := len(pathParts) - 1; i >= 0; i-- {
		if len(pathParts[i]) > 0 {
			spath = filepath.Join(spath, pathParts[i])
		}
	}

	if isGroupSourcetree {
		return filepath.Join(pbx.projectDir, spath)
	}
	return spath
}

func (pbx PBXProject) findSameFilereference(frMap pbxMap) (string, pbxMap) {
	objMap := pbx.project["objects"].(pbxMap)

	for uuid, value := range objMap {
		valMap := value.(pbxMap)
		if valMap.getValueString("isa") == "PBXFileReference" &&
			valMap.getValueString("name") == frMap.getValueString("name") &&
			pbx.getAbsPathByID(uuid) == frMap.getValueString("path") {
			return uuid, valMap
		}
	}

	return "", nil
}
