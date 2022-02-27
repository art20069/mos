package api

import (
	"fmt"
	"main/db"
	_ "main/interceptor"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - call this method to setup product route group
func SetupLottoAPI(router *gin.Engine) {
	LottoAPI := router.Group("/api/v2")
	{
		LottoAPI.POST("/lottoadd" /*interceptor.JwtVerify,*/, createAddLotto)
		LottoAPI.GET("/lotto" /*interceptor.JwtVerify,*/, getLotto)
		// LottoAPI.GET("/product/:id" /*interceptor.JwtVerify,*/, getProductByID)
		// LottoAPI.POST("/product" /*interceptor.JwtVerify,*/, createProduct)
		// LottoAPI.PUT("/product" /*interceptor.JwtVerify,*/, editProduct)
	}
}

func createAddLotto(c *gin.Context) {

	lotto := model.Lottos{}
	lotto.Name = c.PostForm("name")
	lotto.Number_lotto, _ = strconv.ParseInt(c.PostForm("number_lotto"), 10, 64)
	lotto.Multiply, _ = strconv.ParseInt(c.PostForm("multiply"), 10, 64)
	lotto.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	lotto.CreatedAt = time.Now()
	db.GetDB().Create(&lotto)
	c.JSON(http.StatusOK, gin.H{"result": lotto})
}

func getLotto(c *gin.Context) {
	var lotto []model.Lottos

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name like ?", keyword).Find(&lotto)
	} else {
		db.GetDB().Find(&lotto)
	}
	c.JSON(200, lotto)

}

/*
func getProduct(c *gin.Context) {
	var product []model.Product
	db.GetDB().Find(&product)
	c.JSON(200, product)
}
// */

// func getProduct(c *gin.Context) {
// 	var product []model.Product

// 	keyword := c.Query("keyword")
// 	if keyword != "" {
// 		keyword = fmt.Sprintf("%%%s%%", keyword)
// 		db.GetDB().Where("name like ?", keyword).Find(&product)
// 	} else {
// 		db.GetDB().Find(&product)
// 	}
// 	c.JSON(200, product)

// }

// func getProductByID(c *gin.Context) {
// 	var product model.Product
// 	db.GetDB().Where("id = ?", c.Param("id")).First(&product)
// 	c.JSON(200, product)
// }

// func fileExists(filename string) bool {
// 	info, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }

// func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
// 	if image != nil {
// 		runningDir, _ := os.Getwd()
// 		product.Image = image.Filename
// 		extension := filepath.Ext(image.Filename)
// 		fileName := fmt.Sprintf("%d%s", product.ID, extension)
// 		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

// 		if fileExists(filePath) {
// 			os.Remove(filePath)
// 		}
// 		c.SaveUploadedFile(image, filePath)
// 		db.GetDB().Model(&product).Update("image", fileName)
// 	}
// }

// func editProduct(c *gin.Context) {
// 	var product model.Product
// 	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
// 	product.ID = uint(id)
// 	product.Name = c.PostForm("name")
// 	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
// 	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

// 	db.GetDB().Save(&product)
// 	image, _ := c.FormFile("image")
// 	saveImage(image, &product, c)
// 	c.JSON(http.StatusOK, gin.H{"result": product})

// }
