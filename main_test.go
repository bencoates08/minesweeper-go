package main

import (
	. "minesweeper-go/internal/core/domain/board"
	"testing"
)

func BenchmarkNewBoardNew(b *testing.B) {
	boardSettings := BoardSettings{Height: 20, Width: 30, Bombs: 80}
	for i := 0; i < b.N; i++ {
		board := NewBoard(boardSettings)

		board.Reveal(board.GetPosition(0, 0))
		board.Print()
		println()

		board.Reveal(board.GetPosition(0, 4))
		board.Print()
		println()

		board.Reveal(board.GetPosition(1, 2))
		board.Print()
		println()

		board.Reveal(board.GetPosition(15, 15))
		board.Print()
		println()
	}
}
