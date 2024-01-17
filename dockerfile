# # =============== build stage ===============
# FROM golang:1.16.5-buster AS build
# WORKDIR /app

# COPY go.* ./
# RUN go mod download -x all

# COPY . ./
# # ldflags:
# #  -s: disable symbol table
# #  -w: disable DWARF generation
# # run `go tool link -help` to get the full list of ldflags
# RUN go build -ldflags "-s -w" -o the-excellent-app -v path/to/main.go

# =============== final stage ===============
FROM chromedp/headless-shell:93.0.4535.3 AS final

WORKDIR /app

COPY . ./
# COPY --from=build /app/the-excellent-app ./
# ENTRYPOINT ["/app/the-excellent-app", "-other", "flags"]

