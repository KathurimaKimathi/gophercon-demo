package dto

import (
	"io"

	validator "gopkg.in/go-playground/validator.v9"
)

// BusinessProfileInput is used to create a business profile
type BusinessProfileInput struct {
	Name           string    `json:"name"`
	Logo           io.Reader `form:"logo" json:"logo"`
	Description    string    `json:"description"`
	IntroStatement string    `json:"intro_statement"`
	Mission        string    `json:"mission"`
	Vision         string    `json:"vision"`
	Slogan         string    `json:"slogan"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	Facebook       string    `json:"facebook"`
	X              string    `json:"x"`
	Instagram      string    `json:"instagram"`
	TikTok         string    `json:"tiktok"`
	WhatsApp       string    `json:"whats_app"`
	LinkedIn       string    `json:"linked_in"`
	PostalAddress  string    `json:"postal_address"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	Building       string    `json:"building"`
	FloorNumber    string    `json:"floor_number"`
}

type PartnerInput struct {
	Heading              string                  `json:"heading"`
	HeadingSupportText   string                  `json:"heading_support_text"`
	BusinessProfile      string                  `json:"business_profile"`
	BusinessPartnerInput []*BusinessPartnerInput `json:"business_partner"`
}

// BusinessPartnerInput is used to register a business partner
type BusinessPartnerInput struct {
	Name string    `json:"name"`
	Logo io.Reader `form:"file" json:"file"`
}

type ContentInput struct {
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	Body               string      `json:"body"`
	HeroImage          io.Reader   `json:"hero_image"`
	GalleryImages      []io.Reader `json:"gallery_images,omitempty"`
	Category           string      `json:"category"`
	Tags               string      `json:"tags,omitempty"`
	ContentType        string      `json:"content_type"`
	Price              string      `json:"price,omitempty"`
	Media              io.Reader   `json:"media,omitempty"`
	BusinessProfile    string      `json:"business_profile"`
	StoryID            string      `json:"story_id"`
	Heading            string      `json:"heading"`
	HeadingSupportText string      `json:"heading_support_text"`
}

type OurServicesInput struct {
	Heading            string         `json:"heading"`
	HeadingSupportText string         `json:"heading_support_text"`
	BusinessProfile    string         `json:"business_profile"`
	Services           []ServiceInput `json:"services"`
}

// ServiceInput is used to create a service
type ServiceInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CommitmentInput struct {
	Heading            string       `json:"heading"`
	HeadingSupportText string       `json:"heading_support_text"`
	BusinessProfile    string       `json:"business_profile"`
	WhyUs              []WhyUsInput `json:"why_us"`
}

type WhyUsInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// PaginationsInput contains fields required for pagination
type PaginationsInput struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"currentPage" validate:"required"`
}

// Validate helps with validation of PaginationsInput fields
func (f *PaginationsInput) Validate() error {
	v := validator.New()

	err := v.Struct(f)

	return err
}
