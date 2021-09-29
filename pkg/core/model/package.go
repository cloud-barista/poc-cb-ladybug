package model

type Package struct {
	Model
	Namespace string `json:"namespace"`
	Version   string `json:"version"`
}

type PackageList struct {
	ListModel
	Namespace string    `json:"namespace"`
	Items     []Package `json:"items"`
}

func NewPackage(namespace, name, version string) *Package {
	return &Package{
		Model:     Model{Kind: KIND_PACKAGE, Name: name},
		Namespace: namespace,
		Version:   version,
	}
}

func NewPackageList(namespace string) *PackageList {
	return &PackageList{
		ListModel: ListModel{Kind: KIND_PACKAGE_LIST},
		Namespace: namespace,
		Items:     []Package{},
	}
}
