根据题意，我们可以遍历每个格子，以这个格子为起点，往上下左右四个方向前进，如果下一个格子的值比当前格子的值大，则可以前进。

定义 $f[i][j]$ 表示以第 $i$ 行第 $j$ 列的格子为起点的路径数。

由于路径中的数字严格递增，状态无后效性，可以用动态规划解决。

我们把四个方向可以走的格子所对应的状态 $f[i'][j']$ 累加起来，再加上 $1$，即当前格子组成的长度为 $1$ 的路径，即为 $f[i][j]$。

代码实现时可以用记忆化搜索。

本题 [视频讲解](https://www.bilibili.com/video/BV1Yf4y1Z7Ac)。

```py [sol-Python3]
class Solution:
    def countPaths(self, grid: List[List[int]]) -> int:
        MOD = 1_000_000_007
        m, n = len(grid), len(grid[0])
        @cache
        def dfs(i: int, j: int) -> int:
            res = 1
            for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                if 0 <= x < m and 0 <= y < n and grid[x][y] > grid[i][j]:
                    res += dfs(x, y)
            return res % MOD
        return sum(dfs(i, j) for i in range(m) for j in range(n)) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int[][] DIRS = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int countPaths(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] memo = new int[m][n];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }

        long ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                ans += dfs(i, j, grid, memo);
            }
        }
        return (int) (ans % MOD);
    }

    private int dfs(int i, int j, int[][] grid, int[][] memo) {
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        long res = 1;
        for (int[] d : DIRS) {
            int x = i + d[0];
            int y = j + d[1];
            if (0 <= x && x < grid.length && 0 <= y && y < grid[i].length && grid[x][y] > grid[i][j]) {
                res += dfs(x, y, grid, memo);
            }
        }
        return memo[i][j] = (int) (res % MOD); // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;
    const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    int countPaths(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector<int>(n, -1)); // -1 表示没有计算过

        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = 1;
            for (auto& [dx, dy] : dirs) {
                int x = i + dx, y = j + dy;
                if (0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j]) {
                    res = (res + dfs(x, y)) % MOD;
                }
            }
            return res;
        };

        long long ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                ans += dfs(i, j);
            }
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func countPaths(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 1
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j] {
				res = (res + dfs(x, y)) % mod
			}
		}
		*p = res // 记忆化
		return res
	}
	for i, row := range grid {
		for j := range row {
			ans += dfs(i, j)
		}
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。有 $\mathcal{O}(mn)$ 个状态，每个状态有 $\mathcal{O}(1)$ 个转移来源，计算所有状态的时间为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。

## 相似题目

- [329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)

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
