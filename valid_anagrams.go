func isAnagram(s string, t string) bool {
    s1 := []byte(s)
    s2 := []byte(t)
    l1 := len(s)
    l2 := len(t)

    if l1 != l2 {
	return false
    }
    sort.Slice(s1, func (a, b int) bool {
	return s1[a] < s1[b]
    })
    sort.Slice(s2, func (a, b int) bool {
	return s2[a] < s2[b]
    })

    for i := 0; i < len(s1); i++ {
	if s1[i] == s2[i] {
            continue
	}
	return false
    }
    return true
}
