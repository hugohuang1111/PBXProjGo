package pbxproj

import (
	"PBXProjGo/pbxproj/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type pbxTreeShapeListener struct {
	*parser.BasePBXProjListener
	pbxProject *PBXProject

	tStack *stack
}

// NewPBXTreeShapeListener new pbx tree shape listener
func newPBXTreeShapeListener(pbxTree *PBXProject) *pbxTreeShapeListener {
	l := new(pbxTreeShapeListener)
	l.pbxProject = pbxTree
	l.tStack = newStack()

	return l
}

func (tsl *pbxTreeShapeListener) EnterEveryRule(rule antlr.ParserRuleContext) {
	pbxTree := tsl.pbxProject
	tStack := tsl.tStack
	switch rule.(type) {
	case *parser.ProjectContext:
	case *parser.FileEncodeInfoContext:
		{
			prc := rule.(*parser.FileEncodeInfoContext)
			token := prc.GetToken(parser.PBXProjParserNString, 0)
			if nil != token {
				pbxTree.fileEncode = token.GetText()
			}
		}
	// case *parser.SectionNameContext:
	// 	{
	// 		rSection := rule.(*parser.SectionNameContext)
	// 		t := rSection.GetToken(parser.PBXProjParserSectionB, 0)
	// 		sec := newMapSection()
	// 		if nil != t {
	// 			str := t.GetText()
	// 			sum := len(str)
	// 			sec.name = str[9 : sum-11]
	// 		}
	// 		tStack.push(sec)
	// 	}
	// case *parser.SectionNoNameContext:
	// 	{
	// 		tStack.push(newMapSection())
	// 	}
	case *parser.KeyContext:
	case *parser.ValueContext:
	case *parser.ValMapContext:
		{
			tStack.push(make(pbxMap))
		}
	case *parser.ValArrayContext:
		{
			tStack.push(make([]interface{}, 0))
		}
	case *parser.ValStringContext:
		{
			rString := rule.(*parser.ValStringContext)
			t := rString.GetToken(parser.PBXProjParserNString, 0)
			s := t.GetText()
			if '"' == s[0] && '"' == s[len(s)-1] {
				s = s[1 : len(s)-1]
			}
			tStack.push(s)

			t = rString.GetToken(parser.PBXProjParserComment, 0)
			if isUUID(s) && nil != t {
				pbxTree.uuidMap[s] = t.GetText()
			}
		}
	default:
	}
}

func (tsl *pbxTreeShapeListener) ExitEveryRule(rule antlr.ParserRuleContext) {
	pbxTree := tsl.pbxProject
	tStack := tsl.tStack
	switch rule.(type) {
	case *parser.ProjectContext:
		{
			if 1 != tStack.len() {
				fmt.Println("ERROR: something wrong, stack should be include 1 item")
			}
			pbxTree.project = tStack.pop().(pbxMap)
		}
	case *parser.FileEncodeInfoContext:
	// case *parser.SectionNameContext,
	// 	*parser.SectionNoNameContext:
	// 	{
	// 		// fmt.Println(">>> Exit section")
	// 		val := tStack.pop().(pbxMapSection)
	// 		pMap := tStack.pop().(pbxMap)
	// 		pMap = append(pMap, val)
	// 		tStack.push(pMap)
	// 		// tStack.dump()
	// 		break
	// 	}
	case *parser.KeyContext:
	case *parser.ValueContext:
		{
			// fmt.Println(">>> Exit value:", rule.GetText())
			// tStack.dump()
			val1 := tStack.pop()
			val2 := tStack.pop()
			if key, ok := val2.(string); ok {
				valMap := tStack.top().(pbxMap)
				valMap[key] = val1
			} else {
				arr := val2.([]interface{})
				arr = append(arr, val1)
				tStack.push(arr)
				// if arr, ok := val2.(pbxMap); ok {
				// 	arr = append(arr, val1.(pbxMapSection))
				// 	tStack.push(arr)
				// } else {
				// }
			}
		}
	case *parser.ValMapContext:
		{
			// val := tStack.pop()
			// if 0 == tStack.len() {
			// 	pbxTree.projectInfo = val.([]mapSection)
			// }
		}
	case *parser.ValArrayContext:
	case *parser.ValStringContext:
	default:
	}
}

func (tsl *pbxTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
}

func (tsl *pbxTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
}
