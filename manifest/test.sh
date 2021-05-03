#!/bin/bash

kubectl get pods
kubectl run --restart=Never nginx --image=nginx


echo "wait"
sleep 3
echo "start test"

kubectl exec -it nginx -- curl -m 5 http://sample-app
if [ $? -gt 0 ]; then
    echo "sample-app is NOT working correctly"
    kubectl delete pod nginx
    exit 1
else 
    echo "sample-app is working correctly"
    kubectl delete pod nginx
    exit 0
fi