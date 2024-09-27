package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Class  string `json:"class"`
	Gender string `json:"gender"`
}

func main() {
	r := gin.Default()
	studentList := []Student{
		{1, "Student1", 20, "ABC", "Male"},
		{2, "Student2", 20, "ABCD", "Female"},
		{3, "Student3", 21, "ABCDE", "Male"},
		{4, "Student4", 20, "ABCDEF", "Female"},
	}

	r.GET("/get-students", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"students": studentList,
		})
	})

	r.GET("/get-student-detail/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		for _, student := range studentList {
			if student.Id == id {
				c.JSON(http.StatusOK, student)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	})

	r.POST("/add-student", func(c *gin.Context) {
		var newStudent Student

		if err := c.ShouldBindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newStudent.Id = len(studentList) + 1
		studentList = append(studentList, newStudent)

		c.JSON(http.StatusCreated, gin.H{
			"message":     "Student added successfully",
			"studentList": studentList,
		})
	})

	r.PUT("/update-student/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		var updatedStudent Student
		if err := c.ShouldBindJSON(&updatedStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, student := range studentList {
			if student.Id == id {
				updatedStudent.Id = id // Ensure ID remains unchanged
				studentList[i] = updatedStudent
				c.JSON(http.StatusOK, gin.H{
					"message":     "Student updated successfully",
					"studentList": studentList,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	})

	r.DELETE("/delete-student/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		for i, student := range studentList {
			if student.Id == id {
				studentList = append(studentList[:i], studentList[i+1:]...)
				c.JSON(http.StatusOK, gin.H{
					"message":     "Student deleted successfully",
					"studentList": studentList,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
