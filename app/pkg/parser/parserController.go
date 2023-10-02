package parser

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	// "log"
	// "encoding/base64"
	"sample/pkg/helpers"
	utils "sample/pkg/utils"
	// "strconv"
	// "time"
)

func parserPayload(g *gin.Context) {
	var payload ParserPayload

	// 使用 ShouldBindJSON 來解析 JSON 格式的 POST payload
	if err := g.ShouldBindJSON(&payload); err != nil {
		// 如果解析失敗，處理錯誤
		helpers.HandleError(g, nil, err)
		return
	}
	fmt.Println("payload", payload.Data)
	dataList := payload.Data.(map[string]interface{})
	// dataList := payload.Data.(map[string]interface{})

	typeNames := []string{"Temperature", "Relative Humidity", "Folder", "Floor", "Room", "Building"}
	response := convertTopo(dataList, typeNames)
	idLst := convertPayloadToMap(response)
	fmt.Println("idLst Info:", idLst)
	host := "https://xintop.viewlinc.tw"
	userName := "api"
	password := "&kHk1La*TM6#"

	authHeader := utils.GetBase64Auth(userName, password)
	client, err := login(host, userName, password)
	// authHeader := base64.StdEncoding.EncodeToString([]byte(userName + ":" + password))
	authHeaders := map[string]string{
		"Authorization": authHeader,
	}
	// if err != nil {
	// 	fmt.Println("Login error:", err)
	// 	return
	// }

	// client := &http.Client{}
	fmt.Println("client in parser", client)
	locationInfo, err := getlocationInfo(idLst, host, authHeaders)
	if err != nil {
		fmt.Println("Get location info error:", err)
		return
	}

	fmt.Println("Location Info:", locationInfo)
	helpers.HandleError(g, idLst, nil)
}

func convertTopo(topoInfo map[string]interface{}, typeNames []string) []map[string]interface{} {
	templst := []map[string]interface{}{}
	fmt.Println("templst", templst, topoInfo)
	data := topoInfo["data"].([]interface{})
	for _, item := range data {
		if dataMap, isMap := item.(map[string]interface{}); isMap {
			itemTemplst := processItem(dataMap, typeNames)
			templst = append(templst, itemTemplst...)
		}
	}
	return templst
}

func processItem(item map[string]interface{}, typeList []string) []map[string]interface{} {
	tempLst := []map[string]interface{}{}

	if children, ok := item["children"].([]interface{}); ok {
		// 递归处理子项，首先处理子项，然后再检查当前item是否匹配typeList
		for _, child := range children {
			if childMap, isMap := child.(map[string]interface{}); isMap {
				childTemplst := processItem(childMap, typeList)
				tempLst = append(tempLst, childTemplst...)
			}
		}
	}

	if typeName, ok := item["type_name"].(string); ok {
		for _, typeNameInList := range typeList {
			if typeName == typeNameInList {
				// 将匹配的项添加到 tempLst，并从 children 中移除 type_name
				tempLst = append(tempLst, item)
				delete(item, "children")
				break
			}
		}
	}

	return tempLst
}

func convertPayloadToMap(payload []map[string]interface{}) []int {
	tempDict := []int{}

	for _, item := range payload {
		nodeID, ok := item["node_id"].(float64)
		if ok {
			tempDict = append(tempDict, int(nodeID))
		}
	}

	return tempDict
}
