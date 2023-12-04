package utils

func Intersection(a []int, b []int) []int {
	intersection := make([]int, 0)

	set := make(map[int]bool)

	for _, num := range a {
		set[num] = true
	}

	for _, num := range b {
		if set[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection
}
