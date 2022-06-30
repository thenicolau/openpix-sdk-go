# OpenPix SDK in Golang

## Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/GustavoNicolau/openpix-sdk-go/service"
)
func main() {

	op := service.Connect("your-token")

	charge, err := op.GetOneCharge().Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	result, _ := json.Marshal(&charge)
	log.Println(string(result))

	
}

```