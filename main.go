package main

import (
	"PBXProjGo/pbxproj"
	"fmt"
	"path/filepath"
)

func main() {
	pbxFilePath, _ := filepath.Abs("projects/Test/Test.xcodeproj/project.pbxproj")
	pbxProject, e := pbxproj.NewPBXProject(pbxFilePath)
	if nil != e {
		fmt.Println(e.Error())
		return
	}

	// get targets
	targets := pbxProject.GetTargets()
	targetName := targets[0]

	// add source file
	pbxProject.AddSourceFile(targetName, "Test", "Test/T.m")

	// add header file
	pbxProject.AddHeaderFile(targetName, "Test", "Test/T.h")

	// add framework
	pbxProject.AddFramework(targetName, "Accounts.framework")

	// modify build setting
	// val := pbxProject.GetBuildSetting(targetName, "Debug", "OTHER_LDFLAGS")
	// lfs := make([]interface{}, 0)
	// if nil != val {
	// 	lfs = val.([]interface{})
	// }
	// lfs = append(lfs, "kk-kk")
	// pbxProject.SetBuildSetting(targetName, "Debug", "OTHER_LDFLAGS", lfs)

	// save file
	e = pbxProject.Save(pbxFilePath)
	if nil != e {
		fmt.Println(e.Error())
		return
	}
}
