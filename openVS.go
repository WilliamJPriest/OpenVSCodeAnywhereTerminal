package main

import (
	"log"
	"os/exec"
)



func main(){
	path,err := exec.LookPath("C:/Users/Lenovo/AppData/Local/Programs/Microsoft VS Code/code")
	if err != nil {
		if e, ok := err.(*exec.Error);ok && e.Err == exec.ErrNotFound{
			log.Fatal("Executable not found:", e.Name)
		} else {
			log.Fatal(err)
		}}

	 cmd := exec.Command(path)
	 err = cmd.Run()
	 if err != nil {
		 log.Fatal(err)
	 }

}

