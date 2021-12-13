docker build . -t flurach/auth
kind load docker-image flurach/auth:latest
kubectl apply -f k8s.yml
