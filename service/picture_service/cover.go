package picture_service

import (
	"fmt"
	"github.com/goccy/go-json"
	"strconv"
	"web/model"
	"web/model/picture"
	"web/utils"
)

func (ps *PictureService) InsertChinesePictureCover() {
	ps.toStruct()
	return
}

func (ps *PictureService) toStruct() {
	js := ``
	var autoGenerated AutoGenerated
	json.Unmarshal([]byte(js), &autoGenerated)

	list := autoGenerated.Book

	var temp picture.ChineseBook
	var tempList []picture.ChineseBook
	for idx, _ := range list {
		//bytes := []byte(list[idx].BookID)
		// 截取后5位
		//lastFive := bytes[len(bytes)-5:]
		// 将字节切片转换回字符串
		//result := string(lastFive)
		idIndex := "24" + list[idx].BookID
		temp.BookId = utils.MD5String(idIndex)
		bookIdOld, _ := strconv.Atoi(idIndex)
		temp.BookIdOld = bookIdOld
		temp.Title = list[idx].Name
		temp.Icon_1 = "https://qqcdn.qiuqiuhuiben.com" + list[idx].CoverSmall
		temp.Type = 1

		fmt.Printf("%#v \n", temp)
		//db := model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Where("book_id_old = ?", idIndex)
		//db.Updates(&temp)
		tempList = append(tempList, temp)
	}
	model.DefaultPicture().Model(&picture.ChineseBook{}).Debug().Create(&tempList)
}

type AutoGenerated struct {
	Page struct {
		ShowCount   int `json:"showCount"`
		CurrentPage int `json:"currentPage"`
		TotalPage   int `json:"totalPage"`
		TotalResult int `json:"totalResult"`
	} `json:"page"`
	Tags struct {
		Name       string `json:"name"`
		TagID      string `json:"tagId"`
		Describe   string `json:"describe"`
		Cover      string `json:"cover"`
		IsSelected bool   `json:"isSelected"`
		TypeID     int    `json:"typeId"`
	} `json:"tags"`
	Book []struct {
		Name             string `json:"name"`
		BookID           string `json:"bookId"`
		Cover            string `json:"cover"`
		ReadQuantity     string `json:"readQuantity"`
		FavoriteQuantity string `json:"favoriteQuantity"`
		Desc             string `json:"desc"`
		Author           string `json:"author"`
		Publisher        string `json:"publisher"`
		CopyrightStatus  string `json:"copyrightStatus"`
		ChargeStatus     string `json:"chargeStatus"`
		Price            int    `json:"price"`
		LimitFreeDate    string `json:"limitFreeDate"`
		PackSize         string `json:"packSize"`
		BookType         string `json:"bookType"`
		BgMusicPath      string `json:"bgMusicPath"`
		OnlineDate       string `json:"onlineDate"`
		Intro            string `json:"intro"`
		Version          string `json:"version"`
		CoverSmall       string `json:"coverSmall"`
		FirstLetter      string `json:"firstLetter"`
		ClickQuantity    string `json:"clickQuantity"`
		ShareURL         string `json:"shareUrl"`
		Dubber           string `json:"dubber"`
		RecommendText    string `json:"recommendText"`
		AdditionStatus   string `json:"additionStatus"`
		BookGuide        string `json:"bookGuide"`
		BookJSON         string `json:"bookJson"`
		PackPath         string `json:"packPath"`
		Orientation      string `json:"orientation"`
		PageNum          string `json:"pageNum"`
	} `json:"book"`
	ReqCode struct {
		MessageStr string `json:"messageStr"`
		Extra      struct {
		} `json:"extra"`
		MessageCode int `json:"messageCode"`
	} `json:"reqCode"`
}
