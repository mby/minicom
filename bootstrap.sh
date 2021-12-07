for f in `fd deploy.sh`; do
	cd $(dirname $f)
	sh deploy.sh
	cd -
done
