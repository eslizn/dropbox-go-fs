package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"io"
	"io/fs"
	"time"
)

type entry struct {
	files.Client
	files.IsMetadata
}

func (e entry) Name() string {
	switch val := e.IsMetadata.(type) {
	case *files.FileMetadata:
		return val.Name
	case *files.FolderMetadata:
		return val.Name
	default:
		//panic
		return ""
	}
}

func (e entry) Size() int64 {
	switch val := e.IsMetadata.(type) {
	case *files.FileMetadata:
		return int64(val.Size)
	//case *files.FolderMetadata:
	//	return 0
	default:
		return 0
	}
}

func (e entry) Mode() fs.FileMode {
	return 0
}

func (e entry) Type() fs.FileMode {
	return 0
}

func (e entry) ModTime() time.Time {
	switch val := e.IsMetadata.(type) {
	case *files.FileMetadata:
		return val.ServerModified
	default:
		return time.Time{}
	}
}

func (e entry) IsDir() bool {
	_, assert := e.IsMetadata.(*files.FolderMetadata)
	return assert
}

func (e entry) Sys() interface{} {
	return nil
}

func (e entry) Stat() (fs.FileInfo, error) {
	return e, nil
}

func (e entry) Info() (fs.FileInfo, error) {
	return e, nil
}

func (e entry) Read(buff []byte) (int, error) {
	file, assert := e.IsMetadata.(*files.FileMetadata)
	if !assert {
		return 0, io.ErrUnexpectedEOF
	}
	_, reader, err := e.Client.Download(&files.DownloadArg{
		Path: file.PathLower,
	})
	if err != nil {
		return 0, err
	}
	return io.ReadFull(reader, buff)
}

func (e entry) Close() error {
	return nil
}
