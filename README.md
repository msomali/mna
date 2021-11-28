# mna

MNA - stands for mobile number assignment - a small zero external dependency golang library that is used 
to identify mobile number assignment in tanzania. The library is based of TCRA issued document
that is found at https://nms.tcra.go.tz/nms/auvit/National_Numbering_Plan.pdf


## usage

```bash
go get github.com/techcraftt/mna
```

```go

package example

import (
	"fmt"
	"github.com/techcraftlabs/mna"
	"strings"
)

func Example_Details()  {

	info, err := mna.Information("+255788888888")
	if err != nil {
        fmt.Println(err)
    }
	
	operator, err := mna.Get("+25576282735343535")
}
```