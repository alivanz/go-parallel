# Parallel

```Go
func main() {
  p := NewParallel(
    context.Background(),
    parallel.Limit(3),
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
