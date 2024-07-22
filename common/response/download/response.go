package download

type AlbumResponse struct {
	AlbumName   string `json:"album_name"`
	AudioIDList []struct {
		AudioID string `json:"audio_id"`
	} `json:"audio_id_list"`
	AudioInfoList []struct {
		AudioID          string `json:"audio_id"`
		AudioImgURL      string `json:"audio_img_url"`
		AudioLength      string `json:"audio_length"`
		AudioName        string `json:"audio_name"`
		AudioPlayCount   string `json:"audio_play_count"`
		AudioPublishTime string `json:"audio_publish_time"`
		AudioTestState   string `json:"audio_test_state"`
		AudioType        string `json:"audio_type"`
		BuyStatus        string `json:"buy_status"`
		ShowOrder        string `json:"show_order"`
	} `json:"audio_info_list"`
	AudioPrice   string        `json:"audio_price"`
	Balance      string        `json:"balance"`
	BuyWay       string        `json:"buy_way"`
	DownloadList []interface{} `json:"download_list"`
	ErrorInfo    struct {
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"error_info"`
	OrderType  string `json:"order_type"`
	StatusCode string `json:"status_code"`
	Timestamp  string `json:"timestamp"`
}

type AlbumInfoResponse struct {
	AlbumName     string        `json:"album_name"`
	AudioIDList   []interface{} `json:"audio_id_list"`
	AudioInfoList []struct {
		AudioID          string `json:"audio_id"`
		AudioImgURL      string `json:"audio_img_url"`
		AudioLength      string `json:"audio_length"`
		AudioName        string `json:"audio_name"`
		AudioPlayCount   string `json:"audio_play_count"`
		AudioPublishTime string `json:"audio_publish_time"`
		AudioTestState   string `json:"audio_test_state"`
		AudioType        string `json:"audio_type"`
		BuyStatus        string `json:"buy_status"`
		ShowOrder        string `json:"show_order"`
	} `json:"audio_info_list"`
	AudioPrice   string        `json:"audio_price"`
	Balance      string        `json:"balance"`
	BuyWay       string        `json:"buy_way"`
	DownloadList []interface{} `json:"download_list"`
	ErrorInfo    struct {
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"error_info"`
	OrderType  string `json:"order_type"`
	StatusCode string `json:"status_code"`
	Timestamp  string `json:"timestamp"`
}

type AlbumInfoMp3Response struct {
	AdvertisingPath     string `json:"advertisingPath"`
	Advertisingduration int    `json:"advertisingduration"`
	AudioDetailInfoList []struct {
		AlbumBuyState     string   `json:"album_buy_state"`
		AlbumID           string   `json:"album_id"`
		AlbumImgURL       string   `json:"album_img_url"`
		AlbumName         string   `json:"album_name"`
		AlbumRssState     string   `json:"album_rss_state"`
		AlbumType         string   `json:"album_type"`
		AllowDownload     string   `json:"allow_download"`
		AudioBrief        string   `json:"audio_brief"`
		AudioBuyState     string   `json:"audio_buy_state"`
		AudioCommentCount string   `json:"audio_comment_count"`
		AudioID           string   `json:"audio_id"`
		AudioImgList      []string `json:"audio_img_list"`
		AudioImgURL       string   `json:"audio_img_url"`
		AudioLength       string   `json:"audio_length"`
		AudioLikeCount    string   `json:"audio_like_count"`
		AudioLikeState    string   `json:"audio_like_state"`
		AudioName         string   `json:"audio_name"`
		AudioPrice        string   `json:"audio_price"`
		AudioTagInfoList  []struct {
			TagID    string `json:"tag_id"`
			TagName  string `json:"tag_name"`
			TagStyle string `json:"tag_style"`
			TagType  string `json:"tag_type"`
		} `json:"audio_tag_info_list"`
		AudioTestURL            string `json:"audio_test_url"`
		AudioType               string `json:"audio_type"`
		AudioURL                string `json:"audio_url"`
		AuditionTime            string `json:"audition_time"`
		AuthorAuthenticateState string `json:"author_authenticate_state"`
		AuthorFollowState       string `json:"author_follow_state"`
		AuthorID                string `json:"author_id"`
		AuthorImgURL            string `json:"author_img_url"`
		AuthorName              string `json:"author_name"`
		BuyWay                  string `json:"buy_way"`
		Chapter                 string `json:"chapter"`
		Introduction            string `json:"introduction"`
		IsNotText               int    `json:"isNotText"`
		IsAuto                  string `json:"is_auto"`
		IsMember                string `json:"is_member"`
		IsMemberResource        string `json:"is_member_resource"`
		IsYuewen                string `json:"is_yuewen"`
		MemberStatus            string `json:"member_status"`
		PlayLogTime             string `json:"play_log_time"`
		Price                   int    `json:"price"`
		SnyStatus               int    `json:"sny_status"`
		YuewenAcid              string `json:"yuewen_acid"`
		YuewenAdid              string `json:"yuewen_adid"`
	} `json:"audio_detail_info_list"`
	ErrorInfo struct {
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"error_info"`
	H5RegionalRestrictions string      `json:"h5RegionalRestrictions"`
	MemberStatus           string      `json:"member_status"`
	MembersAdvertising     int         `json:"membersAdvertising"`
	RegionPermissionStatus string      `json:"regionPermissionStatus"`
	SnyStatus              interface{} `json:"snyStatus"`
	StatusCode             string      `json:"status_code"`
	Timestamp              string      `json:"timestamp"`
}
