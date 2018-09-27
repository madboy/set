### set

Non-complete set implementation written for exploration and things like adventofcode. If you are looking for something more complete take a look at [golang-set](https://github.com/deckarep/golang-set).

### Examples

```go

package main

import (
    "github.com/madboy/set"
    "fmt"
)

func main() {
    is1 := set.NewIntFromArr([]int{1,1,2,3,5,5,66})
    is2 := set.NewIntFromArr([]int{1,2,3,4})

    fmt.Println("is1 values:", is1.Values())
    fmt.Println("is2 values:", is2.Values())

    diff1 := is1.Difference(&is2)
    diff2 := is2.Difference(&is1)

    fmt.Println("is1 - is2:", diff1.Values())
    fmt.Println("is2 - is1:", diff2.Values())
}
```

```
is1 values: [1 2 3 5 66]
is2 values: [1 2 3 4]
is1 - is2: [5 66]
is2 - is1: [4]
```
