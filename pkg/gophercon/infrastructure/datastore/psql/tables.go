package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Base model contains defines common fields across tables
type Base struct {
	ID        *string    `gorm:"primaryKey;unique;column:id"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	CreatedBy *string    `gorm:"column:created_by"`
	UpdatedBy *string    `gorm:"column:updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Active    bool       `gorm:"column:active"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	b.ID = &id

	b.Active = true

	return
}

type User struct {
	Base

	Username  string `gorm:"column:username"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	UserType  string `gorm:"column:user_type"`
}

func (User) TableName() string {
	return "users_user"
}

type UserCredentials struct {
	Base

	ValidFrom time.Time `gorm:"column:valid_from"`
	ValidTo   time.Time `gorm:"column:valid_to"`
	HashedPin string    `gorm:"column:hashed_pin"`
	Salt      string    `gorm:"column:salt"`
	UserID    *string   `gorm:"column:user_id"`
}

func (UserCredentials) TableName() string {
	return "user_credentials"
}

type BusinessProfile struct {
	Base

	Name           string `gorm:"column:name"`
	Logo           string `gorm:"column:logo"`
	Description    string `gorm:"column:description"`
	IntroStatement string `gorm:"column:intro_statement"`
	Mission        string `gorm:"column:mission"`
	Vision         string `gorm:"column:vision"`
	Slogan         string `gorm:"column:slogan"`
	PhoneNumber    string `gorm:"column:phone_number"`
	Email          string `gorm:"column:email"`
	Facebook       string `gorm:"column:facebook"`
	X              string `gorm:"column:x"`
	Instagram      string `gorm:"column:instagram"`
	TikTok         string `gorm:"column:tiktok"`
	WhatsApp       string `gorm:"column:whats_app"`
	LinkedIn       string `gorm:"column:linkedin"`
	PostalAddress  string `gorm:"column:postal_address"`
	City           string `gorm:"column:city"`
	Country        string `gorm:"column:country"`
	Building       string `gorm:"column:building"`
	FloorNumber    string `gorm:"column:floor_number"`

	Partner    *Partner     `gorm:"ForeignKey:business_profile;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Services   *BaseService `gorm:"ForeignKey:business_profile;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Commitment *Commitment  `gorm:"ForeignKey:business_profile;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (BusinessProfile) TableName() string {
	return "business_profile"
}

type Story struct {
	Base
	CommonModel

	BusinessProfile string     `json:"business_profile"`
	Content         []*Content `gorm:"ForeignKey:story_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Story) TableName() string {
	return "story"
}

type Content struct {
	Base

	Title         string         `gorm:"column:title"`
	Description   string         `gorm:"column:description"`
	Body          string         `gorm:"column:body"`
	HeroImage     string         `gorm:"column:hero_image"`
	GalleryImages pq.StringArray `gorm:"type:text[];column:gallery_images"`
	Category      string         `gorm:"column:category"`
	Tags          pq.StringArray `gorm:"type:text[];column:tags"`
	ContentType   string         `gorm:"column:content_type"`
	Media         string         `gorm:"column:media"`
	Price         float64        `gorm:"column:price"`
	Story         *string        `gorm:"column:story_id"`
}

func (Content) TableName() string {
	return "content"
}

type Partner struct {
	Base
	CommonModel

	BusinessProfile  string             `json:"business_profile"`
	BusinessPartners []*BusinessPartner `gorm:"ForeignKey:PartnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Partner) TableName() string {
	return "partner"
}

// BusinessPartner models a business partner
type BusinessPartner struct {
	Base

	Name      string  `gorm:"column:name"`
	Logo      string  `gorm:"column:logo"`
	PartnerID *string `gorm:"column:partner_id"`
}

func (BusinessPartner) TableName() string {
	return "business_partner"
}

type CommonModel struct {
	Heading            string `gorm:"heading"`
	HeadingSupportText string `json:"heading_support_text" gorm:"heading_support_text"`
}

type BaseService struct {
	Base
	CommonModel

	BusinessProfile string     `json:"business_profile"`
	Services        []*Service `gorm:"ForeignKey:base_service;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (BaseService) TableName() string {
	return "base_service"
}

// Service models the available services
type Service struct {
	Base

	Title       string  `gorm:"title"`
	Description string  `gorm:"description"`
	BaseService *string `gorm:"base_service"`
}

func (Service) TableName() string {
	return "service"
}

type Commitment struct {
	Base
	CommonModel

	BusinessProfile string   `json:"business_profile"`
	WhyUs           []*WhyUs `gorm:"ForeignKey:commitment;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Commitment) TableName() string {
	return "commitment"
}

type WhyUs struct {
	Base

	Title       string  `gorm:"title"`
	Description string  `gorm:"description"`
	Commitment  *string `gorm:"commitment"`
}

func (WhyUs) TableName() string {
	return "why_us"
}
