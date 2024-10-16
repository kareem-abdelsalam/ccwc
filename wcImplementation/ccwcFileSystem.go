package wcImplementation

import (
	"io"
	"os"
)

type OSFileSystem interface {
	Open(name string) (OSFile, error)
	Stat(name string) (os.FileInfo, error)
}

type OSFile interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

type osFS struct{}

func (osFS) Open(name string) (OSFile, error)      { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

func NewOSFS() *osFS {
	return &osFS{}
}
