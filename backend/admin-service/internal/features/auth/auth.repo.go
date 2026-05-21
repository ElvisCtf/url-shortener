package auth

import (
	"gorm.io/gorm"
)

type AdminRepository struct {
	Db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{Db: db}
}

func (r *AdminRepository) Create(admin *Admin) error {
	return r.Db.Create(admin).Error
}

func (r *AdminRepository) GetByID(id uint) (*Admin, error) {
	var admin Admin
	if err := r.Db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) GetByEmail(email string) (*Admin, error) {
	var admin Admin
	if err := r.Db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) Update(admin *Admin) error {
	return r.Db.Save(admin).Error
}

func (r *AdminRepository) Delete(id uint) error {
	return r.Db.Delete(&Admin{}, id).Error
}
