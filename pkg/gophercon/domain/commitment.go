package domain

type Commitment struct {
	Heading            string   `json:"heading"`
	HeadingSupportText string   `json:"heading_support_text"`
	BusinessProfile    string   `json:"business_profile"`
	WhyUs              []*WhyUs `json:"why_us"`
}

type WhyUs struct {
	ID          *string `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}
