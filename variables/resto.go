package variables

import (
	"strconv"
)

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
