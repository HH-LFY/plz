package test

import (
	"testing"
	"github.com/v2pro/plz/reflect2"
)

func Test_slice_map(t *testing.T) {
	t.Run("MakeSlice", testOp(func(api reflect2.API) interface{} {
		valType := api.TypeOf([]map[int]int{}).(reflect2.SliceType)
		obj := valType.MakeSlice(5, 10)
		obj.([]map[int]int)[0] = map[int]int{1:1}
		obj.([]map[int]int)[4] = map[int]int{2:2}
		return obj
	}))
	t.Run("Set", testOp(func(api reflect2.API) interface{} {
		obj := []map[int]int{{1: 1}, nil}
		valType := api.TypeOf(obj).(reflect2.SliceType)
		valType.Set(obj, 0, &map[int]int{10:10})
		valType.Set(obj, 1, &map[int]int{2:2})
		return obj
	}))
	t.Run("Get", testOp(func(api reflect2.API) interface{} {
		obj := []map[int]int{{1:1}, nil}
		valType := api.TypeOf(obj).(reflect2.SliceType)
		return []interface{}{
			valType.Get(&obj, 0),
			valType.Get(&obj, 1),
			valType.Get(obj, 0),
			valType.Get(obj, 1),
		}
	}))
	t.Run("Append", testOp(func(api reflect2.API) interface{} {
		obj := make([]map[int]int, 2, 3)
		obj[0] = map[int]int{1:1}
		obj[1] = map[int]int{2:2}
		valType := api.TypeOf(obj).(reflect2.SliceType)
		obj = valType.Append(obj, map[int]int{3:3}).([]map[int]int)
		// will trigger grow
		obj = valType.Append(obj, map[int]int{4:4}).([]map[int]int)
		return obj
	}))
}