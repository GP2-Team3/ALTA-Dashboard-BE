package data

import (
	"alta-dashboard-be/features/mentee"
	"alta-dashboard-be/features/mentee/models"
	"alta-dashboard-be/utils/helper"
	"errors"

	"gorm.io/gorm"
)

type MenteeData struct {
	db *gorm.DB
}

// Create implements mentee.MenteeData_
func (u *MenteeData) Create(mentee mentee.MenteeCore) error {
	//start transaction
	tx := u.db.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New("internal database transaction error")
	}

	//convert core to models
	m, emergency, edu := convertToModels(&mentee)

	//insert to mentee
	tx = tx.Create(&m)
	if tx.Error != nil || tx.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("internal database create mentee error")
	}

	//set menteeID to lastInsertedID
	if m.ID == 0 {
		tx.Rollback()
		return errors.New("invalid lastInsertedID")
	}
	emergency.MenteeID = int(m.ID)
	edu.MenteeID = int(m.ID)

	//insert to emergency
	tx = tx.Create(&emergency)
	if tx.Error != nil || tx.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("internal database create emergency error")
	}

	//insert to education
	tx = tx.Create(&edu)
	if tx.Error != nil || tx.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("internal database create education error")
	}

	tx = tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New("internal database commit error")
	}

	return nil
}

// Delete implements mentee.MenteeData_
func (u *MenteeData) Delete(id int) error {
	var mdl models.Mentee
	tx := u.db.Where("id = ?", id).Delete(&mdl)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error in database")
	}
	return nil
}

// GetAll implements mentee.MenteeData_
func (u *MenteeData) GetAll(page int, limit int) ([]mentee.MenteeCore, error) {
	mdl := []models.Mentee{}
	limit, offset := helper.LimitOffsetConvert(page, limit)
	tx := u.db.Offset(offset).Limit(limit).Find(&mdl)
	if tx.Error != nil {
		return nil, errors.New("error in database")
	}
	return convertToCoreList(mdl), nil
}

// GetOne implements mentee.MenteeData_
func (u *MenteeData) GetOne(id int) (mentee.MenteeCore, error) {
	mdl := models.Mentee{}
	tx := u.db.Where("id = ?", id).First(&mdl)
	if tx.Error != nil {
		return mentee.MenteeCore{}, errors.New("error in database")
	}
	return convertToCore(&mdl), nil
}

// Update implements mentee.MenteeData_
func (u *MenteeData) Update(id int, mentee mentee.MenteeCore) error {
	//start transaction
	tx := u.db.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New("internal database transaction error")
	}

	//convert core to models
	m, emergency, edu := convertToModels(&mentee)

	//update to mentee
	tx = tx.Where("id = ?", id).Updates(&m)
	if tx.Error != nil || tx.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("internal database create mentee error")
	}

	//set menteeID to lastInsertedID
	if m.ID == 0 {
		tx.Rollback()
		return errors.New("invalid lastInsertedID")
	}
	emergency.MenteeID = int(m.ID)
	edu.MenteeID = int(m.ID)

	//update to emergency
	tx = tx.Where("mentee_id = ?", emergency.MenteeID).Updates(&emergency)
	if tx.Error != nil || tx.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("internal database create emergency error")
	}

	//update to education
	tx = tx.Where("mentee_id = ?", edu.MenteeID).Updates(&edu)
	if tx.Error != nil || tx.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("internal database create education error")
	}

	tx = tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New("internal database commit error")
	}

	return nil
}

func New(db *gorm.DB) mentee.MenteeData_ {
	return &MenteeData{
		db: db,
	}
}