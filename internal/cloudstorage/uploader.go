package cloudstorage

import "context"

type Uploader interface {
	Upload(context.Context, []byte, string) (string, error)
}