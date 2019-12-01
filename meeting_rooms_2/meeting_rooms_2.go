
/*
Input: [[0, 30],[5, 10],[15, 20]]
Output: 2

rooms: [30, 10, ]
sort the input by start time. 
create a room when 
    - there is no room with end time less than equal to my start time. 
        - add the room by append to array and add my end time there. 
    - else if found a room - then update the room availability to end time of this meeting. 

*/ 

func minMeetingRooms(intervals [][]int) int {
    rooms := make([]int, 0)
    sort.Slice(intervals, func (a,b int)bool {
        return intervals[a][0] < intervals[b][0]
    })
    

    for _, meeting := range intervals {
            found := false 
            for index, room := range rooms {
                if room <= meeting[0] {
                    rooms[index] = meeting[1]
                    found = true 
                    break
                }      
            }
            if !found {
                rooms = append(rooms, meeting[1])
            }
    }
    
    return len(rooms)
}
