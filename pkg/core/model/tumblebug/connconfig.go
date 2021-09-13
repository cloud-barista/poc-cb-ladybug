package tumblebug

import (
	"fmt"
	"net/http"
)

type (
	TbConnConfig struct {
		ConfigName     string `json:"configName"`
		CredentialName string `json:"credentialName"`
		DriverName     string `json:"driverName"`
		ProviderName   string `json:"providerName"`
		RegionName     string `json:"regionName"`
	}

	TbConnConfigList struct {
		ConnectionConfig []TbConnConfig `json:"connectionconfig"`
	}
)

func (self *Tumblebug) ListConnConfig() (*TbConnConfigList, error) {
	var resp TbConnConfigList

	_, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("/connConfig"),
		nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
