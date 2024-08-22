package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

func test(useIDs, ids []string) {
	fmt.Printf("Generating counts for %v user IDs....\n", len(useIDs))

	counts := getCounts(useIDs)
	fmt.Println("Counts from select IDs:")
	for _, ok := range ids {
		v := counts[ok]
		fmt.Printf("- %s: %d\n", ok, v)
	}
	fmt.Println("======================================================")
}

func main() {
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	test(userIDs, []string{"00", "ff", "dd"})
	test(userIDs, []string{"aa", "12", "32"})
}
