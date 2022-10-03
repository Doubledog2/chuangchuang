package dao

import (
	"go_boke/models"
	"log"
)

func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_posts where category_id=?;", cId)
	_ = rows.Scan(&count)
	return
}
func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_posts set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println("dao.updatepost fail")
		log.Println(err)
	}
}
func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_posts "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println("dao.savepost fail")
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}
func DeletePost(pid int) {
	_, err := DB.Exec("delete from blog_posts where pid=?", pid)
	if err != nil {
		log.Println("dao.deletepost fail")
		log.Println(err)
	}

}
func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_posts;")
	_ = rows.Scan(&count)
	return
}
func GetPostById(pid int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_posts where pid=?;", pid)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, row.Err()
	}
	return post, nil
}
func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_posts limit ?,?", page, pageSize)
	if err != nil {
		log.Println("blog_post获取失败")
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("blog_post数据赋值失败")
			log.Println(err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPostPageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_posts where category_id = ? limit ?,?", cId, page, pageSize)
	if err != nil {
		log.Println("blog_post fail")
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//搜索
func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_posts where title like ?", "%"+condition+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
