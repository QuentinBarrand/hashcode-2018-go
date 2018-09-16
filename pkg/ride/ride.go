package ride

import (
	"strings"

	"github.com/qubarran/hashcode-2018/pkg/coord"
	"github.com/qubarran/hashcode-2018/pkg/util"
)

type Ride struct {
	Id      int
	Start   *coord.Coord
	End     *coord.Coord
	Earlist int
	latest  int
}

// FromLine builds a Ride struct from a line of an input file.
func FromLine(index int, line string) *Ride {
	splitLine := strings.Split(line, " ")

	return &Ride{
		index,
		coord.New(
			util.ParseInt(splitLine[0]),
			util.ParseInt(splitLine[1]),
		),
		coord.New(
			util.ParseInt(splitLine[2]),
			util.ParseInt(splitLine[3]),
		),
		util.ParseInt(splitLine[4]),
		util.ParseInt(splitLine[5]),
	}
}
