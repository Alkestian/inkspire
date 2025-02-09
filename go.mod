module inkspire

go 1.23.4

require (
	github.com/go-chi/chi/v5 v5.2.1
	inkspire/pkg/models v0.0.0
)

replace inkspire/pkg/models => ./pkg/models
