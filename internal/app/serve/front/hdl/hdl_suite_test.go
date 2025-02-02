package hdl_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/qmute/gi"
	"github.com/qmute/mic-project-layout/internal/app/serve/front/hdl"
	"github.com/qmute/mic-project-layout/internal/app/serve/front/hdl/mid"
	"github.com/qmute/mic-project-layout/internal/domain"
	"github.com/qmute/mic-project-layout/internal/testdata"
	"github.com/qmute/mic-project-layout/pkg/ut"
	"go.uber.org/mock/gomock"
)

func TestHdl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FrontHdl Suite")
}

var _ = BeforeSuite(func() {
	ut.Verbose(false) // 测试期间保持安静
})

var ctl *gomock.Controller
var cleaner func()
var router *gin.Engine
var base hdl.Base
var ctx context.Context
var defaultClientHeader http.Header

var _ = BeforeEach(func() {
	ctx = context.Background()
	ctl = gomock.NewController(GinkgoT())
	cleaner = ctl.Finish

	base = hdl.Base{}

	gin.SetMode(gin.TestMode)

	clientInfo := testdata.ClientInfo(domain.ClientTypeIOS)
	// defaultClientHeader = http.Header{}
	// defaultClientHeader.Set(mid.XClientType, clientInfo.ClientType.String())
	// defaultClientHeader.Set(mid.XClientVersion, clientInfo.ClientVersion)
	// defaultClientHeader.Set(mid.XUserType, strconv.Itoa(int(domain.UserTypeStudent)))
	// defaultClientHeader.Set("X-Real-IP", clientInfo.IP)

	router = gi.New()

	router.Use()

	router.Use(
		gi.MidCookieSession("web-session-front", "web-session-front-secret", sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 7,
		}),

		func(c *gin.Context) {
			c.Set(mid.ClientInfoKey, clientInfo)
		})
})

var _ = AfterEach(func() { cleaner() })
