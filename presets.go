package realip

var presets = map[string][]string{
	// from https://www.cloudflare.com/ips/
	"cloudflare": {
		"103.21.244.0/22",
		"103.22.200.0/22",
		"103.31.4.0/22",
		"104.16.0.0/12",
		"108.162.192.0/18",
		"131.0.72.0/22",
		"141.101.64.0/18",
		"162.158.0.0/15",
		"172.64.0.0/13",
		"173.245.48.0/20",
		"188.114.96.0/20",
		"190.93.240.0/20",
		"197.234.240.0/22",
		"198.41.128.0/17",
		"199.27.128.0/21",
		"2400:cb00::/32",
		"2405:8100::/32",
		"2405:b500::/32",
		"2606:4700::/32",
		"2803:f800::/32",
		"2a06:98c0::/29",
		"2c0f:f248::/32",
	},
	// https://cloud.google.com/compute/docs/load-balancing/http/#firewall_rules
	"gcp": {
		"130.211.0.0/22",
		"35.191.0.0/16",
	},
	// https://support.rackspace.com/how-to/using-cloud-load-balancers-with-rackconnect/
	"rackspace": {
		// DFW region
		"10.189.254.0/24",
		"10.189.252.0/24",
		"10.183.248.0/24",
		"10.187.186.0/24",
		"10.183.250.0/24",
		// IAD region
		"10.187.191.0/24",
		"10.189.255.0/24",
		"10.187.186.0/24",
		"10.189.254.0/24",
		// ORD region
		"10.183.253.0/24",
		"10.183.250.0/24",
		"10.189.246.0/24",
		"10.187.187.0/24",
		"10.187.186.0/24",
		"10.183.252.0/24",
		"10.189.245.0/24",
		"10.183.251.0/24",
		// LON region
		"10.187.191.0/24",
		"10.190.254.0/24",
		"10.189.246.0/24",
		"10.190.255.0/24",
		"10.187.190.0/24",
		"10.189.247.0/24",
		// SYD region
		"10.189.254.0/24",
		// HKG region
		"10.189.254.0/24",
	},
}
