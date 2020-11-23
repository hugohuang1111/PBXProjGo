package pbxproj

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/hugohuang1111/PBXProjGo/pbxproj/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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

	if !isFileExist(pbxPBXProjectPath) {
		return pbx, errors.New("File not exist")
	}

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
	f.writeNewline()

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

// func (pbx PBXProject) AddFile(target string, group string, file string) {
// 	pbx.addFile(target, group, file)
// }

// AddSourceFile add source file
func (pbx PBXProject) AddSourceFile(target string, group string, file string) {
	pbx.addFile(target, group, file)
}

// AddHeaderFile add header file
func (pbx PBXProject) AddHeaderFile(target string, group string, file string) {
	pbx.addFile(target, group, file)
}

// AddFramework add framework
func (pbx PBXProject) AddFramework(target string, file string) {
	pbx.addFile(target, "", file)

	for _, mode := range []string{"Debug", "Release"} {
		bsMap := pbx.getBuildSetting(target, mode)
		bsMap.appendChild("FRAMEWORK_SEARCH_PATHS", "$(inherited)")
		bsMap.appendChild("FRAMEWORK_SEARCH_PATHS", "$(PROJECT_DIR)")
	}
}

// AddResource add resource
func (pbx PBXProject) AddResource(target string, file string) {
	pbx.addFile(target, "", file)
}

// GetBuildSetting get build setting
func (pbx PBXProject) GetBuildSetting(target string, mode string, key string) interface{} {
	bsMap := pbx.getBuildSetting(target, mode)
	return bsMap[key]
}

// SetBuildSetting set build setting
func (pbx PBXProject) SetBuildSetting(target string, mode string, key string, val interface{}) {
	bsMap := pbx.getBuildSetting(target, mode)
	bsMap[key] = val
}
