package paciente

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sebhv/sys-his-clin/backend/internal/domain"
)

type Repository interface {
	Create(p *domain.Paciente) error
	FindByID(id uuid.UUID) (*domain.Paciente, error)
	//FindFullByID(id uuid.UUID) (*domain.Paciente, error)
	FindAll(page, limit int) ([]domain.Paciente, int64, error)
	Update(p *domain.Paciente) error
	Deactivate(id uuid.UUID) error

	ExistsByIdentificacion(identificacionHash string) (bool, error)
	ExistsByEmail(emailHash string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Paciente

func (r *repository) Create(p *domain.Paciente) error {
	return r.db.Create(p).Error
}

func (r *repository) FindByID(id uuid.UUID) (*domain.Paciente, error) {
	var p domain.Paciente

	err := r.db.First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

/*
func (r *repository) FindFullByID(id uuid.UUID) (*domain.Paciente, error) {
	var p domain.Paciente
	err := r.db.
		//Preload("ContactosEmergencia").
		//Preload("Documentos").
		First(&p, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &p, nil
}
*/

func (r *repository) FindAll(page, limit int) ([]domain.Paciente, int64, error) {
	var pacientes []domain.Paciente
	var total int64

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := r.db.Model(&domain.Paciente{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&pacientes).Error; err != nil {
		return nil, 0, err
	}

	return pacientes, total, nil
}

func (r *repository) Update(p *domain.Paciente) error {

	result := r.db.Save(p)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) Deactivate(id uuid.UUID) error {
	result := r.db.Delete(&domain.Paciente{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) ExistsByIdentificacion(identificacionHash string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Paciente{}).
		Where("identificacion_hash = ?", identificacionHash).
		Count(&count).Error
	return count > 0, err
}

func (r *repository) ExistsByEmail(emailHash string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Paciente{}).
		Where("email_hash = ?", emailHash).
		Count(&count).Error
	return count > 0, err
}
