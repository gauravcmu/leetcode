import "fmt"

func generateParenthesis(n int) []string {
	output := ""
	res := make([]string, 0)
	helper(&res, output, 0, 0, n)
	return res
}

func helper(res *[]string, output string, open, closed, n int) {
	fmt.Printf("res:%v output:%v open:%v closed:%v\n", *res, output, open, closed)
	if open == closed && open == n {
		*res = append(*res, output)
		return
	} else {
		if open < n {
			helper(res, output+"(", open+1, closed, n)
		}
		if closed < open {
			helper(res, output+")", open, closed+1, n)
		}
	}
}