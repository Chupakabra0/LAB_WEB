# Лабораторные работы

## Отчёт по лабораторным работам<br>По дисциплине "Разработка и поддержка веб-приложений"<br>студента группы ПА-18-2<br>Сафиюлина Александра

### Лабораторная работа №1

#### Постановка задачи:

1. Разработать гипертекстовую страницу. Использовать минимальное количество гипертекстовых тегов. Создать заголовок, содержание, название подразделов. Предусмотреть возвращение из любой точки страницы к содержанию.
2. Для даной странички создать каскадную таблицу стилей, которая позволяет изменить тип, цвет и размер шрифтов заголовков, подзаголовков и основного текста, выравнивание текста в документе, цвет фона и ссылок.

#### Выполнение:

Создание документа начнем с стандартных тегов `<!DOCKTYPE html>` и `html`. Внутри них создадим теги `<head>` и `<body>`. В первом разместим всю необходимую метаинформацию странички, заголовок сайта, а также ссылку на файл стилей, который в дальнейшем будет использован для написания таблицы стилей. В `<body>` будет размещена сама разметка сайта.

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="css/style.css">
    <title>Word</title>
</head>

<body>
</body>

</html>

```

Для удобства работы с стилями и разбиения на секции создадим внутри теги `<header>` и `<footer>`. В первом разместим следующий код:

```html

    <header>
        <!-- Link bar -->
        <div class="Header">
            <!-- Logo -->
            <div class="Word">
                <a href="#"><span style="color: red;">W</span><span style="color: yellow">O</span><span style="color: green;">R</span><span style="color: blue;">D</span></a>
            </div>

            <!-- Links-buttons -->
            <div class="MenuBar">
                <ul class="Menu">
                    <li><a href="#First">What is HTML?</a></li>
                    <li><a href="#Second">Why do we need Bootstrap?</a></li>
                    <li><a href="#Third">Why Javascript so cool?</a></li>
                    <li><a href="#Fourth">What web protocols do we have?</a></li>
                </ul>
            </div>
        </div>

    </header>

```

Использованные теги:
* `<div>` — блочный элемент, который предназначен для выделения фрагмента документа с целью изменения вида содержания;
* `<a>` — элемент, который предназначен для создания гиперссылок;
* `<span>` — элемент, который предназначен для удобного форматирования малых элементов документа:
* `<ul>` — элемент, который предназначен для создания маркированного списка;
* `<li>` — элемент, который предназначен для создания пунктов списка (необязательно маркированного).

Для того, чтобы в слове WORD каждая буква окрасилась в отдельный цвет, обрамим каждую букву отдельным `<span>` и укажем ей с помощью атрибутов CSS цвета.

Теперь перейдем к `<footer>`. Так принято называть нижнюю часть сайта, в которой размещают информацию про собственника сайта, лицензиях, ссылок и т.д. Код выглядит так:

```html

<footer class="Footer">
        <div class="Source">
            Links:
        </div>
        <div class="Links">
            <ul class="LinksList">
                <a href="http://fastwebstart.com/modern-html-tutorial/"><li>Fastwebstart.com</li></a>
                <a href="https://medium.com/actualize-network/modern-html-explained-for-dinosaurs-65e56af2981"><li>Medium.com</li></a>
                <a href="https://javascript.info/intro"><li>Javascript.info</li></a>
                <a href="https://www.goodcore.co.uk/blog/web-technologies/"><li>Goodcore.co.uk</li></a>
            </ul>
        </div>
</footer>

