package hot2nd

func canJump(nums []int) bool {
	n, reach := len(nums), 0
	for i := 0; i < n; i++ {
		if i > reach || reach >= n-1 {
			break
		}
		reach = max(reach, nums[i]+i)
	}
	return reach >= n-1
}
