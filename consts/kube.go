package consts

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	KubeIngressClassName = "yatai-ingress"

	KubeLabelYataiSelector        = "yatai.ai/selector"
	KubeLabelYataiBentoRepository = "yatai.ai/bento-repository"
	KubeLabelYataiBento           = "yatai.ai/bento"
	KubeLabelYataiModelRepository = "yatai.ai/model-repository"
	KubeLabelYataiModel           = "yatai.ai/model"

	KubeHPAQPSMetric = "http_request"
	KubeHPAGPUMetric = "container_accelerator_duty_cycle"

	KubeNamespaceYataiBentoImageBuilder = "yatai-builders"
	KubeNamespaceYataiModelImageBuilder = "yatai-builders"
	KubeNamespaceYataiDeployment        = "yatai"
	KubeNamespaceYataiOperators         = "yatai-operators"
	KubeNamespaceYataiComponents        = "yatai-components"

	KubeNamespaceYataiSystem              = "yatai-system"
	KubeNamespaceYataiDeploymentComponent = "yatai-deployment"

	KubeLabelYataiBentoDeployment              = "yatai.ai/bento-deployment"
	KubeLabelYataiBentoDeploymentComponentType = "yatai.ai/bento-deployment-component-type"
	KubeLabelYataiBentoDeploymentComponentName = "yatai.ai/bento-deployment-component-name"
	KubeLabelYataiBentoDeploymentTargetType    = "yatai.ai/bento-deployment-target-type"
	KubeLabelYataiBentoDeploymentRunner        = "yatai.ai/bento-deployment-runner"
	KubeLabelBentoRepository                   = "yatai.ai/bento-repository"
	KubeLabelBentoVersion                      = "yatai.ai/bento-version"
	KubeLabelCreator                           = "yatai.ai/creator"
	// nolint: gosec
	KubeLabelYataiDeployToken = "yatai.ai/deploy-token"

	KubeLabelYataiOwnerReference = "yatai.ai/owner-reference"

	KubeLabelGPUAccelerator = "gpu-accelerator"

	KubeLabelHostName = "kubernetes.io/hostname"
	KubeLabelArch     = "kubernetes.io/arch"

	KubeLabelMcdNodePool       = "mcd.io/node-pool"
	KubeLabelAlibabaEdgeWorker = "alibabacloud.com/is-edge-worker"
	KubeLabelMcdEdgeWorker     = "mcd.io/is-edge-worker"
	KubeLabelFalse             = "false"
	KubeLabelTrue              = "true"

	KubeLabelManagedBy    = "app.kubernetes.io/managed-by"
	KubeLabelHelmHeritage = "heritage"
	KubeLabelHelmRelease  = "release"

	KubeAnnotationBentoRepository   = "yatai.ai/bento-repository"
	KubeAnnotationBentoVersion      = "yatai.ai/bento-version"
	KubeAnnotationYataiDeploymentId = "yatai.ai/deployment-id"
	KubeAnnotationHelmReleaseName   = "meta.helm.sh/release-name"

	KubeAnnotationPrometheusScrape = "prometheus.io/scrape"
	KubeAnnotationPrometheusPort   = "prometheus.io/port"
	KubeAnnotationPrometheusPath   = "prometheus.io/path"

	KubeAnnotationARMSAutoEnable = "armsPilotAutoEnable"
	KubeAnnotationARMSAppName    = "armsPilotCreateAppName"

	KubeCreator = "yatai"

	KubeResourceGPUNvidia = "nvidia.com/gpu"

	KubeEventResourceKindPod        = "Pod"
	KubeEventResourceKindHPA        = "HorizontalPodAutoscaler"
	KubeEventResourceKindReplicaSet = "ReplicaSet"

	KubeTaintKeyDedicatedNodeGroup = "mcd.io/dedicated-node-group"
	KubeLabelDedicatedNodeGroup    = "mcd.io/dedicated-node-group"

	KubeImageCSIDriver          = "image.csi.k8s.io"
	KubeImageCSIDriverWarmMetal = "csi-image.warm-metal.tech"

	KubeConfigMapNameNetworkConfig = "network"

	KubeConfigMapKeyNetworkConfigDomainSuffix       = "domain-suffix"
	KubeConfigMapKeyNetworkConfigIngressClass       = "ingress-class"
	KubeConfigMapKeyNetworkConfigIngressAnnotations = "ingress-annotations"

	KubeConfigMapNameYataiConfig = "yatai"

	KubeConfigMapKeyYataiConfigEndpoint           = "endpoint"
	KubeConfigMapKeyYataiConfigClusterName        = "cluster-name"
	KubeConfigMapKeyYataiConfigApiTokenSecretName = "api-token-secret-name"
	KubeConfigMapKeyYataiConfigApiTokenSecretKey  = "api-token-secret-key"

	// nolint: gosec
	KubeSecretNameRegcred = "yatai-regcred"
)

var KubeListEverything = metav1.ListOptions{
	LabelSelector: labels.Everything().String(),
	FieldSelector: fields.Everything().String(),
}
