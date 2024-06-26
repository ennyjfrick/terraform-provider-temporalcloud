package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBasicAccountMetrics(t *testing.T) {
	config := `
provider "temporalcloud" {

}

resource "temporalcloud_account_metrics" "terraform" {
	accepted_client_ca  = base64encode(<<PEM
-----BEGIN CERTIFICATE-----
MIIByTCCAVCgAwIBAgIRAWHkC+6JUf3s9Tq43mdp2zgwCgYIKoZIzj0EAwMwEzER
MA8GA1UEChMIdGVtcG9yYWwwHhcNMjMwODEwMDAwOTQ1WhcNMjQwODA5MDAxMDQ1
WjATMREwDwYDVQQKEwh0ZW1wb3JhbDB2MBAGByqGSM49AgEGBSuBBAAiA2IABCzQ
7DwwGSQKM6Zrx3Qtw7IubfxiJ3RSXCqmcGhEbFVeocwAdEgMYlwSlUiWtDZVR2dM
XM9UZLWK4aGGnDNS5Mhcz6ibSBS7Owf4tRZZA9SpFCjNw2HraaiUVV+EUgxoe6No
MGYwDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFG4N
8lIXqQKxwVs/ixVzdF6XGZm+MCQGA1UdEQQdMBuCGWNsaWVudC5yb290LnRlbXBv
cmFsLlB1VHMwCgYIKoZIzj0EAwMDZwAwZAIwRLfm9S7rKGd30KdQvUMcOcDJlmDw
6/oM6UOJFxLeGcpYbgxQ/bFize+Yx9Q9kNeMAjA7GiFsaipaKtWHy5MCOCas3ZP6
+ttLaXNXss3Z5Wk5vhDQnyE8JR3rPeQ2cHXLiA0=
-----END CERTIFICATE-----
PEM
) 
}`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
		},
	})
}
