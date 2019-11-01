package cmd

import (
	"fmt"

	"github.com/sfuruya0612/snatch/internal/aws"
	"github.com/urfave/cli"
)

func DescribeLogGroups(c *cli.Context) error {
	profile := c.GlobalString("profile")
	region := c.GlobalString("region")
	flag := c.GlobalBool("f")

	fmt.Printf("\x1b[32mAWS_PROFILE: %s , REGION: %s\x1b[0m\n", profile, region)

	err := aws.DescribeLogGroups(profile, region, flag)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}