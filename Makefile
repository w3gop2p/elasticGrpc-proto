protoIngestWorker:
	protoc --go_out=golang --go_opt=paths=source_relative \
			--go-grpc_out=golang --go-grpc_opt=paths=source_relative \
			data_ingest_worker/ingestworker.proto

protoStoreService:
	protoc --go_out=golang --go_opt=paths=source_relative \
			--go-grpc_out=golang --go-grpc_opt=paths=source_relative \
			data_store_service/storeservice.proto

.PHONY: proto