package main

import "fmt"

func main() {
    m := make([]int, 3)
    fmt.Println(m)

    m_t := m
    fmt.Println(m_t)

    m_t[0] = 1
    fmt.Println(m)
    fmt.Println(m_t)
    fmt.Println(&m[0])
    fmt.Println(&m_t[0])

    m_t = append(m_t, 3)
    fmt.Println(m)
    fmt.Println(m_t)
    fmt.Println(&m[0])
    fmt.Println(&m_t[0])

    m_t[1] = 2
    fmt.Println(m)
    fmt.Println(m_t)
    fmt.Println(&m[0])
    fmt.Println(&m_t[0])
}
