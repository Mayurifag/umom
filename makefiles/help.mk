PHONY: help ## Show this help.
help:
	@grep -he '^PHONY:.*##' $(MAKEFILE_LIST) | sed -e 's/ *##/:\t/' | sed -e 's/^PHONY: *//'
