default:
	templ generate
	go build ./
	strip --strip-all citadel

run:
	templ generate
	go run ./
