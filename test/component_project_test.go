package main_test

import (
	"log"
	"net/http"

	a "component_project/controller"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ComponentProject", func() {
	var (
		user a.UserName
	)
	Describe("正确的流程", func() {

		BeforeEach(func() {
			user = a.UserName{
				UserName: "zhangsan",
				PassWord: "123456",
			}
		})
		Context("aaaa", func() {
			It("xxxx", func() {
				b := user.InitInfo()
				log.Println(b)
				By("test1")
				Ω(user.Auth("123456")).Should(Equal(http.StatusOK))
				Ω(user.Auth("123456")).Should(Equal(http.StatusBadRequest))
			})
		})
	})
})

func processInfo(user a.UserName, expect int) {
	param := user.InitInfo()
	log.Println(param)
	Ω(user.Auth("123456")).Should(Equal(expect))
}
