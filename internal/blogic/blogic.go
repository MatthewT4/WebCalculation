package blogic

type IBlogic interface {
	Plus(first float64, second float64) float64
}

type Blogic struct {
}

func (b *Blogic) Plus(first float64, second float64) float64 {
	return first + second
}

func (b *Blogic) substract(first float64, second float64) float64 {
	return first - second
}

func (b *Blogic) division(first float64, second float64) float64 {
	return first / second
}

func (b *Blogic) multiply(first float64, second float64) float64 {
	return first * second
}
