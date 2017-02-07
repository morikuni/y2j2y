package y2j2y

import (
	"io"
)

type Converter interface {
	Convert(io.Reader, io.Writer) error
}
