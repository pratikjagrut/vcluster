package e2enode

import (
	"github.com/loft-sh/vcluster/test/framework"
	"github.com/onsi/ginkgo"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = ginkgo.Describe("Node sync", func() {
	f := framework.DefaultFramework
	ginkgo.It("sync all nodes", func() {
		hostNodes, err := f.HostClient.CoreV1().Nodes().List(f.Context, v1.ListOptions{})
		framework.ExpectNoError(err)

		virtualNodes, err := f.VclusterClient.CoreV1().Nodes().List(f.Context, v1.ListOptions{})
		framework.ExpectNoError(err)

		framework.ExpectEqual(len(hostNodes.Items), len(virtualNodes.Items))
	})
})
