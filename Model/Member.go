package model

import (
	"github.com/zaidanpoin/crud-golang-react/database"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name  string `json:"name" binding:"required" form:"name"`
	Image string `json:"image" form:"image" `
	Url   string `json:"url" form:"url" binding:"required"`
}

func (m *Member) Save() (*Member, error) {

	err := database.Database.Create(&m)

	if err.Error != nil {
		return &Member{}, err.Error
	}

	return m, nil

}

func (m *Member) GetDataMembers(id string) ([]Member, error) {

	if id != "" {

		var member []Member

		if err := database.Database.Find(&member, id).Error; err != nil {
			return []Member{}, err
		}

		return member, nil
	}

	var members []Member

	if err := database.Database.Find(&members).Error; err != nil {
		return []Member{}, err
	}

	return members, nil
}

func (m *Member) DeleteMember(id string) error {
	if err := database.Database.Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *Member) UpdateMember(id string) error {

	if err := database.Database.Model(&Member{}).Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}

	return nil
}
