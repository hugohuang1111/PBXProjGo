package pbxproj

import (
	"PBXProjGo/pbxproj/parser"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// "PBXProjGo/parser"
// "github.com/antlr/antlr4/runtime/Go/antlr"

// PBXProject pbxproj struct
type PBXProject struct {
	fileEncode string
	project    pbxMap

	uuidMap map[string]string

	pbxParser      *parser.PBXProjParser
	pbxTokenStream *antlr.CommonTokenStream
}

// NewPBXProject create pbx project
func NewPBXProject(pbxPBXProjectPath string) (PBXProject, error) {
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
