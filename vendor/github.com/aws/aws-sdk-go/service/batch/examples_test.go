// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package batch_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleBatch_CancelJob() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.CancelJobInput{
		JobId:  aws.String("String"), // Required
		Reason: aws.String("String"), // Required
	}
	resp, err := svc.CancelJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_CreateComputeEnvironment() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.CreateComputeEnvironmentInput{
		ComputeEnvironmentName: aws.String("String"), // Required
		ServiceRole:            aws.String("String"), // Required
		Type:                   aws.String("CEType"), // Required
		ComputeResources: &batch.ComputeResource{
			InstanceRole: aws.String("String"), // Required
			InstanceTypes: []*string{ // Required
				aws.String("String"), // Required
				// More values...
			},
			MaxvCpus: aws.Int64(1), // Required
			MinvCpus: aws.Int64(1), // Required
			SecurityGroupIds: []*string{ // Required
				aws.String("String"), // Required
				// More values...
			},
			Subnets: []*string{ // Required
				aws.String("String"), // Required
				// More values...
			},
			Type:             aws.String("CRType"), // Required
			BidPercentage:    aws.Int64(1),
			DesiredvCpus:     aws.Int64(1),
			Ec2KeyPair:       aws.String("String"),
			SpotIamFleetRole: aws.String("String"),
			Tags: map[string]*string{
				"Key": aws.String("String"), // Required
				// More values...
			},
		},
		State: aws.String("CEState"),
	}
	resp, err := svc.CreateComputeEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_CreateJobQueue() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.CreateJobQueueInput{
		ComputeEnvironmentOrder: []*batch.ComputeEnvironmentOrder{ // Required
			{ // Required
				ComputeEnvironment: aws.String("String"), // Required
				Order:              aws.Int64(1),         // Required
			},
			// More values...
		},
		JobQueueName: aws.String("String"), // Required
		Priority:     aws.Int64(1),         // Required
		State:        aws.String("JQState"),
	}
	resp, err := svc.CreateJobQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DeleteComputeEnvironment() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DeleteComputeEnvironmentInput{
		ComputeEnvironment: aws.String("String"), // Required
	}
	resp, err := svc.DeleteComputeEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DeleteJobQueue() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DeleteJobQueueInput{
		JobQueue: aws.String("String"), // Required
	}
	resp, err := svc.DeleteJobQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DeregisterJobDefinition() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DeregisterJobDefinitionInput{
		JobDefinition: aws.String("String"), // Required
	}
	resp, err := svc.DeregisterJobDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DescribeComputeEnvironments() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DescribeComputeEnvironmentsInput{
		ComputeEnvironments: []*string{
			aws.String("String"), // Required
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.DescribeComputeEnvironments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DescribeJobDefinitions() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DescribeJobDefinitionsInput{
		JobDefinitionName: aws.String("String"),
		JobDefinitions: []*string{
			aws.String("String"), // Required
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
		Status:     aws.String("String"),
	}
	resp, err := svc.DescribeJobDefinitions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DescribeJobQueues() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DescribeJobQueuesInput{
		JobQueues: []*string{
			aws.String("String"), // Required
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.DescribeJobQueues(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_DescribeJobs() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.DescribeJobsInput{
		Jobs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_ListJobs() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.ListJobsInput{
		JobQueue:   aws.String("String"), // Required
		JobStatus:  aws.String("JobStatus"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_RegisterJobDefinition() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.RegisterJobDefinitionInput{
		JobDefinitionName: aws.String("String"),            // Required
		Type:              aws.String("JobDefinitionType"), // Required
		ContainerProperties: &batch.ContainerProperties{
			Image:  aws.String("String"), // Required
			Memory: aws.Int64(1),         // Required
			Vcpus:  aws.Int64(1),         // Required
			Command: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Environment: []*batch.KeyValuePair{
				{ // Required
					Name:  aws.String("String"),
					Value: aws.String("String"),
				},
				// More values...
			},
			JobRoleArn: aws.String("String"),
			MountPoints: []*batch.MountPoint{
				{ // Required
					ContainerPath: aws.String("String"),
					ReadOnly:      aws.Bool(true),
					SourceVolume:  aws.String("String"),
				},
				// More values...
			},
			Privileged:             aws.Bool(true),
			ReadonlyRootFilesystem: aws.Bool(true),
			Ulimits: []*batch.Ulimit{
				{ // Required
					HardLimit: aws.Int64(1),         // Required
					Name:      aws.String("String"), // Required
					SoftLimit: aws.Int64(1),         // Required
				},
				// More values...
			},
			User: aws.String("String"),
			Volumes: []*batch.Volume{
				{ // Required
					Host: &batch.Host{
						SourcePath: aws.String("String"),
					},
					Name: aws.String("String"),
				},
				// More values...
			},
		},
		Parameters: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RegisterJobDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_SubmitJob() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.SubmitJobInput{
		JobDefinition: aws.String("String"), // Required
		JobName:       aws.String("String"), // Required
		JobQueue:      aws.String("String"), // Required
		ContainerOverrides: &batch.ContainerOverrides{
			Command: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Environment: []*batch.KeyValuePair{
				{ // Required
					Name:  aws.String("String"),
					Value: aws.String("String"),
				},
				// More values...
			},
			Memory: aws.Int64(1),
			Vcpus:  aws.Int64(1),
		},
		DependsOn: []*batch.JobDependency{
			{ // Required
				JobId: aws.String("String"),
			},
			// More values...
		},
		Parameters: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.SubmitJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_TerminateJob() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.TerminateJobInput{
		JobId:  aws.String("String"), // Required
		Reason: aws.String("String"), // Required
	}
	resp, err := svc.TerminateJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_UpdateComputeEnvironment() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.UpdateComputeEnvironmentInput{
		ComputeEnvironment: aws.String("String"), // Required
		ComputeResources: &batch.ComputeResourceUpdate{
			DesiredvCpus: aws.Int64(1),
			MaxvCpus:     aws.Int64(1),
			MinvCpus:     aws.Int64(1),
		},
		ServiceRole: aws.String("String"),
		State:       aws.String("CEState"),
	}
	resp, err := svc.UpdateComputeEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBatch_UpdateJobQueue() {
	sess := session.Must(session.NewSession())

	svc := batch.New(sess)

	params := &batch.UpdateJobQueueInput{
		JobQueue: aws.String("String"), // Required
		ComputeEnvironmentOrder: []*batch.ComputeEnvironmentOrder{
			{ // Required
				ComputeEnvironment: aws.String("String"), // Required
				Order:              aws.Int64(1),         // Required
			},
			// More values...
		},
		Priority: aws.Int64(1),
		State:    aws.String("JQState"),
	}
	resp, err := svc.UpdateJobQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
