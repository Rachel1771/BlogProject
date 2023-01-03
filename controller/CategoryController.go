package controller

import (
	"Blog/model"
	"Blog/response"
	"Blog/util"
	"github.com/gin-gonic/gin"
)

func SearchCategory(c *gin.Context) {
	db := util.GetDb()
	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		response.Fail(c, nil, "查找失败")
		return
	}
	response.Success(c, gin.H{"categories": categories}, "查找成功")

}
func SearchCategoryName(c *gin.Context) {
	db := util.GetDb()
	var category model.Category
	categoryId := c.Params.ByName("id")
	if err := db.Where("id=?", categoryId).First(&category).Error; err != nil {
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"categoryName": category.CategoryName}, "查找成功")
}
