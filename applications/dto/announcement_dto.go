package dto

type Announcement struct {
	ID      uint64 `json:"id"`
	From    string `json:"from"`
	Message string `json:"message"`
	Seconds int64  `json:"seconds"`
}
