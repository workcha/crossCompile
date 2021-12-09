package common

import (
	"os"
	"os/exec"
)

func setEnv(system,code string){
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOOS",system)
	if code == "64"{
		os.Setenv("GOARCH","amd"+code)
	}else{
		os.Setenv("GOARCH","386")
	}

}

func commandExec(projectName,system,code string){
	var cmd *exec.Cmd
	if system == "windows"{
		cmd =exec.Command("go","build","-ldflags","-w -s","-trimpath","-o","..\\..\\release\\"+projectName+ "_"+system + "_" + code + ".exe",projectName+".go")
	}else{
		cmd =exec.Command("go","build","-ldflags","-w -s","-trimpath","-o","..\\..\\release\\"+projectName+ "_"+system + "_" + code,projectName+".go")
	}
	cmd.Dir = GetPwd() + "\\project\\" + projectName
	cmd.Run()
}

func FileExists(path string)bool{
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CompileAll(projectName string){
	CompileWin64(projectName)
	CompileWin32(projectName)
	CompileLin64(projectName)
	CompileLin32(projectName)
	CompileMac64(projectName)
	//CompileMac32(projectName,projectMain)
}

//编译win64位
func CompileWin64(projectName string){
	setEnv("windows","64")
	commandExec(projectName,"windows","64")
	if FileExists(".\\release\\"+projectName+ "_windows_64.exe"){
		println("Compile "+projectName+ "_windows_64.exe success")
	}
}

//编译win32位
func CompileWin32(projectName string){
	setEnv("windows","32")
	commandExec(projectName,"windows","32")
	if FileExists(".\\release\\"+projectName+ "_windows_32.exe"){
		println("Compile "+projectName+ "_windows_32.exe success")
	}
}

//编译lin64位
func CompileLin64(projectName string){
	setEnv("linux","64")
	commandExec(projectName,"linux","64")
	if FileExists(".\\release\\"+projectName+ "_linux_64"){
		println("Compile "+projectName+ "_linux_64 success")
	}
}

//编译lin32位
func CompileLin32(projectName string){
	setEnv("linux","32")
	commandExec(projectName,"linux","32")
	if FileExists(".\\release\\"+projectName+ "_linux_32"){
		println("Compile "+projectName+ "_linux_32 success")
	}
}

//编译mac64位
func CompileMac64(projectName string){
	setEnv("darwin","64")
	commandExec(projectName,"mac","64")
	if FileExists(".\\release\\"+projectName+ "_mac_64"){
		println("Compile "+projectName+ "_mac_64 success")
	}
}
//编译mac32位
//func CompileMac32(projectName,projectMain string){
//	setEnv("darwin","32")
//	commandExec(projectName,projectMain,"mac","32")
//}

