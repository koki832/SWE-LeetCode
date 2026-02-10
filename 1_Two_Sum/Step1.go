func twoSum(nums []int, target int) []int {
    //　先頭から全ての要素を順番に見る
	for i := 0; i < len(nums)-1; i++ {
		secondValue := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == secondValue {
				return []int{i, j}
			}
		}
	}
	return nil
}