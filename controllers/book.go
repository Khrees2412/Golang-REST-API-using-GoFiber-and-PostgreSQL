package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/khrees2412/convas/database"
	"github.com/khrees2412/convas/models"
)

type Book struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Author  string `json:"author"`
	UserID uint `json:"user_id"`
}

	var Secretkey = os.Getenv("SECRETKEY")
	
func CreateBook(c *fiber.Ctx) error {
	data := new(Book)
	
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer

	if err := c.BodyParser(data); err != nil {
		return err
	}
	userID, _ := strconv.Atoi(user_ID)

	book := models.Book{
		Name: data.Name,
		Type: data.Type,
		Author: data.Author,
		UserID: userID,
	}
	database.DB.Create(&book)
	fmt.Println("Book created..")
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Book created successfully",
	})
}

func GetBooks(c *fiber.Ctx) error {
		
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer

	var books []models.Book
	database.DB.Where("User_ID = ? ", user_ID).Find(&books)
	fmt.Println("Found books..")
	c.Status(fiber.StatusOK)

	return c.JSON(fiber.Map{
		"message": "Found Books",
		"data":books,
	})
}

func GetBook(c *fiber.Ctx) error {
	paramid := c.Params("id")
		
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer
	id, err := strconv.Atoi(paramid)
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : "Couldn't get param",
			})
		}
	
	var book models.Book
	database.DB.Where("id = ? AND User_ID = ? ", id, user_ID).Find(&book)
	fmt.Println("Found book..")
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Found Book",
		"data":book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	paramid := c.Params("id")

		cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer

	id, err := strconv.Atoi(paramid)
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : "Couldn't get param",
			})
		}
	
	var book models.Book
	database.DB.Where("id = ? AND User_ID = ? ", id, user_ID).Delete(&book)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}

func UpdateBook(c *fiber.Ctx) error {
	paramid := c.Params("id")
	data := new(Book)

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if err := c.BodyParser(data); err != nil {
		return err
	}
	bookname := data.Name
	booktype := data.Type
	bookauthor := data.Author

	claims, _ := token.Claims.(*jwt.StandardClaims)
	user_ID := claims.Issuer

	id, err := strconv.Atoi(paramid)
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : "Couldn't get param",
			})
		}
	
	var book models.Book
	database.DB.Where("id = ? AND User_ID = ? ", id, user_ID).First(&book)
	book.Name = bookname
	book.Type = booktype
	book.Author = bookauthor
	database.DB.Save(&book)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Book Updated successfully",
	})
}
