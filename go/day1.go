package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("../inputs/day1.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    top_elf := 0
    second_elf := 0
    third_elf := 0

    current_elf := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text := scanner.Text()
        if len(text) == 0 {
            if current_elf >= top_elf {
                third_elf = second_elf
                second_elf = top_elf
                top_elf = current_elf
            } else if current_elf >= second_elf {
                third_elf = second_elf
                second_elf = current_elf
            } else if current_elf >= third_elf {
                third_elf = current_elf
            }

            current_elf = 0
        } else {
            i, err := strconv.Atoi(text)
            if err != nil {
                panic(err)
            }
            current_elf += i
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println("Part 1:", top_elf)
    fmt.Println("Part 2:", top_elf + second_elf + third_elf)
}
