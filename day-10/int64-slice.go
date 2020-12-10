package main

// IntSlice is a sortable version of an []int64
type IntSlice []int64

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}