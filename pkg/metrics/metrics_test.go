package metrics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metrics "github.com/rite2nikhil/kubernetes-node-scaler/pkg/metrics"
	"k8s.io/api/core/v1"
)

type FakeKubeClient struct{}

func (f *FakeKubeClient) ListNodes() (objs []*v1.Node, err error) {
	objs := make([]*v1.Node)

	return objs, nil
}

var _ = Describe("Metrics", func() {
	k8s := FakeKubeClient{}
	s := metrics.NewNodeScaller(k8s)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})
