# Using Kubernetes Informes
Informers can react to change of objects nearly in real time and do not require Polling request as watch function.

This demo shows a simple pods Informer can be implemented to watch pods in a namespace 


#Run

````
go run main.go  -n demo -config /Users/xa001/bundle/kube.yml 
````

