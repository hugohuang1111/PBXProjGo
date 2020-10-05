package pbxproj

import (
	"path"
)

const (
	DEFAULT_SOURCETREE         = "\"<group>\""
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

type pbxFileReference struct {
	isa               string
	lastKnownFileType string
	name              string
	path              string
	sourceTree        string
}

func init() {
	FILETYPE_BY_EXTENSION = map[string]string{
		"a":           "archive.ar",
		"app":         "wrapper.application",
		"appex":       "wrapper.app-extension",
		"bundle":      "wrapper.plug-in",
		"dylib":       "compiled.mach-o.dylib",
		"framework":   "wrapper.framework",
		"h":           "sourcecode.c.h",
		"m":           "sourcecode.c.objc",
		"markdown":    "text",
		"mdimporter":  "wrapper.cfbundle",
		"octest":      "wrapper.cfbundle",
		"pch":         "sourcecode.c.h",
		"plist":       "text.plist.xml",
		"sh":          "text.script.sh",
		"swift":       "sourcecode.swift",
		"tbd":         "sourcecode.text-based-dylib-definition",
		"xcassets":    "folder.assetcatalog",
		"xcconfig":    "text.xcconfig",
		"xcdatamodel": "wrapper.xcdatamodel",
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
		"embedded.framework":                     "Embed Frameworks",
		"sourcecode.c.h":                         "Resources",
		"sourcecode.c.objc":                      "Sources",
		"sourcecode.swift":                       "Sources",
	}

	PATH_BY_FILETYPE = map[string]string{
		"compiled.mach-o.dylib":                  "usr/lib/",
		"sourcecode.text-based-dylib-definition": "usr/lib/",
		"wrapper.framework":                      "System/Library/Frameworks/",
	}

	SOURCETREE_BY_FILETYPE = map[string]string{
		"compiled.mach-o.dylib":                  "SDKROOT",
		"sourcecode.text-based-dylib-definition": "SDKROOT",
		"wrapper.framework":                      "SDKROOT",
	}

	ENCODING_BY_FILETYPE = map[string]int{
		"sourcecode.c.h":     4,
		"sourcecode.c.objc":  4,
		"sourcecode.swift":   4,
		"text":               4,
		"text.plist.xml":     4,
		"text.script.sh":     4,
		"text.xcconfig":      4,
		"text.plist.strings": 4,
	}
}

// func unquoted(text string) {
// 	return text == null ? '' : text.replace (/(^")|("$)/g, '')
// }

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

// func defaultExtension(fileRef) {
// var filetype = fileRef.lastKnownFileType && fileRef.lastKnownFileType != DEFAULT_FILETYPE ?
// 	fileRef.lastKnownFileType : fileRef.explicitFileType;

// for(var extension in FILETYPE_BY_EXTENSION) {
// 	if(FILETYPE_BY_EXTENSION.hasOwnProperty(unquoted(extension)) ) {
// 		 if(FILETYPE_BY_EXTENSION[unquoted(extension)] === unquoted(filetype) )
// 			 return extension;
// 	}
// }
// }

// function defaultEncoding(fileRef) {
// var filetype = fileRef.lastKnownFileType || fileRef.explicitFileType,
// 	encoding = ENCODING_BY_FILETYPE[unquoted(filetype)];

// if (encoding) {
// 	return encoding;
// }
// }

// function detectGroup(fileRef, opt) {
// var extension = path.extname(fileRef.basename).substring(1),
// 	filetype = fileRef.lastKnownFileType || fileRef.explicitFileType,
// 	groupName = GROUP_BY_FILETYPE[unquoted(filetype)];

// if (extension === 'xcdatamodeld') {
// 	return 'Sources';
// }

// if (opt.customFramework && opt.embed) {
// 	return GROUP_BY_FILETYPE['embedded.framework'];
// }

// if (!groupName) {
// 	return DEFAULT_GROUP;
// }

// return groupName;
// }

func detectSourcetree(fileRef pbxFileReference) string {
	filetype := fileRef.lastKnownFileType
	sourcetree, exist := SOURCETREE_BY_FILETYPE[filetype]

	if !exist {
		return DEFAULT_SOURCETREE
	}

	return sourcetree
}

// function defaultPath(fileRef, filePath) {
// var filetype = fileRef.lastKnownFileType || fileRef.explicitFileType,
// 	defaultPath = PATH_BY_FILETYPE[unquoted(filetype)];

// if (fileRef.customFramework) {
// 	return filePath;
// }

// if (defaultPath) {
// 	return path.join(defaultPath, path.basename(filePath));
// }

// return filePath;
// }

// function defaultGroup(fileRef) {
// var groupName = GROUP_BY_FILETYPE[fileRef.lastKnownFileType];

// if (!groupName) {
// 	return DEFAULT_GROUP;
// }

// return defaultGroup;
// }

// func newPBXFileRef(filePath string) pbxFileReference {
// 	var fileRef pbxFileReference

// 	fileRef.isa = "PBXFileReference"
// 	fileRef.lastKnownFileType = detectType(filePath)
// 	fileRef.name = path.Base(filePath)
// 	fileRef.path = filePath
// 	fileRef.sourceTree = detectSourcetree(fileRef)

// 	return fileRef
// }

func newPBXFileRef(filePath string) pbxMap {
	m := make(pbxMap, 0)
	s := pbxMapSection{}
	s.value = make(map[string]interface{})
	s.value["isa"] = "PBXFileReference"
	filetype := detectType(filePath)
	s.value["lastKnownFileType"] = filetype
	s.value["name"] = path.Base(filePath)
	s.value["path"] = filePath
	if sourcetree, exist := SOURCETREE_BY_FILETYPE[filetype]; exist {
		s.value["sourceTree"] = sourcetree
	} else {
		s.value["sourceTree"] = DEFAULT_SOURCETREE
	}

	m = append(m, s)

	return m
}
