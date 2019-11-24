package framegram

import (
	"flag"
	"log"
	"os"

	"github.com/estebanborai/framegram/src/task"
	"github.com/estebanborai/framegram/src/util"
)

// Task to execute
type Task string

const (
	// Resize operation
	Resize Task = "resize"
)

// Cli manages all possible framegram operations
type Cli struct {
	Args       []string
	Task       Task
	SourceFile string
	OutputPath string
	Dimensions string
}

func readFlags(cli *Cli) {
	isResizing := flag.Bool("resize", false, "Resizes an Image")
	inputFile := flag.String("src", "", "Source file")
	outputPath := flag.String("out", "", "Output path")
	dimensions := flag.String("dimensions", "", "Dimensions to set the image (Eg: 500x500 [width x height])")

	flag.Parse()

	if *isResizing {
		if *inputFile != "" && *outputPath != "" && *dimensions != "" {
			cli.Task = Resize
			cli.SourceFile = *inputFile
			cli.OutputPath = *outputPath
			cli.Dimensions = *dimensions
		} else {
			log.Fatal("Missing --src and --out arguments for resizing an image")
		}

		return
	}
}

// NewCli creates a new Cli instance and returns a pointer to it
func NewCli() *Cli {
	cli := new(Cli)
	cli.Args = os.Args[1:]

	readFlags(cli)

	return cli
}

// Start run tasks based on arguments
func (cli Cli) Start() {
	switch cli.Task {
	case Resize:
		sizeProfile, err := util.NewSizeProfile(cli.Dimensions)

		if err != nil {
			log.Fatal(err)
		}

		task.ResizeImage(cli.SourceFile, cli.OutputPath, *sizeProfile)
		break
	}
}