package object

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type HashMap map[string]any

// MergeHashMap MergeHashMap. 合并map
func MergeHashMap(toMap *HashMap, subMaps ...*HashMap) *HashMap {
	if toMap == nil {
		toMap = &HashMap{}
	}
	// 拍平subMaps
	for _, subMap := range subMaps {
		if subMap != nil {
			// 迭代每个HashMap
			for k, v := range *subMap {
				toV := (*toMap)[k]
				// if the key is not exist in toMap
				if toV == nil {
					(*toMap)[k] = v
					continue
				}
				// if the toMap by the key is ""
				switch toV.(type) {
				case string:
					if (*toMap)[k] == "" && v != "" {
						(*toMap)[k] = v
					}
				}
			}
		}
	}
	return toMap
}

// ReplaceHashMapRecursive replace
func ReplaceHashMapRecursive(toMap *HashMap, subMaps ...*HashMap) *HashMap {
	if toMap == nil {
		toMap = &HashMap{}
	}
	// 拍平subMaps
	for _, subMap := range subMaps {
		if subMap != nil {
			// 迭代每个HashMap
			for k, v := range *subMap {
				(*toMap)[k] = v
			}
		}
	}
	return toMap
}

func HashMapToStringMap(obj *HashMap) (newMap *StringMap, err error) {
	newMap = &StringMap{}
	if obj == nil {
		return newMap, err
	}
	for k, v := range *obj {
		(*newMap)[k] = fmt.Sprintf("%v", v)
	}
	return newMap, err
}

func StructToHashMapWithXML(obj any) (newMap *HashMap, err error) {
	newMap = &HashMap{}
	if obj == nil {
		return newMap, err
	}
	e := reflect.ValueOf(obj).Elem()
	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i).Interface()
		key := e.Type().Field(i).Tag.Get("xml")
		(*newMap)[key] = field
	}
	return newMap, err
}

func HashMapToStructure(mapObj *HashMap, obj any) (err error) {
	strObj, err := JsonEncode(mapObj)
	if err != nil {
		return err
	}
	err = JsonDecode([]byte(strObj), obj)
	return err
}

func StructToHashMap(obj any) (newMap *HashMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		return
	}
	newMap = &HashMap{}
	err = json.Unmarshal(data, newMap) // Convert to a map
	return
}

func InHash(val any, hash *HashMap) (exists bool, key string) {
	exists = false
	key = ""
	switch reflect.TypeOf(hash).Kind() {
	case reflect.Map:
		//s := reflect.ValueOf(hash)
		for k, v := range *hash {
			if reflect.DeepEqual(val, v) {
				key = k
				return
			}
		}
	default:
		panic("type error")
	}
	return
}

func GetHashMapKV(maps StringMap) (keys []string, values []any) {
	mapLen := len(maps)
	keys = make([]string, 0, mapLen)
	values = make([]any, 0, mapLen)
	for k, v := range maps {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func FilterEmptyHashMap(mapData *HashMap) (filteredMap *HashMap) {
	filteredMap = &HashMap{}
	for k, v := range *mapData {
		if v != nil {
			(*filteredMap)[k] = v
			switch value := v.(type) {
			case HashMap:
				v = FilterEmptyHashMap(&value)
			case *HashMap:
				v = FilterEmptyHashMap(value)
			//case []interface{}:
			//	for _, obj := range v.([]interface{}) {
			//		switch obj.(type) {
			//		case map[string]interface{}:
			//			o, _ := StructToHashMap(obj)
			//			v = FilterEmptyHashMap(o)
			//			break
			//		default:
			//
			//		}
			//	}
			//	break
			case string:
				if value == "" {
					delete(*filteredMap, k)
				}
			}
		}
	}
	return filteredMap
}
