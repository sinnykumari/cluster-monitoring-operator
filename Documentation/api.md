
**NOTE**: The contents of this file are **automatically generated** from source code comments. 
If you wish to make a change or an addition to the content in this document, do so by **changing the code comments**.

# Cluster Monitoring Configuration Reference

Parts of Cluster Monitoring are configurable. Depending on which part of the stack users want to configure, they should edit the following:

- Configuration of OpenShift Container Platform monitoring components lies in a ConfigMap called `cluster-monitoring-config` in the `openshift-monitoring` namespace. Defined by [ClusterMonitoringConfiguration](#clustermonitoringconfiguration).
- Configuration of components that monitor user-defined projects lies in a ConfigMap called `user-workload-monitoring-config` in the `openshift-user-workload-monitoring` namespace. Defined by [UserWorkloadConfiguration](#userworkloadconfiguration).

The configuration file itself is always defined under the `config.yaml` key within the ConfigMap's data.

Monitoring a platform such as OpenShift requires a coordination of multiple components that must work well between themselves.
However, users should be able to customize the monitoring stack in such a way that they end up with a resilient and highly available monitoring solution.
Despite this, to avoid users from misconfiguring the monitoring stack of their clusters not all configuration parameters are exposed.

Configuring Cluster Monitoring is optional. If the config does not exist or is empty or malformed, then defaults will be used.

## Table of Contents
* [AdditionalAlertmanagerConfig](#additionalalertmanagerconfig)
* [AlertmanagerMainConfig](#alertmanagermainconfig)
* [AlertmanagerUserWorkloadConfig](#alertmanageruserworkloadconfig)
* [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)
* [DedicatedServiceMonitors](#dedicatedservicemonitors)
* [K8sPrometheusAdapter](#k8sprometheusadapter)
* [KubeStateMetricsConfig](#kubestatemetricsconfig)
* [OpenShiftStateMetricsConfig](#openshiftstatemetricsconfig)
* [PrometheusK8sConfig](#prometheusk8sconfig)
* [PrometheusOperatorConfig](#prometheusoperatorconfig)
* [PrometheusRestrictedConfig](#prometheusrestrictedconfig)
* [RemoteWriteSpec](#remotewritespec)
* [TLSConfig](#tlsconfig)
* [ThanosQuerierConfig](#thanosquerierconfig)
* [ThanosRulerConfig](#thanosrulerconfig)
* [UserWorkloadConfiguration](#userworkloadconfiguration)

## AdditionalAlertmanagerConfig

#### Description

AdditionalAlertmanagerConfig defines configuration on how a component should communicate with aditional Alertmanager instances.

#### Required
   - ` apiVersion `

<em>appears in: [PrometheusK8sConfig](#prometheusk8sconfig), [PrometheusRestrictedConfig](#prometheusrestrictedconfig), [ThanosRulerConfig](#thanosrulerconfig)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| apiVersion | string | APIVersion defines the api version of Alertmanager. |
| bearerToken | *[v1.SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core) | BearerToken defines the bearer token to use when authenticating to Alertmanager. |
| pathPrefix | string | PathPrefix defines the path prefix to add in front of the push endpoint path. |
| scheme | string | Scheme the URL scheme to use when talking to Alertmanagers. |
| staticConfigs | []string | StaticConfigs a list of statically configured Alertmanagers. |
| timeout | *string | Timeout defines the timeout used when sending alerts. |
| tlsConfig | [TLSConfig](#tlsconfig) | TLSConfig defines the TLS Config to use for alertmanager connection. |

[Back to TOC](#table-of-contents)

## AlertmanagerMainConfig

#### Description

AlertmanagerMainConfig defines configuration related with the main Alertmanager instance.


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| enabled | *bool | Enabled a boolean flag to enable or disable the main Alertmanager instance under openshift-monitoring default: true |
| enableUserAlertmanagerConfig | bool | EnableUserAlertManagerConfig boolean flag to enable or disable user-defined namespaces to be selected for AlertmanagerConfig lookup, by default Alertmanager only looks for configuration in the namespace where it was deployed to. This will only work if the UWM Alertmanager instance is not enabled. default: false |
| logLevel | string | LogLevel defines the log level for Alertmanager. Possible values are: error, warn, info, debug. default: info |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| resources | *[v1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#resourcerequirements-v1-core) | Resources define resources requests and limits for single Pods. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |
| topologySpreadConstraints | []v1.TopologySpreadConstraint | TopologySpreadConstraints defines the pod's topology spread constraints. |
| volumeClaimTemplate | *[monv1.EmbeddedPersistentVolumeClaim](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#embeddedpersistentvolumeclaim) | VolumeClaimTemplate defines persistent storage for Alertmanager. It's possible to configure storageClass and size of volume. |

[Back to TOC](#table-of-contents)

## AlertmanagerUserWorkloadConfig

#### Description

AlertmanagerUserWorkloadConfig defines configuration for the Alertmanager instance for user-defined projects.


<em>appears in: [UserWorkloadConfiguration](#userworkloadconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| enabled | bool | Enabled a boolean flag to enable or disable a dedicated instance of Alertmanager for user-defined projects under openshift-user-workload-monitoring default: false |
| enableAlertmanagerConfig | bool | EnableAlertmanagerConfig a boolean flag to enable or disable user-defined namespaces to be selected for AlertmanagerConfig lookup, by default Alertmanager only looks for configuration in the namespace where it was deployed to default: false |
| logLevel | string | LogLevel defines the log level for Alertmanager. Possible values are: error, warn, info, debug. default: info |
| resources | *[v1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#resourcerequirements-v1-core) | Resources define resources requests and limits for single Pods. |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |
| volumeClaimTemplate | *[monv1.EmbeddedPersistentVolumeClaim](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#embeddedpersistentvolumeclaim) | VolumeClaimTemplate defines persistent storage for Alertmanager. It's possible to configure storageClass and size of volume. |

[Back to TOC](#table-of-contents)

## ClusterMonitoringConfiguration

#### Description

ClusterMonitoringConfiguration defines configuration that allows users to customise the platform monitoring stack through the cluster-monitoring-config ConfigMap in the openshift-monitoring namespace

| Property | Type | Description |
| -------- | ---- | ----------- |
| alertmanagerMain | *[AlertmanagerMainConfig](#alertmanagermainconfig) | AlertmanagerMainConfig defines configuration related with the main Alertmanager instance. |
| enableUserWorkload | *bool | UserWorkloadEnabled boolean flag to enable monitoring for user-defined projects. |
| k8sPrometheusAdapter | *[K8sPrometheusAdapter](#k8sprometheusadapter) | K8sPrometheusAdapter defines configuration related with prometheus-adapter |
| kubeStateMetrics | *[KubeStateMetricsConfig](#kubestatemetricsconfig) | KubeStateMetricsConfig defines configuration related with kube-state-metrics agent |
| prometheusK8s | *[PrometheusK8sConfig](#prometheusk8sconfig) | PrometheusK8sConfig defines configuration related with prometheus |
| prometheusOperator | *[PrometheusOperatorConfig](#prometheusoperatorconfig) | PrometheusOperatorConfig defines configuration related with prometheus-operator |
| openshiftStateMetrics | *[OpenShiftStateMetricsConfig](#openshiftstatemetricsconfig) | OpenShiftMetricsConfig defines configuration related with openshift-state-metrics agent |
| thanosQuerier | *[ThanosQuerierConfig](#thanosquerierconfig) | ThanosQuerierConfig defines configuration related with the Thanos Querier component |

[Back to TOC](#table-of-contents)

## DedicatedServiceMonitors

#### Description

Configuration for prometheus-adapter dedicated Service Monitors. When Enabled is set to true, CMO will deploy and scrape a dedicated Service Monitor, that exposes the kubelet /metrics/resource endpoint. This Service Monitor sets honorTimestamps: true and only keeps metrics that are relevant for the pod resource queries of prometheus-adapter. Additionally prometheus-adapter is configured to use these dedicated metrics. Overall this will improve the consistency of prometheus-adapter based CPU usage measurements used by for example the oc adm top pod command or the Horizontal Pod Autoscaler.


<em>appears in: [K8sPrometheusAdapter](#k8sprometheusadapter)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| enabled | bool |  |

[Back to TOC](#table-of-contents)

## K8sPrometheusAdapter

#### Description

K8sPrometheusAdapter defines configuration related with Prometheus Adapater


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| audit | *Audit | Audit defines the audit configuration to be used by the prometheus adapter instance. Possible profile values are: \"metadata, request, requestresponse, none\". default: metadata |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |
| dedicatedServiceMonitors | *[DedicatedServiceMonitors](#dedicatedservicemonitors) |  |

[Back to TOC](#table-of-contents)

## KubeStateMetricsConfig

#### Description

KubeStateMetricsConfig defines configuration related with the kube-state-metrics agent.


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |

[Back to TOC](#table-of-contents)

## OpenShiftStateMetricsConfig

#### Description

OpenShiftStateMetricsConfig holds configuration related to openshift-state-metrics agent.


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |

[Back to TOC](#table-of-contents)

## PrometheusK8sConfig

#### Description

PrometheusK8sConfig holds configuration related to the Prometheus component.


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| additionalAlertmanagerConfigs | [][AdditionalAlertmanagerConfig](#additionalalertmanagerconfig) | AlertmanagerConfigs holds configuration about how the Prometheus component should communicate with aditional Alertmanager instances. default: nil |
| enforcedBodySizeLimit | string | EnforcedBodySizeLimit enforces body size limit of Prometheus scrapes, if a scrape is bigger than the limit it will fail. 3 kinds of values are accepted:\n 1. empty value: no limit\n 2. a value in Prometheus size format, e.g. \"64MB\"\n 3. string \"automatic\", which means the limit will be automatically calculated based on\n    cluster capacity.\ndefault: 64MB |
| externalLabels | map[string]string | ExternalLabels defines labels to be added to any time series or alerts when communicating with external systems (federation, remote storage, Alertmanager). default: nil |
| logLevel | string | LogLevel defines the log level for Prometheus. Possible values are: error, warn, info, debug. default: info |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| queryLogFile | string | QueryLogFile specifies the file to which PromQL queries are logged. Suports both just a filename in which case they will be saved to an emptyDir volume at /var/log/prometheus, if a full path is given an emptyDir volume will be mounted at that location. Relative paths not supported, also not supported writing to linux std streams. default: \"\" |
| remoteWrite | [][RemoteWriteSpec](#remotewritespec) | RemoteWrite Holds the remote write configuration, everything from url, authorization to relabeling |
| resources | *[v1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#resourcerequirements-v1-core) | Resources define resources requests and limits for single Pods. |
| retention | string | Retention defines the Time duration Prometheus shall retain data for. Must match the regular expression [0-9]+(ms\|s\|m\|h\|d\|w\|y) (milliseconds seconds minutes hours days weeks years). default: 15d |
| retentionSize | string | RetentionSize defines the maximum amount of disk space used by blocks + WAL. default: nil |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |
| topologySpreadConstraints | []v1.TopologySpreadConstraint | TopologySpreadConstraints defines the pod's topology spread constraints. |
| volumeClaimTemplate | *[monv1.EmbeddedPersistentVolumeClaim](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#embeddedpersistentvolumeclaim) | VolumeClaimTemplate defines persistent storage for Prometheus. It's possible to configure storageClass and size of volume. |

[Back to TOC](#table-of-contents)

## PrometheusOperatorConfig

#### Description

PrometheusOperatorConfig holds configuration related to Prometheus Operator.


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration), [UserWorkloadConfiguration](#userworkloadconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| logLevel | string | LogLevel defines the log level for Prometheus Operator. Possible values are: error, warn, info, debug. default: info |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |

[Back to TOC](#table-of-contents)

## PrometheusRestrictedConfig

#### Description

PrometheusRestrictedConfig defines configuration related to the Prometheus component that will monitor user-defined projects.


<em>appears in: [UserWorkloadConfiguration](#userworkloadconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| additionalAlertmanagerConfigs | [][AdditionalAlertmanagerConfig](#additionalalertmanagerconfig) | AlertmanagerConfigs holds configuration about how the Prometheus component should communicate with aditional Alertmanager instances. default: nil |
| enforcedLabelLimit | *uint64 | EnforcedLabelLimit per-scrape limit on the number of labels accepted for a sample. If more than this number of labels are present post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. default: 0 |
| enforcedLabelNameLengthLimit | *uint64 | EnforcedLabelNameLengthLimit per-scrape limit on the length of labels name that will be accepted for a sample. If a label name is longer than this number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. default: 0 |
| enforcedLabelValueLengthLimit | *uint64 | EnforcedLabelValueLengthLimit per-scrape limit on the length of labels value that will be accepted for a sample. If a label value is longer than this number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. default: 0 |
| enforcedSampleLimit | *uint64 | EnforcedSampleLimit defines a global limit on the number of scraped samples that will be accepted. This overrides any SampleLimit set per ServiceMonitor or/and PodMonitor. It is meant to be used by admins to enforce the SampleLimit to keep the overall number of samples/series under the desired limit. Note that if SampleLimit is lower that value will be taken instead. default: 0 |
| enforcedTargetLimit | *uint64 | EnforcedTargetLimit defines a global limit on the number of scraped targets. This overrides any TargetLimit set per ServiceMonitor or/and PodMonitor. It is meant to be used by admins to enforce the TargetLimit to keep the overall number of targets under the desired limit. Note that if TargetLimit is lower, that value will be taken instead, except if either value is zero, in which case the non-zero value will be used. If both values are zero, no limit is enforced. default: 0 |
| externalLabels | map[string]string | ExternalLabels defines labels to be added to any time series or alerts when communicating with external systems (federation, remote storage, Alertmanager). default: nil |
| logLevel | string | LogLevel defines the log level for Prometheus. Possible values are: error, warn, info, debug. default: info |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| queryLogFile | string | QueryLogFile specifies the file to which PromQL queries are logged. Suports both just a filename in which case they will be saved to an emptyDir volume at /var/log/prometheus, if a full path is given an emptyDir volume will be mounted at that location. Relative paths not supported, also not supported writing to linux std streams. default: \"\" |
| remoteWrite | [][RemoteWriteSpec](#remotewritespec) | RemoteWrite Holds the remote write configuration, everything from url, authorization to relabeling |
| resources | *[v1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#resourcerequirements-v1-core) | Resources define resources requests and limits for single Pods. |
| retention | string | Retention defines the Time duration Prometheus shall retain data for. Must match the regular expression [0-9]+(ms\|s\|m\|h\|d\|w\|y) (milliseconds seconds minutes hours days weeks years). default: 15d |
| retentionSize | string | RetentionSize defines the maximum amount of disk space used by blocks + WAL. default: nil |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |
| volumeClaimTemplate | *[monv1.EmbeddedPersistentVolumeClaim](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#embeddedpersistentvolumeclaim) | VolumeClaimTemplate defines persistent storage for Prometheus. It's possible to configure storageClass and size of volume. |

[Back to TOC](#table-of-contents)

## RemoteWriteSpec

#### Description

RemoteWriteSpec is almost a 1to1 copy of monv1.RemoteWriteSpec but with the BearerToken field removed. In the future other fields might be added here.

#### Required
   - ` url `

<em>appears in: [PrometheusK8sConfig](#prometheusk8sconfig), [PrometheusRestrictedConfig](#prometheusrestrictedconfig)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| authorization | *monv1.SafeAuthorization | Authorization defines the authorization section for remote write |
| basicAuth | *[monv1.BasicAuth](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#basicauth) | BasicAuth defines configuration for basic authentication for the URL. |
| bearerTokenFile | string | BearerTokenFile defines the file where the bearer token for remote write resides. |
| headers | map[string]string | Headers custom HTTP headers to be sent along with each remote write request. Be aware that headers that are set by Prometheus itself can't be overwritten. |
| metadataConfig | *[monv1.MetadataConfig](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#metadataconfig) | MetadataConfig configures the sending of series metadata to remote storage. |
| name | string | Name defines the name of the remote write queue, must be unique if specified. The name is used in metrics and logging in order to differentiate queues. |
| oauth2 | *monv1.OAuth2 | OAuth2 configures OAuth2 authentication for remote write. |
| proxyUrl | string | ProxyURL defines an optional proxy URL |
| queueConfig | *[monv1.QueueConfig](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#queueconfig) | QueueConfig allows tuning of the remote write queue parameters. |
| remoteTimeout | string | RemoteTimeout defines the timeout for requests to the remote write endpoint. |
| sigv4 | *monv1.Sigv4 | Sigv4 allows to configures AWS's Signature Verification 4 |
| tlsConfig | *[monv1.SafeTLSConfig](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#safetlsconfig) | TLSConfig defines the TLS configuration to use for remote write. |
| url | string | URL defines the URL of the endpoint to send samples to. |
| writeRelabelConfigs | [][monv1.RelabelConfig](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#relabelconfig) | WriteRelabelConfigs defines the list of remote write relabel configurations. |

[Back to TOC](#table-of-contents)

## TLSConfig

#### Description

TLSConfig configures the options for TLS connections.

#### Required
   - ` insecureSkipVerify `

<em>appears in: [AdditionalAlertmanagerConfig](#additionalalertmanagerconfig)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| ca | *[v1.SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core) | CA defines the CA cert in the Prometheus container to use for the targets. |
| cert | *[v1.SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core) | Cert defines the client cert in the Prometheus container to use for the targets. |
| key | *[v1.SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core) | Key defines the client key in the Prometheus container to use for the targets. |
| serverName | string | ServerName used to verify the hostname for the targets. |
| insecureSkipVerify | bool | InsecureSkipVerify disable target certificate validation. |

[Back to TOC](#table-of-contents)

## ThanosQuerierConfig

#### Description

ThanosQuerierConfig holds configuration related to Thanos Querier component.


<em>appears in: [ClusterMonitoringConfiguration](#clustermonitoringconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| enableRequestLogging | bool | EnableRequestLogging boolean flag to enable or disable request logging default: false |
| logLevel | string | LogLevel defines the log level for Thanos Querier. Possible values are: error, warn, info, debug. default: info |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| resources | *[v1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#resourcerequirements-v1-core) | Resources define resources requests and limits for single Pods. |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |

[Back to TOC](#table-of-contents)

## ThanosRulerConfig

#### Description

ThanosRulerConfig defines configuration for the Thanos Ruler instance for user-defined projects.


<em>appears in: [UserWorkloadConfiguration](#userworkloadconfiguration)</em>

| Property | Type | Description |
| -------- | ---- | ----------- |
| additionalAlertmanagerConfigs | [][AdditionalAlertmanagerConfig](#additionalalertmanagerconfig) | AlertmanagerConfigs holds configuration about how the Thanos Ruler component should communicate with aditional Alertmanager instances. default: nil |
| logLevel | string | LogLevel defines the log level for Thanos Ruler. Possible values are: error, warn, info, debug. default: info |
| nodeSelector | map[string]string | NodeSelector defines which Nodes the Pods are scheduled on. |
| resources | *[v1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#resourcerequirements-v1-core) | Resources define resources requests and limits for single Pods. |
| retention | string | Retention defines the time duration Thanos Ruler shall retain data for. Must match the regular expression [0-9]+(ms\|s\|m\|h\|d\|w\|y) (milliseconds seconds minutes hours days weeks years). default: 15d |
| tolerations | [][v1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#toleration-v1-core) | Tolerations defines the Pods tolerations. |
| topologySpreadConstraints | []v1.TopologySpreadConstraint | TopologySpreadConstraints defines the pod's topology spread constraints. |
| volumeClaimTemplate | *[monv1.EmbeddedPersistentVolumeClaim](https://github.com/prometheus-operator/prometheus-operator/blob/v0.57.0/Documentation/api.md#embeddedpersistentvolumeclaim) | VolumeClaimTemplate defines persistent storage for Thanos Ruler. It's possible to configure storageClass and size of volume. |

[Back to TOC](#table-of-contents)

## UserWorkloadConfiguration

#### Description

UserWorkloadConfiguration defines configuration that allows users to customise the monitoring stack responsible for user-defined projects through the user-workload-monitoring-config ConfigMap in the openshift-user-workload-monitoring namespace

| Property | Type | Description |
| -------- | ---- | ----------- |
| alertmanager | *[AlertmanagerUserWorkloadConfig](#alertmanageruserworkloadconfig) | Alertmanager defines configuration for Alertmanager component. |
| prometheus | *[PrometheusRestrictedConfig](#prometheusrestrictedconfig) | Prometheus defines configuration for Prometheus component. |
| prometheusOperator | *[PrometheusOperatorConfig](#prometheusoperatorconfig) | PrometheusOperator defines configuration for prometheus-operator component. |
| thanosRuler | *[ThanosRulerConfig](#thanosrulerconfig) | ThanosRuler defines configuration for the Thanos Ruler component |

[Back to TOC](#table-of-contents)
