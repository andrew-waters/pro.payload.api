package rpc

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/PayloadPro/pro.payload.api/deps"
	"github.com/PayloadPro/pro.payload.api/models"
)

// CreateBin is a func which takes the incoming request, saves it persistently
// and returns the CreateBinResponse for the consumer
type CreateBin func(context.Context, *http.Request) (*models.Bin, int, error)

// NewCreateBin is the concrete func for CreateBin
func NewCreateBin(services *deps.Services) CreateBin {
	return func(ctx context.Context, r *http.Request) (*models.Bin, int, error) {

		// create the payload
		var bin *models.Bin
		var err error

		if bin, err = models.NewBin(r); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		// save the bin
		if err = services.Bin.Save(bin); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return bin, http.StatusCreated, nil
	}
}
