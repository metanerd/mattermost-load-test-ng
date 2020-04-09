// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package terraform

import (
	"fmt"
)

// Info displays information about the current load-test deployment.
func (t *Terraform) Info() error {
	output, err := t.getOutput()
	if err != nil {
		return err
	}

	t.displayInfo(output)

	return nil
}

func (t *Terraform) displayInfo(output *terraformOutput) {
	fmt.Println("==================================================")
	fmt.Println("Deployment information:")
	fmt.Println("Mattermost URL: http://" + output.Proxy.Value.PublicDNS)
	fmt.Println("App Servers:")
	for _, instance := range output.Instances.Value {
		fmt.Println("- " + instance.Tags.Name + ": " + instance.PublicIP)
	}
	fmt.Println("Load Agents:")
	for _, agent := range output.Agents.Value {
		fmt.Println("- " + agent.Tags.Name + ": " + agent.PublicIP)
	}
	if len(output.Agents.Value) > 0 {
		fmt.Println("Coordinator: " + output.Agents.Value[0].PublicIP)
	}
	fmt.Println("Grafana URL: http://" + output.MetricsServer.Value.PublicIP + ":3000")
	fmt.Println("Prometheus URL: http://" + output.MetricsServer.Value.PublicIP + ":9090")
	fmt.Println("DB reader endpoint: " + output.DBCluster.Value.ReaderEndpoint)
	fmt.Println("DB cluster endpoint: " + output.DBCluster.Value.ClusterEndpoint)
	fmt.Println("==================================================")
}
