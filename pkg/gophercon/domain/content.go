package domain

type Story struct {
	Heading            string     `json:"heading"`
	HeadingSupportText string     `json:"heading_support_text"`
	BusinessProfile    string     `json:"business_profile"`
	Content            []*Content `json:"content"`
}

type Content struct {
	ID            *string  `json:"id"`
	Active        bool     `json:"active"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Body          string   `json:"body"`
	HeroImage     string   `json:"hero_image"`
	GalleryImages []string `json:"gallery_images,omitempty"`
	Category      string   `json:"category"`
	Tags          []string `json:"tags,omitempty"`
	ContentType   string   `json:"content_type"`
	Media         string   `json:"media,omitempty"`
	Price         float64  `json:"price,omitempty"`
	StoryID       string   `json:"story_id,omitempty"`
}

type ContentPage struct {
	Story      *Story     `json:"story"`
	Pagination Pagination `json:"pagination"`
}
