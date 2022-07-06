package main

import (
	"context"
	"fmt"

	//

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientv1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

func main() {

	// config
	// 自动获取配置文件
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	fmt.Println(clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	config.GroupVersion = &clientv1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	clientSet, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// action get data
	svc := &v1.Service{}
	err = clientSet.Get().Namespace("default").Resource("services").Name("sonarqube").Do(context.TODO()).Into(svc)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(svc.Spec.Ports)
	}

}
