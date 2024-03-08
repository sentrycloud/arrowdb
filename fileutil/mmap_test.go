package fileutil

import (
	"fmt"
	"testing"
)

func TestMmap(t *testing.T) {
	mf, err := OpenMmapFile("/tmp/test.txt")
	if err != nil {
		fmt.Printf("OpenMappFile failed: %v\n", err)
		return
	}

	content := string(mf.Bytes())
	fmt.Printf("file content: %s\n", content)

	err = mf.Close()
	if err != nil {
		fmt.Printf("close mmap file failed: %v\n", err)
	} else {
		fmt.Printf("close mmap file success\n")
	}
}
