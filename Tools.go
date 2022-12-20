package codewar

import (
	"fmt"
	"time"
)

func measureTime(funcname string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("使用 %s 函式時間為 %v", funcname, time.Since(start))
	}
}
