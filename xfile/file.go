package xfile

import (
	"ksitigarbha/xstrings"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ParentDir 返回一个路径的父级目录
// 如果path 是一个文件则返回父级目录，如果是目录就返回自己本身
func ParentDir(path string) string {
	isDir, err := IsDirectory(path)
	if err != nil || isDir {
		return path
	}
	return getParentDir(path)
}

func getParentDir(directory string) string {
	if runtime.GOOS == "windows" {
		directory = strings.Replace(directory, "\\", "/", -1)
	}
	return xstrings.Sub(directory, 0, strings.LastIndex(directory, "/"))
}

func IsDirectory(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	switch mode := f.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, nil
	}
	return false, nil
}

// IsEmpty checks whether the given <path> is empty.
// If <path> is a folder, it checks if there's any file under it.
// If <path> is a file, it checks if the file size is zero.
//
// Note that it returns true if <path> does not exist.
func IsEmpty(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		file, err := os.Open(path)
		if err != nil {
			return true
		}
		defer file.Close()
		names, err := file.Readdirnames(-1)
		if err != nil {
			return true
		}
		return len(names) == 0
	} else {
		return stat.Size() == 0
	}
}

// The parameter <path> is suggested to be absolute path.
// the parent dir will be created automatically if not exists
// Note: Remember to close the returning *os.File object to make it unusable for I/O
func Create(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if !Exists(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}
	return os.Create(path)
}

// RealPath converts the given <path> to its absolute path
// and checks if the file path exists.
// If the file does not exist, return an empty string.
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

// Exists checks whether given <path> exist.
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
//
// Note: the result contains symbol '.'.
func Ext(path string) string {
	ext := filepath.Ext(path)
	// 对于 windows 下的shell 重定向到文件(>> file 或者> file)的时候，会多一个^M，显示就是问号在最后
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// IsFile checks whether given <path> a file, which means it's not a directory.
// Note that it returns false if the <path> does not exist.
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}
