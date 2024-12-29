package controllers

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/zaidanpoin/crud-golang-react/Model"
)

func CreateMembers(c *gin.Context) {

	// Bind data form manually

	// Ambil file dari form-data
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to retrieve file",
			"error":   err.Error(),
		})
		return
	}

	// Check file extension
	allowedExtensions := map[string]bool{
		".webp": true,
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid file type. Only .webp, .jpg, and .png are allowed",
		})
		return
	}

	md5Hash := md5.New()
	md5Hash.Write([]byte(file.Filename + fmt.Sprintf("%d", time.Now().UnixNano())))
	hashedFilename := fmt.Sprintf("%x", md5Hash.Sum(nil))

	file.Filename = hashedFilename + ext

	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to retrieve file",
			"error":   err.Error(),
		})
		return
	}

	// Simpan file ke lokasi tertentu (opsional)
	uploadPath := "./uploads/" + file.Filename
	url := "http://localhost:8080/uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to save file",
			"error":   err.Error(),
		})
		return
	}

	// Simpan data ke dalam model
	mem := model.Member{
		Name:  c.PostForm("name"),
		Image: file.Filename,
		Url:   url,
	}

	// Simpan data ke database (implementasi Save() sesuai kebutuhan Anda)
	_, saveErr := mem.Save()
	if saveErr != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to save member",
			"error":   saveErr.Error(),
		})
		return
	}

	// Response sukses
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Member data saved successfully",
		"data": gin.H{
			"name":  mem.Name,
			"image": mem.Image,
		},
	})
}

func GetMembers(c *gin.Context) {

	var member model.Member

	members, err := member.GetDataMembers(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to get members",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Members data",
		"data":    members,
	})
}

func GetMemberByID(c *gin.Context) {

	var member model.Member

	members, err := member.GetDataMembers(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to get members",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Members data",
		"data":    members,
	})
}

func UpdateMembers(c *gin.Context) {
	var memberId model.Member

	// Retrieve the existing member data
	members, err := memberId.GetDataMembers(c.Param("id"))
	if err != nil || len(members) == 0 {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Member not found",
		})
		return
	}

	// Validate the name field
	if c.PostForm("name") == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Name is required",
		})
		return
	}

	// Initialize the member object with existing data
	imagesValue := members[0].Image
	url := "http://localhost:8080/uploads/" + imagesValue
	member := model.Member{
		Name:  c.PostForm("name"),
		Image: imagesValue,
		Url:   url,
	}

	// Check if a new image file is uploaded
	file, err := c.FormFile("image")
	if err == nil { // If there's no error, it means a file was uploaded
		// Check file extension
		allowedExtensions := map[string]bool{
			".webp": true,
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		}

		ext := filepath.Ext(file.Filename)
		if !allowedExtensions[ext] {
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "Invalid file type. Only .webp, .jpg, and .png are allowed",
			})
			return
		}

		// Generate a new hashed filename
		md5Hash := md5.New()
		md5Hash.Write([]byte(file.Filename + fmt.Sprintf("%d", time.Now().UnixNano())))
		hashedFilename := fmt.Sprintf("%x", md5Hash.Sum(nil))

		file.Filename = hashedFilename + ext

		// Save the new file
		uploadPath := "./uploads/" + file.Filename
		url = "http://localhost:8080/uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to save file",
				"error":   err.Error(),
			})
			return
		}

		// Update the member's image field
		member.Image = file.Filename
		member.Url = url

		// Optionally, remove the old image file
		os.Remove("./uploads/" + imagesValue)
	}

	// Update the member in the database
	err = member.UpdateMember(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to update member",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Member data updated successfully",
	})
}
func DeleteMembers(c *gin.Context) {
	var member model.Member
	fmt.Println(c.Param("id"))
	members, _ := member.GetDataMembers(c.Param("id"))

	imagesValue := members[0].Image

	os.Remove("./uploads/" + imagesValue)

	err := member.DeleteMember(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to delete member",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Member data deleted successfully",
	})
}
