package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		inst, err := compute.NewInstance(ctx, "myinstance", &compute.InstanceArgs{
			BootDisk: &compute.InstanceBootDiskArgs{
				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String("ubuntu-os-pro-cloud/ubuntu-pro-2004-lts"),
				},
			},
			MachineType: pulumi.String("n1-standard-1"),
			Zone:        pulumi.String("us-central1-a"),
			Project:     pulumi.String("sopes2-proyecto-348900"),
			NetworkInterfaces: &compute.InstanceNetworkInterfaceArray{
				&compute.InstanceNetworkInterfaceArgs{
					Network: pulumi.String("default"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("instanceName", inst.Name)
		return nil
	})
}
