package main

import (
	"github.com/mattmeyers/level"
	"github.com/mattmeyers/mmdev/app"
	"github.com/mattmeyers/mmdev/http"
)

func main() {
	srv := http.NewServer()
	srv.Logger, _ = level.NewBasicLogger(level.Info, nil)
	srv.Resources = app.Resources

	if err := srv.ListenAndServe(":8080"); err != nil {
		panic(err)
	}
}
