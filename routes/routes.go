package routes

import (
	"../middleware"
	"../models"
	"../sessions"
	"../utilities"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	UrlPath      = "/"
	LoginPath    = "/login"
	RegisterPath = "/register"
	LogoutPath   = "/logout"
	UserPath     = "/{username}"
)

func NewRouter() *mux.Router {
	var (
		router = mux.NewRouter()
		fileServer = http.FileServer(http.Dir("./static/"))
	)
	// Подписываемся на некоторые функции для get и post запросов
	initializeHandler(router, middleware.AuthRequired(indexGetHandler), middleware.AuthRequired(indexPostHandler), UrlPath)
	initializeHandler(router, loginGetHandler, loginPostHandler, LoginPath)
	initializeHandler(router, registerGetHandler, registerPostHandler, RegisterPath)
	initializeHandler(router, logoutGetHandler, nil, LogoutPath)
	initializeHandler(router, middleware.AuthRequired(userGetHandler), nil, UserPath)

	// Устанавливаем префикс static для всех статических ресурсов.
	//Теперь мы можем использовать в этой папке .txt, .css и т. д.
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	return router
}

// Функция-помошник, инициализирующая функции для get и post запросов
func initializeHandler(router *mux.Router, GetHandler func (_ http.ResponseWriter, _ *http.Request), PostHandler func (_ http.ResponseWriter, _ *http.Request), path string)  {
	router.HandleFunc(path, GetHandler).Methods("get")
	router.HandleFunc(path, PostHandler).Methods("post")
}


//---------------------------- INDEX -------------------------------------------------//

// Функция, которая будет вызвана при get-запросе на странице "/"
func indexGetHandler(writer http.ResponseWriter, request *http.Request) {
	// Получаем доступ к первым десяти элементам с ключом key, занеся их в переменную comments
	var (
		session, _ = sessions.CookieStore.Get(request, "session")
		untypedUserId = 	session.Values["user_id"]
		userId, ok = untypedUserId.(int64)
		comments, err1 = models.GetAllComments(0, 10)
	)
	// Если вернулась ошибка
	if err1 != nil || !ok {
		utilities.InternalServerError(writer)
		return
	}

	var user, err2 = models.GetUserById(userId)
	if err2 != nil {
		utilities.InternalServerError(writer)
		return
	}

	// Иначе заносим в файл index.html переменную comments
	if utilities.ExecuteTemplate(writer, "index.html", models.NewIndexViewModel(user, comments)) != nil {
		utilities.InternalServerError(writer)
		return
	}

}

// Функция, которая будет вызвана при post-запросе на странице "/"
func indexPostHandler(writer http.ResponseWriter, request *http.Request) {
	var (
		session, _ = sessions.CookieStore.Get(request, "session")
		untypedUserId = 	session.Values["user_id"]
		userId, ok = untypedUserId.(int64)
	)
	if !ok {
		utilities.ParseFormError(writer)
		return
	}
	// Парсим форму
	if request.ParseForm() != nil {
		utilities.ParseFormError(writer)
		return
	}

	// Пушим в БД новые данные из формы, присваивая ключ key
	var err = models.AddComment(userId, request.PostForm.Get(models.CommentsKeyString))

	if err == utilities.ErrEmptyComment {
		utilities.EmptyCommentError(writer)
		return
	} else if err != nil {
		utilities.InternalServerError(writer)
		return
	}

	// Обновляем сайт
	http.Redirect(writer, request, UrlPath, http.StatusFound)
}

//---------------------------- LOGIN -------------------------------------------------//

// Функция, которая будет вызвана при get-запросе на странице "/login"
func loginGetHandler(writer http.ResponseWriter, _ *http.Request) {
	_ = utilities.ExecuteTemplate(writer, "login.html", nil)
}

// Функция, которая будет вызвана при post-запросе на странице "/login"
func loginPostHandler(writer http.ResponseWriter, request *http.Request) {
	if request.ParseForm() != nil {
		utilities.ParseFormError(writer)
		return
	}

	var (
		session, _ = sessions.CookieStore.Get(request, "session")
		username   = request.PostForm.Get("username")
		user, err1 = models.AuthenticateUser(username, request.PostForm.Get("password"))
	)

	if err1 == utilities.ErrUserNotFound || err1 == utilities.ErrWrongPassword {
		_ = utilities.ExecuteTemplate(writer, "login.html", err1.Error())
		return
	} else if err1 != nil {
		utilities.InternalServerError(writer)
		return
	}

	var userId, err2 = user.GetId()
	if err2 != nil {
		utilities.InternalServerError(writer)
		return
	}

	session.Values["user_id"] = userId
	err1 = session.Save(request, writer)

	if err1 != nil {
		_ = utilities.ExecuteTemplate(writer, "login.html", "save session error")
		return
	}

	http.Redirect(writer, request, "/", http.StatusFound)
}

//---------------------------- REGISTER --------------------------------------------//


func registerGetHandler(writer http.ResponseWriter, _ *http.Request) {
	_ = utilities.ExecuteTemplate(writer, "register.html", nil)
}

func registerPostHandler(writer http.ResponseWriter, request *http.Request) {
	var err = request.ParseForm()

	if err != nil {
		utilities.ParseFormError(writer)
		return
	}

	var emailHide, phoneHide bool
	if request.PostForm.Get("emailHide") == "on" {
		emailHide = true
	} else {
		emailHide = false
	}

	if request.PostForm.Get("phoneHide") == "on" {
		phoneHide = true
	} else  {
		phoneHide = false
	}

	var gender = models.Non
	if "on" == request.Form["gender"][0]  {
		gender = models.Male
	} else if "on" == request.Form["gender"][1] {
		gender = models.Female
	}

	err = models.RegisterUser(request.PostForm.Get("username"), request.PostForm.Get("password"), *models.NewEmail(request.PostForm.Get("email"), emailHide), *models.NewPhone(request.PostForm.Get("phone"), phoneHide), request.PostForm.Get("birth"), gender)

	if err == utilities.ErrUsernameTaken || err == utilities.ErrTooShortPassword || err == utilities.ErrEmptyBirth || err == utilities.ErrEmptyPhone || err == utilities.ErrEmptyEmail || err == utilities.ErrTooShortUsername {
		_ = utilities.ExecuteTemplate(writer, "register.html", err.Error())
		return
	} else if err != nil {
		utilities.InternalServerError(writer)
		return
	}
	http.Redirect(writer, request, "/login", http.StatusFound)
}

//---------------------------- USER -----------------------------------------------//

func logoutGetHandler(writer http.ResponseWriter, request *http.Request) {
	var session, _ = sessions.CookieStore.Get(request, "session")
	delete(session.Values, "user_id")
	var err = session.Save(request, writer)
	if err != nil {
		utilities.InternalServerError(writer)
	}
	http.Redirect(writer, request, LoginPath, http.StatusFound)
}

//---------------------------- USER -----------------------------------------------//

func userGetHandler(writer http.ResponseWriter, request *http.Request) {
	var (
		username = mux.Vars(request)["username"]
		user, err1 = models.GetUserByUsername(username)
	)
	if err1 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("Internal server error"))
		return
	}

	// Иначе заносим в файл index.html переменную comments
	if utilities.ExecuteTemplate(writer, "user.html", models.NewUserViewModel(user)) != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("Internal server error"))
		return
	}
}