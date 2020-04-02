package main

//给你一个整数数组 nums，请你将该数组升序排列。
//
// 
//
//示例 1：
//
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//示例 2：
//
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
// 
//
//提示：
//
//1 <= nums.length <= 50000
//-50000 <= nums[i] <= 50000

func sortArray(nums []int) []int { //插入排序
	length := len(nums)
	for i := 1;i < length;i++{
		tmp := nums[i]
		j := i
		for ;j >= 0 && nums[j-1] > tmp;j--{
			nums[j] = nums[j-1]
		}
		nums[j] = tmp
	}
	return nums
}


//func main() {
//	nums := []int{5,2,3,1}
//	_ = sortArray(nums)
//}