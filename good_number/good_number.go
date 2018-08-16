func rotatedDigits(N int) int {
    count := 0 
    for i:=0; i<= N; i++ {
        if isNumGood(i) {
            count++
        } 
    }
    return count 
}


func isNumGood(N int) bool {
    var r int 
    res := []int{}

    for (N != 0) {
        r = N % 10 
        N = N / 10
        //if any of 3,4,7 are a digit, it can't be good 
        if r == 3 || r == 4 || r == 7 {
            return false 
        } else {
            //append digits in res - validate later
            res = append(res, r)
        }
    }
    for _, v := range res {
        //if the number is made up of only 0,1,8's, then can't be good. 
        if v == 2 || v == 5 || v == 6 || v == 9 {
            //found a non-0,1,8 digit - good enough
            return true
        }
    }
    return false 
}
