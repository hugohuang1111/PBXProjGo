package pbxproj

import (
	"PBXProjGo/pbxproj/parser"
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

type pbxMap []pbxMapSection

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

func (pbx PBXProject) GetGroup(name string) string {
	objMap := findObjs(pbx.project)
	group := findSection(objMap, "PBXGroup")
	for key, value := range group {
		if 0 == strings.Compare(value["name"], name) {
			return key
		}
	}
	return ""
}

func (pbx PBXProject) AddFile() {

}
