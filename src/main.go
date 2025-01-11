package main

import (
    "strings"
    "fmt"
    "os"
    "log"
    "sync"
    flts "lsr/filters"
)

const (
    CURR_DIR = "." + string(os.PathSeparator)
    CURR_DIR_LEN = len(CURR_DIR)
    // TODO Find the optimal number, this one is arbitrary
    MAX_FILES_SIZE = 5000
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
            addFile(path[CURR_DIR_LEN:] + fileName)
        }
    }
}


func addFile(fileName string) {
    lock.Lock()

    files = append(files, fileName)

    if len(files) >= MAX_FILES_SIZE {
        fmt.Println(strings.Join(files, "\n"))
        files = nil
    }

    lock.Unlock()
}


func main() {
    filters.Init()

    wg.Add(1)
    go readDir(CURR_DIR)

    wg.Wait()

    if len(files) > 0 {
        fmt.Println(strings.Join(files, "\n"))
    }
}

