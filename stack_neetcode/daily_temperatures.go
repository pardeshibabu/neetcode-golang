package stackneetcode

func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	r := 1
	stack = append(stack, 0)
	for r < len(temperatures) {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[r] {
			top := stack[len(stack)-1]
			res[top] = r - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, r)
		r++
	}
	for i := 0; i < len(stack); i++ {
		res[stack[i]] = 0
	}
	return res
}

// brute force TLE
// func DailyTemperatures(temperatures []int) []int {
// 	res := make([]int, len(temperatures))
// 	var stack []int
// 	l, r := 0, 1
// 	stack = append(stack, 0)
// 	for l < len(temperatures)-1 {
// 		for r < len(temperatures) && temperatures[stack[0]] >= temperatures[r] {
// 			stack = append(stack, r)
// 			r++
// 		}
// 		// fmt.Println(stack, temperatures)
// 		// fmt.Println(l, r, temperatures[stack[0]], temperatures[r])
// 		if r == len(temperatures) && len(stack) > 0 {
// 			res[l] = 0
// 			l++
// 			r = l + 1
// 			stack = stack[:0]
// 			stack = append(stack, l)
// 			continue
// 		} else if (r - l) >= 1 {
// 			if (r - l) > 1 {
// 				res[l] = r - l
// 				r = l + 1
// 			} else {
// 				res[l] = 1
// 				r++
// 			}
// 			stack = stack[:0]
// 		}
// 		l++
// 		stack = append(stack, l)
// 	}
// 	res[len(temperatures)-1] = 0
// 	return res
// }
