package main

import (
	"os"

	"github.com/kataras/golog"
	"github.com/qubarran/hashcode-2018/pkg/city"
)

func main() {
	golog.SetLevel("debug")

	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	golog.Debugf("Running in %s", dir)

	c := city.FromFile(os.Args[1])

	golog.Debug(city.ToString(&c))
}
