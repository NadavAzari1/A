package A

import (
	"fmt"

	b "github.com/NadavAzari1/B"
)

func HelloWorldA() {
	b.HelloWorldB()
	fmt.Println("Hello Wrold A!")
}
