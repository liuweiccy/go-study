package trans

import (
	"math"
	"fmt"
)

var Pi float64

func init()  {
	Pi = 4 * math.Atan(1)
	fmt.Println(Pi)
}
