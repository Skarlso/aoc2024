package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	id := 0
	disk := []string{}
	free := false
	for i := 0; i < len(content); i++ {
		n := int(content[i]) - '0'
		if free {
			for range n {
				disk = append(disk, ".")
			}
			free = false
		} else {
			for range n {
				disk = append(disk, strconv.Itoa(id))
			}
			id++
			free = true
		}
	}

	// loop from the end backwards
	// keep track of the last free space
	// do a swap
	lastFreeSpaceIndex := 0
	for i, d := range disk {
		if d == "." {
			lastFreeSpaceIndex = i
			break
		}
	}
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != "." {
			disk[lastFreeSpaceIndex], disk[i] = disk[i], disk[lastFreeSpaceIndex]

			for n := lastFreeSpaceIndex; n < len(disk)-1; n++ {
				if disk[n] == "." {
					lastFreeSpaceIndex = n
					break
				}
			}
			// we went over and it would start putting numbers back to back
			if lastFreeSpaceIndex >= i {
				break
			}
			continue
		}
	}

	fmt.Println(disk)
	// printDisk(disk)
	sum := 0
	for index, v := range disk {
		if v == "." {
			break
		}

		n, _ := strconv.Atoi(v)
		sum += index * n
	}

	fmt.Println(sum)
}
