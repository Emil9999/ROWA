package main

import (
	"fmt"
	"time"
)

func main() {
	yesterday := time.Now().AddDate(0, 0, -1)
	fmt.Println("then:", yesterday)

	for i := 0; i < 5; i++ {
		yesterday = yesterday.Add(time.Hour)
		fmt.Println("then:", yesterday)
	}

}
