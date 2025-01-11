package filters

import (
    "strings"
    "os"
    "log"
    "bufio"
    "errors"
)


// Types

type (
    filter struct{
        isFull bool
        value string
    }

    Filters []filter
)


// Functions

func skipDefaults(fileName string) bool {
    return fileName == ".DS_Store" ||
        fileName == ".git" ||
        fileName == ".build" ||
        fileName == ".bin" ||
        fileName == ".idea" ||
        fileName == "_build" ||
        fileName == "elm-stuff" ||
        fileName == "node_modules" ||
        fileName == "stack_work" ||
        fileName == ".qlot" ||
        strings.HasSuffix(fileName, ".app") ||
        strings.HasSuffix(fileName, ".png") ||
        strings.HasSuffix(fileName, ".jpg") ||
        strings.HasSuffix(fileName, ".jpeg") ||
        strings.HasSuffix(fileName, ".gif") ||
        strings.HasSuffix(fileName, ".o") ||
        strings.HasSuffix(fileName, ".ppu")
}


// Receiver funcs

func (filters *Filters) IsFiltered(fileName string) bool {
    if skipDefaults(fileName) {
        return true
    }

    for _, f := range (*filters) {
        if f.isFull {
            if f.value == fileName {
                return true
            }
        } else {
            if strings.HasSuffix(fileName, f.value) {
                return true
            }
        }
    }

    return false
}


func (filters *Filters) Init() {
    filters.buildFromLsrIgnore()
    filters.buildFromGitIgnore()
}

const (
    lsrignore = "." + string(os.PathSeparator) + ".lsrignore"
    gitignore = "." + string(os.PathSeparator) + ".gitignore"
)

func (filters *Filters) buildFromLsrIgnore() {
    if fileExists(lsrignore) {
        f, err := os.Open(lsrignore)

        if err != nil {
            log.Fatal(err)
        }

        defer f.Close()

        scanner := bufio.NewScanner(f)

        isFull := false
        for scanner.Scan() {
            line := scanner.Text()
            if line == "" {
                isFull = true
                continue
            }

            *filters = append(*filters, filter{ isFull: isFull, value: line })
        }
    }
}


func (filters *Filters) buildFromGitIgnore() {
    if !fileExists(gitignore) {
        return;
    }

    f, err := os.Open(gitignore)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        if !folderExists(line) {
            continue
        }

        *filters = append(*filters, filter{ isFull: true, value: line })
    }
}


// UTILS

func folderExists(path string) bool {
    info, err := os.Stat(path)
    return !errors.Is(err, os.ErrNotExist) && info.Mode().IsDir()
}


func fileExists(path string) bool {
    info, err := os.Stat(path)
    return !errors.Is(err, os.ErrNotExist) && !info.Mode().IsDir()
}

