## 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)

## 思路

为方便后面翻译成递推，这里从右往左递归。

定义 $\textit{dfs}(i,j)$ 表示以 $\textit{nums}_j[i]$ 结尾的最长非递减子数组的长度。

用「枚举选哪个」来思考：

- 如果 $\textit{nums}_1[i-1]\le \textit{nums}_j[i]$，那么下一步选 $\textit{nums}_1[i-1]$，有 $\textit{dfs}(i,j) = \textit{dfs}(i-1,0)+1$。
- 如果 $\textit{nums}_2[i-1]\le \textit{nums}_j[i]$，那么下一步选 $\textit{nums}_2[i-1]$，有 $\textit{dfs}(i,j) = \textit{dfs}(i-1,1)+1$。
- 如果都不成立，那么 $\textit{dfs}(i,j)=1$。

这几种情况取最大值，即为 $\textit{dfs}(i,j)$。

递归边界：$\textit{dfs}(0)=1$。

递归入口：$\textit{dfs}(i,j)$。遍历所有 $i,j$ 取 $\textit{dfs}(i,j)$ 的最大值，即为答案。

[视频讲解](https://www.bilibili.com/video/BV1XW4y1f7Wv/) 第三题。

```py [sol-Python3]
class Solution:
    def maxNonDecreasingLength(self, nums1: List[int], nums2: List[int]) -> int:
        nums = (nums1, nums2)
        @cache
        def dfs(i: int, j: int) -> int:
            if i == 0: return 1
            res = 1
            if nums1[i - 1] <= nums[j][i]:
                res = dfs(i - 1, 0) + 1
            if nums2[i - 1] <= nums[j][i]:
                res = max(res, dfs(i - 1, 1) + 1)
            return res
        return max(dfs(i, j) for j in range(2) for i in range(len(nums1)))
```

```go [sol-Go]
func maxNonDecreasingLength(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	nums := [2][]int{nums1, nums2}
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 1
		if nums1[i-1] <= nums[j][i] {
			res = dfs(i-1, 0) + 1
		}
		if nums2[i-1] <= nums[j][i] {
			res = max(res, dfs(i-1, 1)+1)
		}
		*p = res // 记忆化
		return res
	}
	for j := 0; j < 2; j++ {
		for i := 0; i < n; i++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}
```

然后按照 [视频](https://www.bilibili.com/video/BV1Xj411K7oF/) 中讲的，1:1 翻译成递推。

```py [sol-Python3]
class Solution:
    def maxNonDecreasingLength(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        nums = (nums1, nums2)
        f = [[1, 1] for _ in range(n)]
        for i in range(1, n):
            for j in range(2):
                if nums1[i - 1] <= nums[j][i]:
                    f[i][j] = f[i - 1][0] + 1
                if nums2[i - 1] <= nums[j][i]:
                    f[i][j] = max(f[i][j], f[i - 1][1] + 1)
        return max(map(max, f))
```

```go [sol-Go]
func maxNonDecreasingLength(nums1, nums2 []int) int {
	ans, n := 1, len(nums1)
	nums := [2][]int{nums1, nums2}
	f := make([][2]int, n)
	f[0] = [2]int{1, 1}
	for i := 1; i < n; i++ {
		f[i] = [2]int{1, 1}
		for j := 0; j < 2; j++ {
			if nums1[i-1] <= nums[j][i] {
				f[i][j] = f[i-1][0] + 1
			}
			if nums2[i-1] <= nums[j][i] {
				f[i][j] = max(f[i][j], f[i-1][1]+1)
			}
		}
		ans = max(ans, max(f[i][0], f[i][1]))
	}
	return ans
}
```

由于 $f[i]$ 只用到 $f[i-1]$，所以可以去掉第一个维度。

```py [sol-Python3]
class Solution:
    def maxNonDecreasingLength(self, nums1: List[int], nums2: List[int]) -> int:
        ans = f0 = f1 = 1
        for (x0, y0), (x1, y1) in pairwise(zip(nums1, nums2)):
            f = g = 1
            if x0 <= x1: f = f0 + 1
            if y0 <= x1: f = max(f, f1 + 1)
            if x0 <= y1: g = f0 + 1
            if y0 <= y1: g = max(g, f1 + 1)
            f0, f1 = f, g
            ans = max(ans, f0, f1)
        return ans
```

```go [sol-Go]
func maxNonDecreasingLength(nums1, nums2 []int) int {
	ans, n := 1, len(nums1)
	f0, f1 := 1, 1
	for i := 1; i < n; i++ {
		f, g := 1, 1
		if nums1[i-1] <= nums1[i] {
			f = f0 + 1
		}
		if nums2[i-1] <= nums1[i] {
			f = max(f, f1+1)
		}
		if nums1[i-1] <= nums2[i] {
			g = f0 + 1
		}
		if nums2[i-1] <= nums2[i] {
			g = max(g, f1+1)
		}
		f0, f1 = f, g
		ans = max(ans, max(f0, f1))
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 思考题

如果把「子数组」改成「子序列」呢？

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
