

func myAtoi(str string) int {
    mapping := make(map[byte]int, 0)
    mapping['1'] = 1
    mapping['2'] = 2
    mapping['3'] = 3
    mapping['4'] = 4
    mapping['5'] = 5
    mapping['6'] = 6
    mapping['7'] = 7
    mapping['8'] = 8
    mapping['9'] = 9    
    mapping['0'] = 0
    
    str = strings.TrimLeft(str, " ")
    result := 0 
    sign := 1 
    
    if len(str) == 0 {
        return 0
    }
    
    if str[0] == '-' {
        sign = -1 
        str = str[1:]
    }else if str[0] == '+' {
        sign = 1 
        str = str[1:]
    }
    
    for i := 0; i<len(str); i++{
        if digit, ok := mapping[str[i]]; !ok {
            break 
        } else {
                result = result * 10 + digit                 
        }    
    }
    
    if result > math.MaxInt32 {
        return math.MaxInt32
    }
    if result < math.MinInt32 {
        fmt.Printf("returning here.\n")
        return math.MinInt32
    }
    result = result * sign

    return result 
}
