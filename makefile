TASK1 := 1/main_1.go

task1:
	go run $(TASK1)

test1:
	go test -v 1/*.go