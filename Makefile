.PHONY: decrypt-everything
decrypt-everything:
	blackbox_postdeploy

.PHONY: operator-key-link
operator-key-link:
	ln -s ../deploy/env.operator likecoin/.env
	ln -s ../deploy/env.operator likenft/.env

.PHONY: remove-operator-key-link
remove-operator-key-link:
	rm -f likecoin/.env
	rm -f likenft/.env

.PHONY: local-contracts
local-contracts:
	(sleep 1 && make -C operation init-local-state) &
	(sleep 2 && make -C likenft deploy-local) &
	(sleep 7 && make -C likecoin deploy-local)

.PHONY: abigen
abigen:
	make -C likenft build
	make -C likecoin build
	mkdir -p abi
	cp likenft/artifacts/contracts/LikeProtocol.sol/LikeProtocol.json abi/
	cp likenft/artifacts/contracts/BookNFT.sol/BookNFT.json abi/
	cp likecoin/artifacts/contracts/EkilCoin.sol/EkilCoin.json abi/
	make -C likenft-indexer abigen
	make -C migration-backend abigen

.PHONY: setup
setup:
	make -C signer-backend secret
	make -C migration-backend secret
	make -C likenft-indexer secret

.PHONY: start
start:
	docker compose up -d
	docker compose logs -f

.PHONY: stop
stop:
	docker compose stop
	docker compose rm -f

.PHONY: clean-docker-volumes
clean-docker-volumes:
	docker compose down
	docker compose rm -f
	docker volume ls | grep 'likecoin-30' | awk '{ print $$2 }' | xargs docker volume rm

.PHONY: run-migrations
run-migrations:
	docker compose run --rm migration-backend make run-migration
	docker compose run --rm signer-backend make run-migration

.PHONY: docker-images
docker-images:
	DOCKER_BUILD_ARGS=--push make -C migration-backend docker-image
	DOCKER_BUILD_ARGS=--push make -C signer-backend docker-image
	DOCKER_BUILD_ARGS=--push make -C likenft-indexer docker-image
	DOCKER_BUILD_ARGS=--push make -C likenft-migration docker-image
	DOCKER_BUILD_ARGS=--push make -C likecoin-migration docker-image
	DOCKER_BUILD_ARGS=--push make -C migration-admin docker-image

.PHONY: deploy
deploy: decrypt-everything
	make -C deploy deploy
