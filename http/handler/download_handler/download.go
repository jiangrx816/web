package download_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	downloadResq "web/common/response/download"
	downloadReq "web/common/resquest/download"
	"web/model"
	"web/model/story"
	"web/utils"
	"web/utils/errs"
)

type response struct {
	Info string `json:"info"`
}

// ApiAlbumList 获取对应的列表
func (dh *DownloadHandler) ApiAlbumList(ctx *gin.Context) {

	value := ctx.Query("index")
	var responseR response
	albumId := []string{"2314|1127206", "2314|1122188", "2316|1119263", "2316|1114687", "2312|1124536", "2310|1114760", "2312|1111574", "1248|1114673", "2316|1111644", "2314|1123154", "2312|1111551", "2314|1111573", "2314|1122181", "2314|1120004", "2314|1111636", "2314|1120001", "2314|1121256", "2314|1122183", "2310|1112847", "2310|1112858", "2310|1120412", "2310|1120469", "2310|1120470", "2310|1122276", "2310|1121253", "2310|1111635", "2310|1122049", "2310|1120005", "2312|1121292", "2312|1121926", "2316|1117038", "2312|1119997", "2312|1120006", "2310|1114242", "2310|1112496", "2310|1123157", "2310|1113726", "2312|1121294", "2312|1121255", "2310|1121056", "2310|1121254", "2312|1122182", "2312|1123297", "2310|1123405", "12811|1111468", "12811|1111474", "12811|1111479", "12811|1120384", "12811|1120385", "12811|1111472", "12816|1120399", "12816|1114284", "113019|1113444", "113019|1112862"}
	index, _ := strconv.Atoi(value)
	dh.updateBatchMp3(albumId[index])

	responseR.Info = albumId[index]
	ctx.JSON(errs.SucResp(responseR))
}

func (dh *DownloadHandler) updateBatchMp3(albumId string) {

	var albumInfoList []story.SStoryAlbumInfo
	model.DefaultStory().Model(&story.SStoryAlbumInfo{}).Debug().Where("album_id = ?", albumId).Order("id asc").Find(&albumInfoList)
	for idx, _ := range albumInfoList {
		var audioInfoMp3Request = downloadReq.AudioInfoMp3Request{
			AppType:        "01",
			AppVersion:     "4.8.6",
			DeviceID:       "344f7d4d6840db97",
			MessageID:      "838b71d2-45ff-4c94-b6a3-66ea354a4d9b",
			DeviceType:     "2",
			Lon:            "",
			Version:        "3.0",
			DeviceUUID:     "344f7d4d6840db97",
			Token:          "f58d57b9-d04b-4c34-9953-edc9f55c3ca1",
			SysVersion:     "30",
			UserID:         "0",
			SystemVersion:  "30",
			Model:          "Redmi Note 8",
			ChannelType:    "aiyinsitan",
			DeviceUniqueID: "344f7d4d6840db97",
			Brand:          "xiaomi",
			Lat:            "",
		}
		ss := albumInfoList[idx].AudioId
		audioInfoMp3Request.AudioIDList = []struct {
			AudioID string `json:"audio_id"`
		}{
			{AudioID: ss},
		}
		marshal, err2 := json.Marshal(audioInfoMp3Request)
		if err2 != nil {
			fmt.Println(err2)
		}
		requestJson, err := utils.HttpPostRequestJson("https://mobile.aiyinsitanfm.com/audiodetailinfo", marshal)
		if err != nil {
			fmt.Println("HTTP 请求失败")
		}

		var albumInfoMp3Response downloadResq.AlbumInfoMp3Response
		err3 := json.Unmarshal([]byte(requestJson), &albumInfoMp3Response)
		if err3 != nil {
			fmt.Println(err3)
		}
		audioDetailInfoList := albumInfoMp3Response.AudioDetailInfoList

		var albumInfoTemp story.SStoryAlbumInfo
		if len(audioDetailInfoList) > 0 {
			if audioDetailInfoList[0].AudioURL != "" {
				albumInfoTemp.AudioMp3 = audioDetailInfoList[0].AudioURL
			} else {
				if audioDetailInfoList[0].AudioTestURL != "" {
					albumInfoTemp.AudioMp3 = audioDetailInfoList[0].AudioTestURL
				}
			}
		}

		model.DefaultStory().Model(&story.SStoryAlbumInfo{}).Debug().Where("id = ?", albumInfoList[idx].Id).Updates(&albumInfoTemp)
	}
}

