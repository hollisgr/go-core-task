TASK1 := 1/main_1.go
TASK2 := 2/main_2.go
TASK3 := 3/main_3.go
TASK4 := 4/main_4.go


TEST1 := 1/*.go
TEST2 := 2/*.go
TEST3 := 3/*.go
TEST4 := 4/*.go


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

task4:
	go run $(TASK4)

test4:
	go test -v $(TEST4)