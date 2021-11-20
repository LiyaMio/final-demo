package send

import (
	"fmt"
	"sort"
)

func sorted(unsort []int,k int64){
	sort.Ints(unsort)
	fmt.Println(unsort)
}
