package main

import (
	"context"
	"fmt"

	//

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	// "k8s.io/client-go/kubernetes"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// config
	// 自动获取配置文件
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	// config.GroupVersion = &schema.GroupVersion{}
	// config.NegotiatedSerializer = scheme.Codecs
	// config.APIPath = "/api"

	// // Restclient
	// clientSet, err := rest.RESTClientFor(config)
	// if err != nil {
	// 	panic(err)
	// }

	// // action get data
	// svc := &v1.Service{}
	// err = clientSet.Get().Namespace("default").Resource("services").Name("test").Do(context.TODO()).Into(svc)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(svc.Spec.Ports)
	// }

	// clientSet 组合config后调用restClient
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	svcClient := clientSet.CoreV1()
	svc, err := svcClient.Services("default").Get(context.Background(), "kubernetes", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(svc.Name)

	// 调用自定义CRD资源
	dyclinet, err := dynamic.NewForConfig(config)
	sj := dyclinet.Resource(schema.GroupVersionResource{
		Group:    "cilium.io",
		Version:  "v2",
		Resource: "ciliumnodes",
	})
	ul, err := sj.List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, u := range ul.Items {
		fmt.Println(u.GetName())
	}

	// discoveryClient
	// disclient,err := discovery.NewDiscoveryClientForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }

}
