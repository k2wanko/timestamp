package timestamp_test

import (
	"fmt"
	"log"
	"time"

	"github.com/k2wanko/timestamp"
)

func ExampleTimeStamp() {
	o := &struct {
		*timestamp.TimeStamp
	}{
		TimeStamp: &timestamp.TimeStamp{
			Created: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		},
	}

	err := o.Update()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(o.IsZero()) //Output: false
}
