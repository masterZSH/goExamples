package os

import "os"

func getArgs() []string {
	return os.Args
}

func openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_RDWR)
}
