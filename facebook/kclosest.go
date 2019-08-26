import "sort"

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		return (points[a][0]*points[a][0] + points[a][1]*points[a][1]) < (points[b][0]*points[b][0] + points[b][1]*points[b][1])
	})

	return points[:K]
}