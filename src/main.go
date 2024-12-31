package main

import (
    "strings"
    "fmt"
    "os"
    "log"
    "sync"
    flts "lsr/filters"
)


var (
    lock sync.Mutex
    wg sync.WaitGroup
    filters flts.Filters
    files []string
)


func readDir(path string) {
    defer wg.Done()

    entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    for _, entry := range entries {
        fileName := entry.Name()
        if filters.IsFiltered(fileName) {
            continue
        }

        if entry.IsDir() {
            wg.Add(1)
            go readDir(path + fileName + string(os.PathSeparator))
        } else {
            addFile(path[2:] + fileName)
        }
    }
}


func addFile(fileName string) {
    lock.Lock()

    files = append(files, fileName)

    // TODO Find the optimal number, this one is arbitrary
    if len(files) >= 5000 {
        fmt.Println(strings.Join(files, "\n"))
        files = nil
    }

    lock.Unlock()
}


func main() {
    filters.Init()

    wg.Add(1)
    go readDir("." + string(os.PathSeparator))

    wg.Wait()

    if len(files) > 0 {
        fmt.Println(strings.Join(files, "\n"))
    }
}

