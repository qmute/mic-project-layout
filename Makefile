
.PHONY: fmt
fmt:
	go fmt ./...


.PHONY: wire
wire:
	wire gen ./internal/app/serve
	wire gen ./internal/app/daemon

.PHONY: run-serve
run-serve: wire
	go run main.go serve

.PHONY: run
run:
	make run-serve

.PHONY: run-daemon
run-daemon: wire
	go run main.go daemon

.PHONY: test-all # 测试全部，包括repo
test-all: fmt wire
	ginkgo -r .

.PHONY: test # 测试时不必分app，因为所有app都需要保证通过
test: fmt wire
	ginkgo -r --skip-package internal/repo/impl .

.PHONY: swag-front
swag-front:
	swag init -g ./front.go -d ./internal/app/serve/front -o ./doc/front  --instanceName front

.PHONY: swag-admin
swag-admin:
	swag init -g ./admin.go -d ./internal/app/serve/admin -o ./doc/admin  --instanceName admin

.PHONY: swag-all
swag-all: swag-front swag-admin
