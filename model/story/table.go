package story

type SStoryAlbum struct {
	ContentId string `json:"content_id"`
	ItemId    string `json:"item_id"`
	ShowName  string `json:"show_name"`
	AlbumName string `json:"album_name"`
	HostName  string `json:"host_name"`
	ImageAddr string `json:"image_addr"`
	IsFree    int    `json:"is_free"`
}

func (SStoryAlbum) TableName() string {
	return "s_story_album"
}

type SStoryAlbumInfo struct {
	Id          int    `json:"id"`
	AlbumId     string `json:"album_id"`
	AudioId     string `json:"audio_id"`
	AudioName   string `json:"audio_name"`
	AudioMp3    string `json:"audio_mp3"`
	AudioImgUrl string `json:"audio_img_url"`
}

func (SStoryAlbumInfo) TableName() string {
	return "s_story_album_info"
}