```

Весь контент сайта разместим в `<div>`. Ранее можно было заметить использование атрибута `class`. Он помогает обращаться к элементу из таблицы стилей CSS.

```html
    <div class="Content">
        <!-- First paragraph -->
        <div class="FirstParagraph" id="First">
            <div class="H2">
                <h2>What is HTML?</h2>
                <a href="#" class="Back">Back</a>
            </div>
            HTML is a markup language. It marks parts of the document with structural, presentational, and semantic information alongside content. For example, this is simple document with no markup:
            <div class="Image">
                <img src="sources/FirstImage.png">
            </div>
            Whereas, the same document with markup below. The markup enhances the document by adding meaning to parts of the document:
            <div class="Image">
                <img src="sources/SecondImage.png">
            </div>
            Note that parts of the HTML document are marked with constructs called tags to give better meaning to them. <i><b>&lt;h1&gt;</b></i> History <i><b>&lt;/h1&gt;</b></i> means it is a top level heading.
            The h1 tags surrounds the heading. <i><b>&lt;h2&gt;</b></i>Pre-modern<i><b>&lt;/h2&gt;</b></i> means it is a second level heading.
            The word “Ancient China” is Hyper linked to another document. The reader can follow that link to get more information on the topic. <i><b>&lt;img/&gt;</b></i> is used to embed an image in the document.
            This way, an HTML document gives an enhanced experience with hyper links, formatted text with images, video and audio.
            The world wide web is, in fact a collection of several millions of such inter-connected HTML documents hosted on websites.
        </div>

        <!-- Second paragraph -->
        <div class="SecondParagraph" id="Second">
            <div class="Text">
                <div class="H2">
                    <h2>Why do we need Bootstrap?</h2>
                    <a href="#" class="Back">Back</a>
                </div>
                If we take a look at the website that was created just with HTML, well, it seems to be like this one:
                <div class="Image">
                    <img src="sources/ThirdImage.png">
                </div>
                In order to spruce it up, we’ll add a CSS file to apply styling. Now if you’re not particularly good at CSS, it could take you many days to make this website look pretty.
                Instead of writing your own CSS, you could always use a CSS framework, which is essentially CSS that someone else wrote in a reusable manner.
                One popular CSS framework is Bootstrap, which came out in 2011 and quickly became adopted and used by literally millions of websites.
                Let’s see what the code would look like using Bootstrap:
                <div class="Image">
                    <img src="sources/FourthImage.png">
                </div>
                Not bad! Note that in order to use a CSS framework like Bootstrap, you actually don’t need to write any CSS at all to get started
                — you just need to add the appropriate classes to the HTML to take advantage of the CSS that comes with the framework.
                Bootstrap became widely used because it helped with the major pain points of CSS at the time, such as browser inconsistencies and a lack of a proper grid system.
                <br><br>There are some downsides of using CSS frameworks like Bootstrap — in particular, they can be difficult to customize compared to writing CSS from scratch, which can make your website seem generic compared to others.
            </div>

        </div>

        <!-- Third paragraph -->
        <div class="ThirdParagraph" id="Third">
            <div class="Text">
                <div class="H2">
                    <h2>Why Javascript so cool?</h2>
                    <a href="#" class="Back">Back</a>
                </div>
                JavaScript was initially created to “make web pages alive”. The programs in this language are called scripts.
                They can be written right in a web page’s HTML and run automatically as the page loads. Scripts are provided and executed as plain text.
                They don’t need special preparation or compilation to run. In this aspect, JavaScript is very different from another language called Java.
                <br><br>
                When JavaScript was created, it initially had another name: “LiveScript”.
                But Java was very popular at that time, so it was decided that positioning a new language as a “younger brother” of Java would help.
                <br><br>
                Today, JavaScript can execute not only in the browser, but also on the server, or actually on any device that has a special program called the JavaScript engine.
                The browser has an embedded engine sometimes called a “JavaScript virtual machine”. Different engines have different “codenames”. For example:
                <div class="NumList">
                    <ol>
                        <li>V8 - Chrome and Opera</li>
                        <li>SpiderMonkey - Firefox</li>
                        <li>Chakra and Trident - Internet Explorer</li>
                        <li>SquirrelFish and Nitro - Safari</li>
                        <li>ChakraCore - Microsoft Edge</li>
                    </ol>
                </div>
                Modern JavaScript is a “safe” programming language. It does not provide low-level access to memory or CPU, because it was initially created for browsers which do not require it.
                JavaScript’s capabilities greatly depend on the environment it’s running in. For instance, Node.js supports functions that allow JavaScript to read/write arbitrary files, perform network requests, etc.
                JavaScript is the only browser technology that combines fully integration with HTML/CSS, simplicity and support by all major web-browsers. 
            </div>
        </div>

        <!-- Fourth paragraph -->
        <div class="FourthParagraph" id="Fourth">
            <div class="Text">
                <div class="H2">
                    <h2>What web protocols do we have?</h2>
                    <a href="#" class="Back">Back</a>
                </div>
                Web protocols are some predefined rules that must be followed by everyone communicating over the web. HyperText Transfer Protocol, better known as HTTP, is a web protocol that defines two concepts:
                how client requests are relayed to servers and how servers respond to client requests.
                <br><br>
                But there are some another protocols, that we must know such as Transmission Control Protocol (TCP), Internet Protocol (IP), User Datagram Protocol (UDP).
            </div>
        </div>
    </div>

