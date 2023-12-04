package controllers

import (
	"joglo-fiber-gorm/database"
	"joglo-fiber-gorm/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(ctx *fiber.Ctx) error {

	var users []models.User
	var result *gorm.DB
	var searchQuery = ctx.Queries()

	if searchQuery["first_name"] != "" {
		result = database.DB.Where("first_name LIKE ?", "%"+searchQuery["first_name"]+"%").Find(&users)
	} else {
		result = database.DB.Find(&users) // <-- yang di return users
	}

	log.Println(result.Statement.Vars...)
	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": users,
	})

}

func GetUserDetail(ctx *fiber.Ctx) error {

	var params = ctx.AllParams()
	var convertInt, err = strconv.Atoi(params["id"])
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "convert uint error, id param required",
		})
	}
	var id = uint(convertInt)
	var user models.User

	result := database.DB.Where(&models.User{Id: id}).First(&user)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "fetch user success",
		"data":    user,
	})

}

func RegisterUser(ctx *fiber.Ctx) error {

	u := new(models.UserCreateRequest)

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	u.Password = "myPasswordAnything" // encrypt this

	result := database.DB.Table("users").Create(&u)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	log.Println(result.Statement.Vars...)

	return ctx.Status(200).JSON(fiber.Map{
		"message": "register success",
		"data":    u,
	})
}

func LoginUser(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "login success",
		"data":    "",
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "update profile success",
		"data":    "",
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "delete profile success",
		"data":    "",
	})
}
