package digioceanloadbalancer

// DigioceanLoadBalancer represents the configuration to be applied to an existing or a new load balancer.
type DigioceanLoadBalancer struct {
	Name                string           `json:"name,omitempty"`
	Algorithm           string           `json:"algorithm,omitempty"`
	Region              string           `json:"region,omitempty"`
	ForwardingRules     []ForwardingRule `json:"forwarding_rules,omitempty"`
	HealthCheck         HealthCheck      `json:"health_check,omitempty"`
	StickySessions      StickySessions   `json:"sticky_sessions,omitempty"`
	DropletIDs          []int            `json:"droplet_ids,omitempty"`
	Tag                 string           `json:"tag,omitempty"`
	RedirectHTTPToHTTPS bool             `json:"redirect_http_to_https,omitempty"`
}

// ForwardingRule represents load balancer forwarding rules.
type ForwardingRule struct {
	EntryProtocol  string `json:"entry_protocol,omitempty"`
	EntryPort      int    `json:"entry_port,omitempty"`
	TargetProtocol string `json:"target_protocol,omitempty"`
	TargetPort     int    `json:"target_port,omitempty"`
	CertificateID  string `json:"certificate_id,omitempty"`
	TLSPassthrough bool   `json:"tls_passthrough,omitempty"`
}

// HealthCheck represents optional load balancer health check rules.
type HealthCheck struct {
	Protocol               string `json:"protocol,omitempty"`
	Port                   int    `json:"port,omitempty"`
	Path                   string `json:"path,omitempty"`
	CheckIntervalSeconds   int    `json:"check_interval_seconds,omitempty"`
	ResponseTimeoutSeconds int    `json:"response_timeout_seconds,omitempty"`
	HealthyThreshold       int    `json:"healthy_threshold,omitempty"`
	UnhealthyThreshold     int    `json:"unhealthy_threshold,omitempty"`
}

// StickySessions represents optional load balancer session affinity rules.
type StickySessions struct {
	Type             string `json:"typeStickySessions,omitempty"`
	CookieName       string `json:"cookie_name,omitempty"`
	CookieTTLSeconds int    `json:"cookie_ttl_seconds,omitempty"`
}
