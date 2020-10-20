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

	targets := pbxProject.GetTargets()
	targetName := targets[0]

	// add source file
	// pbxProject.AddFile(targetName, "ios", "/Users/hugo/Documents/work/learn/PBXProjGo/projects/Test/bb/A1.h")
	// pbxProject.AddFile(targetName, "ios", "/Users/hugo/Documents/work/learn/PBXProjGo/projects/Test/bb/A1.cpp")
	// pbxProject.AddFile(targetName, "ios", "/Users/hugo/Documents/work/learn/PBXProjGo/projects/Test/bb/A1.mm")

	// add framework
	// pbxProject.AddFramework(targetName, "ios", "/Users/hugo/Documents/work/learn/PBXProjGo/projects/Test/t.framework")
	pbxProject.AddFramework(targetName, "/Users/hugo/Documents/work/learn/PBXProjGo/projects/Test/t.xcframework")

	// modify build setting
	// val := pbxProject.GetBuildSetting(targetName, "Debug", "OTHER_LDFLAGS")
	// lfs := make([]interface{}, 0)
	// if nil != val {
	// 	lfs = val.([]interface{})
	// }
	// lfs = append(lfs, "kk-kk")
	// pbxProject.SetBuildSetting(targetName, "Debug", "OTHER_LDFLAGS", lfs)

	// save file
	e = pbxProject.Save("newpbx.pbxproj")
	if nil != e {
		fmt.Println(e.Error())
		return
	}
}
