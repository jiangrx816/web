package poetry

type DynastyResponse struct {
	List []ResponseDynasty `json:"list"`
}

type ResponseDynasty struct {
	DynastyName string `json:"dynasty_name"`
}

type AuthorResponse struct {
	Page      int          `json:"page"`
	Total     int64        `json:"total"`
	TotalPage int          `json:"total_page"`
	List      []AuthorItem `json:"list"`
}

type AuthorInfoResponse struct {
	Info AuthorItem `json:"info"`
}

type AuthorItem struct {
	Id            int    `json:"id"`
	AuthorId      int    `json:"author_id"`
	Name          string `json:"name"`
	Intro         string `json:"intro"`
	Dynasty       string `json:"dynasty"`
	BirthYear     string `json:"birth_year"`
	DeathYear     string `json:"death_year"`
	Wiki          string `json:"wiki"`
	QuotesCount   int    `json:"quotes_count"`
	WorksCount    int    `json:"works_count"`
	WorksCountShi int    `json:"works_count_shi"`
	WorksCountCi  int    `json:"works_count_ci"`
	WorksCountWen int    `json:"works_count_wen"`
	WorksCountQu  int    `json:"works_count_qu"`
	WorksCountFu  int    `json:"works_count_fu"`
}

type KindResponse struct {
	List []CollectionKindItem `json:"list"`
}

type CollectionKindItem struct {
	CollectionKindId int    `json:"collection_kind_id"`
	Name             string `json:"name"`
}

type CollectionResponse struct {
	List []CollectionItem `json:"list"`
}

type CollectionItem struct {
	CollectionId int    `json:"collection_id"`
	Name         string `json:"name"`
}

type CollectionQuotesResponse struct {
	Page      int                    `json:"page"`
	Total     int64                  `json:"total"`
	TotalPage int                    `json:"total_page"`
	List      []CollectionQuotesItem `json:"list"`
}

type CollectionQuotesItem struct {
	QuoteId          int    `json:"quote_id"`
	Quote            string `json:"quote"`
	QuoteAuthor      string `json:"quote_author"`
	QuoteWork        string `json:"quote_work"`
	QuoteWorkId      int    `json:"quote_work_id"`
	CollectionId     int    `json:"collection_id"`
	CollectionKindId int    `json:"collection_kind_id"`
	Collection       string `json:"collection"`
	Kind             string `json:"kind"`
}

type QuotesResponse struct {
	Info QuotesItem `json:"info"`
}

type QuotesItem struct {
	Id        int    `json:"id"`
	QuoteId   int    `json:"quote_id"`
	Quote     string `json:"quote"`
	Dynasty   string `json:"dynasty"`
	Author    string `json:"author"`
	AuthorId  int    `json:"author_id"`
	Kind      string `json:"kind"`
	WorkId    int    `json:"work_id"`
	WorkTitle string `json:"work_title"`
}

type WorksResponse struct {
	Info WorksItem `json:"info"`
}

type WorksItem struct {
	Id               int    `json:"id"`
	WorkId           int    `json:"work_id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	AuthorId         int    `json:"author_id"`
	Dynasty          string `json:"dynasty"`
	Kind             string `json:"kind"`
	KindCn           string `json:"kind_cn"`
	Wiki             string `json:"wiki"`
	Content          string `json:"content"`
	ShortContent     string `gorm:"-" json:"short_content"`
	Intro            string `json:"intro"`
	Annotation       string `json:"annotation"`
	Translation      string `json:"translation"`
	MasterComment    string `json:"master_comment"`
	AuthorWorksCount int    `json:"author_works_count"`
	QuotesCount      int    `json:"quotes_count"`
	CollectionsCount int    `json:"collections_count"`
}

type CollectionWorksResponse struct {
	Page      int                   `json:"page"`
	Total     int64                 `json:"total"`
	TotalPage int                   `json:"total_page"`
	List      []CollectionWorksItem `json:"list"`
}

type CollectionWorksItem struct {
	Id               int    `json:"id"`
	CollectionWorkId int    `json:"collection_work_id"`
	WorkId           int    `json:"work_id"`
	CollectionId     int    `json:"collection_id"`
	WorkTitle        string `json:"work_title"`
	WorkAuthor       string `json:"work_author"`
	WorkDynasty      string `json:"work_dynasty"`
	WorkContent      string `json:"work_content"`
	WorkKind         string `json:"work_kind"`
	Collection       string `json:"collection"`
}
