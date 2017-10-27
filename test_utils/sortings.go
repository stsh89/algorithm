package test_utils

var SortingsTestSet = []struct {
	In  []int
	Out []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7}},
	{[]int{3, 7, 8, 5, 2, 1, -9}, []int{-9, 1, 2, 3, 5, 7, 8}},
	{[]int{3, 7, 3, 5, 2, 1, -9}, []int{-9, 1, 2, 3, 3, 5, 7}},
}
