build:
	forge build -C lib/gmx-contracts -o out/gmx-contracts --extra-output-files abi --force
	forge build -C lib/gmx-synthetics -o out/gmx-synthetics --extra-output-files abi --force

generate:
	go run cmd/*.go

clean:
	find . -name '*.go' -not -path './example/*' -not -path './cmd/*' -delete
	find . -empty -type d -delete
