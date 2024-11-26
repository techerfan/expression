package contracts

type TextSpan struct {
	Start  int
	Length int
}

func NewTextSpan(start int, length int) *TextSpan {
	return &TextSpan{
		Start:  start,
		Length: length,
	}
}
