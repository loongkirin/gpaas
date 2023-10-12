package server

import (
	"fmt"
)

func Run() {
	fmt.Println("gpaas server start......")
	InitAppContext()

	// fmt.Println(AppContext.APP_DbContext.GetDb() == nil)
	// fmt.Println(AppContext.APP_REDIS == nil)
}
