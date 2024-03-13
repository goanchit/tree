package models

import (
	"time"

	"gorm.io/gorm"
)

type sexType string

const (
	MALE   sexType = "male"
	FEMALE sexType = "female"
)

type UserBase struct {
	gorm.Model
	Username string             `gorm:"unique"`
	Family   []UserFamily       `gorm:"foreignKey:UserBaseID"`
	Relation []UserRelationship `gorm:"foreignKey:UserBaseID"`
}

type UserFamily struct {
	gorm.Model
	UserBaseID uint
	CreatedAt  time.Time
	Name       string
	Sex        sexType
}

type UserRelationship struct {
	gorm.Model
	UserBaseID uint
	CreatedAt  time.Time
	Name       string
}

type Relations struct {
	gorm.Model
	UserBaseID     uint
	Member         UserFamily       `gorm:"foreignKey:MemberID"`      // First Dependent
	Relationship   UserRelationship `gorm:"foreignKey:RelationshipID` // This defines the type of relation between two people
	Relative       UserFamily       `gorm:"foreignKey:RelativeID"`    // Second Dependent
	MemberID       uint
	RelationshipID uint
	RelativeID     uint
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserBase{}, &UserFamily{}, &UserRelationship{}, &Relations{}); err != nil {
		return err
	}
	return nil
}
