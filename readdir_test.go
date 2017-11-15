package watcher

import (
	"io/ioutil"
	"testing"

	"github.com/ShevaXu/web-utils/assert"
)

const currentDir = "."

func TestLs(t *testing.T) {
	names, err := ls(currentDir)
	if err != nil {
		t.Error(err)
	}
	t.Log(names)
}

func TestLa(t *testing.T) {
	names, err := la(currentDir)
	if err != nil {
		t.Error(err)
	}
	t.Log(names)
}

func TestReadDir(t *testing.T) {
	a := assert.NewAssert(t)

	list, err := readDir(currentDir)
	a.NoError(err, "readDir should succeed")

	list2, err := ioutil.ReadDir(currentDir)
	a.NoError(err, "ioutil.ReadDir should succeed")

	a.Equal(list, list2, "Returns the same")
}
