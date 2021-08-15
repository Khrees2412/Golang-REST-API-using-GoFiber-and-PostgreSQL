package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/convas/database"
	"github.com/khrees2412/convas/models"
	"golang.org/x/crypto/bcrypt"
)


type Data struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func Register(c *fiber.Ctx) error {
	 data := new(Data)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := new(Data)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(Secretkey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "strict",
		// Secure: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// 	var user models.User
// 	database.DB.Where("id = ?", claims.Issuer).First(&user)
// 	return c.JSON(user)

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}



func ChangePassword(c *fiber.Ctx) error {
	data := new(Data)
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "user is unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer

	if err := c.BodyParser(data); err != nil {
		return err
	}

	userID, _ := strconv.Atoi(user_ID)

	if userID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var user models.User
	database.DB.Where("id = ?", userID).First(&user)

	newpassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	// Here we might choose to send a link for the user to update the password,
	// since I'm not working with a mailing service yet we do this (below) instead.

	user.Password = newpassword
	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteAccount(c *fiber.Ctx) error {
	data := new(Data)
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "user is unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer

	if err := c.BodyParser(data); err != nil {
		return err
	}

	userID, _ := strconv.Atoi(user_ID)

	if userID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var user models.User
	database.DB.Where("id = ?", userID).First(&user)
	var books []models.Book
	database.DB.Where("User_ID = ?", userID).Find(&books)

	time.Sleep(time.Second * 5)
	go database.DB.Delete(&user)
	go database.DB.Delete(&books)


	update_cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&update_cookie)
	return c.JSON(fiber.Map{
		"message": "User Account deleted successfully",
	})
}