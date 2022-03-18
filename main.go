package main

import (
	"treehole/db"
	_ "treehole/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//db.ArticleAdd()
	db.ArticleRealList()
	beego.Run()

}
