CommentTree

🚀 Установка и запуск
 
 Запустите проект:

    1) make dockerRun - ВСЕ УСТАНОВИТСЯ

🧪 Тестирование

Чтобы затестировать код и узнать покрытие введите команду:

    make testAll

💻 Технический стек

    Язык программирования: Go 1.25
    Web-фреймворк: Ginext
    База данных: PostgreSQL
    Миграции: Migrate
    Логгирование: zerolog
    Документация API: Swagger
    Кэш: Redis 
    Контейнеризация: Docker, Docker Compose

HTTP API

   	r.GET("/home", ren.HomeHandler)
	r.POST("/comments", h.CreateHandler)
	r.GET("/comments", h.GetRootCommentsHandler)
	r.GET("/comments/all", h.GetCommentTreeHandler)               
	r.GET("/comments/:parent_id/children", h.GetChildCommentsHandler)
	r.DELETE("/comments/:id", h.DeleteHandler)







