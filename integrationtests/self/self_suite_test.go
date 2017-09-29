package self_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/lucas-clemente/quic-go/integrationtests/tools/testlog"
	"github.com/lucas-clemente/quic-go/internal/utils"

	"testing"
)

func TestSelf(t *testing.T) {
	utils.SetLogLevel(utils.LogLevelDebug)
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Self integration tests")
}
