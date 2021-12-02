package options

type EtcdOptions struct {
	Endpoints   []string `json:"endpoints"`
	DialTimeout int      `json:"dialTimeout"`
}

func GetEtcdConfig()  {
	
}