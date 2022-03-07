package model

import (
	"BLOGplus/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //文章分类
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"Title"` //标题
	Cid     int    `gorm:"type:int;not null" json:"Cid"`            //文章id
	Desc    string `gorm:"type:varchar(200)" json:"Desc"`           //描述
	Content string `gorm:"type:longtext;not null" json:"Content"`   //文章内容
	Img     string `gorm:"type:varchar(100)" json:"Img"`            //文章图片
}

//新增文章
func CreateArticle(cate *Article) int {
	err := db.Create(cate).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtlist []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArtlist).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtlist, errmsg.SUCCESS, total
}

//查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

//查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int, int) {
	var articlelist []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articlelist).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articlelist, errmsg.SUCCESS, total
}

//编辑文章信息
func EditeArticle(id int, data *Article) int {
	maps := make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Content
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&Article{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int) int {
	err := db.Where("id=?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
