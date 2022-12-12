package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/day01"
	"github.com/alokmenghrajani/adventofcode2022/day02"
	"github.com/alokmenghrajani/adventofcode2022/day03"
	"github.com/alokmenghrajani/adventofcode2022/day04"
	"github.com/alokmenghrajani/adventofcode2022/day05"
	"github.com/alokmenghrajani/adventofcode2022/day06"
	"github.com/alokmenghrajani/adventofcode2022/day07"
	"github.com/alokmenghrajani/adventofcode2022/day08"
	"github.com/alokmenghrajani/adventofcode2022/day09"
	"github.com/alokmenghrajani/adventofcode2022/day10"
	"github.com/alokmenghrajani/adventofcode2022/day11"
	"github.com/alokmenghrajani/adventofcode2022/day12"
	"github.com/alokmenghrajani/adventofcode2022/utils"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.Readfile(d)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(d)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day03.Part2(utils.Readfile(d)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(d)))
	case 5:
		fmt.Printf("part 1: %s\n", day05.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %s\n", day05.Part2(utils.Readfile(d)))
	case 6:
		fmt.Printf("part 1: %d\n", day06.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day06.Part2(utils.Readfile(d)))
	case 7:
		fmt.Printf("part 1: %d\n", day07.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day07.Part2(utils.Readfile(d)))
	case 8:
		fmt.Printf("part 1: %d\n", day08.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day08.Part2(utils.Readfile(d)))
	case 9:
		fmt.Printf("part 1: %d\n", day09.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day09.Part2(utils.Readfile(d)))
	case 10:
		fmt.Printf("part 1: %d\n", day10.Part1(utils.Readfile(d)))
		fmt.Printf("part 2:\n%s\n", day10.Part2(utils.Readfile(d)))
	case 11:
		fmt.Printf("part 1: %d\n", day11.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day11.Part2(utils.Readfile(d)))
	case 12:
		fmt.Printf("part 1: %d\n", day12.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day12.Part2(utils.Readfile(d)))
	default:
		genTree()
		panic(fmt.Errorf("no such day: %d", d))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 12
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}

func genTree() {
	fmt.Println("<pre>")
	n := 30
	len := -1
	for i := 0; i < n; i++ {
		len += 2
		if i%7 == 6 {
			len -= 4
		}
		fmt.Printf("%s", strings.Repeat(" ", n-len/2+1))
		if i == 0 {
			fmt.Println("<span class=\"s\">*</span>")
			continue
		}
		fmt.Printf("/")
		len2 := len
		if i%7 == 5 {
			len2 -= 2
			fmt.Printf("^")
		}
		for j := 0; j < len2-2; j++ {
			t := rand.Intn(5 * i)
			if t < i {
				fmt.Printf("<span class=\"c%d\">•</span>", rand.Intn(3))
			} else if t < int(1.5*float64(i)) {
				fmt.Printf("<span class=\"s\">*</span>")
			} else {
				if i == n-1 {
					fmt.Printf("^")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		if i%7 == 5 {
			fmt.Printf("^")
		}
		fmt.Println("\\")
	}
	fmt.Printf("<span class=\"t\">")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s", strings.Repeat(" ", n-1))
		fmt.Println("===")
	}
	fmt.Printf("</span>")
	fmt.Println("</pre>")
}
