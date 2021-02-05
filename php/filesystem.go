package php

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// Chgrp - 更改文件组
func Chgrp(name string, uid, gid int) error {
	return Chown(name, uid, gid)
}

// Chmod - Changes file mode
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown - Chown changes the numeric uid and gid of the named file.
func Chown(name string, uid int, gid int) error {
	return os.Chown(name, uid, gid)
}

// Copy - Copies file
func Copy(dstName string, srcName string) (written int64, err error) {
	var src *os.File
	src, err = os.Open(srcName)
	if err != nil {
		return
	}
	defer func() {
		_ = src.Close()
	}()

	var dst *os.File
	dst, err = os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer func() {
		_ = dst.Close()
	}()
	return io.Copy(dst, src)
}

// Delete - Deletes a file
func Delete(name string) error {
	return Unlink(name)
}

// Dirname - Returns a parent directory's path
func Dirname(dirPth string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirPth)
}

// FileClose - 关闭打开的文件指针
func FileClose(file *os.File) error {
	return file.Close()
}

// FileSize filesize()
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// FileExists - Checks whether a file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileUpdateTime - 获取文件修改时间
func FileUpdateTime(file string) time.Time {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fi.ModTime()
}

// Glob - 查找与模式匹配的路径名
func Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}

// IsDir - Tells whether the filename is a directory
func IsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

// IsReadable - Tells whether a file exists and is readable
func IsReadable(name string) bool {
	_, err := syscall.Open(name, syscall.O_RDONLY, 0)
	return err == nil
}

// IsWritable - Tells whether the filename is writable
func IsWritable(name string) bool {
	_, err := syscall.Open(name, syscall.O_WRONLY, 0)
	return err == nil
}

// IsWriteable - Alias of IsWritable()
func IsWriteable(name string) bool {
	return IsWritable(name)
}

// Mkdir - Makes directory
func Mkdir(name string, mode os.FileMode) error {
	return os.Mkdir(name, mode)
}

// Realpath - 返回规范化的绝对路径名
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// Rename - Renames a file or directory
func Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// Touch touch()
func Touch(filename string) (bool, error) {
	fd, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, err
	}
	fd.Close()
	return true, nil
}

// Rmdir — Removes directory
func Rmdir(path string) error {
	return os.RemoveAll(path)
}

// Stat - Gives information about a file
func Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Unlink - Deletes a file
func Unlink(name string) error {
	return os.Remove(name)
}

// FileGetContents - Reads entire file into a string
func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// FilePutContents - Write data to a file
func FilePutContents(filename string, data []byte) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// IsFile Tells whether the filename is a regular file
func IsFile(name string) bool {

	fi, err := os.Stat(name)

	return err == nil && fi.Mode().IsRegular()
}

// Chdir - Change directory
func Chdir(dir string) error {

	return os.Chdir(dir)
}

// GetCwd - Get current directory
func GetCwd() (dir string) {

	dir, err := os.Getwd()
	if err != nil {
		dir = err.Error()
	}
	return
}

// ScanDir — List files and directories inside the specified path
func ScanDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}
