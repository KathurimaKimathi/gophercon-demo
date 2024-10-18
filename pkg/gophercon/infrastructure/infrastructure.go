package infrastructure

import (
	"context"
	"io"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/repository"
)

type Upload interface {
	UploadMedia(ctx context.Context, name string, file io.Reader) (string, error)
}

type Infrastructure struct {
	Repository repository.IUser
	Upload     Upload
}

// NewInfrastructureInitializer initializes a new Infrastructure
func NewInfrastructureInitializer(
	db repository.IUser,
	upload Upload,
) *Infrastructure {
	return &Infrastructure{
		Repository: db,
		Upload:     upload,
	}
}
