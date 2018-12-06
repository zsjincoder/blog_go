package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//判断文章id是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

//根据文章标题查询是否存在
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//分页查询文章 预加载
func GetArticles(pageNum int, pageSize int, maps interface{}) (article []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)
	return
}

//根据id查询文章
func GetArtcle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

//修改文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Update(data)
	return true
}

//添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:    data["tag_id"].(int),
		Title:    data["title"].(string),
		Desc:     data["desc"].(string),
		Content:  data["content"].(string),
		CreateBy: data["create_by"].(string),
		State:    data["state"].(int),
	})
	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id=?",id).Delete(&Article{})
	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
