package validSudoku

func Seed1() [][]byte {
	// expect true
	str := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}

	bin := [][]byte{}
	for _, v := range str {
		b := []byte(v)
		bin = append(bin, b)
	}

	return bin
}

func Seed2() [][]byte {
	// expect false
	str := []string{
		"83..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}

	bin := [][]byte{}
	for _, v := range str {
		b := []byte(v)
		bin = append(bin, b)
	}

	return bin
}

func Seed3() [][]byte {
	// expect false
	str := []string{
		"....5..1.",
		".4.3.....",
		".....3..1",
		"8......2.",
		"..2.7....",
		".15......",
		".....2...",
		".2.9.....",
		"..4......",
	}

	bin := [][]byte{}
	for _, v := range str {
		b := []byte(v)
		bin = append(bin, b)
	}

	return bin
}

func Seed4() [][]byte {
	// expect false
	str := []string{
		"..4...63.",
		".........",
		"5......9.",
		"...56....",
		"4.3.....1",
		"...7.....",
		"...5.....",
		".........",
		".........",
	}

	bin := [][]byte{}
	for _, v := range str {
		b := []byte(v)
		bin = append(bin, b)
	}

	return bin
}