func (dh *DownloadHandler) doSaveAlbumData(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("第%d 个goroutine 正在执行 \n", id)
	var albumRequestList []downloadReq.AlbumAudioRequest
	var albumRequest = downloadReq.AlbumAudioRequest{
		AppType:        "01",
		AppVersion:     "4.8.6",
		DeviceID:       "344f7d4d6840db97",
		MessageID:      "ce1ab011-f6be-4712-b628-b618a05585d9",
		DeviceType:     "2",
		Lon:            "",
		Version:        "3.0",
		DeviceUUID:     "344f7d4d6840db97",
		Token:          "f58d57b9-d04b-4c34-9953-edc9f55c3ca1",
		SysVersion:     "30",
		UserID:         "0",
		SystemVersion:  "30",
		AlbumID:        "",
		Model:          "Redmi Note 8",
		ChannelType:    "aiyinsitan",
		DeviceUniqueID: "344f7d4d6840db97",
		Brand:          "xiaomi",
		AudioIDList:    []interface{}{},
		Lat:            "",
	}

	var infoRequest = downloadReq.AlbumAudioInfoRequest{
		AppType:        "01",
		AppVersion:     "4.8.6",
		DeviceID:       "344f7d4d6840db97",
		MessageID:      "bd50ee17-2a9c-41c2-9ede-529427b3d2e7",
		DeviceType:     "2",
		Lon:            "",
		Version:        "3.0",
		DeviceUUID:     "344f7d4d6840db97",
		Token:          "f58d57b9-d04b-4c34-9953-edc9f55c3ca1",
		SysVersion:     "30",
		UserID:         "0",
		SystemVersion:  "30",
		AlbumID:        "2314|1119491",
		Model:          "Redmi Note 8",
		ChannelType:    "aiyinsitan",
		DeviceUniqueID: "344f7d4d6840db97",
		Brand:          "xiaomi",
		//AudioIDList   : []interface{}{},
		Lat: "",
	}
	var albumList []story.SStoryAlbum
	db := model.DefaultStory().Model(&story.SStoryAlbum{}).Debug().Order("id desc")
	db.Find(&albumList)
	for idx, _ := range albumList {
		albumRequest.AlbumID = albumList[idx].ItemId
		albumRequestList = append(albumRequestList, albumRequest)
		infoRequest.AlbumID = albumList[idx].ItemId
	}
	for idx, _ := range albumRequestList {
		var item = albumRequestList[idx]

		marshal, err := json.Marshal(albumRequestList[idx])
		if err != nil {

		}
		requestJson, err := utils.HttpPostRequestJson("https://mobile.aiyinsitanfm.com/albumaudioinfolist", marshal)
		if err != nil {
			fmt.Println("HTTP 请求失败")
		}
		var AlbumResponse downloadResq.AlbumResponse
		if err := json.Unmarshal([]byte(requestJson), &AlbumResponse); err != nil {
			fmt.Println("HTTP 请求失败 json.Unmarshal 错误")
		}
		audioIdList := AlbumResponse.AudioIDList
		if len(audioIdList) < 1 {
			continue
		}
		infoRequest.AudioIDList = audioIdList
		marshal1, err1 := json.Marshal(infoRequest)
		if err1 != nil {

		}
		requestJson1, err1 := utils.HttpPostRequestJson("https://mobile.aiyinsitanfm.com/albumaudioinfolist", marshal1)
		if err1 != nil {
			fmt.Println("HTTP 请求失败")
		}
		var AlbumInfoResponse downloadResq.AlbumInfoResponse
		if err := json.Unmarshal([]byte(requestJson1), &AlbumInfoResponse); err != nil {
			fmt.Println("HTTP 请求失败 json.Unmarshal 错误")
		}
		audioInfoList := AlbumInfoResponse.AudioInfoList
		for index, _ := range audioInfoList {
			var storyAlbumInfo story.SStoryAlbumInfo
			storyAlbumInfo.AlbumId = item.AlbumID
			storyAlbumInfo.AudioId = audioInfoList[index].AudioID
			storyAlbumInfo.AudioName = audioInfoList[index].AudioName
			storyAlbumInfo.AudioImgUrl = audioInfoList[index].AudioImgURL
			model.DefaultStory().Model(&story.SStoryAlbumInfo{}).Debug().Create(&storyAlbumInfo)
		}

	}
}
