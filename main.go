package main

import (
	"fmt"
	"roboBuilder/builder"
)

func main() {
	robotBuilder := builder.GetBuilder("gundam")
	robot := robotBuilder.Build()

	fmt.Printf("%+v", robot)
}
