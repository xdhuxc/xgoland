package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/sirupsen/logrus"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		//Region: aws.String("us-east-1"),
	})
	if err != nil {
		logrus.Fatalf("failed to create session   %v", err)
	}

	client := elbv2.New(sess)

	var NextMarker string
	var params *elbv2.DescribeLoadBalancersInput
	lbs := make([]*elbv2.LoadBalancer, 0)

	for {
		if NextMarker == "" {
			params = &elbv2.DescribeLoadBalancersInput{}
		} else {
			params = &elbv2.DescribeLoadBalancersInput{
				Marker: aws.String(NextMarker),
			}
		}
		resp, err := client.DescribeLoadBalancers(params)
		if err != nil {
			logrus.Println(err.Error())
		}
		loadBalancers := resp.LoadBalancers
		fmt.Println(len(loadBalancers))
		for _, lb := range loadBalancers {
			lbs = append(lbs, lb)
			//fmt.Printf(lb.String())
			// fmt.Println()
		}
		NextMarker = aws.StringValue(resp.NextMarker)
		if NextMarker == "" {
			break
		}
	}
	fmt.Printf("the total load balancers are %d", len(lbs))
}
