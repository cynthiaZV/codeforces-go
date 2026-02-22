在每个下标 $i$ 处，我们有 $3$ 种操作，所以一共有 $3^n\le 3^{19} = 1162261467$ 个不同的操作组合，这太大了，直接暴力搜索会超时。

注意本题 $\textit{nums}[i]\le 6$，如果对 $\textit{nums}[i]$ 做质因数分解，分解出的 $2,3,5$ 的个数均为 $\mathcal{O}(1)$，所以计算乘除后，最终结果的质因数分解中的 $2,3,5$ 的幂次是 $\mathcal{O}(n)$ 的。所以这 $3^n$ 个不同的操作，只会得到 $\mathcal{O}(n^3)$ 个不同的结果。本题涉及到分数，这里把分母的质因子的幂次视作负数，例如 $\dfrac{1}{8} = 2^{-3}$，$\dfrac{25}{6} = 2^{-1}3^{-1}5^2$。

考虑记忆化搜索。我们从 $k$ 和 $n-1$ 开始，倒着乘除，目标是得到 $1$，即质因数分解中的 $2,3,5$ 的幂次均为 $0$。

$\textit{dfs}$ 需要 $4$ 个参数 $i,e_2,e_3,e_5$，分别表示当前要考虑下标 $i$ 怎么操作，以及当前 $k$ 的质因数分解中的 $2,3,5$ 的幂次。

设 $\textit{nums}[i]$ 的质因数分解中有 $x,y,z$ 个 $2,3,5$，那么有如下三种情况：

- 把 $k$ 除以 $\textit{nums}[i]$，$e_2,e_3,e_5$ 分别减少了 $x,y,z$。
- 把 $k$ 乘以 $\textit{nums}[i]$，$e_2,e_3,e_5$ 分别增加了 $x,y,z$。
- 不变。

三种情况继续递归，方案数累加，即为 $\textit{dfs}$ 的返回值。

**递归边界**：当 $i<0$ 时，如果 $e_2 = e_3 = e_5 = 0$，说明此时 $k=1$，我们找到了一个合法的操作组合，返回 $1$；否则返回 $0$。

**递归入口**：$\textit{dfs}(n-1,E_2,E_3,E_5)$，其中 $E_2,E_3,E_5$ 分别是初始 $k$ 的质因数分解中的 $2,3,5$ 的幂次。

**特殊情况**：如果 $k$ 包含大于 $5$ 的质因子，无解，返回 $0$。

