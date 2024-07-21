package main

import (
	"testing"
	"time"
)

func slowStubWebsitesChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// CheckWebsite using a slice of one hundred urls and uses a fake implementation of websitesChecker
// is deliberately slow.
// time.sleep to wait exactly 20 millisecond and then return true.
// ResetTimer is to rest the time the test before it actually run
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsitesChecker, urls)
	}
}
