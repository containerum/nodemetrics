package metrics

import "time"

type Precision string

const (
	FullPrecision Precision = ""
)

func DefaultSeries() SeriesConfig {
	return SeriesConfig{}
}

type SeriesConfig struct {
	Step      *time.Duration
	Precision *Precision
}

func (config SeriesConfig) Merge(another SeriesConfig) SeriesConfig {
	if another.Precision != nil {
		config.Precision = another.Precision
	}
	if another.Step != nil {
		config.Step = another.Step
	}
	return config
}

func MergeSeriesConfigs(configs []SeriesConfig) SeriesConfig {
	var config SeriesConfig
	for _, another := range configs {
		config = config.Merge(another)
	}
	return config
}
