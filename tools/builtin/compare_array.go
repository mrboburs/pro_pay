package builtin

import "github.com/google/uuid"

func CompareTwoUUIDArray(arr1, arr2 []uuid.UUID) (arr1haveArr2Not, arr2haveArr1Not []uuid.UUID) {
	arr1Map := make(map[uuid.UUID]int)
	arr2Map := make(map[uuid.UUID]int)
	for i := range arr1 {
		arr1Map[arr1[i]]++
	}
	for i := range arr2 {
		arr2Map[arr2[i]]++
	}
	for i := range arr2 {
		if _, ok := arr1Map[arr2[i]]; !ok {
			arr2haveArr1Not = append(arr2haveArr1Not, arr2[i])
		}
	}
	for i := range arr1 {
		if _, ok := arr2Map[arr1[i]]; !ok {
			arr1haveArr2Not = append(arr1haveArr2Not, arr1[i])
		}
	}
	return arr1haveArr2Not, arr2haveArr1Not
}
