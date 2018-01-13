build-dir:
	mkdir -p bin

build: build-dir
	go build -o bin/elb-gc ./

test:
	go test ./...

show-tags: build
	./bin/elb-gc -region=us-west-2 -profile=dev -cmd=show-tags -logtostderr -v=4

gc-elbs: build
	./bin/elb-gc -region=us-west-2 -profile=dev -tag-filter=kubernetes.io/cluster/pipeline,kubernetes.io/cluster/publicapi,kubernetes.io/cluster/gke-foo,kubernetes.io/cluster/etcd-test,kubernetes.io/cluster/g-ashisha,kubernetes.io/cluster/repro1,kubernetes.io/cluster/repro,kubernetes.io/cluster/k8s175,kubernetes.io/cluster/k8supgrade,kubernetes.io/cluster/g-test,kubernetes.io/cluster/g-foo,kubernetes.io/cluster/g-master,kubernetes.io/cluster/asdf,kubernetes.io/cluster/g-temp,kubernetes.io/cluster/gke-temp,kubernetes.io/cluster/test,kubernetes.io/cluster/g-gitlab -cmd=gc -logtostderr -v=9