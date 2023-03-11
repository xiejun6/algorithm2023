package main

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	//fmt.Println("after move:", nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	nums = []int{0}
	moveZeroes(nums)
}
