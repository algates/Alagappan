package compare

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"reflect"
	"sort"
	//"fmt"
)

type Compare struct {
	URL1 string
	URL2 string
	Resp1 []byte
	Resp2 []byte
}

func CompareJSON(resp1 []byte, resp2 []byte) bool {
	var av, bv interface{}
	da := json.NewDecoder(bytes.NewReader(resp1))
	da.UseNumber()
	db := json.NewDecoder(bytes.NewReader(resp2))
	db.UseNumber()
	errA := da.Decode(&av)
	errB := db.Decode(&bv)

	if errA != nil || errB != nil {
		return false
	}

	return getDiff(av,bv)
}

func Check(c Compare) bool {
	var av, bv interface{}
	da := json.NewDecoder(bytes.NewReader(c.Resp1))
	da.UseNumber()
	db := json.NewDecoder(bytes.NewReader(c.Resp2))
	db.UseNumber()
	errA := da.Decode(&av)
	errB := db.Decode(&bv)

	if errA != nil || errB != nil {
		return false
	}

	return getDiff(av,bv)
}


func getDiff(a, b interface{}) bool {
	if a == nil || b == nil {
		return false
	}

	ka := reflect.TypeOf(a).Kind()
	kb := reflect.TypeOf(b).Kind()

	if ka != kb {
		return false
	}
	switch ka {
	case reflect.Bool:
		if a.(bool) != b.(bool) {
			return false
		}
	case reflect.String:
		switch aa := a.(type) {
		case json.Number:
			bb, ok := b.(json.Number)

			if !ok || aa != bb {
				return false
			}
		case string:
			bb, ok := b.(string)
			if !ok || aa != bb {
				return false
			}
		}
	case reflect.Slice:
		sa, sb := a.([]interface{}), b.([]interface{})
		salen, sblen := len(sa), len(sb)
		max := salen
		if sblen > max {
			max = sblen
		}
		for i := 0; i < max; i++ {
			if i < salen && i < sblen {
				check:= getDiff(sa[i], sb[i])
				if !check{
					return false
				}
			} else {
				return false
			}
		}
		return true
	case reflect.Map:
		ma, mb := a.(map[string]interface{}), b.(map[string]interface{})
		keysMap := make(map[string]bool)
		for k := range ma {
			keysMap[k] = true
		}
		for k := range mb {
			keysMap[k] = true
		}
		keys := make([]string, 0, len(keysMap))
		for k := range keysMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			va, aok := ma[k]
			vb, bok := mb[k]
			if aok && bok {
				 check:= getDiff(va, vb)
				 if !check{
				 	return false
				 }
			} else {
				return false
			}
		}
		return true
	}
	return true
}