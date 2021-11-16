package main

import (
	"flag"

	"gwi/platform2.0-go-challenge/dbschema"
)

func main() {
	cmdObjects := flag.String("objects", "", "")
	cmdSeed := flag.String("seed", "", "")
	flag.Parse()

	switch *cmdObjects {
	case "up":
		dbschema.DBObjects()
	default:
	}

	switch *cmdSeed {
	case "up":
		dbschema.DBSeed()
	case "down":
		dbschema.DBClear()
	default:
	}
}
