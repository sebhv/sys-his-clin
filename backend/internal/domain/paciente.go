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
	EmailHash          string `json:"-"                   gorm:"column:email_hash;uniqueIndex"`
	Direccion          string `json:"-"                   gorm:"column:direccion"`
	//FotoURL            	string		`json:"foto_url"          gorm:"column:foto_url"`

	CreadoEn      time.Time      `json:"-"     gorm:"column:fecha_registro;autoCreateTime"`
	ActualizadoEn time.Time      `json:"-"     gorm:"column:fecha_actualizacion;autoUpdateTime"`
	EliminadoEn   gorm.DeletedAt `json:"-"     gorm:"column:fecha_eliminacion;index"`
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
