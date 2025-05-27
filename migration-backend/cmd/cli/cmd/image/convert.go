package image

import (
	"fmt"

	"github.com/spf13/cobra"

	imageutil "github.com/likecoin/like-migration-backend/pkg/util/image"
)

var convertCmd = &cobra.Command{
	Use:   "convert src [target]",
	Short: "Convert image format",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			_ = cmd.Usage()
			return
		}

		src := args[0]
		var (
			target string
			err    error
		)

		if len(args) >= 2 {
			target = args[1]
		} else {
			target, err = imageutil.DefaultMimeTypeConversion.GetDefaultConvertedPath(src)
		}

		if err != nil {
			panic(err)
		}

		err = imageutil.Convert(src, target)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Converted image to %s.\n", target)
	},
}

func init() {
	ImageCmd.AddCommand(convertCmd)
}
