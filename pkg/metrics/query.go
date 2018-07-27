package metrics

import (
	"fmt"
	"strings"

	influx "github.com/influxdata/influxdb/client/v2"
)

func (flux *Influx) SelectFromWithLimit(from string, limit uint, columns ...string) influx.Query {
	if limit == 0 {
		limit = 0
	}
	if len(columns) == 0 {
		columns = []string{"*"}
	}
	return influx.NewQuery(fmt.Sprintf("SELECT %s FROM %s LIMIT %d",
		strings.Join(columns, ", "), from, limit), flux.dbName, "1") //TODO: fix precision
}

func (flux *Influx) SelectFrom(from string, columns ...string) influx.Query {
	if len(columns) == 0 {
		columns = []string{"*"}
	}
	return influx.NewQuery(fmt.Sprintf("SELECT %s FROM %s",
		strings.Join(columns, ", "), from), flux.dbName, "1") //TODO: fix precision
}

func (flux *Influx) Query(cmd string, args ...interface{}) ([]influx.Result, error) {
	q := influx.Query{
		Command:  fmt.Sprintf(cmd, args...),
		Database: flux.dbName,
	}
	response, err := flux.Client.Query(q)
	if err != nil {
		return nil, err
	}
	return response.Results, response.Error()
}
