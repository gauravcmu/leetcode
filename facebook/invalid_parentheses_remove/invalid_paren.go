func removeInvalidParentheses(s string) []string {
	result := make(map[string]bool, 0)
	helper(s, result, "", 0, 0, "  ")
	max := 0
	for key, _ := range result {
		if len(key) > max {
			max = len(key)
		}
	}

	retset := make([]string, 0)
	for key, _ := range result {
		if len(key) == max {
			retset = append(retset, key)
		}
	}
	return retset
}

func helper(s string, result map[string]bool, chosen string, open, close int, format string) {
	//fmt.Printf("%ss: %s, chosen: %s, open: %v close: %v\n", format, s, chosen, open, close)
	if s == "" {
		if open == close {
			if _, ok := (result)[chosen]; !ok {
				result[chosen] = true
				//	fmt.Printf("chosen: %s - store in map\n\n", chosen)

			}
		}
		return
	} else {
		if s[0] == '(' {
			//chose opening bracket
			temp := chosen + string('(')
			helper(s[1:], result, temp, open+1, close, format+"  ")
			helper(s[1:], result, chosen, open, close, format+"  ")
		} else if s[0] == ')' {
			if close < open {
				//chose closing bracket
				temp := chosen + string(')')
				helper(s[1:], result, temp, open, close+1, format+"  ")
				helper(s[1:], result, chosen, open, close, format+"  ")
			} else {
				helper(s[1:], result, chosen, open, close, format+"  ")
			}
		} else {
			helper(s[1:], result, chosen, open, close, format+"  ")
		}
	}
}