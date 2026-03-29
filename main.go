package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// 日本標準時 (JST) のロケーションを取得
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst).Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] GitHub Actions: Goプログラムが正常に実行されました。\n", now)
}
