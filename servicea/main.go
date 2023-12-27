package main

import (
	servicea "github.com/wheelergeo/g-otter-gen/servicea/myservice"
	"log"
)

func main() {
	svr := servicea.NewServer(new(MyServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
