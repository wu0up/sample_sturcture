package main

import (
	"fmt"
)

type TopoInfo struct {
	Msg     string         `json:"msg"`
	Data    []TopoNode     `json:"data"`
	Success bool           `json:"success"`
}

type TopoNode map[string]interface{}

func convertTopoFormat(dataArray []interface{}, typeNames []string, tempLst *[]interface{}) {
	for _, item := range dataArray {
		switch v := item.(type) {
		case map[string]interface{}:
			if typeName, ok := v["type_name"].(string); ok {
				for _, typeNameInList := range typeNames {
					if typeName == typeNameInList {
						*tempLst = append(*tempLst, v)
						break
					}
				}
			}
			if children, ok := v["children"].([]interface{}); ok {
				convertTopoFormat(children, typeNames, tempLst)
			}
		case []interface{}:
			convertTopoFormat(v, typeNames, tempLst)
		}
	}
}

func main() {
	// 示例数据，用于测试 convertTopoFormat 函数
	topoInfo := TopoInfo{
		Msg:     "Test Message",
		Success: true,
		Data: []TopoNode{
			{
				"description": "The viewLinc system.",
				"type_id":     0,
				"type_name":   "Folder",
				"children": []TopoNode{
					{
						"description": "Sample Building",
						"type_id":     1,
						"type_name":   "Building",
						"children": []TopoNode{
							{
								"description": "Room 1",
								"type_id":     2,
								"type_name":   "Room",
							},
						},
					},
				},
			},
		},
	}

	typeNames := []string{"Folder", "Building", "Room"}

	// 调用 convertTopoFormat 函数进行测试
	var tempLst []interface{}
	convertTopoFormat(topoInfo.Data, typeNames, &tempLst)

	// 输出测试结果
	fmt.Printf("Temp List: %v\n", tempLst)
}
