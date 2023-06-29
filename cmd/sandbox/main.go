package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, 8, 10, 18, 11, 11, 0, time.UTC)

	fmt.Println(fmt.Sprintf("seconds: %v", t.Unix()))
	fmt.Println(fmt.Sprintf("mili s: %v", t.UnixMilli()))
	fmt.Println(fmt.Sprintf("micro s: %v", t.UnixMicro()))
	fmt.Println(fmt.Sprintf("nanos : %v", t.UnixNano()))

}
