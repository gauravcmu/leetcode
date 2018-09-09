Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

Example 1:

Input: [[0, 30],[5, 10],[15, 20]]
Output: 2
Example 2:

Input: [[7,10],[2,4]]
Output: 1


import "sort"

func minMeetingRooms(intervals []Interval) int {
    rooms := make([]int, 0) 
    if len(intervals) == 0 {
        return 0
    }
    
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a].Start < intervals[b].Start
    })
    
    rooms = append(rooms, 0) // first room available from time = 0
    //rooms store available_at for each room 
    for _, meeting := range intervals {
        found := false 
        for room_number, room := range rooms {
            if meeting.Start >= room {
                //meeting start >= room availability
                
                //update availability to end of this meeting for this room
                rooms[room_number] = meeting.End
                found = true
                break
            } 
        }
        if found == false {
           //no room found - create one, available after this meeting 
            rooms = append(rooms, meeting.End)
        }
    } 
    return len(rooms)
}
