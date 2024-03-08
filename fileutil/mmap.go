package fileutil

import (
	"errors"
	"os"
	"syscall"
)

type MmapFile struct {
	f *os.File
	b []byte
}

func OpenMmapFile(path string) (*MmapFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("open file: " + err.Error())
	}

	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, errors.New("stat: " + err.Error())
	}

	size := int(info.Size())

	b, err := syscall.Mmap(int(f.Fd()), 0, size, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		f.Close()
		return nil, errors.New("mmap failed" + err.Error())
	}

	return &MmapFile{f: f, b: b}, nil
}

func (mf *MmapFile) Close() error {
	err0 := syscall.Munmap(mf.b)
	err1 := mf.f.Close()

	if err0 != nil {
		return err0
	}
	return err1
}

func (mf *MmapFile) File() *os.File {
	return mf.f
}

func (mf *MmapFile) Bytes() []byte {
	return mf.b
}
