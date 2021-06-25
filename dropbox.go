package dropbox

import "io/fs"

type Dropbox struct {
	Token string
}

func (db *Dropbox) Open(name string) (fs.File, error) {
	return nil, nil
}

func New(token string) fs.FS {
	return &Dropbox{Token: token}
}
