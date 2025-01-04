package main

import (
	"github.com/gofiber/fiber/v2"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	type stackFrame struct {
		start int
		end   int
	}
	stack := []stackFrame{{0, len(arr) - 1}}

	for len(stack) > 0 {
		frame := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		start, end := frame.start, frame.end

		if start >= end {
			continue
		}

		pivotIndex := partition(arr, start, end)

		stack = append(stack, stackFrame{start, pivotIndex - 1})
		stack = append(stack, stackFrame{pivotIndex + 1, end})
	}
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}

func main() {
	app := fiber.New()

	app.Post("/sort", func(c *fiber.Ctx) error {
		var arr []int
		if err := c.BodyParser(&arr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		quickSort(arr)

		return c.JSON(arr)
	})

	app.Listen(":8082")
}
