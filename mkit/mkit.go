package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
	"github.com/tal-tech/go-zero/core/load"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stat"
	"github.com/tal-tech/go-zero/tools/goctl/upgrade"
	model "github.com/uncleyeung/devtools/mkit/model/sql/command"
	"github.com/urfave/cli"
)

var (
	buildVersion = "1.0.0"
	commands     = []cli.Command{
		{
			Name:   "upgrade",
			Usage:  "upgrade mkit to latest version",
			Action: upgrade.Upgrade,
		},
		{
			Name:  "dao",
			Usage: "generate model code",
			Subcommands: []cli.Command{
				{
					Name:  "mysql",
					Usage: `generate mysql model`,
					Subcommands: []cli.Command{
						{
							Name:  "ddl",
							Usage: `generate mysql model from ddl`,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "src, s",
									Usage: "the path or path globbing patterns of the ddl",
								},
								cli.StringFlag{
									Name:  "dir, d",
									Usage: "the target dir",
								},
								cli.StringFlag{
									Name:  "style",
									Usage: "the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]",
								},
								cli.StringFlag{
									Name:  "gitlab, g",
									Usage: "the [your.gitlab.com] url",
								},
								cli.StringFlag{
									Name:  "repo, r",
									Usage: "the [your-repository] name",
								},
								cli.BoolFlag{
									Name:  "idea",
									Usage: "for idea plugin [optional]",
								},
								cli.StringFlag{
									Name:  "database, db",
									Usage: "the name of database [optional]",
								},
							},
							Action: model.MysqlDDL,
						},
						{
							Name:  "datasource",
							Usage: `generate model from datasource`,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "url",
									Usage: `the data source of database,like "root:password@tcp(127.0.0.1:3306)/database`,
								},
								cli.StringFlag{
									Name:  "table, t",
									Usage: `the table or table globbing patterns in the database`,
								},
								/*cli.BoolFlag{
									Name:  "cache, c",
									Usage: "generate code with cache [optional]",
								},*/
								cli.StringFlag{
									Name:  "dir, d",
									Usage: "the target dir",
								},
								cli.StringFlag{
									Name:  "gitlab, g",
									Usage: "the [your.gitlab.com] url",
								},
								cli.StringFlag{
									Name:  "repo, r",
									Usage: "the [your-repository] name",
								},
								/*cli.StringFlag{
									Name:  "style",
									Usage: "the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]",
								},*/
								cli.BoolFlag{
									Name:  "idea",
									Usage: "for idea plugin [optional]",
								},
							},
							Action: model.MySqlDataSource,
						},
					},
				},
			},
		},
	}
)

func main() {
	logx.Disable()
	load.Disable()
	stat.DisableLog()

	app := cli.NewApp()
	app.Usage = "a cli tool to generate code"
	app.Version = fmt.Sprintf("%s %s/%s", buildVersion, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	// cli already print error messages
	fmt.Println("beginning to generate...")
	if err := app.Run(os.Args); err != nil {
		fmt.Println(aurora.Red("error: " + err.Error()))
	}
}
