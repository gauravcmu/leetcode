/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "sort"

func merge(intervals []Interval) []Interval {
	result := make([]Interval, 0)
	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].Start < intervals[b].Start
	})

	result = append(result, intervals[0])
	for key, value := range intervals {
		if key == 0 {
			continue
		}
		if value.Start <= result[len(result)-1].End {
			r := Interval{Start: min(result[len(result)-1].Start, value.Start),
				End: max(result[len(result)-1].End, value.End)}
			result[len(result)-1] = r
		} else {
			result = append(result, value)
		}
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}