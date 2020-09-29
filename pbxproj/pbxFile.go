package pbxproj

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type pbxFile struct {
	*os.File
	uuidMap map[string]string
	indent  int
	unfold  bool
}

func newPbxFile(path string, uuidMap map[string]string) (pbxFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return pbxFile{}, err
	}

	return pbxFile{f, uuidMap, 0, true}, nil
}

func (f pbxFile) saveFileEncode(s string) {
	f.writeString("// !$*")
	f.writeString(s)
	f.writeString("*$!")
	f.writeNewline()
}

func (f pbxFile) writeKey(v interface{}) {
	if s, ok := v.(string); ok {
		f.writeString(s)
	} else {
		fmt.Println("WARNING: unexpect key type")
	}
}

func (f pbxFile) writeValue(v interface{}) {
	switch v.(type) {
	case pbxMap:
		{
			f.writeMap(v.(pbxMap))
		}
	case []interface{}:
		{
			f.saveArray(v.([]interface{}))
		}
	case string:
		{
			f.writeString(v.(string))
		}
	default:
		{
			fmt.Println("WARNING: unsupported type", v)
			break
		}
	}

}

func (f pbxFile) writeMap(sections []pbxMapSection) {
	f.writeString("{")
	f.writeNewline()

	f.indent++
	for _, section := range sections {
		f.writeSection(section)
	}
	f.indent--
	f.writeIndent()
	f.writeString("}")
}

func (f pbxFile) saveArray(arr []interface{}) {
	f.writeString("(")

	f.indent++
	for _, item := range arr {
		f.writeNewline()
		f.writeIndent()
		f.writeValue(item)
		f.writeString(",")
	}
	f.indent--

	f.writeNewline()
	f.writeIndent()
	f.writeString(")")
}

func (f pbxFile) writeSection(section pbxMapSection) {
	if 0 != len(section.name) {
		f.writeNewline()
		f.writeString(fmt.Sprintf("/* Begin %s section */", section.name))
		f.writeNewline()
	}

	keys := make(sectionKeyArr, 0, len(section.value))
	isaExist := false
	for key := range section.value {
		if strings.EqualFold(key, "isa") {
			isaExist = true
			continue
		}
		keys = append(keys, key)
	}

	sort.Sort(keys)
	if isaExist {
		keys = append([]string{"isa"}, keys...)
	}

	for _, key := range keys {
		f.writeIndent()
		f.writeKey(key)
		f.writeString(" = ")
		oneline := strings.EqualFold(section.name, "PBXBuildFile") || strings.EqualFold(section.name, "PBXFileReference")
		if oneline {
			f.unfold = false
		}
		f.writeValue(section.value[key])
		if oneline {
			f.unfold = true
		}
		f.writeString(";")
		if !f.unfold {
			f.writeString(" ")
		}
		f.writeNewline()
	}

	if 0 != len(section.name) {
		f.writeString(fmt.Sprintf("/* End %s section */", section.name))
		f.writeNewline()
	}
}

func (f pbxFile) writeString(s string) {
	f.WriteString(s)
	if isUUID(s) {
		if v, ok := f.uuidMap[s]; ok {
			f.writeString(" ")
			f.writeString(v)
		}
	}
}

func (f pbxFile) writeNewline() {
	if !f.unfold {
		return
	}
	f.WriteString("\n")
}

func (f pbxFile) writeIndent() {
	if !f.unfold {
		return
	}
	for i := 0; i < f.indent; i++ {
		f.WriteString("\t")
	}
}
