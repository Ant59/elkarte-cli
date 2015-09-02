package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "elkarte"
	app.Usage = "Create development environments for Elkarte, addons and themes"

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create a new project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "b, branch",
					Value: "master",
					Usage: "Branch of Elkarte to fetch (default: master)",
				},
			},
			Action: func(c *cli.Context) {
				switch c.Args().First() {
				case "addon":
					println("Creating new addon")
					name := c.Args()[1]
					os.Mkdir(name, 0755)
					os.Chdir(name)

					cmd := "git"
					args := []string{"clone", "https://github.com/Ant59/elkarte-vagrant-addon.git", "."}
					if err := exec.Command(cmd, args...).Run(); err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
					println("Successfully fetched environment")
					println("Setting up...")
					os.Mkdir(name, 0755)

					read, err := ioutil.ReadFile("Vagrantfile")
					if err != nil {
						panic(err)
					}

					newContents := strings.Replace(string(read), "addon_name", name, -1)

					err = ioutil.WriteFile("Vagrantfile", []byte(newContents), 0)
					if err != nil {
						panic(err)
					}
					println("Ready!")
				case "theme":
					println("Creating new theme")
					name := c.Args()[1]
					os.Mkdir(name, 0755)
					os.Chdir(name)

					cmd := "git"
					args := []string{"clone", "https://github.com/Ant59/elkarte-vagrant-theme.git", "."}
					if err := exec.Command(cmd, args...).Run(); err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
					println("Successfully fetched environment")
					println("Setting up...")
					os.Mkdir(name, 0755)

					read, err := ioutil.ReadFile("Vagrantfile")
					if err != nil {
						panic(err)
					}

					newContents := strings.Replace(string(read), "theme_name", name, -1)

					err = ioutil.WriteFile("Vagrantfile", []byte(newContents), 0)
					if err != nil {
						panic(err)
					}
					println("Ready!")
				case "elkarte":
					println("Fetching Elkarte and setting up environment")
					name := c.Args()[1]
					os.Mkdir(name, 0755)
					os.Chdir(name)

					cmd := "git"
					args := []string{"clone", "https://github.com/Ant59/elkarte-vagrant.git", "."}
					if err := exec.Command(cmd, args...).Run(); err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
					println("Successfully fetched environment")
					cmd = "git"
					args = []string{"clone", "https://github.com/elkarte/Elkarte.git"}
					if err := exec.Command(cmd, args...).Run(); err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
					println("Successfully fetched Elkarte")
					println("Ready!")
				default:
					println("You must enter 'addon', 'theme' or 'elkarte'")
				}
			},
		},
	}

	app.Run(os.Args)
}
