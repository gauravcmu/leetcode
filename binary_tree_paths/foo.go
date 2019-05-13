package main

import (
	"fmt"
)

func main() {
	s := reports(map[string]string{
		"A": "C", "B": "C", "C": "F", "D": "E", "E": "F", "F": "F"})
	fmt.Printf("S:%v\n", s)

}

func reports(input map[string]string) map[string][]string {
	mapping := make(map[string][]string, 0)
	head := ""
	result := make(map[string][]string, 0)

	for emp, mgr := range input {
		fmt.Printf("here:%v\n", emp)
		if _, ok := mapping[mgr]; ok && emp != mgr {

			mapping[mgr] = append(mapping[mgr], emp)
		} else if emp != mgr {
			mapping[mgr] = make([]string, 0)
			mapping[mgr] = append(mapping[mgr], emp)
		}
		if emp == mgr {
			head = emp
		}
	}

	fmt.Printf("mapping:%v head:%v\n", mapping, head)
	helper(head, mapping, &result)
	return result
}

func helper(head string, mapping map[string][]string, result *map[string][]string) {
	fmt.Printf("head:%v\n", head)
	if emps, ok := mapping[head]; !ok {
		//reached an employee who is at lowest level.
		return
	} else {
		for _, emp := range emps {
			fmt.Printf("here:%v\n", emp)

			(*result)[head] = append((*result)[head], emp)
			helper(emp, mapping, result)
			(*result)[head] = append((*result)[head], (*result)[emp]...)
		}
		fmt.Printf("result: %v\n", *result)
	}
}
