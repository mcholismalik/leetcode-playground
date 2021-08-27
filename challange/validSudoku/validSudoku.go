package validSudoku

import "fmt"

func Run() {
	fmt.Println("start")

	seed := Seed4()
	res := Solution2(seed)
	fmt.Println("res", res)

	fmt.Println("end")
}

func Solution(board [][]byte) bool {
	size := 9
	for i := 0; i < size; i++ {
		ptr := []int{i, i}

		if insPlus := func(a [][]byte, size int, ptr []int) bool {
			for ip, vp := range ptr {
				dup := make(map[string]bool)

				for i := 0; i < size; i++ {
					var v string
					if ip == 0 {
						v = string(a[vp][i])
					} else {
						v = string(a[i][vp])
					}

					if v != "." {
						if dup[v] {
							return false
						}
						dup[v] = true
					}
				}
			}
			return true
		}(board, size, ptr); !insPlus {
			return insPlus
		}

		if (i+1)%3 == 0 {
			if insBox := func(a [][]byte, size int, ptr []int) bool {
				for i := 0; i < 3; i++ {
					n := 1
					for {
						mtx := [][]int{
							{ptr[0] - 2, ptr[1] - 2},
							{ptr[0], ptr[1]},
						}

						if i > 0 {
							mtx[0][i-1] += (3 * n)
							mtx[1][i-1] += (3 * n)

							if mtx[0][i-1] > size-1 || mtx[1][i-1] > size-1 {
								break
							}
						}

						dup := make(map[string]bool)
						for x := mtx[0][0]; x <= mtx[1][0]; x++ {
							for y := mtx[0][1]; y <= mtx[1][1]; y++ {
								v := string(a[x][y])
								if v != "." {
									if dup[v] {
										return false
									}
									dup[v] = true
								}
							}
						}

						if i == 0 {
							break
						}
						n++
					}
				}

				return true
			}(board, size, ptr); !insBox {
				return insBox
			}
		}
	}

	return true
}

func Solution2(board [][]byte) bool {
	var row, col, box []map[byte]bool
	for i := 0; i < 9; i++ {
		row = append(row, make(map[byte]bool))
		col = append(col, make(map[byte]bool))
		box = append(box, make(map[byte]bool))
	}

	for x := range board {
		for y, v := range board[x] {
			if string(v) == "." {
				continue
			}

			if row[x][v] || col[y][v] || box[x/3+(y/3)*3][v] {
				return false
			}

			row[x][v] = true
			col[y][v] = true
			box[x/3+(y/3)*3][v] = true
		}
	}

	return true
}
