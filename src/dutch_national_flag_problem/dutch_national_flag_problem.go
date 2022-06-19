package dutchnationalflagproblem

// 荷兰国旗问题：
// 给定一个数组arr，和一个数num，把小于等于num的数放到数组左边，大于num的数放到右边，时间复杂度为O(N)
func dutchNationalFlagProblem1(arr []int, num int) []int {
	if len(arr) < 2 {
		return arr
	}
	ltIdx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= num {
			arr[i], arr[ltIdx] = arr[ltIdx], arr[i]
			ltIdx++
		}
	}
	return arr
}

// 给定一个数组arr，和一个数num，把小于num的数放到数组左边，等于num的数放到数组中间，大于num的数放到数组右边，要求时间复杂度为O(N)
func dutchNationalFlagProblem2(arr []int, num int) []int {
	if len(arr) < 2 {
		return arr
	}
	ltIdx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= num {
			arr[i], arr[ltIdx] = arr[ltIdx], arr[i]
			ltIdx++
		}
	}

	// 第二次遍历，直接按照第一种方法，再进行一次遍历，找到等于和小于的区间
	eqIdx := 0
	for i := 0; i < ltIdx; i++ {
		if arr[i] < num {
			arr[i], arr[eqIdx] = arr[eqIdx], arr[i]
			eqIdx++
		}
	}

	return arr
}

// 给定一个数组arr，和一个数num，把小于num的数放到数组左边，等于num的数放到数组中间，大于num的数放到数组右边，要求时间复杂度为O(N)
func dutchNationalFlagProblem3(arr []int, num int) []int {
	if len(arr) < 2 {
		return arr
	}
	ltIdx := 0
	mtIdx := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if i == mtIdx {
			// i=mtIdx 说明所有的数据已经比较完了
			break
		}
		if arr[i] < num {
			arr[i], arr[ltIdx] = arr[ltIdx], arr[i]
			ltIdx++
		} else if arr[i] > num {
			arr[i], arr[mtIdx] = arr[mtIdx], arr[i]
			mtIdx--
			i--
		}
	}
	return arr
}
