package onword

import "fmt"

// MemoryGet retrieves an int from a specified position in the OnWord's memory.
func (o *OnWord) MemoryGet(x int) (int, error) {
	if x < 0 || x >= len(o.Memory) {
		return 0, fmt.Errorf(
			"memory get %d: %w",
			x, ErrOutOfBounds,
		)
	}

	return o.Memory[x], nil
}

// MemorySet stores an int at a specified position in the OnWord's memory.
func (o *OnWord) MemorySet(x int, n int) error {
	if x < 0 || x >= len(o.Memory) {
		return fmt.Errorf(
			"memory set %d %d: %w",
			x, n, ErrOutOfBounds,
		)
	}

	o.Memory[x] = n

	return nil
}

// MemoryAdd appends an int to the end of the OnWord's memory, and returns the new position.
func (o *OnWord) MemoryAdd(n int) int {
	o.Memory = append(o.Memory, n)
	return len(o.Memory) - 1
}
