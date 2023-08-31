package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context) {
    defer wg.Done()
    childCtx, _ := context.WithCancel(ctx)
    go memoryInfo(childCtx)
    for {
        select {
        case <- ctx.Done():
            fmt.Println("cpuInfo exit")
            return
        default:
            time.Sleep(time.Second * 1)
            fmt.Println("cpu info")
        }
    }
}

func memoryInfo(ctx context.Context) {
    defer wg.Done()
    for {
        select {
        case <- ctx.Done():
            fmt.Println("memoryInfo exit")
            return
        default:
            time.Sleep(time.Second * 1)
            fmt.Println("memory info")
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(2)
    go cpuInfo(ctx)
    // go memoryInfo(ctx)
    time.Sleep(time.Second * 3)
    cancel()
    wg.Wait()
    fmt.Println("main exit")
}