package pkg

import (
	"fmt"
	"strings"

	"github.com/ashish-amarnath/aws-elb-gc/utils"
	"github.com/golang/glog"
)

const (
	describeELBCmdFmt     = "aws --profile %s --region %s elb describe-load-balancers | jq .LoadBalancerDescriptions[].LoadBalancerName -r"
	describeELBTagsCmdFmt = "aws --profile %s --region %s elb describe-tags --load-balancer-name %s | jq .TagDescriptions[].Tags[].Key -r"
	deleteELBCmdFmt       = "aws --profile %s --region %s elb delete-load-balancer --load-balancer-name %s"
)

func GetAllELBs(profile string, region string) (elbs []string, err error) {
	descELBCmd := fmt.Sprintf(describeELBCmdFmt, profile, region)
	cmdOut, err := utils.RunBashCmd(descELBCmd)
	elbs = nil
	if err == nil {
		elbs = strings.Split(strings.TrimSpace(cmdOut), "\n")
		err = nil
	}
	return
}

func GetELBTags(profile, region, elbName string) (tags []string, err error) {
	getELBTagsCmd := fmt.Sprintf(describeELBTagsCmdFmt, profile, region, elbName)
	cmdOut, err := utils.RunBashCmd(getELBTagsCmd)
	tags = nil
	if err == nil {
		tags = strings.Split(strings.TrimSpace(cmdOut), "\n")
		err = nil
	}
	return
}

func Delete(profile, region, elbName string) {
	deleteELBCmd := fmt.Sprintf(deleteELBCmdFmt, profile, region, elbName)
	_, err := utils.RunBashCmd(deleteELBCmd)
	if err != nil {
		glog.Errorf("Failed to delete ")
	}
	glog.V(5).Infof("ELB %s deleted", elbName)
}

func GetUniqueTags(awsProfile, awsRegion string) {
	tagMap := make(map[string]int)
	allELBs, err := GetAllELBs(awsProfile, awsRegion)
	if err == nil {
		for _, elb := range allELBs {
			tags, err := GetELBTags(awsProfile, awsRegion, elb)
			if err == nil {
				for _, tag := range tags {
					cur := tagMap[tag]
					tagMap[tag] = cur + 1
				}
			} else {
				glog.Errorf("Failed to get tags on ELB [%s]", elb)
			}
		}
	} else {
		glog.Errorf("Failed to describe ELBs using profile=%s and region=%s", awsProfile, awsRegion)
	}

	for tag := range tagMap {
		glog.V(1).Infof("%d ELBs have tag %s", tagMap[tag], tag)
	}
}
