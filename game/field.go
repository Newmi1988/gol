package game

type Field struct {
	arr           [][]bool
	width, heigth int
}

func NewField(width, height int) *Field {
	row := make([][]bool, width)
	for i := range row {
		row[i] = make([]bool, height)
	}
	return &Field{arr: row, width: width, heigth: height}
}

func (f *Field) Set(x, y int, b bool) {
	f.arr[x][y] = b
}

func (f *Field) IsAlive(x, y int) bool {
	x += f.width
	x %= f.width // wrap arround if bigger than width
	y += f.heigth
	y %= f.heigth // wrap around if bigger than height
	return f.arr[x][y]
}

func (f *Field) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i != 0 || j != 0) && f.IsAlive(x+i, y+j) {
				alive++
			}
		}
	}
	return alive == 3 || alive == 2 && f.IsAlive(x, y)
}
