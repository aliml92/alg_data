package array




func CombinationSum2(candidates []int, target int) [][]int {
	m := make(map[int]int)
	for _, v := range candidates {
		// count the number of each element
		m[v]++
	}
	res := [][][]int{}
	for i := 0; i < target + 1 ; i++ {
		res = append(res, [][]int{})
	}
	e := make(map[int]bool)
	for _, c := range candidates {
		// check if c is already used
		if e[c] {
			continue
		}
		e[c] = true
		for i := 1; i < target + 1; i++ {
			if i < c {
				continue
			}
			if i == c {
				res[i] = append(res[i], []int{c})
			} else {
				for _, blist := range res[i - c] {
					// count the number of c in the blist
					count := 0
					for _, b := range blist {
						if b == c {
							count++
						}	
					}
					if count < m[c] {
						a := append([]int{c}, blist...)
						res[i] = append(res[i], a)						
					}
				}
			}
		}
	}
	return res[target]
}

