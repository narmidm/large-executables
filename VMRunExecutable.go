package main

import (
	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
)

// createInstance sends an instance creation request to the Compute Engine API and waits for it to complete.
func spinUpVMAndRunExecutable(region, executableID string) (string, error) {
	projectID := "your_project_id"
	zone := fmt.Sprintf("%s-a", region)
	instanceName := "your_instance_name"
	machineType := "n1-standard-1"
	sourceImage := "projects/debian-cloud/global/images/family/debian-10"
	networkName := "global/networks/default"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return "", fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	req := &computepb.InsertInstanceRequest{
		Project: projectID,
		Zone:    zone,
		InstanceResource: &computepb.Instance{
			Name: proto.String(instanceName),
			Disks: []*computepb.AttachedDisk{
				{
					InitializeParams: &computepb.AttachedDiskInitializeParams{
						DiskSizeGb:  proto.Int64(10),
						SourceImage: proto.String(sourceImage),
					},
					AutoDelete: proto.Bool(true),
					Boot:       proto.Bool(true),
					Type:       proto.String(computepb.AttachedDisk_PERSISTENT.String()),
				},
			},
			MachineType: proto.String(fmt.Sprintf("zones/%s/machineTypes/%s", zone, machineType)),
			NetworkInterfaces: []*computepb.NetworkInterface{
				{
					Name: proto.String(networkName),
				},
			},
		},
	}
	// Define the startup script that runs your executable
	startupScript := fmt.Sprintf(`#!/bin/bash
echo "Hello, World! The VM has started successfully."
echo "Executable ID: %s"`, executableID)

	// Add the startup script to your instance's metadata
	req.InstanceResource.Metadata = &computepb.Metadata{
		Items: []*computepb.Items{
			{
				Key:   proto.String("startup-script"),
				Value: proto.String(startupScript),
			},
		},
	}

	op, err := instancesClient.Insert(ctx, req)
	if err != nil {
		return "", fmt.Errorf("unable to create instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return "", fmt.Errorf("unable to wait for the operation: %w", err)
	}

	return instanceName, nil

}
