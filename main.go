package main

import (
	"fmt"

	"github.com/Tsuryu/bff-basu/routers"
)

func main() {
	routers.Router()
	fmt.Println("Server up and running")
}
