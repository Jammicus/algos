.PHONY: plan


run:
	go run cmd/main.go
# If the score is > blast radius this will return false
test: plan
	go test ./... -v -count 1 

bench: 
	cd internal; go test -bench=. -v

bench-long: 
	cd internal; go test -bench=. -v -count 5
