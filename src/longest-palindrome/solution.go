package longest_palindrome

import (
	"strings"
)

// 给你一个字符串 s，找到 s 中最长的回文子串。

// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：

// 输入：s = "cbbd"
// 输出："bb"
//

// 提示：

// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func LongestPalindrome(s string) string {

	arr := strings.Split(s,"")

	var palindrome string
	for i := 0; i < len(arr); i++ {
		// 判断回文字符串
		if i==0 {
			palindrome = arr[i]
		}
		if i==1 && arr[0]==arr[1] {
			palindrome = arr[0]+arr[1]
		}
		
	}
	l := len(arr)
	if len(arr)%2 == 0 {
		// 偶数情况
		for i := 0; i < len(arr); i++ {
			if arr[i]==arr[l-1]{
				
			}
		}
	}else {
		// 奇数情况
	}
	return palindrome
}
