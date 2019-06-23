package array

// [1,2,3,4]
//  1		1		1*2 1*2*3
//2*3*4	3*4 4		1
// 24   12  8   6

func ProductExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			res[i] = 1
		} else {
			res[i] = res[i-1] * nums[i-1]
		}
	}

	p := 1
	for i := l; i > 0; i-- {
		if i == l {
			res[i-1] *= p
		} else {
			p *= nums[i]
			res[i-1] *= p
		}
	}

	return res

}
