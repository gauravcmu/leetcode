package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type userinfo struct {
	name  string
	spent int
}

func readInput() (map[string]int, error) {
	//use a scanner to scan the input line by line
	scanner := bufio.NewScanner(os.Stdin)

	//create a userinfo map which holds {name: total expense} records
	mapping := make(map[string]int, 0)

	for scanner.Scan() == true {
		line := strings.Split(scanner.Text(), "$")
		//        fmt.Printf("line - %v\n", line)

		//val - expense value for this entry for this user
		val, err := strconv.Atoi(line[1])
		if err != nil {
			return mapping, err
		}

		if _, ok := mapping[line[0]]; !ok {
			//userinfo not in map - create and add first expense
			mapping[line[0]] = val
		} else {
			//add expense for this user as this user already has entry in the mapping
			mapping[line[0]] += val
		}
	}
	return mapping, nil
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	mapping, err := readInput()
	if err != nil {
		return
	}

	//mapping has user and expenses - go over mapping and create slice of userinfo
	users := make([]userinfo, 0)

	for key, value := range mapping {
		users = append(users, userinfo{name: key, spent: value})
	}

	//sort this slice by expenses only when expenses (spent value) is not same, else sort by name.
	// func Less() should be working as more when comparing spent values but not names. (more spent comes first than less...)
	sort.Slice(users, func(i int, j int) bool {
		if users[i].spent == users[j].spent {
			return users[i].name < users[j].name
		} else {
			return users[i].spent > users[j].spent

		}
	})

	//print output as expected - username $spent
	for _, user := range users {
		fmt.Printf("%v $%v\n", user.name, user.spent)
	}
}
