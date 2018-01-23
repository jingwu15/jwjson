package main

import(
    "fmt"
    "reflect"
    json "github.com/jingwu15/jwjson/json"
)

func main() {
    //读取配置json文件
    jhandle, err := json.JParseFile("./test.json")
    if err != nil {
        fmt.Println(err)
    }

    //解析json
    resultString, err := json.JGet(jhandle, "log_error")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resultString, reflect.TypeOf(resultString))

    resultBool, err := json.JGet(jhandle, "test")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resultBool, reflect.TypeOf(resultBool))

    resultList, err := json.JGetList(jhandle, "levels")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resultList)
    for _, row := range resultList {
        fmt.Println(row, reflect.TypeOf(row))
    }

    resultMap, err := json.JGetMap(jhandle, "")
    fmt.Println(resultMap, reflect.TypeOf(resultMap))
    //fmt.Println(resultMap["es_api"], reflect.TypeOf(resultMap["es_api"]))
}

