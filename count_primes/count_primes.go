func countPrimes(n int) int {
    if n <=2 {
        return 0
    }
    
    if n == 3 {
        return 1
    }
    count := 1 
    for i:=3; i<n; i+=2 {
        if isprime(i) {
            count +=1
        } 
    }
    return count 
}

func isprime(m int) bool {
    if m <= 1 {
        return false 
    }
    if m <=3 {
        return true
    }
    
    if m%2 == 0 || m%3 == 0 {
        return false 
    }
    for i:=5; i*i<= m; i+=6 {
        if m%i == 0 || m%(i+2) == 0 {
            return false 
        }
    } 
    return true
}
