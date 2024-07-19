package picture_service

import (
	"bytes"
	"encoding/json"
	"github.com/jiangrx816/gopkg/server/api"
	"io/ioutil"
	"log"
	"net/http"
	pictureResp "web/common/response/picture"
	pictureReq "web/common/resquest/picture"
)

func (ps *PictureService) FindEnglishBookList(typeId, offset int) (response pictureResp.EnglishBookResponse, apiErr api.Error) {
	// 创建要发送的数据
	data := pictureReq.RequestEnglishParam{
		HTs:        "1685021025740e",
		HM:         28800,
		Zone:       0,
		HLc:        "zh",
		Uid:        0,
		Token:      "",
		HCn:        "miniprogram",
		HDt:        3,
		Cate:       1,
		Atype:      3,
		Source:     4,
		Offset:     int32(offset),
		Difficulty: int32(typeId),
	}

	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// 创建 POST 请求
	url := "https://www.ipalfish.com/klian/ugc/picturebook/level/list"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var contentResult pictureResp.ResponseEnglishBook
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	response.Data = contentResult.Data
	return
}

func (ps *PictureService) FindEnglishBookInfo(bookId int) (response pictureResp.EnglishBookInfoResponse, apiErr api.Error) {

	// 创建要发送的数据
	data := pictureReq.RequestEnglishInfoParam{
		HDt:    3,
		HDid:   "16851689799320000",
		Did:    "16851689799320000",
		HCH:    "miniprogram",
		HTs:    "1685168991352",
		HM:     0,
		Zone:   28800,
		Token:  "",
		HLc:    "zh",
		Cate:   1,
		Atype:  3,
		BookId: bookId,
		Limit:  200,
	}

	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// 创建 POST 请求
	url := "https://www.ipalfish.com/klian/ugc/picturebook/official/product/bookid/page/list"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var contentResult pictureResp.ResponseEnglishBookInfo
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	response.Data = contentResult.Data

	return
}
