run:
	# vars
	# export NODE1=192.168.1.21
	# export REGISTRY=gcr.io/etcd-development/etcd
	# create etcd mount
	docker volume create --name etcd-data
	export DATA_DIR=etcd-data
	# start containers
	docker compose up

kill:
	docker compose down
	docker image rm example_syncer-example
