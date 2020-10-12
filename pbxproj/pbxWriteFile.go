package pbxproj

import (
	"fmt"
	"os"
	"strings"
)

type pbxWriteFile struct {
	*os.File
	uuidMap          map[string]string
	level            int
	unfold           bool
	writeUUIDComment bool
}

func newPbxWriteFile(path string, uuidMap map[string]string) (pbxWriteFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return pbxWriteFile{}, err
	}

	return pbxWriteFile{f, uuidMap, 0, true, true}, nil
}

func (f pbxWriteFile) writeFileEncode(s string) {
	f.writeString("// !$*")
	f.writeString(s)
	f.writeString("*$!")
	f.writeNewline()
}

func (f pbxWriteFile) writeObjectsMap(m pbxMap) {
	f.writeString("{")
	f.writeNewline()

	sections := []string{
		"PBXBuildFile", "PBXContainerItemProxy", "PBXFileReference",
		"PBXFrameworksBuildPhase", "PBXGroup", "PBXNativeTarget",
		"PBXProject", "PBXReferenceProxy", "PBXResourcesBuildPhase",
		"PBXSourcesBuildPhase", "XCBuildConfiguration", "XCConfigurationList"}

	f.level++

	keys := getPbxMapSortedKeys(m)

	for _, section := range sections {
		f.writeNewline()
		f.writeString(fmt.Sprintf("/* Begin %s section */", section))
		f.writeNewline()
		for _, key := range keys {
			valMap := m[key].(pbxMap)

			if 0 != strings.Compare(valMap["isa"].(string), section) {
				continue
			}

			f.writeIndent()
			f.writeKey(key)
			f.writeString(" = ")
			oneline := strings.EqualFold(section, "PBXBuildFile") || strings.EqualFold(section, "PBXFileReference")
			if oneline {
				f.unfold = false
			}
			f.writeValue(valMap)
			if oneline {
				f.unfold = true
			}
			f.writeString(";")
			if !f.unfold {
				f.writeString(" ")
			}
			f.writeNewline()
		}
		f.writeString(fmt.Sprintf("/* End %s section */", section))
		f.writeNewline()
	}
	f.level--
	f.writeIndent()
	f.writeString("}")
}

func (f pbxWriteFile) writeKey(v interface{}) {
	if s, ok := v.(string); ok {
		b := isStringNeedQuote(s)
		if b {
			f.WriteString("\"")
		}
		f.writeString(s)
		if b {
			f.WriteString("\"")
		}
	} else {
		fmt.Println("WARNING: unexpect key type")
	}
}

func (f pbxWriteFile) writeValue(v interface{}) {
	switch v.(type) {
	case pbxMap:
		{
			f.writeMap(v.(pbxMap))
		}
	case []interface{}:
		{
			f.writeArray(v.([]interface{}))
		}
	case string:
		{
			s := v.(string)
			b := isStringNeedQuote(s)
			if b {
				f.WriteString("\"")
			}
			f.writeString(s)
			if b {
				f.WriteString("\"")
			}
		}
	default:
		{
			fmt.Println("WARNING: unsupported type", v)
			break
		}
	}

}

func (f pbxWriteFile) writeMap(m pbxMap) {
	f.writeString("{")
	f.writeNewline()

	f.level++
	keys := getPbxMapSortedKeys(m)

	for _, key := range keys {
		f.writeIndent()
		f.writeKey(key)
		f.writeString(" = ")
		if 3 == f.level && 0 == strings.Compare(key, "attributes") {
			f.writeUUIDComment = false
		}
		if 1 == f.level && 0 == strings.Compare(key, "objects") {
			f.writeObjectsMap(m[key].(pbxMap))
		} else {
			f.writeValue(m[key])
		}
		if 3 == f.level && 0 == strings.Compare(key, "attributes") {
			f.writeUUIDComment = true
		}
		f.writeString(";")
		if !f.unfold {
			f.writeString(" ")
		}
		f.writeNewline()
	}

	f.level--
	f.writeIndent()
	f.writeString("}")
}

func (f pbxWriteFile) writeArray(arr []interface{}) {
	f.writeString("(")

	f.level++
	for _, item := range arr {
		f.writeNewline()
		f.writeIndent()
		f.writeValue(item)
		f.writeString(",")
	}
	f.level--

	f.writeNewline()
	f.writeIndent()
	f.writeString(")")
}

func (f pbxWriteFile) writeString(s string) {
	f.WriteString(s)
	if f.writeUUIDComment && isUUID(s) {
		if v, ok := f.uuidMap[s]; ok {
			f.writeString(" ")
			f.writeString(v)
		}
	}
}

func (f pbxWriteFile) writeNewline() {
	if !f.unfold {
		return
	}
	f.WriteString("\n")
}

func (f pbxWriteFile) writeIndent() {
	if !f.unfold {
		return
	}
	for i := 0; i < f.level; i++ {
		f.WriteString("\t")
	}
}
