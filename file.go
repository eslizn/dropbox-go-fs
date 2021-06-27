package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"io"
	"io/fs"
	"time"
)

type file struct {
	files.Client
	files.IsMetadata
}

func (fs file) Name() string {
	switch val := fs.IsMetadata.(type) {
	case *files.FileMetadata:
		return val.Name
	case *files.FolderMetadata:
		return val.Name
	default:
		//panic
		return ""
	}
}

func (fs file) Size() int64 {
	switch val := fs.IsMetadata.(type) {
	case *files.FileMetadata:
		return int64(val.Size)
	//case *files.FolderMetadata:
	//	return 0
	default:
		return 0
	}
}

func (fs file) Mode() fs.FileMode {
	//@todo
	return 0
}

func (fs file) ModTime() time.Time {
	switch val := fs.IsMetadata.(type) {
	case *files.FileMetadata:
		return val.ServerModified
	default:
		return time.Time{}
	}
}

func (fs file) IsDir() bool {
	_, assert := fs.IsMetadata.(*files.FolderMetadata)
	return assert
}

func (fs file) Sys() interface{} {
	return nil
}

func (fs file) Stat() (fs.FileInfo, error) {
	return fs, nil
}

func (fs file) Read(buff []byte) (int, error) {
	file, assert := fs.IsMetadata.(*files.FileMetadata)
	if !assert {
		return 0, io.ErrUnexpectedEOF
	}
	_, reader, err := fs.Client.Download(&files.DownloadArg{
		Path: file.PathLower,
	})
	if err != nil {
		return 0, err
	}
	return io.ReadFull(reader, buff)
}

func (fs file) Close() error {
	return nil
}
