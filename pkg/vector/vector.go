package vector

import "log"

type Vec []float64

func (vec Vec) Len() int {
	return len(vec)
}

func (vec Vec) Sum() float64 {
	var sum float64
	for _, x := range vec {
		sum += x
	}
	return sum
}

func (vec Vec) New() Vec {
	return make(Vec, 0, vec.Len())
}

func (vec Vec) Copy() Vec {
	return append(vec.New(), vec...)
}

func (vec Vec) DivideScalar(scalar float64) Vec {
	var divided = vec.New()
	for _, x := range vec {
		divided = append(divided, x/scalar)
	}
	return divided
}

func (vec Vec) Average() float64 {
	return vec.Sum() / float64(vec.Len())
}

func (vec Vec) Map(op func(x float64) float64) Vec {
	var mapped = vec.New()
	for _, x := range vec {
		mapped = append(mapped, op(x))
	}
	return mapped
}

// alpha determines the degree of smoothing
func (vec Vec) Smooth(alpha float64) Vec {
	if alpha < 0 || alpha > 1 {
		log.Panicf("[nodeMetrics.pkg.vector.Vec.Smooth] the alpha parameter in the range of 0..1 is expected, got %v", alpha)
	}
	if vec.Len() == 0 {
		return vec.New()
	}
	var average = vec.Average()
	return vec.Map(func(x float64) float64 {
		return (1-alpha)*x + alpha*average
	})
}

func (vec Vec) Uints() []uint64 {
	var uints = make([]uint64, 0, vec.Len())
	for _, x := range vec {
		uints = append(uints, uint64(x))
	}
	return uints
}
