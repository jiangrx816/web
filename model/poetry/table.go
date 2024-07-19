package poetry

type RXDynasty struct {
	Id        int    `json:"id"`
	DynastyId int    `json:"dynasty_id"`
	Name      string `json:"name"`
	Intro     string `json:"intro"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
}

func (RXDynasty) TableName() string {
	return "rx_dynasty"
}

type RXAuthors struct {
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

func (RXAuthors) TableName() string {
	return "rx_authors"
}

type RXCollectionKinds struct {
	Id               int    `json:"id"`
	CollectionKindId int    `json:"collection_kind_id"`
	Name             string `json:"name"`
	Limit            int    `json:"limit"`
	Sort             int    `json:"sort"`
}

func (RXCollectionKinds) TableName() string {
	return "rx_collection_kinds"
}

type RXCollections struct {
	Id           int    `json:"id"`
	CollectionId int    `json:"collection_id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	ShortDesc    string `json:"short_desc"`
	Cover        string `json:"cover"`
	Kind         string `json:"kind"`
	KindId       int    `json:"kind_id"`
	QuotesCount  int    `json:"quotes_count"`
	WorksCount   int    `json:"works_count"`
	Sort         int    `json:"sort"`
}

func (RXCollections) TableName() string {
	return "rx_collections"
}

type RXCollectionQuotes struct {
	Id               int    `json:"id"`
	Sort             int    `json:"sort"`
	QuoteId          int    `json:"quote_id"`
	Quote            string `json:"quote"`
	QuoteAuthor      string `json:"quote_author"`
	QuoteWork        string `json:"quote_work"`
	QuoteWorkId      int    `json:"quote_work_id"`
	CollectionId     int    `json:"collection_id"`
	CollectionKindId int    `json:"collection_kind_id"`
}

func (RXCollectionQuotes) TableName() string {
	return "rx_collection_quotes"
}

type RXQuotes struct {
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

func (RXQuotes) TableName() string {
	return "rx_quotes"
}

type RXCollectionWorks struct {
	Id               int    `json:"id"`
	Sort             int    `json:"sort"`
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

func (RXCollectionWorks) TableName() string {
	return "rx_collection_works"
}

type RXWorks struct {
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

func (RXWorks) TableName() string {
	return "rx_works"
}
