package main

import (
	"fmt"
	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

func main() {

	accesskey := "<access key of your jdcloud account>"
	secretkey := "<secret key of your jdcloud account>"

	credentials := NewCredentials(accesskey, secretkey)
	myClient := NewVmClient(credentials)

	myRequest := NewDescribeInstancesRequest("cn-east-2")

	resp, err := myClient.DescribeInstances(myRequest)

	if err != nil {
		return
	}

	fmt.Println(resp.RequestID)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(resp.Result.Instances)
}

/*
bhwnpi3gg73qjai0a80u007fw10q021g
5
[{i-9ah3wjio3g k8s-node-vmzgta-h9n2fhydae g.n2.medium vpc-kthawbwukv subnet-o5h45bnfj1 10.0.32.3   running kubernetes cluster node img-e02k8n1218 {local true { 100} {     0 0 0  []  false false false  {    } []} vda attached} [] {1 true {port-l0ga0cyotg fa:16:3e:45:ef:0f vpc-kthawbwukv subnet-o5h45bnfj1 [{sg-j3t7zb1vd1 默认安全组开放全部端口}] 0 {10.0.32.3  } []}} [{2 true {port-tokseyzkcv fa:16:3e:c7:43:26 vpc-kthawbwukv subnet-0pfavmq0zp [{sg-j3t7zb1vd1 默认安全组开放全部端口}] 0 {10.0.0.3  } [{10.0.0.4  } {10.0.0.5  } {10.0.0.6  } {10.0.0.7  } {10.0.0.8  } {10.0.0.9  } {10.0.0.10  } {10.0.0.11  } {10.0.0.12  } {10.0.0.13  } {10.0.0.14  } {10.0.0.15  } {10.0.0.16  } {10.0.0.17  } {10.0.0.18  } {10.0.0.19  } {10.0.0.20  } {10.0.0.21  } {10.0.0.22  } {10.0.0.24  }]}}] 2019-02-27 13:18:42 cn-east-2a [] {postpaid_by_duration normal 2019-02-27T05:15:52Z  } {k8s-vm-ag-h9n2fhydae ag-bzpwtl76a0} 1 []} {i-ib5gsel9yr k8s-node-vm5p6l-h9n2fhydae g.n2.medium vpc-kthawbwukv subnet-o5h45bnfj1 10.0.32.4   running kubernetes cluster node img-e02k8n1218 {local true { 100} {     0 0 0  []  false false false  {    } []} vda attached} [] {1 true {port-dpte5zr0cm fa:16:3e:7d:5e:aa vpc-kthawbwukv subnet-o5h45bnfj1 [{sg-j3t7zb1vd1 默认安全组开放全部端口}] 0 {10.0.32.4  } []}} [{2 true {port-9h6ld8i997 fa:16:3e:0c:18:06 vpc-kthawbwukv subnet-0pfavmq0zp [{sg-j3t7zb1vd1 默认安全组开放全部端口}] 0 {10.0.0.23  } [{10.0.0.25  } {10.0.0.26  } {10.0.0.27  } {10.0.0.28  } {10.0.0.29  } {10.0.0.30  } {10.0.0.31  } {10.0.0.32  } {10.0.0.33  } {10.0.0.34  } {10.0.0.35  } {10.0.0.36  } {10.0.0.37  } {10.0.0.38  } {10.0.0.39  } {10.0.0.40  } {10.0.0.41  } {10.0.0.42  } {10.0.0.43  } {10.0.0.44  }]}}] 2019-02-27 13:18:43 cn-east-2b [] {postpaid_by_duration normal 2019-02-27T05:15:52Z  } {k8s-vm-ag-h9n2fhydae ag-bzpwtl76a0} 1 []} {i-j7wcghap08 k8s-cpfqkygk15-nat-vm-10a g.n2.medium vpc-kthawbwukv subnet-qb5ybcnwfu 10.0.48.3 fip-yyus1n5spg 114.67.88.224 running NodePodNat img-8zlcxtnjlo {local true { 40} {     0 0 0  []  false false false  {    } []} vda attached} [] {1 true {port-pq1fkyen79 fa:16:3e:1a:3c:a7 vpc-kthawbwukv subnet-qb5ybcnwfu [{sg-j3t7zb1vd1 默认安全组开放全部端口}] 0 {10.0.48.3 fip-yyus1n5spg 114.67.88.224} []}} [] 2019-02-27 13:16:12 cn-east-2a [] {postpaid_by_duration normal 2019-02-27T05:15:48Z  } { }  []} {i-69swuggnsx libing-centos74-workstation g.n2.large vpc-hevrjgh95e subnet-pw2kfw651v 10.0.0.3 fip-46rhvdt26h 114.67.88.149 running  img-2wul0u50re {cloud false { 0} {vol-j1shanhsmo cn-east-2a libing-centos74-workstation2  ssd.gp1 200 0 0 in-use []  false false false 2019-02-26 11:37:29 {    } []} vda attached} [] {1 true {port-zxsgq1vtzt fa:16:3e:4e:de:c4 vpc-hevrjgh95e subnet-pw2kfw651v [{sg-gvcqlm6mkw 默认安全组开放全部端口}] 0 {10.0.0.3 fip-46rhvdt26h 114.67.88.149} []}} [] 2019-02-26 11:37:53 cn-east-2a [] {prepaid_by_duration normal 2019-02-26T03:37:22Z 2019-03-26T15:59:59Z } { }  []} {i-lac1ba4v7j libing-centos74-workstation-2 g.n2.large vpc-hevrjgh95e subnet-pw2kfw651v 10.0.0.4 fip-1cbjoddpm2 114.67.88.150 running  img-2wul0u50re {cloud false { 0} {vol-vt8rud6r27 cn-east-2a libing-centos74-workstation1  ssd.gp1 200 0 0 in-use []  false false false 2019-02-26 11:37:29 {    } []} vda attached} [] {1 true {port-2ik2119d24 fa:16:3e:4d:93:79 vpc-hevrjgh95e subnet-pw2kfw651v [{sg-gvcqlm6mkw 默认安全组开放全部端口}] 0 {10.0.0.4 fip-1cbjoddpm2 114.67.88.150} []}} [] 2019-02-26 11:37:52 cn-east-2a [] {prepaid_by_duration normal 2019-02-26T03:37:22Z 2019-03-26T15:59:59Z } { }  []}]
*/