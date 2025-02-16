OPENAPI_SPEC = openapi/v1/openapi.yaml
GENERATED_DIR = internal/gen
GENERATED_API = $(GENERATED_DIR)/api.gen.go
GENERATED_MODELS = $(GENERATED_DIR)/models.gen.go

generate-api: $(GENERATED_API)
	@echo "Generating API code..."

$(GENERATED_API): $(OPENAPI_SPEC)
	@mkdir -p $(GENERATED_DIR)
	oapi-codegen -generate server,spec -package generated $(OPENAPI_SPEC) > $(GENERATED_API)
	@echo "API code generated successfully."

generate-models: $(GENERATED_MODELS)
	@echo "Generating Models code..."

$(GENERATED_MODELS): $(OPENAPI_SPEC)
	@mkdir -p $(GENERATED_DIR)
	oapi-codegen -generate types,models -package generated $(OPENAPI_SPEC) > $(GENERATED_MODELS)
	@echo "Models code generated successfully."

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(GENERATED_DIR)
	@echo "Generated files cleaned."

all: generate-api generate-models
	@echo "All generation tasks completed."