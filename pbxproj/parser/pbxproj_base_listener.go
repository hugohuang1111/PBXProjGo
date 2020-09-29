// Code generated from PBXProj.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PBXProj

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePBXProjListener is a complete listener for a parse tree produced by PBXProjParser.
type BasePBXProjListener struct{}

var _ PBXProjListener = &BasePBXProjListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePBXProjListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePBXProjListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePBXProjListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePBXProjListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProject is called when production project is entered.
func (s *BasePBXProjListener) EnterProject(ctx *ProjectContext) {}

// ExitProject is called when production project is exited.
func (s *BasePBXProjListener) ExitProject(ctx *ProjectContext) {}

// EnterFileEncodeInfo is called when production fileEncodeInfo is entered.
func (s *BasePBXProjListener) EnterFileEncodeInfo(ctx *FileEncodeInfoContext) {}

// ExitFileEncodeInfo is called when production fileEncodeInfo is exited.
func (s *BasePBXProjListener) ExitFileEncodeInfo(ctx *FileEncodeInfoContext) {}

// EnterValMap is called when production valMap is entered.
func (s *BasePBXProjListener) EnterValMap(ctx *ValMapContext) {}

// ExitValMap is called when production valMap is exited.
func (s *BasePBXProjListener) ExitValMap(ctx *ValMapContext) {}

// EnterValArray is called when production valArray is entered.
func (s *BasePBXProjListener) EnterValArray(ctx *ValArrayContext) {}

// ExitValArray is called when production valArray is exited.
func (s *BasePBXProjListener) ExitValArray(ctx *ValArrayContext) {}

// EnterValString is called when production valString is entered.
func (s *BasePBXProjListener) EnterValString(ctx *ValStringContext) {}

// ExitValString is called when production valString is exited.
func (s *BasePBXProjListener) ExitValString(ctx *ValStringContext) {}

// EnterKey is called when production key is entered.
func (s *BasePBXProjListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasePBXProjListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BasePBXProjListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePBXProjListener) ExitValue(ctx *ValueContext) {}

// EnterSectionName is called when production sectionName is entered.
func (s *BasePBXProjListener) EnterSectionName(ctx *SectionNameContext) {}

// ExitSectionName is called when production sectionName is exited.
func (s *BasePBXProjListener) ExitSectionName(ctx *SectionNameContext) {}

// EnterSectionNoName is called when production sectionNoName is entered.
func (s *BasePBXProjListener) EnterSectionNoName(ctx *SectionNoNameContext) {}

// ExitSectionNoName is called when production sectionNoName is exited.
func (s *BasePBXProjListener) ExitSectionNoName(ctx *SectionNoNameContext) {}
