package prometheus

import (
	"context"
	"fmt"
	"time"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/prometheus/client_golang/api"
	prometheusAPI_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type API struct {
	timeout time.Duration
	prometheusAPI_v1.API
}

type Config struct {
	Timeout time.Duration
	Addr    string
}

func New(config Config) *API {
	var apiClient, err = api.NewClient(api.Config{
		Address: config.Addr,
	})
	if err != nil {
		panic(err)
	}
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}
	return &API{
		timeout: config.Timeout,
		API:     prometheusAPI_v1.NewAPI(apiClient),
	}
}

func (api *API) ctxWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), api.timeout)
}

func (api *API) Query(query string, args ...interface{}) (model.Value, error) {
	var ctx, done = api.ctxWithTimeout()
	defer done()
	return api.API.Query(ctx, fmt.Sprintf(query, args...), time.Now())
}

func (api *API) QueryRange(r metrics.Range, query string, args ...interface{}) (model.Value, error) {
	var ctx, done = api.ctxWithTimeout()
	defer done()
	return api.API.QueryRange(ctx, fmt.Sprintf(query, args...), prometheusAPI_v1.Range{
		Start: r.From,
		End:   r.To,
		Step:  r.Step,
	})
}
