package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// env := os.Environ()
	// procAttr := &os.ProcAttr{
	// 	Env: env,
	// 	Files: []*os.File{
	// 		os.Stdin,
	// 		os.Stdout,
	// 		os.Stderr,
	// 	},
	// }
	// 1st example: list files
	// pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	// if err != nil {
	// 	fmt.Printf("Error %v starting process!", err) //
	// 	os.Exit(1)
	// }
	// fmt.Printf("The process id is %v", pid)
	env := os.Environ()

	cmd := exec.Command("ls", env...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}
