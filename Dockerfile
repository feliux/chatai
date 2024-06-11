FROM golang:1.22-alpine as builder
WORKDIR /app
RUN apk add --no-cache nodejs npm git
COPY . ./
RUN go install github.com/a-h/templ/cmd/templ@latest && \
	go get ./... && \
	go mod vendor && \
	go mod tidy && \
	go mod download && \
	npm install -D tailwindcss && \
	npm install -D daisyui@latest
RUN npx tailwindcss -i view/css/app.css -o public/styles.css && \
    templ generate view && \
	go build -ldflags "-s -w" -tags pro -o /chatai main.go static_pro.go

FROM scratch
COPY --from=builder /chatai /chatai
ENTRYPOINT [ "/chatai" ]
