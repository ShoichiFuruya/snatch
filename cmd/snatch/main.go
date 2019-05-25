package main

import (
	"fmt"
	"os"

	"github.com/ShoichiFuruya/snatch/internal/aws"
	"github.com/urfave/cli"
)

var (
	date      string
	hash      string
	goversion string
)

func main() {
	//snatch := Exec(date, goversion)
	snatch := Exec()
	if err := snatch.Run(os.Args); err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		os.Exit(1)
	}
}

//func Exec(date, goversion) *cli.App {
func Exec() *cli.App {
	app := cli.NewApp()

	app.Name = "snatch"
	app.Usage = "Show AWS resources cli command. (Made in Golang)"
	//	app.Version = fmt.Printf("%s %s(%s), date, hash, goversion")
	app.Version = "0.1.0"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:   "ec2",
			Usage:  "Get EC2 list",
			Action: aws.DescribeEc2,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "profile, p",
					Value: "default",
					Usage: "Choose AWS credential.",
				},
				cli.StringFlag{
					Name:  "region",
					Value: "ap-northeast-1",
					Usage: "Select Region.",
				},
				// extra, -e フラグ追加予定
			},
		},
		{
			Name:   "rds",
			Usage:  "Get RDS list",
			Action: aws.DescribeRds,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "profile, p",
					Value: "default",
					Usage: "Choose AWS credential.",
				},
				cli.StringFlag{
					Name:  "region",
					Value: "ap-northeast-1",
					Usage: "Select Region.",
				},
			},
		},
	}
	return app
}