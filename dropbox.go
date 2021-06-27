package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func New(token string) *FileSystem {
	return &FileSystem{files.New(dropbox.Config{Token: token})}
}
