// Code generated from PBXProj.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PBXProj

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PBXProjListener is a complete listener for a parse tree produced by PBXProjParser.
type PBXProjListener interface {
	antlr.ParseTreeListener

	// EnterProject is called when entering the project production.
	EnterProject(c *ProjectContext)

	// EnterFileEncodeInfo is called when entering the fileEncodeInfo production.
	EnterFileEncodeInfo(c *FileEncodeInfoContext)

	// EnterValMap is called when entering the valMap production.
	EnterValMap(c *ValMapContext)

	// EnterValArray is called when entering the valArray production.
	EnterValArray(c *ValArrayContext)

	// EnterValString is called when entering the valString production.
	EnterValString(c *ValStringContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterSectionName is called when entering the sectionName production.
	EnterSectionName(c *SectionNameContext)

	// EnterSectionNoName is called when entering the sectionNoName production.
	EnterSectionNoName(c *SectionNoNameContext)

	// ExitProject is called when exiting the project production.
	ExitProject(c *ProjectContext)

	// ExitFileEncodeInfo is called when exiting the fileEncodeInfo production.
	ExitFileEncodeInfo(c *FileEncodeInfoContext)

	// ExitValMap is called when exiting the valMap production.
	ExitValMap(c *ValMapContext)

	// ExitValArray is called when exiting the valArray production.
	ExitValArray(c *ValArrayContext)

	// ExitValString is called when exiting the valString production.
	ExitValString(c *ValStringContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitSectionName is called when exiting the sectionName production.
	ExitSectionName(c *SectionNameContext)

	// ExitSectionNoName is called when exiting the sectionNoName production.
	ExitSectionNoName(c *SectionNoNameContext)
}
