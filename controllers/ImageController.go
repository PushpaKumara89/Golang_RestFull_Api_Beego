package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math"

	beego "github.com/beego/beego/v2/server/web"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) UploadImage() {
	// Retrieve the uploaded file from the request
	file, _, err := c.GetFile("image")
	if err != nil {
		// Error handling
		c.Data["json"] = map[string]string{"error": "Failed to retrieve the image"}
		c.ServeJSON()
		return
	}
	defer file.Close()

	// Save the uploaded file to a desired location
	// You can customize the file storage path and name as per your requirements
	filePath := "path/to/save/" + randomBase16String(12) + ".jpg"
	err = c.SaveToFile("image", filePath)
	if err != nil {
		// Error handling
		log.Println(err.Error())
		c.Data["json"] = map[string]string{"error": "Failed to save the image"}
		c.ServeJSON()
		return
	}

	// Image uploaded successfully
	c.Data["json"] = map[string]string{"message": "Image uploaded successfully"}
	c.ServeJSON()
}

func (c *ImageController) GetImage() {
	// Retrieve the image ID from the URL parameter
	imageID := c.GetString(":id")

	// Load image data from storage or database based on the imageID

	// Set the appropriate headers
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename="+imageID)
	c.Ctx.Output.Header("Content-Type", "image/jpeg") // Set the correct content type

	// Serve the image using Beego's Output.Download method
	c.Ctx.Output.Download("path/to/save/" + imageID)
}

func randomBase16String(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)

	return str[:l] // strip 1 extra character we get from odd length results
}
