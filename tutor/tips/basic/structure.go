package basic

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(reflect.TypeOf(Vertex{1, 2}))
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}