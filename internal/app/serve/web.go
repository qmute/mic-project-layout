package serve

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"easyslip.cc/mic-project-layout/internal/app/serve/admin"
	"easyslip.cc/mic-project-layout/internal/app/serve/front"
	"easyslip.cc/mic-project-layout/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qmute/gi"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Web struct {
	Config config.Config
	Front  *front.App // app用户模块
	Admin  *admin.App // 管理后台模块
}

func (p *Web) Mount(version string) *gin.Engine {
	router := gi.New(gi.WithPprof())
	router.Use(
		func(c *gin.Context) {
			agent := c.Request.Header.Get("User-Agent")
			if strings.Contains(strings.ToLower(agent), strings.ToLower("DingTalkBot")) {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		},

		gi.MidRequestId(),
		gi.MidHSTS(),
		gi.MidCORS(),
		gi.MidRecovery(),
		gi.MidLogger(
			gi.LogWithField("platform", func(r *http.Request) any { return r.Header.Get("X-Platform") }),
			gi.LogWithField("clientVersion", func(r *http.Request) any { return r.Header.Get("X-Client-Version") }),
			gi.LogWithField("header", func(r *http.Request) any { return r.Header }),
			gi.LogWithField("query", func(r *http.Request) any { return r.URL.RawQuery }),
		),
	)

	root := router.Group("")

	root.HEAD("/v", gi.HdlVersion(version))
	root.GET("/v", gi.HdlVersion(version))

	api := router.Group("/api")

	// 用户
	p.Front.Run(api.Group("/front"))

	// 管理后台
	p.Admin.Run(api.Group("/admin"))

	return router
}

func NewReverseProxyHandler(target *url.URL) gin.HandlerFunc {
	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(target)
			uri := r.Out.URL
			if !strings.HasSuffix(uri.Path, "apple-app-site-association") {
				lastPath, _ := lo.Last(strings.Split(uri.Path, "/"))
				// path最后一部分没有"."，说明不是文件，是文件夹，则为其添加index.html（未作详细文件类型匹配）
				if !strings.Contains(lastPath, ".") {
					uri.Path, _ = url.JoinPath(uri.Path, "index.html")
					uri.RawPath, _ = url.JoinPath(uri.RawPath, "index.html")
				}
			}
		},
		ModifyResponse: func(r *http.Response) error {
			// 不正常
			if r.StatusCode >= 400 {
				b, _ := io.ReadAll(r.Body)
				log.Warningln("proxy fail", string(b))
				return errors.Errorf("not found %s", target.String())
			}
			return nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("not found"))
		},
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorln("proxy err", err)
				c.AbortWithStatus(http.StatusServiceUnavailable) // 当作临时不可用
			}
		}()
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
