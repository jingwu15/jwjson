package json

/*
//用法示例
jsondata := utils.JsonParse(respRaw)
fmt.Println(reflect.TypeOf(utils.JsonGet(jsondata, "took")))
fmt.Printf("%.1f", utils.JsonGet(jsondata, "hits.max_score"))
fmt.Println(reflect.TypeOf(utils.JsonGet(jsondata, "hits.max_score")))
fmt.Println(reflect.TypeOf(utils.JsonGet(jsondata, "aggregations.NAME.buckets")))
fmt.Println(reflect.TypeOf(utils.JsonGet(jsondata, "aggregations.NAME.buckets.0.key")))
fmt.Println(reflect.TypeOf(utils.JsonGet(jsondata, "aggregations.NAME.buckets.0.doc_count")))
*/

import (
	"errors"
	"strconv"
	"strings"
    "io/ioutil"
	"encoding/json"
)

func JType(v interface{}) string {
	var vtype string
	switch v.(type) {
	case string:
		vtype = "string"
	case float64:
		vtype = "float64"
	case bool:
		vtype = "bool"
	case int:
		vtype = "int"
	case []interface{}:
		vtype = "[]interface{}"
	default:
		vtype = "map[string]interface{}"
	}
	return vtype
}

func convert(raw interface{}) (map[string]interface{}, []interface{}, interface{}, string) {
	vtype := JType(raw)
	var empMap map[string]interface{}
	var empList []interface{}
	var empUnit interface{}
	if vtype == "[]interface{}" {
		return empMap, raw.([]interface{}), empUnit, "list"
	}
	if vtype == "map[string]interface{}" {
		return raw.(map[string]interface{}), empList, empUnit, "map"
	}
	return empMap, empList, raw, "unit"
}

func JParse(raw []byte) (interface{}, error) {
	var jsondata interface{}
    err := json.Unmarshal(raw, &jsondata)
    if err == nil {
	    return jsondata,nil
    } else {
        return jsondata, err
    }
}

func JParseFile(jsonfile string) (interface{}, error) {
	var jsondata interface{}
    raw, err := ioutil.ReadFile("./test.json")
    if err != nil {
        return jsondata, err
    }
    err = json.Unmarshal(raw, &jsondata)
    if err == nil {
	    return jsondata,nil
    } else {
	    return jsondata,err
    }
}

func gotoKey(jsondata interface{}, key string) interface{} {
	var response interface{}
	response = jsondata
	t1, t2, t3, t4 := convert(jsondata)
	keys := strings.Split(key, ".")
	if key == "" {
		if t4 == "list" {
			response = t2
		} else if t4 == "map" {
			response = t1
		} else {
			response = t3
		}
		return response
	}
	for _, k := range keys {
		if t4 == "list" {
			kInt, _ := strconv.Atoi(k)
			if kInt < len(t2) {
				response = t2[kInt]
				t1, t2, t3, t4 = convert(t2[kInt])
			} else {
				return nil
			}
		} else if t4 == "map" {
			if _, ok := t1[k]; ok {
				response = t1[k]
				t1, t2, t3, t4 = convert(t1[k])
			}
		} else {
			response = t3
		}
	}
	return response
}

func JGet(jsondata interface{}, key string) (interface{}, error) {
    response := gotoKey(jsondata, key)
    if response == nil {
	    return false, errors.New("not exists")
    } else {
	    return response, nil
    }
}

func JGetList(jsondata interface{}, key string) ([]interface{}, error) {
    var rows []interface{}
    response := gotoKey(jsondata, key)
    if response == nil {
	    return rows, errors.New("not exists")
    } else {
	    return response.([]interface{}), nil
	}
}

func JGetMap(jsondata interface{}, key string) (map[string]interface{}, error) {
    var rows map[string]interface{}
    response := gotoKey(jsondata, key)
    if response == nil {
	    return rows, errors.New("not exists")
    } else {
	    return response.(map[string]interface{}), nil
	}
}

