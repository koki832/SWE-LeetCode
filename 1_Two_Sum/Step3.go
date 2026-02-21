func twoSum(nums []int, target int) []int {
    // マップを作成
    indexMap := make(map[int]int)

    for i,num := range nums {
        complement := target - num
        
        if partnerIdx,isExist := indexMap[complement] ;isExist{
			return []int{partnerIdx, i}
        }
        indexMap[num]=i
    }
    return nil
}