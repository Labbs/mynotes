package models

import (
	"database/sql/driver"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gosimple/slug"
	"github.com/labbs/mynotes/internal/shortuuid"
	"gorm.io/gorm"
)

// Document is a model for a document
type Document struct {
	Id   string       `json:"id"`
	Name string       `json:"name"`
	Slug string       `json:"slug"`
	Type DocumentType `json:"type"`

	Favorite bool `gorm:"-" json:"favorite"`

	Config     DocumentConfig `json:"config"`
	Metadata   JSONB          `json:"metadata"`
	ParentId   string         `json:"parent_id"`
	Properties Properties     `json:"properties"`

	SpaceId string `json:"space_id"`

	Content string `json:"content"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// TableName returns the name of the table
func (d Document) TableName() string {
	return "document"
}

// BeforeCreate is a hook that runs before creating a document
func (d *Document) BeforeCreate(tx *gorm.DB) error {
	d.Id = utils.UUIDv4()
	d.Slug = slug.Make(d.Name + "-" + shortuuid.GenerateShortUUID())
	return nil
}

// DocumentType is the type of document
type DocumentType string

// DocumentType constants
const (
	DocumentTypeDocument DocumentType = "document"
)

// Properties is a list of properties for a document
type Properties []Propertie

// Propertie is a property for a document
type Propertie struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Order int    `json:"order"`
}

// Properties implements the driver.Valuer interface
func (p Properties) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Properties implements the sql.Scanner interface
func (p *Properties) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		// PostgreSQL usually returns []byte
		return json.Unmarshal(v, p)
	case string:
		// SQLite often returns string
		return json.Unmarshal([]byte(v), p)
	case nil:
		// Handle null case
		*p = Properties{}
		return nil
	default:
		// Fall back to string conversion
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, p)
	}
}

// DocumentConfig is the configuration for a document
type DocumentConfig struct {
	FullWidth        bool   `json:"full_width"`
	Icon             string `json:"icon,omitempty"`
	Lock             bool   `json:"lock"`
	HeaderBackground string `json:"header_background"`
}

// Value implements the driver.Valuer interface
func (dc DocumentConfig) Value() (driver.Value, error) {
	return json.Marshal(dc)
}

// Scan implements the sql.Scanner interface
func (dc *DocumentConfig) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		// PostgreSQL usually returns []byte
		return json.Unmarshal(v, dc)
	case string:
		// SQLite often returns string
		return json.Unmarshal([]byte(v), dc)
	case nil:
		// Handle null case
		*dc = DocumentConfig{}
		return nil
	default:
		// Fall back to string conversion
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, dc)
	}
}

// DocumentRepository is the repository for documents
type DocumentRepository interface {
	CreateDocument(document Document) (Document, error)
	GetDocumentsFirstLevelForSpace(spaceId string) ([]Document, error)
	GetDocumentBySlug(slug string, preloadBlock bool) (Document, error)
	GetDocumentById(id string, preloadBlock bool) (Document, error)
	UpdateDocument(document Document) (Document, error)
}

// DocumentService is the service for documents
type DocumentService interface {
	CreateDocument(document Document) (Document, error)
	GetDocumentsFirstLevelForSpace(spaceId string) ([]Document, error)
	GetDocumentBySlug(slug string, preloadBlock bool) (Document, error)
	GetDocumentById(id string, preloadBlock bool) (Document, error)
	UpdateDocument(document Document) (Document, error)
}