```

Новые теги:
* `<br>` — элемент, который определяет отступ. Равно как и `<DOCKTYPE>` явяется однострочным тегом, то бишь его не нужно закрывать `</br>`;
* `<hi>` — элемент, который определяет строчный заголовок (i определяет уровень заголовка от 1 до 6);
* `<ol>` — элемент, который определяет нумерованный список;
* `<omg>` — элемент, который определяет изображение. Является однострочным.

Имеем в итоге такую страничку:
![Первый скриншот](Screenshots/Screenshot11.png)

Добавим к файлу таблицы следующие код:

```css
/* Body style */
    body {
        background: #333333;
        color: #EEEEEE;
        margin: 0px;
        font-size: 20px;
        font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
        text-align: justify;
    }
    /* Footer style */
    footer {
        background: #000000;
        padding: 25px;
        font-family: cursive;
    }
    /* Hyperlinks style */
    a {
        text-decoration: none;
        color: #EEEEEE;
    }
    /* Style for logo div */
    div.Word {
        display: inline-block;
        margin: 10px 0px 0px 30px;
        text-align: center;
    }
    /* Style for header div of headers */
    div.H2 h2 {
        display: inline-block;
    }
    /* Style for back-link div of headers */
    div.H2 a {
        text-decoration: inherit;
        float: right;
        /* margin of h2 + 10px/2 of font-size difference */
        margin: 29.900px 0px;
    }
    /* Menu list style */
    ul.Menu {
        list-style: none;
        text-align: center;
    }
    /* Elements of menu style */
    ul.Menu li {
        display: inline-block;
        padding: 8px;
    }
    /* Links' elements style */
    ul.Menu li a {
        color: white;
        text-decoration: inherit;
    }
    /* Menu bar div style */
    div.MenuBar {
        display: inline-block;
    }
    /* Header div style */
    div.Header {
        background-color: black;
        font-weight: bold;
        font-size: 20px;
        font-family: cursive;
    }
    /* Images div style */
    div.Image {
        margin: 20px 0px;
        text-align: center;
    }
    /* Content div style */
    div.Content {
        margin: 25px;
    }
    /* Footer sources div style */
    footer div.Source {
        display: inline-block;
    }
    /* Footer links style */
    footer div.Links {
        display: inline-block;
    }
    /* Link lists style */
    ul.LinksList {
        list-style: none;
        /* to remove 40px padding-inline-start, that actually created from nothing */
        padding-inline-start: 0px;
    }
    /* Style for elements of link list */
    ul.LinksList li {
        display: inline-block;
        padding: 0px 0px 0px 8px;
    }

```

Для изменения цвета элемента мы используем атрибут `color`, для изменения цвета фона —  `background`. Чтобы поменять шрифт достаточно воспользоваться тегом `font-family`, для указания размера используют тег `font-size`. Для выравнивания текста —  `text-aligment`.

Теперь страница выглядит так:
![Второй скриншот](Screenshots/Screenshot12.png)

### Лабораторная работа №2

#### Постановка задачи:

1. Создать гипертекстовую страницу, которая будет включать несколько абзацев и одно растровое изображение.
2. Растровое изображение поместить непосредственно в html-документ с помощью протокола data.
3. Создать смену стилей абзацев при наведении графического курсора. При наведении мыши на графическое изображение оно должно увеличиваться вдвое.

#### Выполнение:

Возьмем разметку из предыдущей работы. Воспользуемся [этим сайтом](https://www.cy-pr.com/tools/datauri) для того, чтобы сконвертировать все изображения на страничке в формат data. Одно из изображений будет выглядить в коде так:

```html
    <div class="Image">
        <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAqwAAAJwCAIAAADDXKQz...==g">
    </div>
```

Для смены стилей при событии наведения на элемент мышки воспользуемся псевдоклассом CSS `:hover`. Например, стиля элемента `<a>` можно поменять так:

```css
    /* Hyperlinks style */
    a {
        transition: color 0.5s ease-out;
        text-decoration: none;
        color: #EEEEEE;
    }
    a:hover {
        color: yellow;
    }
