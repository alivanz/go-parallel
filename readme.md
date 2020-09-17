# Parallel

```Go
package main

import (
	"context"
	"log"
	"time"

	"github.com/alivanz/go-parallel"
)

func main() {
	p := parallel.NewParallel(
		context.Background(),
		parallel.NewLimit(3),
	)
	for i := 0; i < 10; i++ {
		p.Do(func(ctx context.Context) {
			log.Print("hi")
			time.Sleep(1 * time.Second)
		})
	}
	log.Print("wait")
	p.Wait()
	log.Print("done")
}

```

```txt
2020/09/17 17:19:09 hi
2020/09/17 17:19:09 hi
2020/09/17 17:19:09 wait
2020/09/17 17:19:09 hi
2020/09/17 17:19:10 hi
2020/09/17 17:19:10 hi
2020/09/17 17:19:10 hi
2020/09/17 17:19:11 hi
2020/09/17 17:19:11 hi
2020/09/17 17:19:11 hi
2020/09/17 17:19:12 hi
2020/09/17 17:19:13 done
```
