package routes

import (
	"lms/config/db"
	"lms/controller"
	"lms/controller/answer"
	"lms/controller/category"
	"lms/controller/content"
	"lms/controller/course"
	"lms/controller/knowledge"
	"lms/controller/menu"
	"lms/controller/question"
	"lms/controller/quiz"
	"lms/controller/reference"
	"lms/controller/register"
	"lms/controller/role"
	"lms/controller/section"
	"lms/controller/user"
	"lms/docs"
	"lms/handler"
	"lms/middleware"

	ranswer "lms/repository/r_answer"
	rcategory "lms/repository/r_category"
	rcontent "lms/repository/r_content"
	rcourse "lms/repository/r_course"
	rknowledge "lms/repository/r_knowledge"
	rquestion "lms/repository/r_question"
	rquiz "lms/repository/r_quiz"
	rsection "lms/repository/r_section"
	ruser "lms/repository/r_user"
	"lms/service/s_answer"
	"lms/service/s_category"
	"lms/service/s_content"
	"lms/service/s_course"
	"lms/service/s_knowledge"
	"lms/service/s_question"
	"lms/service/s_quiz"
	"lms/service/s_section"
	"lms/service/s_user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Secure JSON prefix
	r.SecureJsonPrefix(")]}',\n")

	// create repository instance
	sectionRepo := rsection.NewSectionRepository()
	sectionGetRepo := rsection.NewSectionGetRepository()
	sectionGetService := s_section.NewSectionGetService(sectionGetRepo)

	userRepo := ruser.NewUserRepository()
	userUpdateRepo := ruser.UpdateUserRepository()

	knowledgeGetRepo := rknowledge.NewKnowledgeGetRepository()
	knowledgeGetService := s_knowledge.NewKnowledgeGetService(knowledgeGetRepo)
	knowledgeRepo := rknowledge.NewKnowledgeRepository()
	knowledgeUpdateRepo := rknowledge.UpdateKnowledgeRepository()
	knowledgeDeleteRepo := rknowledge.DeleteKnowledgeRepository()
	knowledgeService := s_knowledge.NewKnowledgeService(knowledgeRepo)
	knowledgeDeleteService := s_knowledge.DeleteKnowledgeService(knowledgeDeleteRepo)

	courseRepo := rcourse.NewCourseRepository(db.Server())
	courseGetRepo := rcourse.NewCourseGetRepository()
	courseGetService := s_course.NewCourseGetService(courseGetRepo)

	contentRepo := rcontent.NewContentRepository()
	contentService := s_content.NewContentService(contentRepo)
	contentDeleteRepo := rcontent.DeleteContentRepository()
	contentDeleteService := s_content.DeleteContentService(contentDeleteRepo)
	categoryRepo := rcategory.NewCategoryRepository()
	categoryService := s_category.NewCategoryService(categoryRepo)

	quizRepo := rquiz.NewQuizRepository(db.Server())
	quizService := s_quiz.NewQuizService(quizRepo)
	quizDeleteRepo := rquiz.DeleteQuizRepository()
	quizDeleteService := s_quiz.DeleteQuizService(quizDeleteRepo)
	quizGetRepo := rquiz.NewQuizGetRepository()
	quizGetService := s_quiz.NewQuizGetService(quizGetRepo)
	quizUpdateRepo := rquiz.UpdateQuizRepository()

	questionGetRepo := rquestion.NewQuestionGetRepository()
	questionGetService := s_question.NewQuestionGetService(questionGetRepo)
	questionRepo := rquestion.NewQuestionRepository(db.Server())
	questionService := s_question.NewQuestionService(questionRepo)
	questionUpdateRepo := rquestion.UpdateQuestionRepository()
	questionDeleteRepo := rquestion.DeleteQuestionRepository()
	questionDeleteService := s_question.DeleteQuestionService(questionDeleteRepo)

	answerGetRepo := ranswer.NewAnswerGetRepository()
	answerGetService := s_answer.NewAnswerGetService(answerGetRepo)
	answerRepo := ranswer.NewAnswerRepository(db.Server())
	answerService := s_answer.NewAnswerService(answerRepo)
	answerUpdateRepo := ranswer.UpdateAnswerRepository()
	answerDeleteRepo := ranswer.DeleteAnswerRepository()
	answerDeleteService := s_answer.DeleteAnswerService(answerDeleteRepo)

	// create controller instance
	sectionPostController := section.NewSectionController(sectionRepo)
	sectionGetController := section.NewSectionGetController(sectionGetService)

	courseController := course.NewCourseController(courseRepo)
	contentController := content.NewContentController(contentService)
	contentDeleteController := content.DeleteContentController(contentDeleteService)

	categoryController := category.NewCategoryController(categoryService)

	userAddController := user.NewUserController(userRepo)
	userUpdateController := user.UpdateUserController(userUpdateRepo)
	userDeleteRepo := ruser.NewUserDeleteRepository()
	userDeleteService := s_user.NewUserDeleteService(userDeleteRepo)
	userDeleteController := user.NewUserDeleteController(userDeleteService)

	knowledgeController := knowledge.NewKnowledgeController(knowledgeService)
	knowledgeUpdateController := knowledge.UpdateKnowledgeController(knowledgeUpdateRepo)
	knowledgeDeleteController := knowledge.DeleteKnowledgeController(knowledgeDeleteService)
	knowledgeGetController := knowledge.NewKnowledgeGetController(knowledgeGetService)

	courseGetController := course.NewCourseGetController(courseGetService)

	quizController := quiz.NewQuizController(quizService)
	quizDeleteController := quiz.DeleteQuizController(quizDeleteService)
	quizGetController := quiz.NewQuizGetController(quizGetService)
	quizUpdateController := quiz.UpdateQuizController(quizUpdateRepo)

	questionController := question.NewQuestionController(questionService)
	questionGetController := question.NewQuestionGetController(questionGetService)
	questionDeleteController := question.DeleteQuestionController(questionDeleteService)
	questionUpdateController := question.UpdateQuestionController(questionUpdateRepo)

	answerController := answer.NewAnswerController(answerService)
	answerGetController := answer.NewAnswerGetController(answerGetService)
	answerDeleteController := answer.DeleteAnswerController(answerDeleteService)
	answerUpdateController := answer.UpdateAnswerController(answerUpdateRepo)

	// Setup auth middleware
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		panic(err)
	}

	//Routes Swagger
	docs.SwaggerInfo.BasePath = "/"

	// Apply to public routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", controller.Helloworld)
	r.POST("/login", handler.LoginHandler)
	r.POST("/register", register.Register)
	r.POST("/users", userAddController.RegisterUser)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Apply auth middleware to routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//Users
		auth.GET("/users", user.GetUsers)
		auth.GET("/users/:id_user", user.GetUserByID)
		// auth.POST("/users", userController.AddUser)
		auth.PUT("/users/:id_user", userUpdateController.UpdateUser)
		auth.DELETE("/users/:id_user", userDeleteController.DeleteUser)

		//Knowledge
		auth.GET("/knowledge", knowledgeGetController.GetAllKnowledge)
		auth.GET("/knowledge/:id_knowledge", knowledgeGetController.GetByIdKnowledge)
		auth.POST("/knowledge/", knowledgeController.AddKnowledge)
		auth.PUT("/knowledge/:id_knowledge", knowledgeUpdateController.UpdateKnowledgeByID)
		auth.DELETE("/knowledge/:id_knowledge", knowledgeDeleteController.DeleteKnowledge)

		//Category
		auth.GET("/category", category.GetCategory)
		auth.POST("/category/", categoryController.AddCategory)
		auth.PUT("/category/:idcategory", category.UpdateCategory)
		auth.DELETE("/category/:idcategory", category.DeleteCategory)

		//Section
		auth.GET("/section", sectionGetController.GetAllSection)
		auth.GET("/section/:idsection", sectionGetController.GetByIdSection)
		auth.POST("/section/", sectionPostController.AddSection)
		auth.PUT("/section/:idsection", section.UpdateSection)
		auth.DELETE("/section/:idsection", section.DeleteSection)

		//Content
		auth.GET("content", content.GetContent)
		auth.GET("content/:idcontent", content.GetContentByID)
		auth.POST("content/", contentController.AddContent)
		auth.PUT("content/:idcontent", content.UpdateContent)
		auth.DELETE("content/:idcontent", contentDeleteController.DeleteContent)

		//Roles
		auth.GET("role", role.GetRoles)
		auth.POST("role/", role.AddRoles)
		auth.PUT("/role/:id_role", role.UpdateRole)
		auth.DELETE("/role/:id_role", role.DeleteRole)

		//Menu
		auth.GET("menu", menu.GetMenu)

		//Course
		auth.GET("course", courseGetController.GetAllCourse)
		auth.GET("course/:idcourse", courseGetController.GetByIdCourse)
		// auth.GET("course/:id_user", course.GetCoursesByUserID)
		auth.POST("course/", courseController.AddCourse)
		auth.PUT("course/:idcourse", course.UpdateCourse)
		auth.DELETE("course/:idcourse", course.DeleteCourse)

		//Quiz
		auth.GET("quiz", quizGetController.GetAllQuiz)
		auth.GET("/quiz/:id_quiz", quizGetController.GetByIdQuiz) // Endpoint baru
		auth.POST("quiz/", quizController.AddQuiz)
		auth.PUT("quiz/:id_quiz", quizUpdateController.UpdateQuizByID)
		auth.DELETE("quiz/:id_quiz", quizDeleteController.DeleteQuiz)

		//Question
		auth.GET("question", questionGetController.GetAllQuestion)
		auth.POST("question/", questionController.AddQuestion)
		auth.PUT("question/:id_question", questionUpdateController.UpdateQuestionByID)
		auth.DELETE("question/:id_question", questionDeleteController.DeleteQuestion)

		//Answer
		auth.GET("answer", answerGetController.GetAllAnswer)
		auth.POST("answer/", answerController.AddAnswer)
		auth.PUT("answer/:id_answer", answerUpdateController.UpdateAnswerByID)
		auth.DELETE("answer/:id_answer", answerDeleteController.DeleteAnswer)

		//Reference
		auth.GET("reference", reference.GetReference)
		auth.POST("reference/", reference.AddReference)
	}

	return r
}
