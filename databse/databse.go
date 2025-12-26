package databse

type ItemList struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
}
