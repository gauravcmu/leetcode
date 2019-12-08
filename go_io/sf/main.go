package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DEPEND, INSTALL, REMOVE, and LIST
/*
1. DEPEND item1 item2 - depend always comes before the install for the item, say item1.
2. INSTALL item1 - install item1 and its dependencies.
3. REMOVE item1 - remove item2 and its dependencies.
4. LIST - list all installed packages.
5. END - end of input.
 */
func main() {
	dep, installed, err := buildDependency("input.txt")
	fmt.Printf("--- Main -- Dep:%v installed:%v err:%s \n", dep, installed, err)
}

// build a map with key the name of package and its dependencies as list of strings.
func buildDependency(filename string) (*map[string][]string, *map[string]bool, error) {
	dep := make(map[string][]string, 0)
	installed := make(map[string]bool, 0)

	f, err := os.Open(filename)
	if err != nil {
		return &dep, &installed, err
	}

	rdr := bufio.NewReader(f)

	scnr := bufio.NewScanner(rdr)

	for scnr.Scan() != false {
		cmd := scnr.Text()
		cmdfields := strings.Fields(cmd)
		//fmt.Printf("%v\n", cmdfields)
		if cmdfields[0] == "DEPEND" {
			d := make([]string, 0)
			for index, pkg := range cmdfields {
				if index > 1 {
					d = append(d, pkg)
				}
			}
			dep[cmdfields[1]] = d
			//fmt.Printf("result:%v\n", dep)
		} else if cmdfields[0] == "INSTALL" {
			/* */
			install(cmdfields[1], &dep, &installed)
		} else if cmdfields[0] == "LIST" {
			fmt.Printf("installed:%v\n", installed)
		} else if cmdfields[0] == "END" {
			break
		}
	}
	return &dep, &installed, err
}

func install(pkg string, dep * map[string][]string, installed * map[string]bool) {
	fmt.Printf("installing:%v\n", pkg)
	if (*installed)[pkg] == true {
		fmt.Printf("WARN: pkg already installed:%v\n", pkg)
		return
	}
	depsList := (*dep)[pkg]
	for _, p := range depsList {
		if (*installed)[p] != true {
			install(p, dep, installed)
		}
	}
	(*installed)[pkg] = true
}