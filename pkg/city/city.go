package city

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kataras/golog"
	"github.com/qubarran/hashcode-2018/pkg/ride"
	"github.com/qubarran/hashcode-2018/pkg/util"
)

type city struct {
	Rows     int
	Columns  int
	Vehicles int
	Bonus    int
	Steps    int

	RidesArr []*ride.Ride
}

func addRide(c *city, r *ride.Ride) {
	c.RidesArr = append(c.RidesArr, r)
}

// FromFile This function reads a City file and returns a City struct instance with is characteristics.
func FromFile(path string) (c city) {
	golog.Debugf("Reading file %s", path)

	fd, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	// Read the header
	scanner.Scan()

	header := strings.Split(scanner.Text(), " ")

	rides := util.ParseInt(header[3])

	c = city{
		Rows:     util.ParseInt(header[0]),
		Columns:  util.ParseInt(header[1]),
		Vehicles: util.ParseInt(header[2]),
		Bonus:    util.ParseInt(header[4]),
		Steps:    util.ParseInt(header[5]),
		RidesArr: make([]*ride.Ride, rides),
	}

	for r := 0; r < rides; r++ {
		scanner.Scan()

		ride := ride.FromLine(r, scanner.Text())

		addRide(&c, ride)
	}

	return
}

func ToString(c *city) string {
	return fmt.Sprintf(
		"City with %d rows | %d cols | %d rides",
		c.Rows,
		c.Columns,
		len(c.RidesArr),
	)
}
