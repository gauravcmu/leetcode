/*

Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.
n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/


import "strconv"
func fizzBuzz(n int) []string {
    res := make([]string, 0)
    for i:=1; i<=n; i++ {
        res = append(res, fb(i))
    }
    return res
}

func fb(n int) string {
    if n % 3 == 0 {
        if n%5 == 0 {
            return "FizzBuzz"
        }
        return "Fizz"
    } 
    
    if n%5 == 0 {
        return "Buzz"
    }
    return strconv.Itoa(n)
}
