package visualization

const (
	minWeight   = 0.0
	maxWeight   = 500.0
	smoothedMin = 2.0
	smoothedMax = 20.0
)

func smoothWeight(weight float64) float64 {
	// 线性插值
	smoothedWeight := smoothedMin + (weight-minWeight)*(smoothedMax-smoothedMin)/(maxWeight-minWeight)

	return smoothedWeight
}
