package util

import (
	set "github.com/senseisub/collegepluscollegefinderbackend-backend-api/util/Set"
	"github.com/senseisub/collegepluscollegefinderbackend-backend-api/util/Set/hashset"
)

func ArrToSet(arr *[]string) *set.Set[string] {
	ret := hashset.NewWithInput[string](arr)
	return &ret
}
