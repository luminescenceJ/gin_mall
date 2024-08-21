//package main
//
//import (
//	"gin_mal_tmp/conf"
//	"gin_mal_tmp/routes"
//)
//
//func main() {
//	conf.Init()
//	r := routes.NewRouter()
//	_ = r.Run(conf.HttpPort)
//}

package main

import "fmt"

func main() {
	var C, N int
	fmt.Scan(&C, &N)
	weights := make([]int, N)
	values := make([]int, N)
	k := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&values[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&k[i])
	}

	dp := make([]int, C+1) //dp[i]在容量为i时能够获得的最大价值
	dp[0] = 0

	for i := 1; i <= C; i++ {
		index := -1
		for idx, weight := range weights {
			if k[idx] <= 0 || weight > i {
				continue
			}
			if dp[i] < dp[i-weight]+values[idx] {
				index = idx
				dp[i] = dp[i-weight] + values[idx]
			}
		}
		if index != -1 {
			k[index]--
		}
	}

	fmt.Println(dp[C])

}
