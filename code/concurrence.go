package main

import (
    "fmt"
    "sync"
)

func main() {
    sum := 0

    var wg sync.WaitGroup
    var mu sync.Mutex

    wg.Add(1)
    go foo(&wg, &sum, &mu)

    wg.Add(1)
    go bar(&wg, &sum, &mu)

    wg.Wait()
    fmt.Println("sum: ", sum)
}

func foo(wg *sync.WaitGroup, sum *int, mu *sync.Mutex) {
    defer wg.Done()
    defer mu.Unlock()
    mu.Lock()
    for i := 0; i < 100000; i++ {
        *sum += 1
    }
}

func bar(wg *sync.WaitGroup, sum *int, mu *sync.Mutex) {
    defer wg.Done()
    defer mu.Unlock()
    mu.Lock()
    for i := 0; i < 100000; i++ {
        *sum += 1
    }
}


