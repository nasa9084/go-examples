package example_test

import (
	"fmt"
	"time"
	// _ "time/tzdata"
)

// ExampleTZData is an example using time/tzdata. time/tzdata is imported
// so timezone data is embedded and this function won't cause panic even
// if timezone info is not installed in the runtime environment.
func ExampleTZData() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// caused panic if time/tzdata is not imported and
		// timezone info is not installed in the runtime env.
		panic(err)
	}

	fmt.Println(time.Now().In(jst).Format("-0700"))

	// Output:
	// +0900
}
