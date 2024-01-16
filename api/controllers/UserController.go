package controllers

import (
	"brittola-api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db: db}
}

func (u *userController) Create(ctx *gin.Context) {
	var user entities.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := u.db.Create(&user).Error; err != nil {
		if isDuplicateKeyError(err) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Usuário ou e-mail já existente"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": user.Username + " created"})
}

// função auxiliar para verificar se um erro é relacionado à violação de unicidade no MySQL
func isDuplicateKeyError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok && mysqlErr.Number == 1062 {
		return true
	}
	return false
}
