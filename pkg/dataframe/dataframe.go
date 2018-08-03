package dataframe

import (
	"bytes"
	"fmt"
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

func (dataframe Dataframe) String() string {
	var buf = &bytes.Buffer{}
	for i := 0; i < dataframe.Len(); i++ {
		var label, value = dataframe.Get(i)
		fmt.Fprintf(buf, "%s : %v%s\n", label, value, dataframe.Units)
	}
	return buf.String()
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

func MakeDataframe(units string, l int, source func(index int) (string, float64)) Dataframe {
	var dataframe = Dataframe{
		Units:  units,
		Values: make(vector.Vec, 0, l),
		Labels: make([]string, 0, l),
	}
	for index := 0; index < l; index++ {
		var label, value = source(index)
		dataframe.Labels = append(dataframe.Labels, label)
		dataframe.Values = append(dataframe.Values, value)
	}
	return dataframe
}
