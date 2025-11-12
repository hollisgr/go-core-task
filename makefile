TASKS := 1 2 3 4 5 6

MAIN_FILES = $(addprefix $(1)/main_,$(1).go)

TEST_FILES = $(wildcard $(1)/*.go)

define TASK_RULE
task$1:
	go run $$(call MAIN_FILES,$1)

test$1:
	go test -v $$(call TEST_FILES,$1)
endef

$(foreach t,$(TASKS),$(eval $(call TASK_RULE,$t)))