

func numberToWords(num int) string {
	res := make([]byte, 0)
	thousands := []string{"", "Thousand ", "Million ", "Billion "}

	count := 0

	if num == 0 {
		return "Zero"
	}

	for num != 0 {
		r := num % 1000
		//read the 3 digit remainder - now treat this as 3 digit to word + thousands[count]
        
        if r != 0 {
		    s := threeDigitToString(r)    
            res = append([]byte(thousands[count]), res...)
            res = append([]byte(s), res...)
        }
		num = num / 1000
		count++
	}

	return strings.TrimSuffix(string(res), " ") 
}

func threeDigitToString(num int) string {

	tens := []string{"", "Ten", "Twenty ", "Thirty ", "Forty ", "Fifty ", "Sixty ", "Seventy ", "Eighty ", "Ninety "}
	lessthantwenty := []string{"", "One ", "Two ", "Three ", "Four ", "Five ", "Six ", "Seven ", "Eight ", "Nine ", "Ten ", "Eleven ",
		"Twelve ", "Thirteen ", "Fourteen ", "Fifteen ", "Sixteen ", "Seventeen ", "Eighteen ", "Nineteen "}

	if num == 0 {
		return ""
	} else if num < 20 {
		return lessthantwenty[num]
	} else if num < 100 {
		return tens[num/10] + threeDigitToString(num%10)
	} else {
		//number between 100 - 999

		return lessthantwenty[num/100] + "Hundred " + threeDigitToString(num%100)
	}

}

