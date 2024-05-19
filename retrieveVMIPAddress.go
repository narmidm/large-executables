package main

import (
	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"google.golang.org/api/option"
)

func retrieveVMIPAddress(instanceName string) (string, error) {
	ctx := context.Background()
	client, err := compute.NewInstancesRESTClient(ctx, option.WithCredentialsFile("path/to/your/service-account-file.json"))
	if err != nil {
		return "", fmt.Errorf("failed to create compute client: %v", err)
	}
	defer client.Close()

	projectID := "your-project-id"
	zone := "us-central1" // Make sure this matches the zone you used

	instance, err := client.Get(ctx, &computepb.GetInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get instance details: %v", err)
	}

	// Fetch the external IP address
	for _, iface := range instance.NetworkInterfaces {
		for _, config := range iface.AccessConfigs {
			if *config.Type == "ONE_TO_ONE_NAT" && *config.NatIP != "" {
				return *config.NatIP, nil
			}
		}
	}

	return "", fmt.Errorf("no external IP found for instance %s", instanceName)
}
