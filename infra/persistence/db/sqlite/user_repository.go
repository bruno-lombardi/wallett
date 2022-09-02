package sqlite

import (
	"math"
	"wallett/domain/models"
	"wallett/infra/generators"

	"gorm.io/gorm"
)

type SQLiteUser struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:VARCHAR(12);not null;unique"`
	Email    string `gorm:"unique"`
	Name     string `gorm:"type:VARCHAR(255)"`
	Password string `gorm:"type:VARCHAR(128)"`
}

type SQLiteUserRepository struct {
	db *gorm.DB
}

func NewSQLiteUserRepository(db *gorm.DB) *SQLiteUserRepository {
	db.AutoMigrate(&SQLiteUser{})
	return &SQLiteUserRepository{
		db: db,
	}
}

func (r *SQLiteUserRepository) Create(createUserDto *models.CreateUserDTO) (*models.User, error) {
	user := &SQLiteUser{
		ID:       generators.ID("u"),
		Email:    createUserDto.Email,
		Name:     createUserDto.Name,
		Password: createUserDto.Password,
	}
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
		}, nil
	}
}

func (r *SQLiteUserRepository) Get(ID string) (*models.User, error) {
	var user SQLiteUser
	result := r.db.Model(&SQLiteUser{ID: ID}).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
		}, nil
	}
}

func (r *SQLiteUserRepository) List(listUsersDto *models.ListUsersDTO) (*models.PaginatedUserResultDTO, error) {
	switch {
	case listUsersDto.Limit > 100:
		listUsersDto.Limit = 100
	case listUsersDto.Limit <= 0:
		listUsersDto.Limit = 10
	}
	var users []SQLiteUser
	var count int64

	result := db.Model(&SQLiteUser{}).Offset((listUsersDto.Page - 1) * listUsersDto.Limit).Limit(listUsersDto.Limit).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	countResult := db.Model(&SQLiteUser{}).Count(&count)

	if countResult.Error != nil {
		return nil, countResult.Error
	}

	data := MapSQLiteUsersListToUsersList(&users)

	return &models.PaginatedUserResultDTO{
		Page:       listUsersDto.Page,
		PerPage:    listUsersDto.Limit,
		Data:       *data,
		Count:      count,
		TotalPages: int(math.Ceil(float64(count) / float64(listUsersDto.Limit))),
	}, nil
}

func (r *SQLiteUserRepository) Update(updateUserDto *models.UpdateUserDTO) (*models.User, error) {
	return nil, nil
}

func (r *SQLiteUserRepository) Delete(ID string) error {
	return nil
}

func MapSQLiteUserToUserModel(sqliteUser SQLiteUser) *models.User {
	return &models.User{
		ID:       sqliteUser.ID,
		Email:    sqliteUser.Email,
		Name:     sqliteUser.Name,
		Password: sqliteUser.Password,
	}
}

func MapSQLiteUsersListToUsersList(sqliteUsers *[]SQLiteUser) *[]models.User {
	data := &[]models.User{}
	for _, user := range *sqliteUsers {
		*data = append(*data, *MapSQLiteUserToUserModel(user))
	}
	return data
}
