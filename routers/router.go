package routers

import (
	"art_system/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:ShowGet;post:Post")

	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandlePost")

	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	//文章列表页访问
	beego.Router("/article/showArticleList", &controllers.ArticleController{}, "get:ShowArticleList")
	//添加文章
	beego.Router("/article/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	//显示文章详情
	beego.Router("/article/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
	//编辑文章内容
	beego.Router("/article/updateArticle", &controllers.ArticleController{}, "get:ShowUpdateArticle;post:HandleUpdate")
	//删除文章
	beego.Router("/article/deleteArticle", &controllers.ArticleController{}, "get:HandleDelete")
	//添加类型界面显示
	beego.Router("/article/addArticleType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")
	//退出登陆
	beego.Router("/article/logout", &controllers.UserController{}, "get:Logout")
}
