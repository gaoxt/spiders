package storage

import (
	"encoding/json"
	"singleTaskCrawfer/dao"
	"singleTaskCrawfer/engine"
	"strconv"
)

func BookInsert(c interface{}) engine.StorageResult {
	client := dao.RedisClient()
	mapBookList := make(map[string]interface{})
	jsonBookList, _ := json.Marshal(c)
	json.Unmarshal(jsonBookList, &mapBookList)
	jsonDetail, _ := json.Marshal(mapBookList["Detail"])
	delete(mapBookList, "Detail")
	mapBookList["Detail"] = string(jsonDetail)

	bookID := strconv.FormatFloat(mapBookList["Id"].(float64), 'f', -1, 64)
	err := client.HMSet(bookID, mapBookList).Err()
	if err != nil {
		return engine.StorageResult{Err: err}
	}
	return engine.StorageResult{Items: []interface{}{bookID}, Err: err}
}
