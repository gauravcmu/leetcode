import "sort"

type info struct {
	distance int
	point    []int
}

func kClosest(points [][]int, K int) [][]int {
	result := make([]info, 0)

	for _, point := range points {
		result = append(result, info{distance: findDistance(point), point: point})
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a].distance < result[b].distance
	})

	res := make([][]int, 0)

	for _, r := range result {
		res = append(res, r.point)
	}
	return res[:K]
}

func findDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}