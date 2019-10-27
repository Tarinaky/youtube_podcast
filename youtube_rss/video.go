package youtube_rss

type Video struct {
	Title       string `json:"title"`
	VideoID     string `json:"display_id"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	Uploader    string `json:"uploader"`
	UploadDate  string `json:"upload_date"`
}
