package main

import (
	"fmt"
	"os"
	"strconv"
)

type fileBlock struct {
	length   int
	startsAt int // so we can replace all the values with `...`
}

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
	m := map[int]fileBlock{} // map that has the ID and the block
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
			m[id] = fileBlock{
				length:   n,
				startsAt: len(disk) - n,
			}
			id++
			free = true
		}
	}

	// fmt.Println(disk)
	// we don't store the last ++
	id--
loop:
	for i := id; i >= 0; i-- {
		// find a fitting space
		block := m[i]
		// fmt.Println("considering block: ", block)
		begin, end := 0, 0
		for j := 0; j <= block.startsAt; j++ {
			if disk[j] == "." {
				begin = j
				for k := j; k <= block.startsAt; k++ {
					if disk[k] != "." {
						end = k
						// fmt.Println("segment: ", disk[begin:end])

						if end-begin >= block.length {
							// do move
							for c := 0; c < block.length; c++ {
								disk[begin+c], disk[block.startsAt+c] = disk[block.startsAt+c], disk[begin+c]
							}
							// fmt.Println(disk)

							// moved, next block
							continue loop
						}

						// skip this entire section, no point in trying to see
						// if the current block fits.. move on.
						j = k
						break
					}
				}
			}
		}
	}
	// fmt.Println(disk)
	// id will be the last highest ID, so we just loop that backwards.

	// loop from the end backwards
	// keep track of the last free space
	// do a swap
	// lastFreeSpaceIndex := 0
	// for i, d := range disk {
	// 	if d == "." {
	// 		lastFreeSpaceIndex = i
	// 		break
	// 	}
	// }
	// for i := len(disk) - 1; i >= 0; i-- {
	// 	if disk[i] != "." {
	// 		disk[lastFreeSpaceIndex], disk[i] = disk[i], disk[lastFreeSpaceIndex]

	// 		for n := lastFreeSpaceIndex; n < len(disk)-1; n++ {
	// 			if disk[n] == "." {
	// 				lastFreeSpaceIndex = n
	// 				break
	// 			}
	// 		}
	// 		// we went over and it would start putting numbers back to back
	// 		if lastFreeSpaceIndex >= i {
	// 			break
	// 		}
	// 		continue
	// 	}
	// }

	fmt.Println(disk)
	// printDisk(disk)
	sum := 0
	for index, v := range disk {
		if v != "." {
			n, _ := strconv.Atoi(v)
			sum += index * n
		}
	}

	fmt.Println(sum)
}
