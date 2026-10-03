// Modified from https://github.com/arina999999997/elichika/blob/463b37a0d69dbe68fff7112fe87105cfb2e260d6/router/router.go
package router

import (
	"honoka-chan/config"
	honokautils "honoka-chan/internal/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type HandlerInfo struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
}
type SpecialGroupSetup = func(*gin.RouterGroup)
type GroupInfo struct {
	InitialHandlers []gin.HandlerFunc
	Handlers        map[string]HandlerInfo
	SpecialSetups   []SpecialGroupSetup
}

var (
	groups    = map[string]*GroupInfo{}
	templates []string
)

func initGroup(g string) {
	_, exist := groups[g]
	if exist {
		return
	}
	groups[g] = &GroupInfo{
		InitialHandlers: []gin.HandlerFunc{},
		Handlers:        map[string]HandlerInfo{},
		SpecialSetups:   []SpecialGroupSetup{},
	}
}

func (g *GroupInfo) AddInitialHandler(handler gin.HandlerFunc) {
	g.InitialHandlers = append(g.InitialHandlers, handler)
}

func (g *GroupInfo) AddHandler(method, path string, handlers ...gin.HandlerFunc) {
	_, exist := g.Handlers[method+path]
	if exist {
		panic("Multiple handler for path and method: " + method + " " + path)
	}
	g.Handlers[method+path] = HandlerInfo{
		Method:   method,
		Path:     path,
		Handlers: handlers,
	}
}

func (g *GroupInfo) AddSpecialSetup(specialGroupSetup SpecialGroupSetup) {
	g.SpecialSetups = append(g.SpecialSetups, specialGroupSetup)
}

func AddInitialHandler(group string, handler gin.HandlerFunc) {
	initGroup(group)
	groups[group].AddInitialHandler(handler)
}
func AddHandler(group, method, path string, handlers ...gin.HandlerFunc) {
	initGroup(group)
	groups[group].AddHandler(method, path, handlers...)
}

func AddSpecialSetup(group string, specialGroupSetup SpecialGroupSetup) {
	initGroup(group)
	groups[group].AddSpecialSetup(specialGroupSetup)
}

func AddTemplates(path string) {
	templates = append(templates, path)
}

func SifRouter(r *gin.Engine) {
	// Static
	r.Static("/static", "static")

	var files []string
	_ = filepath.Walk("static/templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	r.LoadHTMLFiles(files...)

	// favicon
	r.StaticFile("/favicon.ico", "static/images/favicon.ico")

	// session
	store := cookie.NewStore([]byte(config.Conf.Settings.SessionKey))
	r.Use(sessions.Sessions("llsif", store))

	// manga
	r.GET("/manga", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "common/manga.html", gin.H{})
	})

	// webui
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index")
	})

	notFoundHandler := func(ctx *gin.Context) {
		if !honokautils.IsMainPHPRequest(ctx.Request.URL.Path) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		honokautils.AbortMaintenanceJSON(ctx, http.StatusNotFound, honokautils.NewNotFoundContent(ctx.Request.URL.Path))
	}
	r.NoRoute(notFoundHandler)
	r.NoMethod(notFoundHandler)

	// router
	for groupPath, groupInfo := range groups {
		groupApi := r.Group(groupPath, groupInfo.InitialHandlers...)
		for _, handlerInfo := range groupInfo.Handlers {
			switch handlerInfo.Method {
			case "POST":
				groupApi.POST(handlerInfo.Path, handlerInfo.Handlers...)
			case "GET":
				groupApi.GET(handlerInfo.Path, handlerInfo.Handlers...)
			case "PUT":
				groupApi.PUT(handlerInfo.Path, handlerInfo.Handlers...)
			case "DELETE":
				groupApi.DELETE(handlerInfo.Path, handlerInfo.Handlers...)
			default:
				panic("Must be GET, POST, PUT or DELETE")
			}
		}
		for _, specialSetup := range groupInfo.SpecialSetups {
			specialSetup(groupApi)
		}
	}
}
