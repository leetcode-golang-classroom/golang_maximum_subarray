package sol

func maxSubArray(nums []int) int {
	nLen := len(nums)
	// sum represent max sum that include nums[i] end on index i
	// initial with nums[0]
	sum := nums[0]
	// maxSum represent max sum end on index i
	maxSum := nums[0]
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for end := 1; end < nLen; end++ {
		sum = max(sum+nums[end], nums[end]) // positive sum + nums[end] <= nums[end], when before negative number
		// currentMaxSum = max (previous i-1 maxSum , include num[i-1] max sum)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
