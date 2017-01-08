package yampatch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestYampatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Yampatch Suite")
}
