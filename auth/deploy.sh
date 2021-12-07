docker build . -t flurach/$(basename `pwd`)
docker push flurach/auth
kubectl apply -f k8s.yml
