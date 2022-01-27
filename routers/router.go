package routers

import (
	"gin-blog/handler"
	"gin-blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	// 管理后台
	// 管理员登陆
	r.POST("/login", handler.AuthLogin)

	// 管理员路由组
	userGroup := r.Group("/user/v1")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("/create-user", handler.CreateUser)
		userGroup.GET("/get-users", handler.GetUsers)
		userGroup.POST("/del-user-by-id", handler.DeleteUserById)
		userGroup.GET("/get-user-by-id", handler.GetUserById)
		userGroup.POST("/update-user-by-id", handler.UpdateUserById)
		userGroup.POST("/disable-user-by-id", handler.DisableUserById)
		userGroup.POST("/enable-user-by-id", handler.EnableUserById)

	}

	// 标签路由组
	tagsGroup := r.Group("/tags/v1")
	tagsGroup.Use(middleware.AuthMiddleware())
	{
		tagsGroup.POST("/create-tags", handler.CreateTags)
		tagsGroup.GET("/get-tags-list", handler.GetTagsList)
		tagsGroup.POST("/update-Tags-by-id", handler.UpdateTagsById)
		tagsGroup.POST("/del-tags-by-id", handler.DeleteTagsById)
	}

	// 分类路由组
	cateGroup := r.Group("/cate/v1")
	cateGroup.Use(middleware.AuthMiddleware())
	{
		cateGroup.POST("/create-cate", handler.CreateCate)
		cateGroup.GET("/get-cate-list", handler.GetCateList)
		cateGroup.POST("/update-cate-by-id", handler.UpdateCateById)
		cateGroup.POST("/del-cate-by-id", handler.DeleteCateById)
	}

	// 文章路由组
	postsGroup := r.Group("/posts/v1")
	postsGroup.Use(middleware.AuthMiddleware())
	{
		postsGroup.POST("/create-post", handler.CreatePost)
		postsGroup.GET("/get-posts-list", handler.GetPostsList)
		postsGroup.POST("/update-post-by-id", handler.UpdatePostById)
		postsGroup.POST("/del-post-by-id", handler.DeletePostById)
	}

	// 文章评论路由组
	commentGroup := r.Group("/comment/v1")
	{
		commentGroup.POST("/create-comment", handler.CreateComment)
		commentGroup.GET("/get-comment-list", handler.GetCommentList)
		commentGroup.POST("/del-comment-by-id", handler.DelCommentById)
	}

	// 友联路由组
	linkGroup := r.Group("/link/v1")
	{
		linkGroup.POST("/create-link", handler.CreateLink)
		linkGroup.GET("/get-link-list", handler.GetLinkList)
		linkGroup.POST("/update-link-by-id", handler.UpdateLinkById)
		linkGroup.POST("/del-link-by-id", handler.DeleteLinkById)
	}


	r.Run(":8080")
}
