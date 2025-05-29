package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	return slices.Max(data)
}

var s sync.WaitGroup

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	var maxValuesFromChunks []int

	s.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		j := i * chunkSize
		x := j + chunkSize

		if i == CHUNKS-1 {
			x = len(data)
		}

		func(chunk []int) {
			defer s.Done()
			maxValuesFromChunks = append(maxValuesFromChunks, maximum(chunk))
		}(data[j:x])
	}
	s.Wait()

	return maximum(maxValuesFromChunks)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	randomNumbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeFrom := time.Now()
	maxNumber := maximum(randomNumbers)
	timeTo := time.Now()
	elapsed := timeTo.Sub(timeFrom).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumber, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	timeFrom = time.Now()
	maxNumber = maxChunks(randomNumbers)
	timeTo = time.Now()
	elapsed = timeTo.Sub(timeFrom).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumber, elapsed)
}
