package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	recStack := make([]bool, numCourses)

	canFinishRecursive(numCourses, c[0])
}

//2, [[1,0]]
func canFinishRecursive(numCourses int, course int, visited []bool, recStack []bool, prerequisites [][]int) bool {
	if visited[course] == true {
		return false
	}
	if recStack[course] == true {
		return false
	}

	visited[course] = true
	recStack[course] = true
	temp := false
	for _, c := range prerequisites {
		// [course, depends on]
		if c[0] == course {
			if visited[c[1]] == true {
				continue
			} else if recStack[c[1]] == true {
				return false
			} else {
				temp = canFinishRecursive(numCourses, c[1], visited, recStack, prerequisites)
			}
		}
	}
	recStack[course] = false
	return temp
}