```

Использованные атрибуты:
* `transition` — используется для задержки смены свойства (вид свойства, время задержки, функция, по которой производится изменения);
* `text-decoration` — используется для указания украшения текста;
* `color` — используется для указания цвета. Можно использовать стандартные названия цветов, RGB, RGBA, HSL;

### Лабораторная работа №3-4

#### Постановка задачи:

1. Создать гипертекстовую страницу с формой, куда поместить как можно больше доступных элементов, которые позволяют ввода данных.
2. Написать CGI-скрипт. С помощью скрипта обеспечить контроль данных перед отправкой их на сервер.
3. Для обработки данных, пересланых на сервер, создать CGI-скрипт. Раскодированые данные вернуть клиенту.
4. Написать скрипт для авторизации клиента. Данные о клиентах сохранять в БД.

#### Выполнение:

Возьмем страницу из предыдущих лабораторных работ. Добавим html-шаблоны для регистрации и авторизации, в которых и будут использован `<form>`.

```html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
    <link rel="stylesheet" type="text/css" href="../static/style.css">
</head>
<body class="logBody">
    <div class="logDiv">

        {{ if . }}
            <div class="error">
                {{ . }}
            </div>
        {{ end }}

        <div class="H2" style="display: flex; justify-content: center">
            <h2>Login:</h2>
        </div>

        <div>
            <form method="post">
                <div class="formDiv">
                    <p>Username:</p><input name="username">
                </div>
                <div class="formDiv">
                    <p>Password:</p><input name="password" type="password">
                </div>
                <div style="display: flex; justify-content: center; flex-direction: column">
                    <div class="formDiv" style="display: flex; justify-content: center; flex-direction: column">
                        <button type="submit">Submit</button>
                    </div>
                    <div class="formDiv" style="display: flex; justify-content: center">
                        <a href="/register">Register</a>
                    </div>
                </div>
            </form>
        </div>

    </div>
</body>
</html>

```

``` html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Register</title>
    <link rel="stylesheet" type="text/css" href="../static/style.css">
</head>
<body class="logBody">
    <div class="LogDiv">
        {{ if . }}
            <div class="error">
                {{ . }}
            </div>
        {{ end }}

        <div class="H2" style="display: flex; justify-content: center">
            <h2>Register:</h2>
        </div>

        <div>
            <form method="post">
                <div class="formDiv">
                    <p>Username:</p><input name="username" type="text">
                </div>
                <div class="formDiv">
                    <p>Password:</p><input name="password" type="password">
                </div>
                <div class="formDiv">
                    <p>Email:</p><input name="email" type="email">
                    <input type="checkbox" name="emailHide">Hide?</input>
                </div>
                <div class="formDiv">
                    <p>Phone:</p><input name="phone" type="tel" pattern="(\+?\d[- .]*){7,13}">
                    <input name="phoneHide" type="checkbox">Hide?</input>
                </div>
                <div class="formDiv">
                    <p>Birthdate:</p><input name="birth" type="date">
                </div>
                <div class="formDiv">
                    <input name="gender" type="radio" checked="checked">M</input>
                    <input name="gender" type="radio">W</input>
                    <input name="gender" type="radio">N</input>
                </div>
                <div class="formDiv" style="display: flex; justify-content: center;">
                    <input name="button" type="submit" value="Submit">
                </div>
            </form>
        </div>
    </div>
</body>
</html>

```

Также создадим шаблон странички пользователя, на которой будут отображаться данные о нем.

```html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .User.GetUsername }}</title>
</head>
<body style="display: flex; justify-content: center; flex-direction: column">

    <div>

        <p>Username: {{ .User.GetUsername }}</p>

        <p>Birthday: {{ .User.GetBirth }}</p>

        <p>Gender: {{ .User.GetGender }}</p>

        <p>Email: {{ .User.GetEmail }}</p>

        <p>Phone: {{ .User.GetPhone }}</p>

    </div>

</body>
</html>

