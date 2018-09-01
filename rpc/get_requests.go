package rpc

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/PayloadPro/pro.payload.api/deps"
	"github.com/PayloadPro/pro.payload.api/models"
	"github.com/gorilla/mux"
)

// GetRequestsResponse is the response from the GetPayload endpoint
type GetRequestsResponse struct {
	Requests []models.Request `json:"requests"`
}

// GetRequests is a func which takes the incoming request for the bin requests and returns the GetRequestsResponse
type GetRequests func(context.Context, *http.Request) (*GetRequestsResponse, int, error)

// NewGetRequestsForBin is the concrete func for GetRequests
func NewGetRequestsForBin(services *deps.Services) GetRequests {
	return func(ctx context.Context, r *http.Request) (*GetRequestsResponse, int, error) {

		// get the bin from the DB based on ID in the URL to check it exists
		vars := mux.Vars(r)
		bin, err := services.Bin.GetByID(vars["id"])
		if err != nil {
			return nil, http.StatusNotFound, err
		}

		// bin exists, get the requests
		var requests = make([]models.Request, 0)
		if requests, err = services.Request.GetRequestsForBin(bin.ID); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return &GetRequestsResponse{
			Requests: requests,
		}, http.StatusOK, nil
	}
}
