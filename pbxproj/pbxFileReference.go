package pbxproj

import (
	"path"
	"path/filepath"
	"strconv"
)

//group
//absolute
//SOURCE_ROOT

const (
	DEFAULT_SOURCETREE         = "<group>"
	DEFAULT_PRODUCT_SOURCETREE = "BUILT_PRODUCTS_DIR"
	DEFAULT_FILEENCODING       = 4
	DEFAULT_GROUP              = "Resources"
	DEFAULT_FILETYPE           = "unknown"
)

var FILETYPE_BY_EXTENSION map[string]string
var GROUP_BY_FILETYPE map[string]string
var PATH_BY_FILETYPE map[string]string
var SOURCETREE_BY_FILETYPE map[string]string
var ENCODING_BY_FILETYPE map[string]int
var buildPhaseByFileType map[string]string

func init() {
	FILETYPE_BY_EXTENSION = map[string]string{
		"a":           "archive.ar",
		"app":         "wrapper.application",
		"appex":       "wrapper.app-extension",
		"bundle":      "wrapper.plug-in",
		"cpp":         "sourcecode.cpp.cpp",
		"dylib":       "compiled.mach-o.dylib",
		"framework":   "wrapper.framework",
		"h":           "sourcecode.c.h",
		"m":           "sourcecode.c.objc",
		"markdown":    "text",
		"mdimporter":  "wrapper.cfbundle",
		"mm":          "sourcecode.c.objc",
		"octest":      "wrapper.cfbundle",
		"pch":         "sourcecode.c.h",
		"plist":       "text.plist.xml",
		"sh":          "text.script.sh",
		"swift":       "sourcecode.swift",
		"tbd":         "sourcecode.text-based-dylib-definition",
		"xcassets":    "folder.assetcatalog",
		"xcconfig":    "text.xcconfig",
		"xcdatamodel": "wrapper.xcdatamodel",
		"xcframework": "wrapper.xcframework",
		"xcodeproj":   "wrapper.pb-project",
		"xctest":      "wrapper.cfbundle",
		"xib":         "file.xib",
		"strings":     "text.plist.strings",
	}

	GROUP_BY_FILETYPE = map[string]string{
		"archive.ar":                             "Frameworks",
		"compiled.mach-o.dylib":                  "Frameworks",
		"sourcecode.text-based-dylib-definition": "Frameworks",
		"wrapper.framework":                      "Frameworks",
		"wrapper.xcframework":                    "Frameworks",
		"embedded.framework":                     "Embed Frameworks",
		"sourcecode.c.h":                         "Resources",
		"sourcecode.c.objc":                      "Sources",
		"sourcecode.cpp.cpp":                     "Sources",
		"sourcecode.swift":                       "Sources",
	}

	PATH_BY_FILETYPE = map[string]string{
		"compiled.mach-o.dylib":                  "usr/lib/",
		"sourcecode.text-based-dylib-definition": "usr/lib/",
		"wrapper.framework":                      "System/Library/Frameworks/",
		"wrapper.xcframework":                    "System/Library/Frameworks/",
	}

	SOURCETREE_BY_FILETYPE = map[string]string{
		"compiled.mach-o.dylib":                  "SDKROOT",
		"sourcecode.text-based-dylib-definition": "SDKROOT",
		"wrapper.framework":                      "SDKROOT",
		"wrapper.xcframework":                    "SDKROOT",
	}

	ENCODING_BY_FILETYPE = map[string]int{
		"sourcecode.c.h":     4,
		"sourcecode.c.objc":  4,
		"sourcecode.cpp.cpp": 4,
		"sourcecode.swift":   4,
		"text":               4,
		"text.plist.xml":     4,
		"text.script.sh":     4,
		"text.xcconfig":      4,
		"text.plist.strings": 4,
	}

	buildPhaseByFileType = map[string]string{
		"wrapper.framework":   "Frameworks",
		"wrapper.xcframework": "Frameworks",
		"embedded.framework":  "Frameworks",
		"sourcecode.c.objc":   "Sources",
		"sourcecode.cpp.cpp":  "Sources",
		"sourcecode.swift":    "Sources",
		"text.plist.strings":  "Resources",
	}
}

func detectType(filePath string) string {
	ext := path.Ext(filePath)
	if len(ext) > 1 {
		ext = ext[1:]
	}
	if val, ok := FILETYPE_BY_EXTENSION[ext]; ok {
		return val
	}

	return DEFAULT_FILETYPE
}

func detectGroup(fr pbxMap) string {
	filetype := fr["lastKnownFileType"].(string)
	if group, ok := GROUP_BY_FILETYPE[filetype]; ok {
		return group
	}

	return DEFAULT_GROUP
}

func detectBuildPhase(fr pbxMap) string {
	filetype := fr["lastKnownFileType"].(string)
	if bp, ok := buildPhaseByFileType[filetype]; ok {
		return bp
	}

	return ""
}

func (pbx PBXProject) newPBXFileRef(frPath string) pbxMap {
	m := make(pbxMap, 0)
	m["isa"] = "PBXFileReference"
	filetype := detectType(frPath)
	m["lastKnownFileType"] = filetype
	m["name"] = path.Base(frPath)
	if v, ok := ENCODING_BY_FILETYPE[filetype]; ok {
		m["fileEncoding"] = strconv.Itoa(v)
	}

	if filepath.IsAbs(frPath) {
		m.setPath(frPath, "<absolute>")
	} else {
		frFullPath := filepath.Join(pbx.projectDir, frPath)
		if isFileExist(frFullPath) {
			m.setPath(frPath, "<SOURCE_ROOT>")
		} else {
			frPath1, ok := PATH_BY_FILETYPE[filetype]
			if !ok {
				frPath1 = frPath
			}
			frst, ok := SOURCETREE_BY_FILETYPE[filetype]
			if !ok {
				frst = DEFAULT_SOURCETREE
			}
			m.setPath(frPath1, frst)
		}
	}

	return m
}
