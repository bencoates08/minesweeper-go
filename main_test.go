package main

import (
	. "minesweeper-go/internal/core/domain/board"
	"testing"
)

func BenchmarkNewBoardNew(b *testing.B) {
	boardSettings := BoardSettings{Height: 20, Width: 30, Bombs: 300}
	for i := 0; i < b.N; i++ {
		_ = NewBoardState(boardSettings)
	}
}
