package commands

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var cmdHelp = &Command{
	Usage: "help [command]",
	Short: "Show help",
	Long:  `Shows usage for a command.`,
}

func init() {
	cmdHelp.Run = runHelp // break init loop
}

func runHelp(cmd *Command, args *Args) {
	if args.IsEmpty() {
		PrintUsage()
		return // not os.Exit(2); success
	}
	if args.Size() != 1 {
		log.Fatal("too many arguments")
	}

	for _, cmd := range All() {
		if cmd.Name() == args.First() {
			cmd.PrintUsage()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic: %q. Run 'gh help'.\n", args.First())
	os.Exit(2)
}

var usageTemplate = template.Must(template.New("usage").Parse(`Usage: gh [command] [options] [arguments]

Branching Commands:{{range .BranchingCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-16s"}}  {{.Short}}{{end}}{{end}}{{end}}

Remote Commands:{{range .RemoteCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-16s"}}  {{.Short}}{{end}}{{end}}{{end}}

GitHub Commands:{{range .GitHubCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-16s"}}  {{.Short}}{{end}}{{end}}{{end}}

See 'gh help [command]' for more information about a command.
`))

func PrintUsage() {
	usageTemplate.Execute(os.Stdout, struct {
		BranchingCommands []*Command
		RemoteCommands    []*Command
		GitHubCommands    []*Command
	}{
		Branching,
		Remote,
		GitHub,
	})
}

func Usage() {
	PrintUsage()
	os.Exit(2)
}
