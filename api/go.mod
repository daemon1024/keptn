module github.com/keptn/keptn/api

go 1.13

require (
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/runtime v0.19.20
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/go-test/deep v1.0.7
	github.com/google/uuid v1.1.2
	github.com/jessevdk/go-flags v1.4.0
	github.com/keptn/go-utils v0.8.0-alpha.0.20210203161317-67ac0f2ba06d
	github.com/keptn/kubernetes-utils v0.8.0-alpha
	golang.org/x/net v0.0.0-20210224082022-3d97a244fca7
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible
