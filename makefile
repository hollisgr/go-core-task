TASK1 := 1/main_1.go
TASK2 := 2/main_2.go
TASK3 := 3/main_3.go

TEST1 := 1/*.go
TEST2 := 2/*.go
TEST3 := 3/*.go

task1:
	go run $(TASK1)

test1:
	go test -v $(TEST1)

task2:
	go run $(TASK2)

test2:
	go test -v $(TEST2)

task3:
	go run $(TASK3)

test3:
	go test -v $(TEST3)