func mostCommonWord(paragraph string, banned []string) string {
	p := []string {"!", "?", "'", ",", ";", "."}
   	for _, v := range p {
		paragraph = strings.Replace(paragraph, v, "", -1)
    }
    mapping := make(map[string]int, 0)

    s2 := strings.Fields(paragraph)
	//banned words hash map
	bmap := make(map[string]int, len(banned))
	for _, v := range banned {
		if _, ok := bmap[v]; ok {
			continue
        }
        bmap[v] = 1
    }

    //s2 has all punctuations removed words.
    //maximize using max
    var max int 
    res := ""
    for _, u := range s2 {
        v := strings.ToLower(u)
        if bmap[v] != 0 {
            //banned word. ignore
            continue
        }
        if _, ok := mapping[v]; ok {
            mapping[v] += 1 
        } else {
            mapping[v] = 1
        }
        if mapping[v] > max {
            max = mapping[v]
            res = v
        }
    }
    return res 
}
