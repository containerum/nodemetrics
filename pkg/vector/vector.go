package vector

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

func (vec Vec) Smooth(alpha float64) Vec {
	if vec.Len() == 0 {
		return vec.New()
	}
	var ema = vec[0]
	return vec.Map(func(x float64) float64 {
		return alpha*x + (1-alpha)*ema
	})
}
