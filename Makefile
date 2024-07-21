.PHONY: build
build:
	@templ generate 
	@go build -o bin/api 
	
	
	
	

run: build
	@./bin/api
test:
	@go test -v ./...



templWatch:
	@templ generate --watch --proxy="http://localhost:8000" --open-browser=false

templ:
	@templ generate





tailwind:
	@npx tailwindcss -i ./static/css/input.css -o ./static/css/styles.css --watch





tailwind-build:
	@npx tailwindcss -i ./views/css/input.css -o ./views/css/style.min.css --minify