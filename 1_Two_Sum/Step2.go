func twoSum(nums []int, target int) []int {
	// 値をキー、インデックスを値として保持するマップ
	indexMap := make(map[int]int)
	for i, num := range nums {
		// ターゲットに到達するために必要な残りの数
		complement := target - num

		// マップの中に「残りの数」がすでに存在するかチェック
		if partnerIdx, isExist := indexMap[complement];isExist{
			// 見つかったら、そのインデックスと現在のインデックスを返す
			return []int{partnerIdx, i}
		}
		// 今見ている値をマップに登録
		indexMap[num] = i
	}
	return nil
}