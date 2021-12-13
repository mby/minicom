[ -d operator ] || git clone --depth=1 https://github.com/mongodb/mongodb-kubernetes-operator.git operator
cd operator
git reset --hard HEAD
git pull

kubectl apply -f config/crd/bases/mongodbcommunity.mongodb.com_mongodbcommunity.yaml
kubectl get crd/mongodbcommunity.mongodbcommunity.mongodb.com

kubectl create namespace mongodb-operator
kubectl apply -k config/rbac/ --namespace mongodb-operator
kubectl create -f config/manager/manager.yaml --namespace mongodb-operator

sed -i '' 's/\<your\-password\-here\>/password/g' config/samples/mongodb.com_v1_mongodbcommunity_cr.yaml
sed -i '' 's/members:\ 3/members:\ 1/g' config/samples/mongodb.com_v1_mongodbcommunity_cr.yaml

kubectl apply -f config/samples/mongodb.com_v1_mongodbcommunity_cr.yaml --namespace mongodb-operator
