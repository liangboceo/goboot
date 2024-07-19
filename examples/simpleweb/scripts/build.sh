docker build -t yuanbootdemo-dev:v0.0.1 .

docker run --rm -p 8080:8080  --name yoyodemov1 yuanbootdemo:v1

kubectl apply -f k8s-deploy.yaml

kubectl apply -f k8s-service.yaml

kubectl expose deployment yuanbootdemo --type="LoadBalancer"
