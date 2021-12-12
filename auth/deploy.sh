docker build . -t flurach/$(basename `pwd`)
docker push flurach/$(basename `pwd`)
kubectl apply -f k8s.yml
