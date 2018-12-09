package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"os"
	"time"
)

var err error

type User struct {
	/**
		如果字段名称为 Id，而且类型为 int64并没有定义tag，则会被 xorm 视为主键并拥有自增属性。
		如果想用 Id 以外的名字或非 int64 类型作为主键名，必须在对应的 Tag 上加上 xorm:"pk" 来定义主键，加上 xorm:"autoincr" 作为自增。
	 */
	UserId int32 `xorm:"int(11) pk autoincr 'user_id'"`
	UserName string `xorm:"varchar(150) notnull unique 'user_name'"`
	Password string `xorm:"varchar(150) notnull 'password'"`
	Email string `xorm:"varchar(150) notnull 'email'"`
	Age int8 `xorm:"int(4) 'age'"`
	CreateTime time.Time `xorm:"DateTime created 'create_time'"`
	UpdateTime time.Time `xorm:"DateTime updated 'update_time'"`
}

/**
	处理 404 错误
 */
func NotFound(ctx iris.Context) {
	// 对大文件启用 gzip 压缩
	ctx.Gzip(true)
	// 从 views 目录下寻找该文件
	ctx.View("404.html")
}

/**
	处理 500 错误
 */
func InternalServerError(ctx iris.Context) {
	ctx.View("500.html")
}

type UserController struct {

}

func MainApp() *iris.Application {
	app := iris.New()
	// 设置日志级别为 debug
	app.Logger().SetLevel("debug")
	/**
		禁用日志
		app.Logger().SetLevel("disable")
	 */

	app.Use(recover.New())
	app.Use(logger.New())

	/**
		使用标准的 html/template 引擎定义模板，
		从 ./views 目录下加载所有扩展名为 .html 的模板并解析它们，
		在每次请求时加载模板，仅用于开发模式
	 */
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))
	// 错误处理
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, InternalServerError)

	mvc.New(app).Handle(new(UserController))

	return app
}


func SqlEngine() *xorm.Engine {
	/**
		一般情况下，如果只操作一个数据库，只需要创建一个 engine 即可，engine 是 Goroutine 安全的
		engine 可以通过 engine.Close 来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭。
	 */
	engine, err := xorm.NewEngine("mysql", "root:19940423@tcp(127.0.0.1:3306)/xgolang?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 在控制台打印出生成的 SQL 语句
	engine.ShowSQL(true)
	// 在控制台打印调试及以上的信息
	engine.Logger().SetLevel(core.LOG_DEBUG)
	// 显示执行时间
	engine.ShowExecTime(true)

	sqlLogName := os.Getenv("SQL_LOG_NAME")
	/**
		字符串判空不能使用 str != nil
	 */
	if sqlLogName == "" {
		sqlLogName = "sql.log"
	}

	// 将信息保存为文件
	f, err := os.Create(sqlLogName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	return engine
}

func MainLogger() *log.Logger {
	logFileName := os.Getenv("LOG_FILE_NAME")
	if logFileName == "" {
		logFileName = "xgoland"
	}
	logFile, err := os.Create(logFileName + "_" + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		fmt.Println(err)
	}

	xlogger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	return xlogger
}


func main() {
	app := MainApp()
	mainLogger := MainLogger()

	// 采用同步的方式初始化数据库表
	engine := SqlEngine()
	if engine != nil {
		err = engine.Sync2(new(User))
		if err != nil {
			mainLogger.Fatalf("%v", err)
		}
	}

	user := User{UserName: "xdhuxc", Password: "Xdhuxc163", Email: "xdhuxc@163.com", Age: 24}
	_, err := engine.Insert(&user)
	if err != nil {
		mainLogger.Printf("%v", err)
	}

	xuser := User{UserName: "xdhuxc"}
	_, err = engine.Get(&xuser)
	if err != nil {
		mainLogger.Printf("%v", err)
	}
	mainLogger.Println(xuser)



	app.Handle("GET", "/", func(ctx iris.Context) {
		// ctx.HTML("<h1>Welcome To Use Iris!</h1>")

		ctx.ViewData("Name", "xdhuxc")
		ctx.View("index.html")


	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/user/{name}", func(ctx iris.Context){
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	// app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./conf/iris.tml")), iris.WithoutServerError(iris.ErrServerClosed))
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./conf/iris.tml")))

}


func (u *UserController) Get() string {
	return "abc"
}

func (u *UserController) GetHello() interface{} {
	return map[string]string{"message": "Hello World!"}
}