package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf585E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, v int
	Fscan(in, &n)
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	const mx int = 1e7
	var cnt, f [mx + 1]int32
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}

	ans := 0
	for i := mx; i > 1; i-- {
		tot := int(cnt[i])
		sumF := 0
		for j := i * 2; j <= mx; j += i {
			tot += int(cnt[j])
			sumF += int(f[j])
		}
		res := ((pow2[tot]-1)*(n-tot) - sumF) % mod
		f[i] = int32(res)
		ans += res
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf585E(bufio.NewReader(os.Stdin), os.Stdout) }
