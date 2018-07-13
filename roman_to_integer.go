func romanToInt(s string) int {
    l := len(s)
    if len(s) == 0 {
        return 0
    }
    sum := 0
    
    for i:=l-1; i >= 0; i-- {
        sum = sum + value(string(s[i]))
        if i != 0 && value(string(s[i])) > (value(string(s[i-1]))) {
            sum = sum - value(string(s[i-1]))
            i--
        }
    }
    return sum 
}

var mapping map[string]int = map[string]int {
    "I":1, "V":5, "X":10, "L":50, "C":100, "M":1000, "D":500, 
}

func value (str string) int {
    return mapping[str] 
}
