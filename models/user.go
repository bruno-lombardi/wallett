package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PaginatedUserResultDTO struct {
	TotalPages int    `json:"total_pages"`
	PerPage    int    `json:"per_page"`
	Page       int    `json:"page"`
	Data       []User `json:"data"`
}
