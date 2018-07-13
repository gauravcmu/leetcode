var paran [][]string =  [][]string{ {"{", "}"}, {"(", ")"}, {"[", "]"}}

func isOpenParan(value string) bool {
    for _, paran := range paran {
        if value  == paran[0] {
            return true
        }
    }
    return false 
}

func isCloseParan(value string) bool {
    for _, paran := range paran {
        if value  == paran[1] {
            return true
        }
    }
    return false 
}

func findparan (open, cls string) bool {
    for _, paran := range paran {
        if (open  == paran[0])  && (cls == paran[1]){
            return true
        }
    }
    return false 
}

func isValid(s string) bool {
    stack := make([]string, 0)    
    
    for _, value := range s {
        if isOpenParan(string(value)) {
            push(&stack, string(value))
        }
        if isCloseParan(string(value)) {
            //closing paran 
            //Find stack top in paran array and match corresponding closing and check if that is same as this. 
            if (!notEmpty(&stack) || !findparan(peek(&stack), string(value))){
                return false
            } else {
                pop(&stack)
            }  
        }
    } 
    
    return notEmpty(&stack) == false 
}

func notEmpty (stack * []string) bool {
    return len(*stack) != 0 
}

func pop(stack *[]string) string {
    temp := (*stack)[0]
    *stack = (*stack)[1:] 
    return temp 
}

func peek(stack *[]string) string {
    return (*stack)[0]
}

func push(stack *[]string, value string) {
    *stack = append([]string{value}, *stack...) 
}
