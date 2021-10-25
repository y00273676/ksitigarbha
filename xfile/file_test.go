package xfile_test

import (
	"io/ioutil"
	"ksitigarbha/xfile"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func testpath() string {
	return strings.TrimRight(os.TempDir(), "\\/")
}

func createDir(paths string) {
	TempDir := testpath()
	_ = os.Mkdir(TempDir+paths, 0777)
}

func createTestFile(filename, content string) error {
	TempDir := testpath()
	err := ioutil.WriteFile(TempDir+filename, []byte(content), 0666)
	return err
}

func deleteDir(filenames string) {
	_ = os.RemoveAll(testpath() + filenames)
}

func TestRealPath(t *testing.T) {
	result := xfile.RealPath("./")
	expected, _ := filepath.Abs("./")
	assert := tassert.New(t)
	assert.Equal(result, expected)
}

func TestIsDir(t *testing.T) {
	path := "/testfile"
	createDir(path)
	defer deleteDir(path)
	assert := tassert.New(t)
	var isDirA, _ = xfile.IsDirectory(testpath() + path)
	var isDirB, _ = xfile.IsDirectory(testpath() + path + "1")

	assert.Equal(isDirA, true)
	assert.Equal(isDirB, false)
}

func TestIsEmpty(t *testing.T) {
	assert := tassert.New(t)
	path1 := "/testdir_" + "1"
	createDir(path1)
	defer deleteDir(path1)

	assert.Equal(xfile.IsEmpty(testpath()+path1), true)

	path2 := "/testdir_" + "2"

	_ = createTestFile(path2, "hi")
	defer deleteDir(path2)

	assert.Equal(xfile.IsEmpty(testpath()+path2), false)
}

func TestExtName(t *testing.T) {
	assert := tassert.New(t)
	path1 := "/testdir_" + ".txt"
	createDir(path1)
	defer deleteDir(path1)

	assert.Equal(xfile.Ext(path1), ".txt")
}

func TestCreate(t *testing.T) {
	assert := tassert.New(t)
	fileObject, err := xfile.Create(testpath() + "/testfile_cc1.txt")
	assert.Nil(err)
	defer deleteDir("/testfile_cc1.txt")
	_ = fileObject.Close()
}

func TestIsFile(t *testing.T) {
	assert := tassert.New(t)
	file1 := "/testfile_tt.txt"
	_ = createTestFile(file1, "123")
	defer deleteDir(file1)
	ok := xfile.IsFile(testpath() + file1)
	assert.Equal(ok, true)
}
