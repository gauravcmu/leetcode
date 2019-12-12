package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return []string{}
	}

	sort.Slice(tickets, func(a, b int) bool {
		return tickets[a][1] < tickets[b][1]
	})

	return helper(tickets)
}


func helper(tickets [][]string) []string {
	visited := make([]bool, len(tickets))
	result := make([]string, 0)

	fmt.Printf("tickets: %v\n\n", tickets)
	q := &queue{}

	for index, tkt := range tickets {
		if tkt[0] == "JFK" {
			q.enqueue(index)
			result = append(result, tickets[index][0])
			result = append(result, tickets[index][1])
			break
		}
	}

	for q.isEmpty() != true {
		u := q.dequeue()
		visited[u] = true
		for i := 0; i< len(tickets); i++ {
			if visited[i] != true && tickets[i][0] == tickets[u][1] {
				q.enqueue(i)
				fmt.Printf("append:%v\n", tickets[i])
				result = append(result, tickets[i][1])
				break
			}
		}
	}
	return result
}


type queue []int

func (q *queue) isEmpty() bool {
	return len(*q) ==0
}

func (q *queue) enqueue(v int) {
	fmt.Printf("enqueue: %v\n", v)
	*q = append(*q, v)
}

func (q *queue) dequeue() int {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}
