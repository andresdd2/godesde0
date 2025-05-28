package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int     // Por declaración, no se define como nulas, sino con su valor más bajo
	intde32 := int32(10) // Por asignación
	intde64 := int64(100)
	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
