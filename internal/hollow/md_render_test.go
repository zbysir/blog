package hollow

import (
	"testing"
)

func TestMd(t *testing.T) {
	s := renderMd([]byte(`## h2
![在 Golang 中尝试“干净架构”](../../static/img/在%20Golang%20中尝试干净架构_1.png)
![在 Golang 中尝试“干净架构”](/static/img/在%20Golang%20中尝试干净架构_1.png)

在[文中](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)，他提出的干净架构是这样的：

`))

	t.Logf("%s", s)
}

func TestMdRender(t *testing.T) {
	m := newMdRender(func(s string) string {
		return s
	})
	a := m.Render([]byte(`## h2
![在 Golang 中尝试“干净架构”](../../static/img/在%20Golang%20中尝试干净架构_1.png)
![在 Golang 中尝试“干净架构”](/static/img/在%20Golang%20中尝试干净架构_1.png)

在[文中](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)，他提出的干净架构是这样的：

`))
	t.Logf("%s", a)
}
