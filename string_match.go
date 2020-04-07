package main

import "fmt"

func find1(target, match string) {
    fmt.Println(target)
    fmt.Println(match)
    next := make([]int, len(match))
    next[0] = -1
    //next compute
    for idx, c := range match {
        if idx == 0 {
            next[0] = -1
            continue
        }
        j := next[idx-1]+1
        for true {
            if byte(c) == match[j] {
                next[idx] = j
                break
            } else if j == 0 {
                next[idx] = -1
                break
            }else {
                j = next[j-1] + 1
            }
        }
    }
    fmt.Println("next:", next)

    i:=0
    j:=0
    for i < len(target) && j < len(match) {
        if target[i] == match[j] {
            i++
            j++
        } else if j == 0 { // || next[j-1] == -1 {
            i++
        } else {
            fmt.Println("next[", j-1, "]:", next[j-1])
            j = next[j-1] + 1
        }
        fmt.Println("i:", i, "j:", j)
    }
    if j == len(match) {
        fmt.Println(target[i-len(match): i])
    }
}

func find2(target, match string) {
    next := make([]int, len(match))
    //next compute
    for idx, c := range match {
        if idx == 0 {
            next[0] = 0
            continue
        }
        j := next[idx-1]
        for true {
            if byte(c) == match[j] {
                next[idx] = j + 1
                break
            } else if j == 0 {
                next[idx] = 0
                break
            } else {
                j = next[j-1]
            }
        }
    }

    i:=0
    j:=0
    for i < len(target) && j < len(match) {
        if target[i] == match[j] {
            i++
            j++
        } else if j == 0 {
            i++
        } else {
            j = next[j-1]
        }
    }
    fmt.Println("i:", i, "j:", j)
    if j == len(match) {
        fmt.Println(target[i-len(match): i])
    }
}

func main() {
    // find1("mississippi", "pi")
    find1("aabaaabaaac", "aabaaac")
    find2("aabaaabaaac", "aabaaac")
    // find2("abcdefadaba", "fiad")
    // find2("abcdefadaba", "afad")
    // find2("abcdefadaba", "defad")
}
