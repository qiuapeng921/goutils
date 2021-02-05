package php

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Exec(s string) error {
	return exec.Command(s).Run()
}

// Basename - Returns trailing name component of path
func Basename(path string) string {
	return filepath.Base(path)
}

// Exit — Output a message and terminate the current script
func Exit(code int) {
	os.Exit(code)
}

// Die — Equivalent to exit
func Die(code int) {
	Exit(code)
}

// PutEnv
func GetEnv(name string) string {
	return os.Getenv(name)
}

// PutEnv
// The setting, like "FOO=BAR"
func PutEnv(setting string) error {
	s := strings.Split(setting, "=")
	if len(s) != 2 {
		panic("setting: invalid")
	}
	return os.Setenv(s[0], s[1])
}

// PrintR - Prints human-readable information about a variable
func PrintR(v interface{}) {
	fmt.Print(v)
}

// Print - Output a string
func Print(v interface{}) {
	print(v)
}

// Echo - Output one or more strings
func Echo(args ...interface{}) {
	fmt.Println(args...)
}

// SysGetTempDir - Returns directory path used for temporary files
func SysGetTempDir() string {

	return os.TempDir()
}
