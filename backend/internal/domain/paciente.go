package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Paciente struct {
	ID uuid.UUID `json:"id"    gorm:"type:uuid;primaryKey"`

	Nombres         string     `json:"nombres"            gorm:"column:nombres;not null"`
	Apellidos       string     `json:"apellidos"          gorm:"column:apellidos;not null"`
	Nacionalidad    string     `json:"nacionalidad"       gorm:"column:nacionalidad;not null"`
	Genero          string     `json:"genero"             gorm:"column:genero;not null"`
	FechaNacimiento *time.Time `json:"fecha_nacimiento"   gorm:"column:fecha_nacimiento"`

	TipoIdentificacion string `json:"tipo_identificacion" gorm:"column:tipo_identificacion;not null"`
	Identificacion     string `json:"-"                   gorm:"column:identificacion;not null"`
	IdentificacionHash string `json:"-"                   gorm:"column:identificacion_hash;uniqueIndex"`
	TipoSangre         string `json:"tipo_sangre"         gorm:"column:tipo_sangre"`
	Telefono           string `json:"-"                   gorm:"column:telefono"`
	Email              string `json:"-"                   gorm:"column:email;not null"`
	EmailHash          string `json:"-"                   gorm:"column:email_hash;index:,unique,where:email_hash IS NOT NULL AND email_hash != ''"`
	Direccion          string `json:"-"                   gorm:"column:direccion"`
	//FotoURL            	string		`json:"foto_url"          gorm:"column:foto_url"`

	CreatedAt time.Time      `json:"created_at"           gorm:"column:created_at; not null"`
	UpdatedAt time.Time      `json:"updated_at"           gorm:"column:updated_at; not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"           gorm:"column:deleted_at;index"`
}

// BeforeCreate intercepta la inserción en GORM para generar un UUID v7 si el ID está vacío.
func (p *Paciente) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}
		p.ID = newUUID
	}
	return nil
}

func (p *Paciente) SensitiveFields() []*string {
	return []*string{
		&p.Nombres,
		&p.Apellidos,
		&p.Identificacion,
		&p.Email,
		&p.Telefono,
		&p.Direccion,
	}
}

func (p *Paciente) EntityName() string { return "paciente" }
