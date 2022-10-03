package dao

import (
	"go_boke/models"
	"log"
)

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_categories where cid=?", cId)
	if row.Err() != nil {
		log.Println("blog_category获取失败")
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_categories")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
