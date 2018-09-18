import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	image_info := make([][]*info, len(image))

	for rowid, rowval := range image {
		image_info[rowid] = make([]*info, len(image[0]))
		for colid, colval := range rowval {
			fmt.Printf("%+v image_info %v %v %v rowval colval%v \n", image_info, rowid, colid, rowval, colval)
			image_info[rowid][colid] = &info{
				sr:      rowid,
				sc:      colid,
				color:   colval,
				visited: false,
			}
		}
	}

	var q *queue = &queue{}
	q.Enqueue(image_info[sr][sc])

	for !q.IsEmpty() {
		v := q.Dequeue()
		if v.sr > 0 && image_info[v.sr-1][v.sc].color == v.color && image_info[v.sr-1][v.sc].visited != true {
			q.Enqueue(image_info[v.sr-1][v.sc])
		}
		if v.sr < len(image_info)-1 && image_info[v.sr+1][v.sc].color == v.color && image_info[v.sr+1][v.sc].visited != true {
			q.Enqueue(image_info[v.sr+1][v.sc])
		}
		if v.sc > 0 && image_info[v.sr][v.sc-1].color == v.color && image_info[v.sr][v.sc-1].visited != true {
			q.Enqueue(image_info[v.sr][v.sc-1])
		}
		if v.sc < len(image_info[0])-1 && image_info[v.sr][v.sc+1].color == v.color && image_info[v.sr][v.sc+1].visited != true {
			q.Enqueue(image_info[v.sr][v.sc+1])
		}
		v.visited = true
		image[v.sr][v.sc] = newColor
	}
	return image
}

type info struct {
	sr      int
	sc      int
	color   int
	visited bool
}

type queue []*info

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Enqueue(val *info) {
	if *q == nil {
		return
	}
	*q = append((*q), val)
}

func (q *queue) Dequeue() *info {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}
