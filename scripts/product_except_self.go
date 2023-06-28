package scripts

func ProductExceptSelf(nums []int) []int {
	var outPut []int
	var preiFix, postFix = 1, 1
	for _, v := range nums {
		outPut = append(outPut, preiFix)
		preiFix = preiFix * v
	}
	for i := len(nums) - 1; i >= 0; i-- {
		outPut[i] = outPut[i] * postFix
		postFix = postFix * nums[i]
	}
	return outPut
}
