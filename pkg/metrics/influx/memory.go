package influx

import (
	"encoding/json"
)

const (
	MEM_COEFF = 1e10
)

func (flux *Influx) MemoryCurrent() (uint64, error) {
	var result, err = flux.Query("SELECT MEAN(value) FROM memory_usage WHERE time > now()-5m")
	if err != nil {
		return 0, err
	}
	if len(result) < 1 {
		return 0, ErrEmptyResult
	}
	if len(result[0].Series) < 1 {
		return 0, ErrNoSeriesFound
	}
	if len(result) < 1 {
		return 0, ErrEmptyResult
	}
	if len(result[0].Series) < 1 {
		return 0, ErrNoSeriesFound
	}
	if len(result[0].Series[0].Values) < 1 {
		return 0, ErrNoValuesFound
	}
	if len(result[0].Series[0].Values[0]) < 2 {
		return 0, ErrInvalidDataPointFormat
	}
	var average, _ = result[0].Series[0].Values[0][1].(json.Number).Float64()
	average /= MEM_COEFF / 32 // TODO: remove hardcoded value
	return uint64(100 * average), nil
}
