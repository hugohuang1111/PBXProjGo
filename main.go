package main

import (
	"PBXProjGo/pbxproj"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("args wrong")
		return
	}
	pbxFilePath := os.Args[1]
	pbxProject, e := pbxproj.NewPBXProject(pbxFilePath)
	if nil != e {
		fmt.Println(e.Error())
		return
	}

	// uuid := pbxProject.FindGroupUUID("ios")
	uuid := pbxProject.AddFile("ios", "./projects/AppDelegate.h")
	if 0 == len(uuid) {
		fmt.Println("Error, add file failed")
		return
	}
	uuid = pbxProject.AddFramework("./project/a.framework")
	if 0 == len(uuid) {
		fmt.Println("Error, add framework failed")
		return
	}

	e = pbxProject.Save("newpbx.pbxproj")
	if nil != e {
		fmt.Println(e.Error())
		return
	}
}
