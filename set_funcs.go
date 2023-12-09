package util

import (
	set "github.com/CollegePlusCo/GoUtil/Set"
	"github.com/CollegePlusCo/GoUtil/Set/hashset"
)

func ArrToSet(arr *[]string) *set.Set[string] {
	ret := hashset.NewWithInput[string](arr)
	return &ret
}
