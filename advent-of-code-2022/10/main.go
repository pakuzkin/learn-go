package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func cycles() []int {
	return []int{20, 60, 100, 140, 180, 220}
}

type Cpu struct {
	x     int
	ticks int
	total map[int]int
}

func (c *Cpu) doOp(cmd string) {
	words := strings.Split(cmd, " ")
	if words[0] == "noop" {
		c.tick()
	} else {
		num, _ := strconv.Atoi(words[1])
		c.tick()
		c.tick()
		c.x += num
	}

	// fmt.Println(cmd, " result ", c)
}

func (c *Cpu) tick() {
	c.ticks++
	if slices.Contains(cycles(), c.ticks) {
		c.total[c.ticks] = c.x
	}
}

func main() {
	fmt.Println("qwe")
	// Open the file
	file, err := os.Open("./real.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	cpu := Cpu{1, 0, make(map[int]int)}

	// Read and print each line
	for scanner.Scan() {
		line := scanner.Text()
		cpu.doOp(line)
	}

	fmt.Println(cpu.total)
	var ans int
	for k, v := range cpu.total {
		ans += k * v
	}
	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}
