package picture_service

import (
	"encoding/json"
	"github.com/jiangrx816/gopkg/server/api"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	pictureResp "web/common/response/picture"
	"web/utils"
)

func (ps *PictureService) FindPoetryBookList(page, limit, typeId int) (response pictureResp.PoetryBookResponse, apiErr api.Error) {
	utils.DefaultIntOne(&page)
	utils.DefaultIntFifty(&limit)
	utils.DefaultIntOne(&typeId)

	// 设置表单数据
	data := url.Values{
		"org_id":  {"0"},
		"user_id": {"1"},
		"diff":    {"all"},
		"is_read": {"all"},
		"sort":    {"id"},
		"limit":   {"" + strconv.Itoa(limit) + ""},
		"page":    {"" + strconv.Itoa(page) + ""},
		"type_id": {"" + strconv.Itoa(typeId) + ""},
	}

	// 将表单数据转换为字符串
	dataString := data.Encode()

	// 转换为字节流
	dataBytes := strings.NewReader(dataString)

	// 设置请求
	req, err := http.NewRequest("POST", "https://mzbook.com/api/book.Book/getBookList", dataBytes)
	if err != nil {
		panic(err)
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 打印响应
	var contentResult pictureResp.ResponsePoetryBookJson
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	response.Data = contentResult.Data
	return
}

func (ps *PictureService) FindPoetryBookInfo(id int) (response pictureResp.PoetryBookInfoResponse, apiErr api.Error) {
	utils.DefaultIntOne(&id)
	// 设置表单数据
	data := url.Values{
		"id": {"" + strconv.Itoa(id) + ""},
	}

	// 将表单数据转换为字符串
	dataString := data.Encode()

	// 转换为字节流
	dataBytes := strings.NewReader(dataString)

	// 设置请求
	req, err := http.NewRequest("POST", "https://mzbook.com/api/book.Book/getSentence", dataBytes)
	if err != nil {
		panic(err)
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 打印响应
	var contentResult pictureResp.ResponsePoetryBookInfo
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	response.Data = contentResult.Data
	return
}
