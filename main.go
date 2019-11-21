package main

import (
	"fmt"
	"os"

	"github.com/sfuruya0612/snatch/cmd"
	"github.com/urfave/cli"
)

const version = "19.11.1"

var (
	date      string
	hash      string
	goversion string
)

func main() {
	snatch := New(date, hash, goversion)
	if err := snatch.Run(os.Args); err != nil {
		fmt.Printf("\n[ERROR]: %v\n", err)
		os.Exit(1)
	}
}

// New returns cli.App
func New(date, hash, goversion string) *cli.App {
	app := cli.NewApp()

	app.EnableBashCompletion = true

	app.Name = "snatch"
	app.Usage = "Cli command to get and display Amazon Web Services resources"

	if date != "" || hash != "" || goversion != "" {
		app.Version = fmt.Sprintf("%s %s (Build by: %s)", date, hash, goversion)
	} else {
		app.Version = version
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "profile, p",
			EnvVar: "AWS_PROFILE",
			Value:  "default",
			Usage:  "AWS credential (~/.aws/config) or read AWS_PROFILE environment variable",
		},
		cli.StringFlag{
			Name:  "region, r",
			Value: "ap-northeast-1",
			Usage: "Specify a valid AWS region",
		},
	}

	app.Before = cmd.Before

	app.Commands = []cli.Command{
		{
			Name:   "ec2",
			Usage:  "Get a list of EC2 resources",
			Action: cmd.ListEc2,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "tag, t",
					Usage: "The Key-Value of the tag to filter",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:   "log",
					Usage:  "Get the console output for the specified instance",
					Action: cmd.GetEc2Log,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "instanceid, i",
							Usage: "Set EC2 instance id",
						},
					},
				},
			},
		},
		{
			Name:   "rds",
			Usage:  "Get a list of RDS resources",
			Action: cmd.ListRds,
		},
		{
			Name:    "elasticache",
			Aliases: []string{"ec"},
			Usage:   "Get a list of ElastiCache Cluster resources",
			Action:  cmd.ListElasticache,
			Subcommands: []cli.Command{
				{
					Name:   "node",
					Usage:  "Get a list of ElastiCache Node resources",
					Action: cmd.ListReplicationGroups,
				},
			},
		},
		{
			Name:   "elb",
			Usage:  "Get a list of ELB(Classic) resources.",
			Action: cmd.ListElb,
		},
		{
			Name:   "elbv2",
			Usage:  "Get a list of ELB(Application & Network) resources",
			Action: cmd.ListElbV2,
		},
		{
			Name:   "route53",
			Usage:  "Get a list of Rotue53 Record resources",
			Action: cmd.ListHostedZones,
		},
		{
			Name:   "acm",
			Usage:  "Get a list of ACM resources",
			Action: cmd.ListCertificates,
		},
		{
			Name:   "s3",
			Usage:  "Get Objects in selected S3 Bucket at interactive prompt",
			Action: cmd.ListBuckets,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "l",
					Usage: "Get Objects list",
				},
			},
		},
		{
			Name:   "ssm",
			Usage:  "Start a session on your instances by launching bash or shell terminal",
			Action: cmd.StartSession,
			Subcommands: []cli.Command{
				{
					Name:   "run",
					Usage:  "Runs commands on one target instance",
					Action: cmd.SendCommand,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file, f",
							Usage: "Set execute file",
						},
						cli.StringFlag{
							Name:  "tag, t",
							Usage: "Set Key-Value of the tag (e.g. -t Name:test-ec2)",
						},
						cli.StringFlag{
							Name:  "instanceid, i",
							Usage: "Set EC2 instance id",
						},
					},
				},
			},
		},
		{
			Name:   "logs",
			Usage:  "Display messages for selected log groups and streams at interactive prompt",
			Action: cmd.DescribeLogGroups,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Like `tail -f`.",
				},
			},
		},
		{
			Name:   "cfn",
			Usage:  "Display a list of stacks",
			Action: cmd.DescribeStacks,
		},
	}

	return app
}
