package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
)

// CreateUserSchema -> create user schema validation
type CreateUserSchema struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Age      int64
	Address  string
}

// CreateUser -> validation to create user
func CreateUser(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	userValidate := &CreateUserSchema{
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
	}
	Validate(userValidate)
}

// GetUserParamSchema -> check schema param validation
type GetUserParamSchema struct {
	ID uint `uri:"id" binding:"required"`
}

// GetUser -> validation to get user by id
func GetUser(c *gin.Context) {
	param := GetUserParamSchema{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Param must be of type integer, required", "INVALID_BODY")
	}
}

// UpdateUserSchema -> update user schema validation
type UpdateUserSchema struct {
	Name    string
	Email   string `validate:"email"`
	Age     int64
	Address string
}

// UpdateUser -> validation to update user by id with body
func UpdateUser(c *gin.Context) {
	param := GetUserParamSchema{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Param must be of type integer, required", "INVALID_BODY")
	}

	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	userValidate := &UpdateUserSchema{
		Name:    user.Name,
		Address: user.Address,
		Age:     user.Age,
		Email:   user.Email,
	}
	Validate(userValidate)
}

// DeleteUser -> validation to delete user by id
func DeleteUser(c *gin.Context) {
	param := GetUserParamSchema{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Param must be of type integer, required", "INVALID_BODY")
	}
}
