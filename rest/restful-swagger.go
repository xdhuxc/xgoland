package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"log"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	restful.Add(ws)

	ws.Route(ws.GET("/{medium}").
		To(noop).
		Doc("Search all books").
		Param(ws.PathParameter("medium", "digital or paperback").DataType("string")).
		Param(ws.QueryParameter("language", "en,nl,de").DataType("string")).
		Param(ws.HeaderParameter("If-Modified-Since", "last known timestamp").DataType("datetime")).
		Do(returns200, returns500))

	ws.Route(ws.PUT("/{medium}").
		To(noop).Doc("Add a new book").
		Param(ws.PathParameter("medium", "digital or paperback").DataType("string")).
		Reads(Book{}))

	config := swagger.Config{
		WebServices:     restful.DefaultContainer.RegisteredWebServices(),
		WebServicesUrl:  "http://localhost:8080",
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Users/wanghuan/GolandProjects/GoPath/src/github.com/xdhuxc/xgoland/dist",
		ApiVersion:      "2.0",
	}
	swagger.RegisterSwaggerService(config, restful.DefaultContainer)
	log.Print("Start listening on localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: restful.DefaultContainer,
	}
	log.Fatal(server.ListenAndServe())

}

func noop(req *restful.Request, resp *restful.Response) {

}

func returns200(b *restful.RouteBuilder) {
	b.Returns(http.StatusOK, "OK", Book{})
}

func returns500(b *restful.RouteBuilder) {
	b.Returns(http.StatusInternalServerError, "There is something wrong", nil)
}
