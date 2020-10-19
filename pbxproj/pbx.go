package pbxproj

import (
	"PBXProjGo/pbxproj/parser"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// "PBXProjGo/parser"
// "github.com/antlr/antlr4/runtime/Go/antlr"

// PBXProject pbxproj struct
type PBXProject struct {
	fileEncode string
	project    pbxMap

	projectDir string
	uuidMap    map[string]string

	pbxParser      *parser.PBXProjParser
	pbxTokenStream *antlr.CommonTokenStream
}

// NewPBXProject create pbx project
func NewPBXProject(pbxPBXProjectPath string) (PBXProject, error) {
	var pbx PBXProject
	pbx.projectDir, _ = filepath.Abs(filepath.Join(pbxPBXProjectPath, "..", ".."))
	input, _ := antlr.NewFileStream(pbxPBXProjectPath)
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
	f, err := newPbxWriteFile(filePath, pbx.uuidMap)
	if err != nil {
		return err
	}
	defer f.Close()

	f.writeFileEncode(pbx.fileEncode)
	f.writeMap(pbx.project)

	f.Sync()

	return nil
}

// GetTargets return tagrets name
func (pbx PBXProject) GetTargets() []string {
	targets := make([]string, 0)
	objMap := pbx.project["objects"].(pbxMap)
	var project pbxMap
	for _, value := range objMap {
		valMap := value.(pbxMap)
		if 0 == strings.Compare(valMap["isa"].(string), "PBXProject") {
			project = valMap
		}
	}

	if nil == project {
		return targets
	}

	for _, vTID := range project["targets"].([]interface{}) {
		tid := vTID.(string)
		valMap := objMap[tid].(pbxMap)
		targets = append(targets, valMap["name"].(string))
	}

	return targets
}

func (pbx PBXProject) AddFile(target string, group string, file string) {
	pbx.addFile(target, group, file)
}

func (pbx PBXProject) AddFramework(target string, file string) {
	pbx.addFile(target, "", file)
}

func (pbx PBXProject) AddResource(target string, file string) {
	pbx.addFile(target, "", file)
}

func (pbx PBXProject) GetBuildSetting(target string, mode string, key string) interface{} {
	bsMap := pbx.getBuildSetting(target, mode)
	return bsMap[key]
}

func (pbx PBXProject) SetBuildSetting(target string, mode string, key string, val interface{}) {
	bsMap := pbx.getBuildSetting(target, mode)
	bsMap[key] = val
}
