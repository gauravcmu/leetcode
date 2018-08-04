func isPowerOfThree(n int) bool {
    return pwOfThree(n)
}
func pwOfThree(n int) bool {
    if n == 3 || n == 1{
        return true
    }
    if n == 0 {
        return false
    }
    if n % 3 != 0 {
        return false 
    }
    return pwOfThree(n/3)
}
