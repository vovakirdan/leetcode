package uniquepaths

func uniquePaths(m int, n int) int {
	return comb(m+n-2, m-1)
}

func comb(n, k int) int {
	if k > n-k {
		k = n - k
	}

	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - i + 1) / i
	}
	return res
}
