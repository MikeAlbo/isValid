## isValid
a validation package for simple text inputs

* character strings
* int strings
* street address 
* apartment numbers

## usage
Currently, for each method call you must enter a reference to a string value `&value`, and a min and max length value.

```go
package main 
import (
	"isValid"
	"fmt"
)

var t = "HelloWorld"

func ValidateSomething(){
	if err:= isValid.Length(&t,1,10); err != nil {
    	fmt.Println(err.Error())
    }
}
```
