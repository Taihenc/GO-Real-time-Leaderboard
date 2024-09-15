# Define variables
GO_BUILD_CMD = go build -o ./tmp/main.exe .
TAILWIND_CMD = .\bin\tailwindcss.exe -i .\public\index.css -o .\public\output.css --minify

# Default target
all: build tailwind

# Build the Go executable
build:
	$(GO_BUILD_CMD)
	@echo "Go build completed, main.exe created in ./tmp/"

# Compile Tailwind CSS
tailwind:
	$(TAILWIND_CMD)
	@echo "Tailwind CSS compiled to ./public/output.css"

# Clean target (optional)
clean:
	rm -rf ./tmp/main.exe ./public/output.css
	@echo "Cleaned build artifacts"

# Phony targets to prevent conflicts with actual file names
.PHONY: all build tailwind clean
