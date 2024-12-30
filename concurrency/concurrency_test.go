package myconcurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebSites(t *testing.T) {
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://idontknow.com"}
	mockChecker := func(_ string) bool { return true }
	got := CheckWebSites(mockChecker, urls)
	want := map[string]bool{"http://www.google.com": true, "http://www.yahoo.com": true, "http://idontknow.com": true}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func SlowChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebSitesPerf(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebSites(SlowChecker, urls)
	}

}
