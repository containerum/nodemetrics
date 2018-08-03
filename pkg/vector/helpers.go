package vector

func MakeVec(size int, source func(index int) float64) Vec {
	var vec = make(Vec, 0, size)
	for i := 0; i < size; i++ {
		vec = append(vec, source(i))
	}
	return vec
}

func FromInts(ints []int) Vec {
	return MakeVec(len(ints), func(index int) float64 {
		return float64(ints[index])
	})
}
