package flavors

type Flavor_v2_1 struct {
	// ID is the flavor's unique ID.
	ID string `json:"id" hello:"id"`

	// Disk is the amount of root disk, measured in GB.
	Disk int `json:"disk" hello:"disk"`

	// RAM is the amount of memory, measured in MB.
	RAM int `json:"ram" hello:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name" hello:"name"`

	// RxTxFactor describes bandwidth alterations of the flavor.
	RxTxFactor float64 `json:"rxtx_factor" hello:"rx_tx_factor"`

	// Swap is the amount of swap space, measured in MB.
	Swap int `json:"-" hello:"swap"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	VCPUs int `json:"vcpus" hello:"vcp_us"`

	// IsPublic indicates whether the flavor is public.
	IsPublic bool `json:"os-flavor-access:is_public" hello:"is_public"`

	// Ephemeral is the amount of ephemeral disk space, measured in GB.
	Ephemeral int `json:"OS-FLV-EXT-DATA:ephemeral" hello:"ephemeral"`

	// Description is a free form description of the flavor. Limited to
	// 65535 characters in length. Only printable characters are allowed.
	// New in version 2.55
	Description string `json:"description" hello:"description"`
}
