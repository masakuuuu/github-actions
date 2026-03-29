package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] GitHub Actions: Goプログラムが正常に実行されました。\n", now)
}
