package utils

import (
	"bytes"
	"github.com/kubebuilder-demo/api/v1beta1"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netingress "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"text/template"
)

func parseTemplate(templateName string, app *v1beta1.App) []byte {
	tmpl, err := template.ParseFiles("controllers/template/" + templateName + ".yml")
	if err != nil {
		println("创建temp文件失败")
		panic(err)
	}
	b := new(bytes.Buffer)
	err = tmpl.Execute(b, app)
	if err != nil {
		panic(err)

	}
	return b.Bytes()

}

func NewDeployment(app *v1beta1.App) *appv1.Deployment {
	d := &appv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", app), d)
	if err != nil {
		println("创建出现异常")
		panic(err)
	}
	return d
}

func NewIngress(app *v1beta1.App) *netingress.Ingress {
	i := &netingress.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress", app), i)
	if err != nil {
		panic(err)
	}
	return i
}

func NewService(app *v1beta1.App) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", app), s)
	if err != nil {
		panic(err)
	}
	return s
}
