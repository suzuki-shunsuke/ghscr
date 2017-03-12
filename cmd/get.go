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

	"github.com/spf13/cobra"
)

var option_o = ""
var option_O = false

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the file from GitHub",
    Long: "Get the file from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
        var url_part = strings.Split(args[0], "/")
        var url string
        const url_head = "https://raw.githubusercontent.com/"
        switch c := len(url_part); true {
        case c == 2:
            if len(args) != 2 {
                fmt.Fprint(os.Stderr, "[ghscr][Error] invalid arguments\n")
                os.Exit(1)
            }
            url = url_head + args[0] + "/master/" + args[1]
        case c == 3:
            if len(args) != 2 {
                fmt.Fprint(os.Stderr, "[ghscr][Error] invalid arguments\n")
                os.Exit(1)
            }
            url = url_head + args[0] + "/" + args[1]
        case c > 3:
            url = url_head + args[0]
        default:
            fmt.Fprint(os.Stderr, "[ghscr][Error] invalid arguments\n")
            os.Exit(1)
        }

        if option_O {
            err := exec.Command("curl", url, "-O").Run()
            if err != nil {
                fmt.Fprintf(os.Stderr, "[ghscr][Error] %s\n", err)
                os.Exit(1)
            }
            os.Exit(0)
        }

        if len(option_o) > 0 {
            err := exec.Command("curl", url, "-o", option_o).Run()
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
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    getCmd.Flags().BoolVarP(&option_O, "remote-name", "O", false, "curl -O option")
    getCmd.Flags().StringVarP(&option_o, "", "o", "", "curl -o option")
}
