up:
	docker-compose up --build -d

down:
	docker-compose down

test:
	cd account && \
	go test -v ./...

postgres:
	docker-compose up postgres-account

.PHONY: up down create-keypair migrate-create migrate-up migrate-down migrate-force

ACCTPATH = ./account
MPATH = $(ACCTPATH)/migrations
PORT = 5432

# 默认 number 执行迁移 up 或者 down
N = 1

create-keypair:
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable force $(VERSION)