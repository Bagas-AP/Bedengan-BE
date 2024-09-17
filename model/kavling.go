package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kavling struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey;default:(UUID())"`
	Nama      string         `json:"nama" gorm:"type:varchar(255);not null"`
	GroundID  string         `json:"ground_id" gorm:"type:varchar(36);not null"`
	Harga     int            `json:"harga" gorm:"type:bigint;not null"`                   // Sesuaikan tipe harga ke bigint
	Status    string         `json:"status" gorm:"type:text;not null;default:'tersedia'"` // GORM default
	Reservasi []Reservasi    `json:"reservasi" gorm:"foreignKey:KavlingID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Kavling) TableName() string {
	return "kavlings"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Kavling) BeforeCreate() error {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (p *Kavling) BeforeSave() error {
	p.UpdatedAt = time.Now()
	return nil
}

type KavlingInput struct {
	Nama     string `json:"nama" binding:"required"`
	GroundID string `json:"ground_id" binding:"required"`
	Harga    int    `json:"harga" binding:"required"`
}

func (p *KavlingInput) ToKavling() *Kavling {
	return &Kavling{
		Nama:     p.Nama,
		GroundID: p.GroundID,
		Harga:    p.Harga,
	}
}
