package bit_arithmetic

//一个数组中，只有一种数出现奇数次，另外的数出现偶数次，求出现奇数次的数？
//（要求时间复杂度为O(n)，空间复杂度为 O(1)）
func printOddTimesNum1(arr []int) int {
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	return eor
}

//一个数组中，有两种数出现奇数次，另外的数出现偶数次，求出现奇数次的数？
//（要求时间复杂度为O(n)，空间复杂度为 O(1)）
func printOddTimesNum2(arr []int) (a int, b int) {
	eor1 := 0
	for _, v := range arr {
		eor1 ^= v
	}
	// eor1 = a^b ，由于a!=b, 所以eor1!=0, 那么eor1 32 位中至少会有一位为1,
	// 假设第N为是1，基于第N位，能将数组分为两个集合,
	// a 和 b 会分别分到两个集合中
	ro := eor1 & (^eor1 + 1) // eor1 找到最右边为1的数

	// eor2 只和两个集合中的一个异或，就能得到 a|b
	eor2 := 0 // 只和ro与运算为0的数异或，则结果是a|b
	for _, v := range arr {
		if v&ro == 0 {
			eor2 ^= v
		}
	}
	return eor2, eor1 ^ eor2
}
