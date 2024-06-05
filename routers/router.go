package routers

import (
	_ "gin-test/docs" // 千万不要忘了导入把你上一步生成的docs
	"gin-test/middleware/jwt"
	"gin-test/pkg/export"
	"gin-test/pkg/setting"
	"gin-test/pkg/upload"
	"gin-test/routers/api"
	v1 "gin-test/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	//router.GET创建不同的HTTP方法绑定到Handlers中
	//Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，
	//在gin中包含大量Context的方法

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	r.GET("/test", api.Test)
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