[本题视频讲解](https://www.bilibili.com/video/BV1HeZfB7EBt/?t=22m54s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    # 返回 k 中的质因子 2,3,5 的个数，以及 k 是否只包含 <= 5 的质因子
    def primeFactorization(self, k: int) -> Tuple[Tuple[int, int, int], bool]:
        e2 = (k & -k).bit_length() - 1
        k >>= e2

        e3 = 0
        while k % 3 == 0:
            e3 += 1
            k //= 3

        e5 = 0
        while k % 5 == 0:
            e5 += 1
            k //= 5

        return (e2, e3, e5), k == 1

    def countSequences(self, nums: List[int], k: int) -> int:
        (e2, e3, e5), ok = self.primeFactorization(int(k))
        if not ok:  # k 有大于 5 的质因子
            return 0

        es = [self.primeFactorization(x)[0] for x in nums]

        @cache
        def dfs(i: int, e2: int, e3: int, e5: int) -> int:
            if i < 0:
                return 1 if e2 == e3 == e5 == 0 else 0

            x, y, z = es[i]
            res1 = dfs(i - 1, e2 - x, e3 - y, e5 - z)  # k 除以 nums[i]
            res2 = dfs(i - 1, e2 + x, e3 + y, e5 + z)  # k 乘以 nums[i]
            res3 = dfs(i - 1, e2, e3, e5)  # k 不变
            return res1 + res2 + res3

        return dfs(len(nums) - 1, e2, e3, e5)  # 从 k 开始，目标是变成 1
```

```java [sol-Java]
class Solution {
    private record Exp(int e2, int e3, int e5) {
    }

    // 返回 k 中的质因子 2,3,5 的个数，以及 k 是否只包含 <= 5 的质因子
    private Pair<Exp, Boolean> primeFactorization(long k) {
        int e2 = Long.numberOfTrailingZeros(k);
        k >>= e2;

        int e3 = 0;
        while (k % 3 == 0) {
            e3++;
            k /= 3;
        }

        int e5 = 0;
        while (k % 5 == 0) {
            e5++;
            k /= 5;
        }

        return new Pair<>(new Exp(e2, e3, e5), k == 1);
    }

    public int countSequences(int[] nums, long k) {
        Pair<Exp, Boolean> res = primeFactorization(k);
        Exp e = res.getKey();
        boolean ok = res.getValue();
        if (!ok) { // k 有大于 5 的质因子
            return 0;
        }

        int n = nums.length;
        Exp[] es = new Exp[n];
        for (int i = 0; i < n; i++) {
            es[i] = primeFactorization(nums[i]).getKey();
        }

        Map<Integer, Integer> memo = new HashMap<>();
        return dfs(n - 1, e.e2, e.e3, e.e5, es, memo);
    }

    private int dfs(int i, int e2, int e3, int e5, Exp[] es, Map<Integer, Integer> memo) {
        if (i < 0) {
            return e2 == 0 && e3 == 0 && e5 == 0 ? 1 : 0; // k 变成 1
        }

        // 把 i,e2,e3,e5 拼成一个 int（每个数至多 6 位）
        int n = es.length;
        int key = i << 18 | (e2 + n * 2) << 12 | (e3 + n) << 6 | (e5 + n);
        if (memo.containsKey(key)) {
            return memo.get(key);
        }

        Exp e = es[i];
        int x = e.e2, y = e.e3, z = e.e5;
        int res1 = dfs(i - 1, e2 - x, e3 - y, e5 - z, es, memo); // k 除以 nums[i]
        int res2 = dfs(i - 1, e2 + x, e3 + y, e5 + z, es, memo); // k 乘以 nums[i]
        int res3 = dfs(i - 1, e2, e3, e5, es, memo); // k 不变
        int res = res1 + res2 + res3;

        memo.put(key, res);
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回 k 中的质因子 2,3,5 的个数，以及 k 是否只包含 <= 5 的质因子
    pair<tuple<int, int, int>, bool> primeFactorization(long long k) {
        int e2 = countr_zero((uint64_t) k);
        k >>= e2;

        int e3 = 0;
        while (k % 3 == 0) {
            e3++;
            k /= 3;
        }

        int e5 = 0;
        while (k % 5 == 0) {
            e5++;
            k /= 5;
        }

        return {tuple(e2, e3, e5), k == 1};
    }

public:
    int countSequences(vector<int>& nums, long long k) {
        auto [e, ok] = primeFactorization(k);
        if (!ok) { // k 有大于 5 的质因子
            return 0;
        }

        int n = nums.size();
        vector<tuple<int, int, int>> es(n);
        for (int i = 0; i < n; i++) {
            es[i] = primeFactorization(nums[i]).first;
        }

        unordered_map<int, int> memo;

        auto dfs = [&](this auto&& dfs, int i, int e2, int e3, int e5) -> int {
            if (i < 0) {
                return e2 == 0 && e3 == 0 && e5 == 0; // k 变成 1
            }

            // 把 i,e2,e3,e5 拼成一个 int（每个数至多 6 位）
            int key = i << 18 | (e2 + n * 2) << 12 | (e3 + n) << 6 | (e5 + n);
            auto it = memo.find(key);
            if (it != memo.end()) {
                return it->second;
            }

            auto [x, y, z] = es[i];
            int res1 = dfs(i - 1, e2 - x, e3 - y, e5 - z); // k 除以 nums[i]
            int res2 = dfs(i - 1, e2 + x, e3 + y, e5 + z); // k 乘以 nums[i]
            int res3 = dfs(i - 1, e2, e3, e5); // k 不变
            int res = res1 + res2 + res3;

            memo[key] = res;
            return res;
        };

        auto [e2, e3, e5] = e;
        return dfs(n - 1, e2, e3, e5); // 从 k 开始，目标是变成 1
    }
};
```

```go [sol-Go]
// 返回 k 中的质因子 2,3,5 的个数，以及 k 是否只包含 <= 5 的质因子
func primeFactorization(k int) ([3]int, bool) {
	e2 := bits.TrailingZeros(uint(k))
	k >>= e2

	e3 := 0
	for k%3 == 0 {
		e3++
		k /= 3
	}

	e5 := 0
	for k%5 == 0 {
		e5++
		k /= 5
	}

	return [3]int{e2, e3, e5}, k == 1
}

func countSequences(nums []int, k int64) int {
	e, ok := primeFactorization(int(k))
	if !ok { // k 有大于 5 的质因子
		return 0
	}

	n := len(nums)
	es := make([][3]int, n)
	for i, x := range nums {
		es[i], _ = primeFactorization(x)
	}

	type args struct{ i, e2, e3, e5 int }
	memo := map[args]int{}
	var dfs func(int, int, int, int) int
	dfs = func(i, e2, e3, e5 int) int {
		if i < 0 {
			if e2 == 0 && e3 == 0 && e5 == 0 { // k 变成 1
				return 1
			}
			return 0
		}
		p := args{i, e2, e3, e5}
		if res, ok := memo[p]; ok {
			return res
		}

		e := es[i]
		res1 := dfs(i-1, e2-e[0], e3-e[1], e5-e[2]) // k 除以 nums[i]
		res2 := dfs(i-1, e2+e[0], e3+e[1], e5+e[2]) // k 乘以 nums[i]
		res3 := dfs(i-1, e2, e3, e5)                // k 不变
		res := res1 + res2 + res3

		memo[p] = res
		return res
	}
	return dfs(n-1, e[0], e[1], e[2]) // 从 k 开始，目标是变成 1
}
```

#### 复杂度分析

状态个数：由于 $\textit{nums}[i]\le 6$，分解出的质因子 $2,3,5$ 的个数均为 $\mathcal{O}(1)$，所以至多有 $\mathcal{O}(n)$ 个不同的 $i,e_2,e_3,e_5$，所以有 $\mathcal{O}(n^4)$ 个状态。

- 时间复杂度：$\mathcal{O}(n^4 + \log k)$，其中 $n$ 是 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^4)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(n^4)$。分解 $k$ 需要 $\mathcal{O}(\log k)$ 的时间。
- 空间复杂度：$\mathcal{O}(n^4)$。保存多少状态，就需要多少空间。

## 写法二

也可以用**最简分数** $\dfrac{p}{q}$ 表示 $\textit{val}$。从 $\dfrac{1}{1}$ 开始，通过若干操作，得到 $\dfrac{k}{1}$。

这种写法更简单，缺点是需要在 $\textit{dfs}$ 的过程中计算 $\gcd$，效率不如预处理质因数分解的做法。

```py [sol-Python3]
class Solution:
    def countSequences(self, nums: List[int], k: int) -> int:
        @cache
        def dfs(i: int, p: int, q: int) -> int:
            if i < 0:
                return 1 if p == k and q == 1 else 0

            x = nums[i]
            g = gcd(p, q * x)
            res1 = dfs(i - 1, p // g, q * x // g)  # 除以 nums[i]
            g = gcd(p * x, q)
            res2 = dfs(i - 1, p * x // g, q // g)  # 乘以 nums[i]
            res3 = dfs(i - 1, p, q)  # 不变
            return res1 + res2 + res3

        return dfs(len(nums) - 1, 1, 1)  # 从 1/1 开始，目标是变成 k/1
```

```go [sol-Go]
func countSequences(nums []int, k int64) int {
	type args struct{ i, p, q int }
	memo := map[args]int{}
	var dfs func(int, int, int) int
	dfs = func(i, p, q int) int {
		if i < 0 {
			if p == int(k) && q == 1 {
				return 1
			}
			return 0
		}
		t := args{i, p, q}
		if res, ok := memo[t]; ok {
			return res
		}

		x := nums[i]
		g := gcd(p, q*x)
		res1 := dfs(i-1, p/g, q*x/g) // 除以 nums[i]
		g = gcd(p*x, q)
		res2 := dfs(i-1, p*x/g, q/g) // 乘以 nums[i]
		res3 := dfs(i-1, p, q)       // 不变
		res := res1 + res2 + res3

		memo[t] = res
		return res
	}
	return dfs(len(nums)-1, 1, 1) // 从 1/1 开始，目标是变成 k/1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

**注**：本题也可以用折半搜索做，时间复杂度 $\mathcal{O}(3^{n/2})$。

见下面回溯题单的「**§4.8 折半搜索**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
