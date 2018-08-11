func IsNeighbor(a, b string) bool {
   //     fmt.Printf("IsNeighbors: %v, %v\n", a, b)
    count := 0 
    for _, value := range a {
        for _, value2 := range b {
            if value == value2 {
                count++
            } 
        }
    }  
    if count == (len(a) -1)  {
        fmt.Printf("neighbors: %v, %v\n", a, b)
        return true
    }
    return false 
}

type wordInfo struct {
    value string
    visited bool
    path string
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    Q := &queue{ }
    
    //maintain hash for visited. 
    wordmap := make(map[string]wordInfo, len(wordList))
    for _, value := range wordList {
        wordmap[value]= wordInfo {
            visited: false, 
            value: value, 
            path: "",   
        }
    }
    wordmap[beginWord] = wordInfo {
            visited: false, 
            value: beginWord, 
            path: "", 
        }
    
    //insert the begin word to queue
    add(Q, wordInfo {
            visited: false, 
            value: beginWord, 
            path: "", 
        }) 
    
    for (!isEmpty(Q)) {
        // dequeue from queue. 
        // if node already visited - continue
        // else 
        // for this node v, find all children that are neighbor and not visited and queue
        v := remove(Q)
        if v.visited == false {
            //Node not visited before
            v.visited = true
            wordmap[v.value] = v
        } else {
            //node already visited, move on
            fmt.Printf(" here 2, v:%v\n", v)
            continue
        }
        
        if v.value == endWord {
            //stop - word reached
            w := v
            count :=0
            for (w.value!= beginWord) {
                count++
                fmt.Printf("w:%v\t", w)
                w = wordmap[w.path] 
            }
            return count  
        } 
        
        for _, u := range wordmap {
            //each word that is neighbor (1 edit distance) - enqueue
            if (IsNeighbor(v.value, u.value) == true)  {
                if u.visited != true {
                    u.path = v.value 
                    
                    fmt.Printf("u:%v\t", u)
                    add(Q, u)         
                }
            } 
        } 
    }
    return 0  
}

type queue struct {
    q []wordInfo 
}

func isEmpty(q *queue) bool {
    //fmt.Printf("Length of queue:%v\n", len(q.q))
    return len(q.q) == 0 
}

func add(q *queue, value wordInfo) {
    q.q = append(q.q, value) 
} 
    
func remove(q * queue) wordInfo {
    temp := q.q[len(q.q)-1]
    q.q = q.q[0:len(q.q)-1] 
    return temp
}
