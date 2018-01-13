package main

import (
	"flag"
	"strings"

	"github.com/ashish-amarnath/aws-elb-gc/elb"
	"github.com/golang/glog"
)

type cliFlags struct {
	awsProfile *string
	awsRegion  *string
	tagKey     *string
	cmd        *string
}

const (
	gcCmd       = "gc"
	showTagsCmd = "show-tags"
)

func isELBTagged(tags []string, tagKeys []string) bool {
	tagged := false
	for _, tag := range tags {
		for _, key := range tagKeys {
			if tag == key {
				tagged = true
				break
			}
		}
		if tagged {
			break
		}
	}
	return tagged
}

func getCLIFlags() cliFlags {
	var flags cliFlags
	flags.awsProfile = flag.String("profile", "", "AWS profile to use in CLI commands")
	flags.awsRegion = flag.String("region", "", "AWS region to use in CLI commands")
	flags.tagKey = flag.String("tag-filter", "", "Coma separated list of tags by which ELBs need to be filtered")
	flags.cmd = flag.String("cmd", "show-tags", "Command that this GC tool should run.")
	flag.Parse()
	return flags
}

func showAllTags(awsProfile, awsRegion string) {
	pkg.GetUniqueTags(awsProfile, awsRegion)
}

func gcELB(awsProfile, awsRegion string, tagKeys []string) {
	allELBs, err := pkg.GetAllELBs(awsProfile, awsRegion)
	if err == nil {
		for _, elb := range allELBs {
			tags, err := pkg.GetELBTags(awsProfile, awsRegion, elb)
			if err == nil {
				if isELBTagged(tags, tagKeys) {
					pkg.Delete(awsProfile, awsRegion, elb)
				} else {
					glog.V(5).Infof("ELB [%s] tags are [%s]. Tag [%s] not found.", elb, strings.Join(tags, ";"), strings.Join(tagKeys, ", "))
				}
			} else {
				glog.Errorf("Failed to get tags on ELB [%s]", elb)
			}
		}
	} else {
		glog.Errorf("Failed to describe ELBs using profile=%s and region=%s", awsProfile, awsRegion)
	}
}

func main() {
	flags := getCLIFlags()
	glog.V(1).Infof("Running command %s, awsProfile=%s, awsRegion=%s", *flags.cmd, *flags.awsProfile, *flags.awsRegion)

	switch *flags.cmd {
	case gcCmd:
		gcELB(*flags.awsProfile, *flags.awsRegion, strings.Split(*flags.tagKey, ","))
		break
	case showTagsCmd:
		showAllTags(*flags.awsProfile, *flags.awsRegion)
		break
	default:
		break
	}
}
