# golang_maximum_subarray

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return *its sum*.

A **subarray** is a **contiguous** part of an array.

## Examples

**Example 1:**

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

```

**Example 2:**

```
Input: nums = [1]
Output: 1

```

**Example 3:**

```
Input: nums = [5,4,-1,7,8]
Output: 23

```

**Constraints:**

- `1 <= nums.length <= 105`
- `104 <= nums[i] <= 104`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer**

## 解析

給一個整數陣列 nums

要求寫一個演算法找出一個具有子陣列所具有得最大的合

子陣列必須每個相鄰元素都必須在原本陣列 nums 位置是連續的

定義 dp[i] = 從 0 到 i 為止是包含 nums[i] 所可能行成的最大和

如果陣列所有值都是正數

則最大合就是把所有值做相加

dp[i] = dp[i-1] + nums[i]

如果陣列具都是負數

因為遇加會愈小, 所以要包含 nums[i] 的最大和會是 nums[i] 本身

dp[i] = nums[i] 

如果陣列可能有負值

則包含 nums[i] 的最大和

dp[i] = max(dp[i-1] + nums[i], nums[i])

然而整體的最大和可能包含 nums[i] 或是不包含 nums[i]

所以需每次需要把包含 nums[i]的最大和或是不包含 nums[i] 的最大和做比較

maxSum 初始化為 nums[0] 因為至少要包含一個元素

算完 dp[i] 包含 nums[i] 可能的最大和

在來利用 maxSum 這個前 i-1 個可能形成的最大和

再來算最大和 maxSum = max(dp[i], maxSum) 

![](https://i.imgur.com/PjUhgXg.png)

## 程式碼
```go
package sol

func maxSubArray(nums []int) int {
  nLen := len(nums)
  sum := nums[0]
  maxSum := nums[0]
  var max = func(a, b int) int {
    if a > b {
      return a
    }
    return b
  }
  for end := 1; end < nLen; end++ {
    sum = max(sum + nums[end], nums[end])
    maxSum = max(maxSum, sum)
  }
  return maxSum
}

```
## 困難點

1. 需要理解每個位置與最大和個關係
2. 每個子陣列的最大和之間具有遞迴關係

## Solve Point

- [x]  初始化最大和 maxSum 為 nums[0], sum 為 nums[0] 因為至少要包含一個元素
- [x]  其中 maxSum 為到位置 i 為止最大和, sum 為到位置 i 為止包含 nums[i]最大和
- [x]  從 i =2 開始逐步計算  sum = max(sum + nums[i], nums[i]) 還有 maxSum = max(maxSum, sum)
- [x]  最後結果就是 maxSum