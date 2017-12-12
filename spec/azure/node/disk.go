package node

type Disk struct {
	// Name is the disk's name
	Name string `json:"name" yaml:"name"`
	// Caching is the disk's caching mode; None, ReadOnly, ReadWrite
	Caching string `json:"caching" yaml:"caching"`
	// SizeGB id the disk's size in Gb
	SizeGB int `json:"sizeGB" yaml:"sizeGB"`
	// StorageAccountType is the disk's storage account type; Standard_LRS or Premium_LRS
	StorageAccountType string `json:"storageAccountType" yaml:"storageAccountType"`
}
