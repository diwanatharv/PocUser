package logger

import (
	"os"
)

func Createloggerfile() *os.File {
	var Logfile, _ = os.Create("logfile.txt")
	return Logfile
}
func Closeloggerfile() {
	defer Createloggerfile().Close()
}
