package utility

import (
	"net/url"

	"github.com/yaihub/ndid-go-sdk/pkg/v5/utility/swagger"
)

type Utility struct {
	APIClient   *swagger.APIClient
	endpointUrl string
}

func NewUtility(endpointUrl string) (*Utility, error) {

	u, err := url.Parse(endpointUrl)
	if err != nil {
		return nil, err
	}

	u.JoinPath("/utility")

	swaggerCfg := swagger.NewConfiguration()
	swaggerCfg.BasePath = u.Path
	swaggerCfg.UserAgent = "ndid-go-sdk/0.1"

	return &Utility{
		APIClient:   swagger.NewAPIClient(swaggerCfg),
		endpointUrl: endpointUrl,
	}, nil
}
