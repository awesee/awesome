package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/awesee/awesome/util"
	"github.com/spf13/cobra"
)

var gitHubProxy string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Install the latest version of awesome",
	Run: func(cmd *cobra.Command, args []string) {
		targetFile, err := os.Executable()
		util.CheckErr(err)
		verbosePrintln(targetFile)
		binaryName := fmt.Sprintf("awesome-%s-%s%s", runtime.GOOS, runtime.GOARCH, exeSuffix)
		downloadUrl := fmt.Sprintf("https://github.com/awesee/awesome/releases/latest/download/%s", binaryName)
		if gitHubProxy > "" {
			parsedUrl, err := url.ParseRequestURI(gitHubProxy)
			util.CheckErr(err)
			parsedUrl.Path = downloadUrl
			downloadUrl = parsedUrl.String()
		}
		verbosePrintln(downloadUrl)
		resp, err := http.Get(downloadUrl)
		util.CheckErr(err)
		if resp.StatusCode != http.StatusOK {
			resp.Status = strings.TrimPrefix(resp.Status, fmt.Sprintf("%d ", resp.StatusCode))
			if resp.Status == "" {
				resp.Status = http.StatusText(resp.StatusCode)
			}
			fmt.Printf("HTTP/%d.%d %03d %s\n", resp.ProtoMajor, resp.ProtoMinor, resp.StatusCode, resp.Status)
			return
		}
		data, err := io.ReadAll(resp.Body)
		util.CheckErr(err)
		sourceFileDir := filepath.Join(os.TempDir(), "awesee")
		util.CheckErr(os.MkdirAll(sourceFileDir, fs.ModePerm))
		sourceFile := filepath.Join(sourceFileDir, binaryName)
		util.CheckErr(os.WriteFile(sourceFile, data, fs.ModePerm))
		verbosePrintln(sourceFile)
		util.CheckErr(os.Chmod(sourceFileDir, fs.ModePerm))
		util.CheckErr(os.Rename(sourceFile, targetFile))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&gitHubProxy, "proxy", "", "github proxy")
}
