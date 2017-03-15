// Copyright © 2017 Suzuki Shunsuke <suzuki.shunsuke.1989@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

const Description = `
  Args:
  
    <user>/<repo>[/<branch or tag or revision>] <path> [-o <path>] [-O]
    <user>/<repo>/<branch or tag or revision>/<path> [-o <path>] [-O]
  
    user     - GitHub repository's user name
    repo     - GitHub repository's repository name
    branch   - GitHub repository's branch name
    tag      - GitHub repository's tag name
    revision - GitHub repository's revision
    path     - Relative path from repository's root directory to the target file
`

var GetCommand = cli.Command{
	Name:        "get",
	Usage:       "Get the file from GitHub",
	Description: Description,
	Action: func(c *cli.Context) error {
		arg0 := c.Args().Get(0)
		url_part := strings.Split(arg0, "/")
		var url string
		const url_head = "https://raw.githubusercontent.com/"
		L := c.NArg()
		switch n := len(url_part); true {
		case n == 2:
			if L != 2 {
				fmt.Fprint(os.Stderr, "[ghscr][Error] invalid arguments\n")
				os.Exit(1)
			}
			url = url_head + arg0 + "/master/" + c.Args().Get(1)
		case n == 3:
			if L != 2 {
				fmt.Fprint(os.Stderr, "[ghscr][Error] invalid arguments\n")
				os.Exit(1)
			}
			url = url_head + arg0 + "/" + c.Args().Get(1)
		case n > 3:
			url = url_head + arg0
		default:
			fmt.Fprint(os.Stderr, "[ghscr][Error] invalid arguments\n")
			os.Exit(1)
		}

		if c.Bool("O") {
			err := exec.Command("curl", url, "-O").Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "[ghscr][Error] %s\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}

		if len(c.String("o")) > 0 {
			err := exec.Command("curl", url, "-o", c.String("o")).Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "[ghscr][Error] %s\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}

		out, err := exec.Command("curl", url).Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ghscr][Error] %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
		return nil
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "O",
			Usage: "curl's -O option",
		},
		cli.StringFlag{
			Name:  "o",
			Value: "",
			Usage: "curl's -o option",
		},
	},
}
