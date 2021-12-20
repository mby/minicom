cluster_exists=`kubectl cluster-info --context kind-kind 2> /dev/null`
[ "$cluster_exists" != 0 ] && kind create cluster

cd kubernetes-dashboard
sh deploy.sh
cd -

cd auth
sh deploy.sh
cd -
