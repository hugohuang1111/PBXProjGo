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
	e = pbxProject.Save("newpbx.pbxproj")
	if nil != e {
		fmt.Println(e.Error())
		return
	}
}
