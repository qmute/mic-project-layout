package hdl_test

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/front/hdl"
	"easyslip.cc/mic-project-layout/pkg/ut"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pub", func() {
	var app *hdl.PubHdl

	BeforeEach(func() {
		app = &hdl.PubHdl{
			Base: base,
		}
		app.Mount(router)
	})

	It("hello", func() {
		path := "/hello"
		r := ut.Req(ut.GET, path, nil, nil)
		w := ut.Serve(router, r)
		Ω(w).To(HaveHTTPStatus(200))
		Ω(w.Body.String()).To(MatchJSON(`{
          "id": 1,
          "name": "a",
          "mobile": "13088888888",
          "status": 1,
          "super": true,
          "ct": 1
        }`))
	})

})
