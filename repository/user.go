package repository

import (
	"errors"
	"family-tree/config"
	"family-tree/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func NewUserRepository(username string) *UserRepository {
	return &UserRepository{
		Username: username,
	}
}

type UserRepository struct {
	Username string
}

func (u UserRepository) LoginIntoUser() error {

	db := config.OpenDb()
	defer func() {
		dbSql, _ := db.DB()
		dbSql.Close()
	}()

	err := models.AutoMigrate(db)
	if err != nil {
		panic(err)
	}

	user := models.UserBase{
		Username: u.Username,
	}

	result := db.FirstOrCreate(&user, models.UserBase{Username: u.Username})

	if result.RowsAffected == 1 {
		log.Println("New user created")
	} else {
		log.Println("Logged in into current user")
	}
	return nil
}

func (u UserRepository) AddPerson(personName string, sex string) error {
	db := config.OpenDb()
	defer func() {
		dbSql, _ := db.DB()
		dbSql.Close()
	}()

	// Fetch the UserBase record based on username
	var user models.UserBase
	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return err
	}

	sex_type := models.FEMALE
	if sex == "male" {
		sex_type = models.MALE
	}

	family := models.UserFamily{
		UserBaseID: user.ID,
		Name:       personName,
		Sex:        sex_type,
		CreatedAt:  time.Now(),
	}

	if err := db.Create(&family).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) AddRelation(relationName string) error {
	db := config.OpenDb()
	defer func() {
		dbSql, _ := db.DB()
		dbSql.Close()
	}()

	// Fetch the UserBase record based on username
	var user models.UserBase
	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return err
	}

	relations := models.UserRelationship{
		UserBaseID: user.ID,
		Name:       relationName,
		CreatedAt:  time.Now(),
	}

	if err := db.Create(&relations).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) AttachRelationship(member string, relationtype string, relative string) error {
	db := config.OpenDb()
	defer func() {
		dbSql, _ := db.DB()
		dbSql.Close()
	}()

	var (
		m1   models.UserFamily
		m2   models.UserFamily
		rel  models.UserRelationship
		user models.UserBase
	)

	// Fetch the UserBase record based on username
	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return err
	}

	if err := db.Where("name = ?", member).First(&m1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("member with specific name not found/listed under the user provided")
		}
		return err
	}
	if err := db.Where("name = ?", relative).First(&m2).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("member with specific name not found/listed under the user provided")
		}
		return err
	}

	if err := db.Where("name = ?", relationtype).First(&rel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("relation with specific name not found/listed under the user provided")
		}
		return err
	}

	relationship := models.Relations{
		MemberID:       m1.ID,
		RelativeID:     m2.ID,
		RelationshipID: rel.ID,
		UserBaseID:     user.ID,
	}

	if err := db.Create(&relationship).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) FindRecords(relation string, member string, byDependent bool) ([]models.Relations, error) {
	db := config.OpenDb()

	defer func() {
		dbSql, _ := db.DB()
		dbSql.Close()
	}()

	var (
		m1       models.UserFamily
		rel      models.UserRelationship
		user     models.UserBase
		userRels []models.Relations
	)

	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userRels, errors.New("user not found")
		}
		return userRels, err
	}

	if err := db.Where("name = ? AND user_base_id = ?", member, user.ID).First(&m1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userRels, errors.New("member with specific name not found/listed under the user provided")
		}
		return userRels, err
	}

	if err := db.Where("name = ? AND user_base_id = ?", relation, user.ID).First(&rel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userRels, errors.New("relation with specific name not found/listed under the user provided")
		}
		return userRels, err
	}

	qry := &models.Relations{
		UserBaseID:     user.ID,
		RelationshipID: rel.ID,
		MemberID:       m1.ID,
	}

	if byDependent {
		qry = &models.Relations{
			UserBaseID:     user.ID,
			RelationshipID: rel.ID,
			RelativeID:     m1.ID,
		}
	}

	err := db.Where(qry).Find(&userRels).Error
	if err != nil {
		return userRels, err
	}

	return userRels, nil

}

func (u UserRepository) FindById(id_list []uint) ([]models.UserFamily, error) {
	db := config.OpenDb()

	defer func() {
		dbSql, _ := db.DB()
		dbSql.Close()
	}()

	var (
		m1   []models.UserFamily
		user models.UserBase
	)

	if err := db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return m1, errors.New("user not found")
		}
		return m1, err
	}

	err := db.Where("user_base_id = ? AND id in ?", user.ID, id_list).Find(&m1).Error
	if err != nil {
		return m1, errors.New("failed to get results")
	}

	return m1, nil
}
