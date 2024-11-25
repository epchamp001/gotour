# Exercise: Web Crawler

In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the `Crawl` function to fetch URLs in parallel without fetching the same URL twice.

_Hint:_ you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!


## My solution
Created a `Cache` structure in which the `map` of visited `urls` will be stored.
The `Visit` method:
- Checks if the `URL` has already been visited.
- Adds the `URL` to the cache if it has not been visited yet.
- Returns `true` if the URL is new, and `false' if it has already been visited.

The `NewSafeCache` function:
- Used to initialize the `Cache` structure
- Creates a `Cache` object with the necessary initial parameters.
- Initializes the `visited` field as an empty map `map[string]bool` to store visited `URLs`.

I also use a `sync.Wait Group` to synchronize the completion of all goroutines.

### Results
```text
found: https://golang.org/ "The Go Programming Language"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/pkg/fmt/ "Package fmt"
```



