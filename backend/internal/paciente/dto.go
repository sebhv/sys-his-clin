package paciente

import "time"

type CreatePacienteRequest struct {
	Nombres            string `json:"nombres"                    validate:"required,min=2,max=80"`
	Apellidos          string `json:"apellidos"                  validate:"required,min=2,max=80"`
	Genero             string `json:"genero"                     validate:"required,oneof=masculino femenino otro"`
	Nacionalidad       string `json:"nacionalidad"               validate:"required,min=2,max=30"`
	TipoIdentificacion string `json:"tipo_identificacion"        validate:"required,oneof=cedula pasaporte otro"`
	Identificacion     string `json:"identificacion"             validate:"required,min=5,max=20"`
	Telefono           string `json:"telefono,omitempty"         validate:"omitempty,min=10,max=15"`
	Email              string `json:"email"                      validate:"required,email,max=120"`
	Direccion          string `json:"direccion,omitempty"        validate:"omitempty,max=240"`
	TipoSangre         string `json:"tipo_sangre,omitempty"      validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O- desconocido"`
	FechaNacimiento    string `json:"fecha_nacimiento,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

type UpdatePacienteRequest struct {
	Nombres         *string `json:"nombres,omitempty"          validate:"omitempty,min=2,max=80"`
	Apellidos       *string `json:"apellidos,omitempty"        validate:"omitempty,min=2,max=80"`
	Genero          *string `json:"genero,omitempty"           validate:"omitempty,oneof=masculino femenino otro"`
	Nacionalidad    *string `json:"nacionalidad,omitempty"     validate:"omitempty,min=2,max=30"`
	Telefono        *string `json:"telefono,omitempty"         validate:"omitempty,min=10,max=15"`
	Direccion       *string `json:"direccion,omitempty"        validate:"omitempty,max=240"`
	TipoSangre      *string `json:"tipo_sangre,omitempty"      validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O- desconocido"`
	FechaNacimiento *string `json:"fecha_nacimiento,omitempty" validate:"omitempty,datetime=2006-01-02"`
	//FotoURL         *string    `json:"foto_url"`

}

type PacienteResponse struct {
	ID                 string    `json:"id"`
	Nombres            string    `json:"nombres"`
	Apellidos          string    `json:"apellidos"`
	Genero             string    `json:"genero"`
	Nacionalidad       string    `json:"nacionalidad"`
	TipoIdentificacion string    `json:"tipo_identificacion"`
	Identificacion     string    `json:"identificacion,omitempty"`
	Telefono           string    `json:"telefono,omitempty"`
	Email              string    `json:"email,omitempty"`
	Direccion          string    `json:"direccion,omitempty"`
	TipoSangre         string    `json:"tipo_sangre,omitempty"`
	FechaNacimiento    string    `json:"fecha_nacimiento,omitempty"`
	CreatedAt          time.Time `json:"creado_en"`
	UpdatedAt          time.Time `json:"actualizado_en"`
}

//type PacienteDetailResponse struct {
//	PacienteResponse
//
//	Contactos []contacto.ContactoResponse   `json:"contactos,omitempty"`
//}
