**前置知识**：[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

对于本题，新加入元素 $x=\textit{nums}[\textit{right}]$ 后，如果 $x$ 的出现次数超过 $k$，则不断右移左指针 $\textit{left}$，直到窗口内的 $x$ 的出现次数等于 $k$ 为止，然后用窗口大小 $\textit{right}-\textit{left}+1$ 更新答案的最大值。

```py [sol-Python3]
class Solution:
    def maxSubarrayLength(self, nums: List[int], k: int) -> int:
        ans = left = 0
        cnt = defaultdict(int)
        for right, x in enumerate(nums):
            cnt[x] += 1
            while cnt[x] > k:
                cnt[nums[left]] -= 1
                left += 1
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxSubarrayLength(int[] nums, int k) {
        int ans = 0;
        int left = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int right = 0; right < nums.length; right++) {
            cnt.merge(nums[right], 1, Integer::sum); // cnt[nums[right]]++
            while (cnt.get(nums[right]) > k) {
                cnt.merge(nums[left], -1, Integer::sum); // cnt[nums[left]]--
                left++;
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSubarrayLength(vector<int>& nums, int k) {
        int ans = 0, left = 0;
        unordered_map<int, int> cnt;
        for (int right = 0; right < nums.size(); right++) {
            cnt[nums[right]]++;
            while (cnt[nums[right]] > k) {
                cnt[nums[left]]--;
                left++;
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, x := range nums {
		cnt[x]++
		for cnt[x] > k {
			cnt[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
