package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func init() {
	godotenv.Load()
}

func TestDropbox(t *testing.T) {
	client := files.New(dropbox.Config{Token: os.Getenv("ACCESS_TOKEN")})
	list, err := client.ListFolder(files.NewListFolderArg(""))
	if err != nil {
		t.Error(err)
		return
	}
	for k := range list.Entries {
		//switch entry := list.Entries[k].(type) {
		//case *files.DeletedMetadata:
		//
		//case *files.FileMetadata:
		//
		//case *files.FolderMetadata:
		//
		//}
		t.Logf("%+v\n", list.Entries[k])
	}
	entry, err := client.GetMetadata(&files.GetMetadataArg{Path: "/dropbox 快速入门.pdf"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", entry)
}
