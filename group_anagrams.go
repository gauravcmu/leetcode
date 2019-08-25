import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mapping := make(map[string][]string, 0)
	res := make([][]string, 0)
	for _, v := range strs {
		r := []rune(v)
		sort.Sort(sortRunes(r))
		if _, ok := mapping[string(r)]; ok {
			mapping[string(r)] = append(mapping[string(r)], v)
		} else {
			mapping[string(r)] = make([]string, 0)
			mapping[string(r)] = append(mapping[string(r)], v)
		}

		fmt.Printf("r:%v\n", string(r))
	}

	for _, val := range mapping {
		res = append(res, val)
	}

	fmt.Printf("%v\n", res)
	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
