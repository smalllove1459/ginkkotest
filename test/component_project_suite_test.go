package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComponentProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ComponentProject Suite")
}
