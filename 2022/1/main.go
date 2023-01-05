package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	// TODO: Fix input path.
	f, err := os.Open("2022/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var (
		max1, max2, max3, acc int
	)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			max1, max2, max3 = calcMax3(acc, max1, max2, max3)
			acc = 0
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			acc += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	max1, max2, max3 = calcMax3(acc, max1, max2, max3)

	log.Printf("first answer is %d\n", max1)
	log.Printf("second answer is %d\n", max1+max2+max3)
}

func calcMax3(curr, m1, m2, m3 int) (int, int, int) {
	if curr > m1 {
		m3 = m2
		m2 = m1
		m1 = curr
	} else if curr > m2 {
		m3 = m2
		m2 = curr
	} else if curr > m3 {
		m3 = curr
	}
	return m1, m2, m3
}
