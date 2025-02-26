module github.com/nicolasverle/slugifyer

go 1.22.0

toolchain go1.22.1

replace (
	github.com/argoproj/gitops-engine => github.com/argoproj/gitops-engine v0.7.1-0.20240122213038-792124280fcc
	// https://github.com/golang/go/issues/33546#issuecomment-519656923
	github.com/go-check/check => github.com/go-check/check v0.0.0-20180628173108-788fd7840127

	github.com/golang/protobuf => github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway => github.com/grpc-ecosystem/grpc-gateway v1.16.0

	// Avoid CVE-2022-3064
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0

	// Avoid CVE-2022-28948
	gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1

	k8s.io/api => k8s.io/api v0.26.11
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.26.11
	k8s.io/apimachinery => k8s.io/apimachinery v0.26.11
	k8s.io/apiserver => k8s.io/apiserver v0.26.11
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.26.11
	k8s.io/client-go => k8s.io/client-go v0.26.11
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.26.11
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.26.11
	k8s.io/code-generator => k8s.io/code-generator v0.26.11
	k8s.io/component-base => k8s.io/component-base v0.26.11
	k8s.io/component-helpers => k8s.io/component-helpers v0.26.11
	k8s.io/controller-manager => k8s.io/controller-manager v0.26.11
	k8s.io/cri-api => k8s.io/cri-api v0.26.11
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.26.11
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.26.11
	k8s.io/kms => k8s.io/kms v0.26.11
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.26.11
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.26.11
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.26.11
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.26.11
	k8s.io/kubectl => k8s.io/kubectl v0.26.11
	k8s.io/kubelet => k8s.io/kubelet v0.26.11
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.26.11
	k8s.io/metrics => k8s.io/metrics v0.26.11
	k8s.io/mount-utils => k8s.io/mount-utils v0.26.11
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.26.11
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.26.11
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.26.11
	k8s.io/sample-controller => k8s.io/sample-controller v0.26.11

)

require (
	github.com/onsi/ginkgo/v2 v2.22.2
	github.com/rs/zerolog v1.33.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20241210010833-40e02aabc2ad // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/swaggo/swag v1.8.12 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/tools v0.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/dn365/gin-zerolog v0.0.0-20171227063204-b43714b00db1
	github.com/gin-gonic/gin v1.10.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lib/pq v1.10.9
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/gomega v1.36.2
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.19.0
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
