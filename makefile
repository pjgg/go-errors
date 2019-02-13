REPORT_PATH ?= report
VERSION = 0.0.1

all: sync test

$(REPORT_PATH):
	mkdir -p $(REPORT_PATH)

test: | $(REPORT_PATH)
	go test -v -tags=unit -count=1 -coverprofile $(REPORT_PATH)/cover.out ./...
	go tool cover -html=$(REPORT_PATH)/cover.out -o $(REPORT_PATH)/cover.html

clean:
	rm -rf $(REPORT_PATH)

version:
	@echo $(VERSION)-v$$(basename $(JOB_NAME))$(BUILD_NUMBER)	

sync:
	go mod tidy