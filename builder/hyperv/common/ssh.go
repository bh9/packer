package common

import (
	"context"
	"log"

	"github.com/hashicorp/packer/helper/multistep"
)

func CommHost(host string) func(multistep.StateBag) (string, error) {
	return func(state multistep.StateBag) (string, error) {
		ctx := context.TODO()

		// Skip IP auto detection if the configuration has an ssh host configured.
		if host != "" {
			log.Printf("Using ssh_host value: %s", host)
			return host, nil
		}

		vmName := state.Get("vmName").(string)
		driver := state.Get("driver").(Driver)

		mac, err := driver.Mac(ctx, vmName)
		if err != nil {
			return "", err
		}

		ip, err := driver.IpAddress(ctx, mac)
		if err != nil {
			return "", err
		}

		return ip, nil
	}
}
