package compare_versions
import (
	"fmt"
	"strings"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	l1 := len(v1)
	l2 := len(v2)

	l := max(l1,l2)
	v1int := make([]int, l)
	v2int := make([]int, l)


	for i:=0; i < l1; i++  {
		vv1, _ := strconv.Atoi(v1[i])
		v1int[i] = vv1
	}
	for i:=0; i < l2; i++  {
		vv2, _ := strconv.Atoi(v2[i])
		v2int[i] = vv2
	}

	for i:=0; i<l; i++ {
		if v1int[i] < v2int[i] {
			return -1
		} else  if v1int[i] > v2int[i] {
			return 1
		}
	}

	return 0
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

