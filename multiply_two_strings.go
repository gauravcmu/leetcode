func multiply(num1 string, num2 string) string {
	atemp := num1
	btemp := num2
    
    if (num1[0] == 48) || (num2[0] == 48){
        return "0" 
    }

	var m int
	var n int
	var prod int
	var carry int
	var num int

	carry = 0

	result := make([][]int, 0)
	count := 0
	for len(atemp) != 0 {
		m = int(atemp[len(atemp)-1] - 48)
		atemp = atemp[:len(atemp)-1]

		res := make([]int, 0)
		for i := 0; i < count; i++ {
			res = append(res, 0)
		}

		for len(btemp) != 0 {
			n = int(btemp[len(btemp)-1] - 48)
			btemp = btemp[:len(btemp)-1]
			//fmt.Printf("m:%v n:%v\n", m, n)
			prod = int(m*n) + carry
			carry = prod / 10
			num = prod % 10
			res = append(res, num)
			if len(btemp) == 0 && carry != 0 {
				//fmt.Printf("HERE I AM \n\n")
				res = append(res, carry)
				carry = 0
			}
		}
		//count increment
		count++
		result = append(result, res)
		btemp = num2
	}

	final := make([]int, 0)
	carry = 0
	for _, value := range result {
		//fmt.Printf("value = %v\n", value)
		for i := 0; i < len(value); i++ {
			if i < len(final) {
				//fmt.Printf("final[i]: %v  value[i]:%v carry: %v\n", final[i], value[i], carry)
				sum := (final[i] + value[i] + carry)
				final[i] = sum % 10
				carry = sum / 10
			} else {
				//fmt.Printf("value[i]:%v carry: %v\n", value[i], carry)
				sum := (value[i] + carry)
				final = append(final, sum%10)
				carry = sum / 10
			}
		}
		if carry != 0 {
			final = append(final,carry)
			carry = 0 
		}
	}

	//fmt.Printf("res %v\n", result)
	//fmt.Printf("final %v\n", final)

	ans := make([]byte, 0)
	for _, val := range final {
		ans = append([]byte{byte(val + 48)}, ans...)
	}
	//fmt.Printf("ans %s\n", ans)

	return string(ans)
}
