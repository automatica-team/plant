package tmpl

import "text/template"

var Build = func() *template.Template {
	t := template.New("Build").Funcs(funcs)
	t = template.Must(t.New("Dockerfile").Parse(buildDockerfile))
	return t
}()

const buildDockerfile = "#" + genHeader + `
FROM golang:alpine as builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/bot .

FROM alpine

WORKDIR /app

RUN apk add --no-cache tzdata

COPY locales locales

COPY bot.yml .

COPY plant.yml .

COPY --from=builder /go/bin/bot .

ENTRYPOINT ["/app/bot"]
`