```

Для этих шаблонов была модифицирована таблица стилей:

```css

    /* Body style */
    body {
        background: #333333;
        /*color: #EEEEEE;*/
        margin: 0px;
        font-size: 20px;
        font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
        text-align: justify;
    }
    /* Footer style */
    footer {
        background: #000000;
        padding: 25px;
        font-family: cursive;
    }
    /* Hyperlinks style */
    a {
        transition: color 0.5s ease-out;
        text-decoration: none;
        color: #EEEEEE;
    }
    a:hover {
        color: yellow;
    }
    /* Style for logo div */
    div.Word {
        display: inline-block;
        margin: 10px 0px 0px 30px;
        text-align: center;
    }
    /* Style for header div of headers */
    div.H2 h2 {
        color: #EEEEEE;
        display: inline-block;
    }
    div.Text {
        color: white;
    }
    /* Style for back-link div of headers */
    div.H2 a {
        text-decoration: inherit;
        float: right;
        /* margin of h2 + 10px/2 of font-size difference */
        /* TODO remove this disgusting piece of madness and make some shared div margin */
        margin: 29.9px 0px;
    }
    /* Menu list style */
    ul.Menu {
        list-style: none;
        text-align: center;
    }
    /* Elements of menu style */
    ul.Menu li {
        display: inline-block;
        padding: 8px;
    }
    /* Links' elements style */
    ul.Menu li a {
        color: white;
        text-decoration: inherit;
    }
    /* Menu bar div style */
    div.MenuBar {
        display: inline-block;
    }
    /* Header div style */
    div.Header {
        background-color: black;
        font-weight: bold;
        font-size: 20px;
        font-family: cursive;
    }
    /* Images div style */
    div.Image {
        margin: 20px 0px;
        text-align: center;
    }

    /* Hover image resize */
    div.Image img {
        transition: all 0.5s ease-out;
        height: 300px;
        width: auto;
    }
    div.Image img:hover {
        height: 600px;
        width: auto;
        cursor: pointer;
    }
    /* Content div style */
    div.Content {
        margin: 25px;
    }
    /* Footer sources div style */
    footer div.Source {
        display: inline-block;
        color: #EEEEEE;
    }
    /* Footer links style */
    footer div.Links {
        display: inline-block;
    }
    /* Link lists style */
    ul.LinksList {
        list-style: none;
        /* to remove 40px padding-inline-start, that actually created from nothing */
        padding-inline-start: 0px;
    }
    /* Style for elements of link list */
    ul.LinksList li {
        display: inline-block;
        padding: 0px 0px 0px 8px;
    }



    div.commentsSection {
        padding: 25px;
        background: #222222;
    }

    div.commentsSection div.addCommentForm {
        text-align: right;
    }

    div.commentsSection div.addCommentForm textarea {
        min-height: 50px;
        width: 100%;
        margin: 0.5em 0em;
        padding: 0px;
        resize: vertical;
    }

    div.commentsSection > div.comments {
        border: 1px solid #AAAAAA;
        padding: 0.5em;
        margin: 1em 0em;
        color: #CCCCCC;
    }

    div.error {
        color: red;
    }

    body.logBody {
        display: flex;
        flex-direction: inherit;
        justify-content: center;
        color: #DDDDDD;
    }

    div.formDiv {
        margin: 10px;
        font-size: 16px;
    }

    div.logoutSection {
        display: inline-block;
        float: right;
        margin: 16px 3px;
        padding: 0.5em;
        color: #EEEEEE;
    }

```

Писать серверную часть будем на языке программирования Go. В `main` пропишем следующее строки:

```Golang

package main

import (
	"./models"
	"./routes"
	"./utilities"
	"fmt"
	"net/http"
)

func main() {
	models.InitializeDB()
	utilities.ParseTemplates("templates/*.html")
	http.Handle(routes.UrlPath, routes.NewRouter())
	fmt.Println("Server has been launched...")
	_ = http.ListenAndServe(":8888", nil)
}

```

Здесь инициализируется БД (использован Redis), парсятся html-шаблоны и размещается сервер на `localhost:8888`. Функции, которые будут вызываться при post и get запросах инициализируются в `routes`. Код этого модуля ниже:

```Golang

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

```

Теперь при заходе на стартовую страчнику будет проверяться текущая сессия. Если вы не инициализированы, то вас перекинет на страничку с логированием и ссылкой на регистрацию. Перейдя по ссылке и указав необходимые данные, вас перекинет на страницу, где вы сможете вести логин и пароль и зайти на страницу с контентом. Внизу будет размещена секция с комментариями, в которой вы сможете оставить свой или прочитать чужой комментарий такого же зарегестрированного пользователя, а также перейти на его страничку и почтитать о нем.

![Первый скриншот](Screenshots/Screenshot31.png)
![Второй скриншот](Screenshots/Screenshot32.png)
![Третий скриншот](Screenshots/Screenshot33.png)
![Четвертый скриншот](Screenshots/Screenshot34.png)
