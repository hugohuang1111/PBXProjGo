package pbxproj

import (
	"PBXProjGo/pbxproj/parser"
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// "PBXProjGo/parser"
// "github.com/antlr/antlr4/runtime/Go/antlr"

type pbxMapSection struct {
	name  string
	value map[string]interface{}
}

// PBXProject pbx project tree
type PBXProject struct {
	fileEncode string
	project    pbxMap
	uuidMap    map[string]string

	pbxParser      *parser.PBXProjParser
	pbxTokenStream *antlr.CommonTokenStream
}

func newMapSection() pbxMapSection {
	v := pbxMapSection{}
	v.value = make(map[string]interface{})

	return v
}

// NewPBXProject create pbx project
func NewPBXProject(pbxProjectPath string) (PBXProject, error) {
	var pbx PBXProject

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewPBXProjLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	pbx.pbxTokenStream = stream
	pbx.pbxParser = parser.NewPBXProjParser(stream)
	pbx.pbxParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	pbx.pbxParser.BuildParseTrees = true
	tree := pbx.pbxParser.Project()
	pbx.uuidMap = make(map[string]string)
	antlr.ParseTreeWalkerDefault.Walk(newPBXTreeShapeListener(&pbx), tree)

	return pbx, nil
}

// Save save project
func (pbx PBXProject) Save(filePath string) error {
	f, err := newPbxFile(filePath, pbx.uuidMap)
	if err != nil {
		return err
	}
	defer f.Close()

	f.saveFileEncode(pbx.fileEncode)
	f.writeMap(pbx.project)

	f.Sync()

	return nil
}

// FindGroupUUID find group uuid by name
func (pbx PBXProject) FindGroupUUID(name string) string {
	objMap := pbx.getObject()
	if nil == objMap {
		return ""
	}
	k, _ := objMap.searchWithSection("PBXGroup", func(key string, val interface{}) bool {
		vMap := val.(pbxMap)
		k, _ := vMap.search(func(key string, val interface{}) bool {
			if 0 != strings.Compare(key, "name") {
				return false
			}
			sName := val.(string)
			if strings.Compare(sName, name) == 0 {
				return true
			}
			return false
		})
		if 0 != len(k) {
			return true
		}
		return false
	})
	return k
}

//AddFile add file to group
func (pbx PBXProject) AddFile(group string, file string) string {
	uuid := group
	if !isUUID(group) {
		uuid = pbx.FindGroupUUID(group)
	}

	fid, _ := pbx.addFileReference(file)
	val := pbx.getValueByUUID(uuid)
	vMap := val.(pbxMap)
	vMap.addValue("", "children", fid)

	return fid
}

// AddGroup add group
func (pbx PBXProject) AddGroup(group string) string {
	key, _ := pbx.addGroup(group)

	return key
}

//AddFramework add framework to group
func (pbx PBXProject) AddFramework(file string) error {
	uuid := pbx.FindGroupUUID("Frameworks")
	if 0 == len(uuid) {
		uuid = pbx.addGroup("Frameworks")
	}

	fid, _ := pbx.addFileReference(file)
	val := pbx.getValueByUUID(uuid)
	vMap := val.(pbxMap)
	vMap.addValue("", "children", fid)

	return nil
}

func (pbx PBXProject) addFileReference(filePath string) (string, pbxMap) {
	frSection := pbx.getSection("PBXFileReference")
	key := genUUID()
	val := newPBXFileRef(filePath)
	frSection.value[key] = val

	pbx.uuidMap[key] = fmt.Sprintf("/* %v */", val.getValue("name"))

	return key, val
}

func (pbx PBXProject) addGroup(group string) (string, pbxMap) {
	gpSection := pbx.getSection("PBXGroup")
	key := genUUID()
	val := newPBXGroup(group)
	gpSection.value[key] = val

	pbx.uuidMap[key] = fmt.Sprintf("/* %v */", val.getValue("name"))

	return key, val
}

func (pbx PBXProject) getObject() pbxMap {
	_, v := pbx.project.search(func(key string, val interface{}) bool {
		if 0 == strings.Compare(key, "objects") {
			return true
		}
		return false
	})

	return v.(pbxMap)
}

func (pbx PBXProject) getValueByUUID(uuid string) interface{} {
	objMap := pbx.getObject()
	for _, section := range objMap {
		if val, ok := section.value[uuid]; ok {
			return val
		}
	}

	return nil
}

func (pbx PBXProject) getSection(name string) pbxMapSection {
	objMap := pbx.getObject()
	for _, section := range objMap {
		if 0 == strings.Compare(section.name, name) {
			return section
		}
	}

	return pbxMapSection{}
}
