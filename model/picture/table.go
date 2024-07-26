package picture

type ChineseBookNavName struct {
	Id         int    `json:"-"`
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
	SSort      int    `json:"s_sort"`
	SType      int    `json:"s_type"`
}

func (ChineseBookNavName) TableName() string {
	return "s_book_name"
}

type ChineseBook struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Type     uint8  `json:"type"`
	Position uint8  `json:"position"`
}

func (ChineseBook) TableName() string {
	return "s_chinese_picture"
}

type ChineseBookInfo struct {
	Id       int    `json:"id"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
}

func (ChineseBookInfo) TableName() string {
	return "s_chinese_picture_info"
}

type AlbumPicture struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Position uint8  `json:"position"`
}

func (AlbumPicture) TableName() string {
	return "s_album_picture"
}

type AlbumPictureInfo struct {
	Id       int    `json:"id"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Duration string `json:"duration"`
}

func (AlbumPictureInfo) TableName() string {
	return "s_album_picture_info"
}

type ChineseBookTemp struct {
	Id        int    `json:"-"`
	BookId    string `json:"book_id"`
	BookIdOld int    `json:"book_id_old"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Type      int8   `json:"type"`
	Position  uint8  `json:"position"`
}

func (ChineseBookTemp) TableName() string {
	return "s_chinese_picture_temp"
}

type ChineseBookInfoTemp struct {
	Id        int    `json:"id"`
	BookId    string `json:"book_id"`
	BookIdOld int    `json:"book_id_old"`
	Mp3       string `json:"mp3"`
	Pic       string `json:"pic"`
	Position  int    `json:"position"`
}

func (ChineseBookInfoTemp) TableName() string {
	return "s_chinese_picture_info_temp"
}
