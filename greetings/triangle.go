package greetings

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) changeBase(f float64) {

	t.base = f

	return

}

func (t *Triangle) changeBasePnt(f float64) {

	t.base = f

	return

}