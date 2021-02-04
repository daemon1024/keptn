module github.com/keptn/keptn/configuration-service

go 1.13

require (
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/frankban/quicktest v1.9.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/loads v0.19.4
	github.com/go-openapi/runtime v0.19.4
	github.com/go-openapi/spec v0.19.3
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.5
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/gophercloud/gophercloud v0.9.0 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/keptn/go-utils v0.8.0-alpha.0.20210118153845-86a445dfcb86
	github.com/keptn/kubernetes-utils v0.8.0-alpha
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/otiai10/copy v1.4.2
	github.com/pierrec/lz4 v2.3.0+incompatible // indirect
	github.com/stretchr/testify v1.6.1
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/yaml v1.2.0 // indirect
)

// Transitive requirement from Helm: See https://github.com/helm/helm/blob/v3.1.2/go.mod
replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible
