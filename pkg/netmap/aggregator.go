package netmap

import (
	"sort"
)

type (
	// aggregator can calculate some value across all netmap
	// such as median, minimum or maximum.
	aggregator interface {
		Add(float64)
		Compute() float64
	}

	// normalizer normalizes weight.
	normalizer interface {
		Normalize(w float64) float64
	}

	meanSumAgg struct {
		sum   float64
		count int
	}

	meanAgg struct {
		mean  float64
		count int
	}

	minAgg struct {
		min float64
	}

	maxAgg struct {
		max float64
	}

	meanIQRAgg struct {
		k   float64
		arr []float64
	}

	reverseMinNorm struct {
		min float64
	}

	maxNorm struct {
		max float64
	}

	sigmoidNorm struct {
		scale float64
	}

	constNorm struct {
		value float64
	}

	// weightFunc calculates n's weight.
	weightFunc = func(n *Node) float64
)

var (
	_ aggregator = (*meanSumAgg)(nil)
	_ aggregator = (*meanAgg)(nil)
	_ aggregator = (*minAgg)(nil)
	_ aggregator = (*maxAgg)(nil)
	_ aggregator = (*meanIQRAgg)(nil)

	_ normalizer = (*reverseMinNorm)(nil)
	_ normalizer = (*maxNorm)(nil)
	_ normalizer = (*sigmoidNorm)(nil)
	_ normalizer = (*constNorm)(nil)
)

// newWeightFunc returns weightFunc which multiplies normalized
// capacity and price.
func newWeightFunc(capNorm, priceNorm normalizer) weightFunc {
	return func(n *Node) float64 {
		return capNorm.Normalize(float64(n.Capacity)) * priceNorm.Normalize(float64(n.Price))
	}
}

// newMeanAgg returns an aggregator which
// computes mean value by recalculating it on
// every addition.
func newMeanAgg() aggregator {
	return new(meanAgg)
}

// newMinAgg returns an aggregator which
// computes min value.
func newMinAgg() aggregator {
	return new(minAgg)
}

// newMeanIQRAgg returns an aggregator which
// computes mean value of values from IQR interval.
func newMeanIQRAgg() aggregator {
	return new(meanIQRAgg)
}

// newReverseMinNorm returns a normalizer which
// normalize values in range of 0.0 to 1.0 to a minimum value.
func newReverseMinNorm(min float64) normalizer {
	return &reverseMinNorm{min: min}
}

// newSigmoidNorm returns a normalizer which
// normalize values in range of 0.0 to 1.0 to a scaled sigmoid.
func newSigmoidNorm(scale float64) normalizer {
	return &sigmoidNorm{scale: scale}
}

func (a *meanSumAgg) Add(n float64) {
	a.sum += n
	a.count++
}

func (a *meanSumAgg) Compute() float64 {
	if a.count == 0 {
		return 0
	}

	return a.sum / float64(a.count)
}

func (a *meanAgg) Add(n float64) {
	c := a.count + 1
	a.mean = a.mean*(float64(a.count)/float64(c)) + n/float64(c)
	a.count++
}

func (a *meanAgg) Compute() float64 {
	return a.mean
}

func (a *minAgg) Add(n float64) {
	if a.min == 0 || n < a.min {
		a.min = n
	}
}

func (a *minAgg) Compute() float64 {
	return a.min
}

func (a *maxAgg) Add(n float64) {
	if n > a.max {
		a.max = n
	}
}

func (a *maxAgg) Compute() float64 {
	return a.max
}

func (a *meanIQRAgg) Add(n float64) {
	a.arr = append(a.arr, n)
}

func (a *meanIQRAgg) Compute() float64 {
	l := len(a.arr)
	if l == 0 {
		return 0
	}

	sort.Slice(a.arr, func(i, j int) bool { return a.arr[i] < a.arr[j] })

	var min, max float64

	const minLn = 4

	if l < minLn {
		min, max = a.arr[0], a.arr[l-1]
	} else {
		start, end := l/minLn, l*3/minLn-1
		iqr := a.k * (a.arr[end] - a.arr[start])
		min, max = a.arr[start]-iqr, a.arr[end]+iqr
	}

	count := 0
	sum := float64(0)

	for _, e := range a.arr {
		if e >= min && e <= max {
			sum += e
			count++
		}
	}

	return sum / float64(count)
}

func (r *reverseMinNorm) Normalize(w float64) float64 {
	if w == 0 {
		return 0
	}

	return r.min / w
}

func (r *maxNorm) Normalize(w float64) float64 {
	if r.max == 0 {
		return 0
	}

	return w / r.max
}

func (r *sigmoidNorm) Normalize(w float64) float64 {
	if r.scale == 0 {
		return 0
	}

	x := w / r.scale

	return x / (1 + x)
}

func (r *constNorm) Normalize(_ float64) float64 {
	return r.value
}
