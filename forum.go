package esforum

type Post struct {
	Id      int64  `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Thread  int64  `json:"thread"`
}

type Thread struct {
	Id    int64  `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title string `json:"title"`
}
