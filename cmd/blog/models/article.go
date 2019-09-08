package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_mod-project/cmd/blog/utils"
	"strconv"
)

type Article struct {
	Id int
	Title string
	Tags string
	Short string
	Content string
	Author string
	Createtime int64
	//Status int	// Status=0 为正常， 1为删除， 2为冻结
}

func AddArticle(article Article) (int64, error) {
	return insertArticle(article)
}

// 添加文章
func insertArticle(article Article) (int64, error)  {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values (?, ?, ?, ?, ?, ?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime,
	)
}
// ---------- 查询文章 -----------

// 根据页码查询
func FindArticleWithPage(page int) ([]Article, error)  {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page --
	return QueryArticleWithPage(page, num)
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面没有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string) ([]Article, error)  {
	sql = "select id, title, tags, short, content, author, createtime from article " + sql

	rows, err := utils.QueryDB(sql)

	if err != nil {
		return nil, err
	}

	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content,author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

// ----- 翻页 ------
//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return  num
}

func QueryArticleWithId (id int) Article  {
	row := utils.QueryRowDB("select id, title, tags, short, content, author, createtime from article where id = " + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

func UpdateArticle(article Article) (int64, error)  {
	return utils.ModifyDB("update article set title=?, tags=?, short=?, content=? where id=?",
			article.Title, article.Tags, article.Short, article.Content, article.Id,
		)
}

func SetArticleRowsNum()  {
	articleRowsNum = QueryArticleRowNum()
}

func deleteArticle(id int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", id)
}

func DeleteArticle(id int) (int64, error)  {
	i, err := deleteArticle(id)
	SetArticleRowsNum()
	return i, err
}

func QueryArticleWithParam(param string) []string {
	sql := fmt.Sprintf("select %s from article", param)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		fmt.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticleWithCon(sql)
}