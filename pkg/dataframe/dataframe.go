package dataframe

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/vector"
	"github.com/containerum/nodeMetrics/utils/metatime"
)

type Dataframe struct {
	Units  string     `json:"units"`
	Values vector.Vec `json:"values"`
	Labels []string   `json:"labels"`
}

func (dataframe Dataframe) WithUnits(units string) Dataframe {
	return Dataframe{
		Units:  units,
		Labels: append([]string{}, dataframe.Labels...),
		Values: dataframe.Values.Copy(),
	}
}

func (dataframe Dataframe) Get(index int) (label string, value float64) {
	return dataframe.Labels[index], dataframe.Values[index]
}

func (dataframe Dataframe) Len() int {
	return len(dataframe.Values)
}

func NewFromTicks(units string, from, to time.Time, step time.Duration, values vector.Vec) Dataframe {
	var N = to.Sub(from) / step
	var labels = make([]string, 0, N)
	var tick = from
	for tick.Before(to) {
		labels = append(labels, tick.Format(metatime.ISO8601))
		tick = tick.Add(step)
	}
	return Dataframe{
		Units:  units,
		Labels: labels,
		Values: values.Copy(),
	}
}
