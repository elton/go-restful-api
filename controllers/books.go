// Copyright 2020 Elton Zheng
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"net/http"

	"github.com/elton/go-restful-api/models"
	"github.com/gin-gonic/gin"
)

// CreateBookInput 创建图书的输入对象
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required`
}

// UpdateBookInput 更新图书的输入对象
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks 得到所有的书籍
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook 根据ID获取对应的图书
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook 创建图书
func CreateBook(c *gin.Context) {
	// 验证输入
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建图书
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook 更新book数据
// curl -X PATCH -H 'Content-Type: application/json' -d '{"title":"The Infinite Game"}' http://localhost:8080/books/1
func UpdateBook(c *gin.Context) {
	// 得到一个图书的模型
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// 验证输入
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建待更新的图书
	newBook := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	models.DB.Model(&book).Updates(&newBook)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
