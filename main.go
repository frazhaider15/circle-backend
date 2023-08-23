package main

import (
	"fmt"

	"github.com/qisst/ms-nadra-verification/rest"
	"github.com/qisst/ms-nadra-verification/service"
)

func main() {

	fmt.Println("#==================================#")
	fmt.Println("#===========Starting Server =======#")
	fmt.Println("#==================================#")

	/*
	* Initiate Service Layer Container
	 */

	serviceContainer := service.NewServiceContainer()

	/*
	* Initiate Rest Server
	 */
	rest.StartServer(serviceContainer)

	fmt.Println("========== Rest Server Started ============")
	fmt.Println("========== Server Running ============")

	select {}

}
