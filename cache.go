package intersdk

import "sync"

// localCache is a cache that stores the oauth locally.
// It is used to avoid making requests to the oauth server every time
// because the rate limit is 5 requests per minute. This cache is
// used to store the oauth token for 1 minute.
type localCache struct {
	stop chan struct{}

	wg sync.WaitGroup
	mu sync.RWMutex
}
