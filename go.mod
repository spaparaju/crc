module github.com/code-ready/crc

go 1.13

require (
	github.com/Masterminds/semver v1.5.0
	github.com/YourFin/binappend v0.0.0-20181105185800-0add4bf0b9ad
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/cavaliercoder/grab v2.0.0+incompatible
	github.com/cheggaaa/pb/v3 v3.0.5
	github.com/code-ready/clicumber v0.0.0-20200728062640-1203dda97f67
	github.com/code-ready/gvisor-tap-vsock v0.0.0-20201105131011-9258bacc7a6c
	github.com/code-ready/machine v0.0.0-20201109095558-3e6386d4f69e
	github.com/cucumber/godog v0.9.0
	github.com/cucumber/messages-go/v10 v10.0.3
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/go-units v0.4.0
	github.com/fatih/color v1.10.0 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/h2non/filetype v1.1.0
	github.com/hectane/go-acl v0.0.0-20190604041725-da78bae5fc95
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/libvirt/libvirt-go-xml v6.10.0+incompatible
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mattn/go-colorable v0.1.8
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/openshift/api v0.0.0-20200930075302-db52bc4ef99f
	github.com/openshift/oc v0.0.0-alpha.0.0.20201126035554-299b6af535d1
	github.com/pbnjay/memory v0.0.0-20201129165224-b12e5d931931
	github.com/pborman/uuid v1.2.1
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/browser v0.0.0-20201207095918-0426ae3fba23
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/oauth2 v0.0.0-20201203001011-0b49973bad19 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1
	golang.org/x/text v0.3.4
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.8
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	howett.net/plist v0.0.0-20201203080718-1454fab16a06
	k8s.io/api v0.19.0
	k8s.io/client-go v0.19.0
)

replace (
	github.com/Microsoft/hcsshim => github.com/Microsoft/hcsshim v0.8.10
	github.com/apcera/gssapi => github.com/openshift/gssapi v0.0.0-20161010215902-5fb4217df13b
	github.com/containers/image => github.com/openshift/containers-image v0.0.0-20190130162819-76de87591e9d
	// Taking changes from https://github.com/moby/moby/pull/40021 to accomodate new version of golang.org/x/sys.
	// Although the PR lists c3a0a3744636069f43197eb18245aaae89f568e5 as the commit with the fixes,
	// d1d5f6476656c6aad457e2a91d3436e66b6f2251 is more suitable since it does not break fsouza/go-clientdocker,
	// yet provides the same fix.
	github.com/docker/docker => github.com/docker/docker v1.4.2-0.20191121165722-d1d5f6476656

	k8s.io/api => k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/apimachinery => github.com/openshift/kubernetes-apimachinery v0.0.0-20200831185207-c0eb43ac4a3e
	k8s.io/apiserver => k8s.io/apiserver v0.19.0
	k8s.io/cli-runtime => github.com/openshift/kubernetes-cli-runtime v0.0.0-20200831185531-852eec47b608
	k8s.io/client-go => github.com/openshift/kubernetes-client-go v0.0.0-20200908071752-9409de4c95e0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.0
	k8s.io/code-generator => k8s.io/code-generator v0.19.0
	k8s.io/component-base => k8s.io/component-base v0.19.0
	k8s.io/cri-api => k8s.io/cri-api v0.19.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.0
	k8s.io/kubectl => github.com/openshift/kubernetes-kubectl v0.0.0-20200922135455-1f5b2cd472a9
	k8s.io/kubelet => k8s.io/kubelet v0.19.0
	k8s.io/kubernetes => github.com/openshift/kubernetes v1.20.0-alpha.0.0.20200922142336-4700daee7399
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.0
	k8s.io/metrics => k8s.io/metrics v0.19.0
	k8s.io/node-api => k8s.io/node-api v0.19.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.0
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.19.0
	k8s.io/sample-controller => k8s.io/sample-controller v0.19.0
)
