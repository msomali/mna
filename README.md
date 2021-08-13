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
	"github.com/techcraftt/mna"
	"strings"
)

func Example_CheckNumber()  {

	phone := "0765992100"
	response, err := mna.CheckNumber(phone)
	if err != nil{
		fmt.Printf("error occurred: %v\n",err)
		return
	}
	
	if response.CommonName == mna.AirtelCommonName{
		fmt.Printf("Yes the Provided Number belongs to Airtel")
		return
	}else {
		fmt.Printf("The Provided Number does not belong to Airtel but to %s\n", response.CommonName)
		fmt.Printf("operationsal status of the number is %s\n", response.Status)
		fmt.Printf("the official registered name of the provider is %s\n", response.OperatorName)
		fmt.Printf("All the prefixes owned by the operator are %s\n", strings.Join(response.Prefixes, ","))
		return
	}
}
```