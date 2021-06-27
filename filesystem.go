package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"io/fs"
	"io/ioutil"
)

type FileSystem struct {
	files.Client
}

func (fs FileSystem) Open(name string) (fs.File, error) {
	meta, err := fs.Client.GetMetadata(&files.GetMetadataArg{Path: name})
	if err != nil {
		return nil, err
	}
	return entry{
		Client:     fs.Client,
		IsMetadata: meta,
	}, nil
}

func (fs FileSystem) Stat(name string) (fs.FileInfo, error) {
	meta, err := fs.Client.GetMetadata(&files.GetMetadataArg{Path: name})
	if err != nil {
		return nil, err
	}
	return entry{
		Client:     fs.Client,
		IsMetadata: meta,
	}, nil
}

func (fs FileSystem) Glob(name string) ([]string, error) {

	return nil, nil
}

//func (fs FileSystem) Sub(dir string) (fs.FS, error) {
//
//	return nil, nil
//}

func (fs FileSystem) ReadDir(name string) ([]fs.DirEntry, error) {

	return nil, nil
}

func (fs FileSystem) ReadFile(name string) ([]byte, error) {
	_, reader, err := fs.Client.Download(&files.DownloadArg{
		Path: name,
	})
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}
