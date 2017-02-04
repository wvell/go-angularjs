package main

import (
	"log"

	angularjs "github.com/wvell/go-angularjs"
)

func main() {
	app := angularjs.NewModule("test", []string{}, nil)
	app.NewController("TestCtrl", TestCtrl)

	log.Print("running app")
}

func TestCtrl(scope *angularjs.Scope) {
	scope.Set("test", "Im a variable that is set on angular's $scope")

	ticked := 0
	scope.Set("tick", func() {
		ticked++
		log.Printf("ticking: %d", ticked)
	})
}
