API Documentation

OpenAPI specification is available at docs/openapi.yaml.

Quick ways to view it locally:

- Using ReDoc (recommended):

1. Install redoc-cli via npm: npm i -g redoc-cli
2. Serve the spec: redoc-cli serve docs/openapi.yaml

- Using Swagger UI (Docker):

1. Run: docker run --rm -p 8081:8080 -e SWAGGER_JSON=/tmp/openapi.yaml -v $(pwd)/docs:/tmp swaggerapi/swagger-ui
2. Open http://localhost:8081 in browser

Or paste docs/openapi.yaml into https://editor.swagger.io/ to view/edit.
