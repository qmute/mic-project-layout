package hdl_test

import (
	"testing"

	"easyslip.cc/mic-project-layout/internal/app/serve/admin/hdl"
	"easyslip.cc/mic-project-layout/internal/app/serve/admin/hdl/mid"
	"easyslip.cc/mic-project-layout/pkg/ut"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/qmute/gi"
	"go.uber.org/mock/gomock"
)

func TestHdl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AdminHdl Suite")
}

var _ = BeforeSuite(func() {
	ut.Verbose(false) // 测试期间保持安静
})

var ctl *gomock.Controller
var cleaner func()
var router *gin.Engine
var base hdl.Base

var _ = BeforeEach(func() {
	ctl = gomock.NewController(GinkgoT())
	cleaner = ctl.Finish

	base = hdl.Base{
		Mid: &mid.Mid{},
	}

	gin.SetMode(gin.TestMode)

	router = gi.New()

	router.Use(gi.MidCookieSession("web-session-cinema", "web-session-cinema-secret", sessions.Options{
		Path:   "/api/admin",
		MaxAge: 86400 * 7,
	}))

})

var _ = AfterEach(func() { cleaner() })
