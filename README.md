# oras-explorer-mini

<img width="700" alt="image" src="https://github.com/VasuDevrani/oras-explorer-mini/assets/101383635/8476cd0f-64a6-41b3-85c1-ee4d436d9eee">

## Run locally
- `git clone https://github.com/VasuDevrani/oras-explorer-mini.git`
- `cd oras-explorer-mini`
- `npm i`
- `hugo server` to start the hugo frontend (:1313)
- `cd /go` > `go run .` to start the go backend server (:8080)
- visit `localhost:1313`
- Example oci artifacts:
  ```json
  {
    registry: docker.io,
    repository: /library/nginx,
    tag: latest
  },
  {
    registry: gcr.io,
    repository: /distroless/static,
    tag: latest
  },
  ```
