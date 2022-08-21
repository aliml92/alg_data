package array



func CombinationSum(candidates []int, target int) [][]int {
	res := [][][]int{}
	for i := 0; i < target + 1 ; i++ {
		res = append(res, [][]int{})
	}
	for _, c := range candidates {
		for i := 1; i < target + 1; i++ {
			if i < c {
				continue
			}
			if i == c {
				res[i] = append(res[i], []int{c})
			} else {
				for _, blist := range res[i - c] {
					a := append([]int{c}, blist...)
					res[i] = append(res[i], a)
				}
			}
		}
	}
	return res[target]
}


