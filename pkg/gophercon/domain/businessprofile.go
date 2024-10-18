package domain

// BusinessProfile is used to create a business profile
type BusinessProfile struct {
	ID             *string `json:"id"`
	Name           string  `json:"name"`
	Logo           string  `json:"logo"`
	Description    string  `json:"description"`
	IntroStatement string  `json:"intro_statement"`
	Mission        string  `json:"mission"`
	Vision         string  `json:"vision"`
	Slogan         string  `json:"slogan"`
	PhoneNumber    string  `json:"phone_number"`
	Email          string  `json:"email"`
	Facebook       string  `json:"facebook"`
	X              string  `json:"x"`
	Instagram      string  `json:"instagram"`
	TikTok         string  `json:"tiktok"`
	WhatsApp       string  `json:"whats_app"`
	LinkedIn       string  `json:"linked_in"`
	PostalAddress  string  `json:"postal_address"`
	City           string  `json:"city"`
	Country        string  `json:"country"`
	Building       string  `json:"building"`
	FloorNumber    string  `json:"floor_number"`

	Partners   Partner     `json:"partners,omitempty"`
	Services   BaseService `json:"services,omitempty"`
	Commitment Commitment  `json:"commitment,omitempty"`
	Solutions  Story       `json:"solutions,omitempty"`
}

type Partner struct {
	Heading            string             `json:"heading,omitempty"`
	HeadingSupportText string             `json:"heading_support_text,omitempty"`
	BusinessProfile    string             `json:"business_profile,omitempty"`
	BusinessPartner    []*BusinessPartner `json:"business_partner,omitempty"`
}

// BusinessPartner creates a business partner
type BusinessPartner struct {
	ID      *string `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Logo    string  `json:"logo,omitempty"`
	Partner string  `json:"partner,omitempty"`
}

// Pagination contains the struct fields for performing pagination.
type Pagination struct {
	Limit        int   `json:"limit"`
	CurrentPage  int   `json:"currentPage"`
	Count        int64 `json:"count"`
	TotalPages   int   `json:"totalPages"`
	NextPage     *int  `json:"nextPage"`
	PreviousPage *int  `json:"previousPage"`
}

// GetOffset calculates the deviation in pages that come before
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit calculates the maximum number of items to be shown per page
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

// GetPage gets the current page
func (p *Pagination) GetPage() int {
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}

	return p.CurrentPage
}
