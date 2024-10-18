package domain

type BaseService struct {
	Heading            string     `json:"heading,omitempty"`
	HeadingSupportText string     `json:"heading_support_text,omitempty"`
	BusinessProfile    string     `json:"business_profile,omitempty"`
	Services           []*Service `json:"services,omitempty"`
}

// Service models the services offered by the business
type Service struct {
	ID          *string `json:"id,omitempty"`
	Active      bool    `json:"active,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
}
