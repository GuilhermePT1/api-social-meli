package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/infra/http/controllers"
	"github.com/GuilhermePT1/api-social-meli/internal/infra/repositories"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Repositories
	productRepo := repositories.NewProductRepository(db)
	postRepo := repositories.NewPostRepository(db)
	followRepo := repositories.NewFollowRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Services
	productService := services.NewProductService(productRepo)
	postService := services.NewPostService(postRepo)
	followService := services.NewFollowService(followRepo)
	userService := services.NewUserService(userRepo)

	// Controllers
	productController := controllers.NewProductController(productService)
	postController := controllers.NewPostController(postService)
	followController := controllers.NewFollowController(followService)
	userController := controllers.NewUserController(userService)

	api := router.Group("/api")
	{
		registerFollowRoutes(api, followController)
		registerProductRoutes(api, productController)
		registerPostRoutes(api, postController)
		registerUserRoutes(api, userController)
	}
}

func registerUserRoutes(router *gin.RouterGroup, c *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", c.CreateUser)
		userRoutes.GET("/:user_id", c.GetUserById)
		userRoutes.GET("", c.GetAllUsers)
	}
}

func registerFollowRoutes(router *gin.RouterGroup, c *controllers.FollowController) {
	followRoutes := router.Group("/users")
	{
		followRoutes.POST("/follow", c.Follow)
		followRoutes.POST("/unfollow", c.Unfollow)
		followRoutes.GET("/:user_id/followers/count", c.CountFollowers)
		followRoutes.GET("/:user_id/followers/list", c.GetFollowers)
		followRoutes.GET("/:user_id/followed/list", c.GetFollowed)
	}
}

func registerProductRoutes(router *gin.RouterGroup, c *controllers.ProductController) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", c.CreateProduct)
		productRoutes.GET("/:product_id", c.GetProductById)
		productRoutes.GET("", c.GetAllProducts)
	}
}

func registerPostRoutes(router *gin.RouterGroup, c *controllers.PostController) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("", c.CreatePost)
		postRoutes.GET("/users/:user_id", c.FindByUser)
		postRoutes.GET("/promo", c.FindPromoPosts)
		postRoutes.GET("/promo/count", c.CountPromoProducts)
	}
}
