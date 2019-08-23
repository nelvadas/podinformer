package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	var namespace, configFile string

	flag.StringVar(&namespace, "n", "default", "Your Namespace ")
	//read config file or pick the default kubeconfig file
	flag.StringVar(&configFile, "config", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "Config file ")
	flag.Parse()

	fmt.Printf("Running with kube config file %v\n", configFile)

	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err.Error())
	}

	//Create a client set from config
	clientset, err := kubernetes.NewForConfig(config)

	//Create an informer from client
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute*1)
	podInformer := informerFactory.Core().V1().Pods()

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{

		AddFunc: func(obj interface{}) {
			newPod := obj.(*v1.Pod)
			fmt.Printf("new pod added %v\n", newPod.GetName())
		},

		UpdateFunc: func(old, new interface{}) {
			//fmt.Printf("new pod updated  %s \n", old)
		},

		DeleteFunc: func(new interface{}) {
			//fmt.Println("new pod deleted ")
		},
	})

	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)

	stopper := make(chan struct{})
	defer close(stopper)
	podInformer.Informer().Run(stopper)

}
