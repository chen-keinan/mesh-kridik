package cli

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chen-keinan/go-command-eval/eval"
	"github.com/chen-keinan/go-user-plugins/uplugin"
	"github.com/chen-keinan/mesh-kridik/internal/cli/commands"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/internal/hooks"
	"github.com/chen-keinan/mesh-kridik/internal/logger"
	"github.com/chen-keinan/mesh-kridik/internal/startup"
	"github.com/chen-keinan/mesh-kridik/pkg/models"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/mitchellh/cli"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"plugin"
	"strings"
)

// StartCLI start ldx-prob audit tester
func StartCLI() {
	app := fx.New(
		// dependency injection
		fx.NopLogger,
		fx.Provide(NewLxdResultChan),
		fx.Provide(NewCompletionChan),
		fx.Provide(NewArgFunc),
		fx.Provide(NewCliArgs),
		fx.Provide(utils.NewKFolder),
		fx.Provide(initBenchmarkSpecData),
		fx.Provide(NewCliCommands),
		fx.Provide(NewCommandArgs),
		fx.Provide(createCliBuilderData),
		fx.Provide(logger.GetLog),
		fx.Invoke(StartCLICommand),
	)
	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
}

//initBenchmarkSpecData initialize benchmark spec file and save if to file system
func initBenchmarkSpecData(fm utils.FolderMgr, ad ArgsData) []utils.FilesInfo {
	err := utils.CreateHomeFolderIfNotExist(fm)
	if err != nil {
		panic(err)
	}
	err = utils.CreateBenchmarkFolderIfNotExist(ad.SpecType, ad.SpecVersion, fm)
	if err != nil {
		panic(err)
	}
	var filesData []utils.FilesInfo
	switch ad.SpecType {
	case "lxd":
		if ad.SpecVersion == "v1.0.0" {
			filesData, err = startup.GenerateLxdBenchmarkFiles()
		}
	}
	if err != nil {
		panic(err)
	}
	err = startup.SaveBenchmarkFilesIfNotExist(ad.SpecType, ad.SpecVersion, filesData)
	if err != nil {
		panic(err)
	}
	return filesData
}

//initBenchmarkSpecData initialize benchmark spec file and save if to file system
func initPluginFolders(fm utils.FolderMgr) {
	err := utils.CreatePluginsSourceFolderIfNotExist(fm)
	if err != nil {
		panic(err)
	}
	err = utils.CreatePluginsCompiledFolderIfNotExist(fm)
	if err != nil {
		panic(err)
	}
}

//loadAuditBenchPluginSymbols load API call plugin symbols
func loadAuditBenchPluginSymbols(log *zap.Logger) hooks.LxdBenchAuditResultHook {
	fm := utils.NewKFolder()
	sourceFolder, err := utils.GetPluginSourceSubFolder(fm)
	if err != nil {
		panic(fmt.Sprintf("failed tpo get plugin source sourceFolder %s", err.Error()))
	}
	compliledFolder, err := utils.GetCompilePluginSubFolder(fm)
	if err != nil {
		panic(fmt.Sprintf("failed to get plugin compiled sourceFolder %s", err.Error()))
	}

	pl := uplugin.NewPluginLoader(sourceFolder, compliledFolder)
	names, err := pl.Plugins(uplugin.CompiledExt)
	if err != nil {
		panic(fmt.Sprintf("failed to get plugin compiled plugins %s", err.Error()))
	}
	apiPlugin := hooks.LxdBenchAuditResultHook{Plugins: make([]plugin.Symbol, 0), Plug: pl}
	for _, name := range names {
		sym, err := pl.Load(name, common.MeshSecurityCheckResultHook)
		if err != nil {
			log.Error(fmt.Sprintf("failed to load sym %s error %s", name, err.Error()))
			continue
		}
		apiPlugin.Plugins = append(apiPlugin.Plugins, sym)
	}
	return apiPlugin
}

// init new plugin worker , accept audit result chan and audit result plugin hooks
func initPluginWorker(plChan chan models.MeshCheckResults, completedChan chan bool) {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	lxdHooks := loadAuditBenchPluginSymbols(log)
	pluginData := hooks.NewPluginWorkerData(plChan, lxdHooks, completedChan)
	worker := hooks.NewPluginWorker(pluginData, log)
	worker.Invoke()
}

//StartCLICommand invoke cli lxd command mesh-kridik cli
func StartCLICommand(fm utils.FolderMgr, plChan chan models.MeshCheckResults, completedChan chan bool, ad ArgsData, cmdArgs []string, commands map[string]cli.CommandFactory, log *logger.MeshKridikLogger) {
	// init plugin folders
	initPluginFolders(fm)
	// init plugin worker
	initPluginWorker(plChan, completedChan)
	if ad.Help {
		cmdArgs = cmdArgs[1:]
	}
	status, err := invokeCommandCli(cmdArgs, commands)
	if err != nil {
		log.Console(err.Error())
	}
	os.Exit(status)
}

//NewCommandArgs return new cli command args
// accept cli args and return command args
func NewCommandArgs(ad ArgsData) []string {
	cmdArgs := []string{"a"}
	cmdArgs = append(cmdArgs, ad.Filters...)
	return cmdArgs
}

//NewCliCommands return cli lxd obj commands
// accept cli args data , completion chan , result chan , spec files and return artay of cli commands
func NewCliCommands(ad ArgsData, plChan chan models.MeshCheckResults, completedChan chan bool, fi []utils.FilesInfo) []cli.Command {
	cmds := make([]cli.Command, 0)
	// invoke cli
	evaluator := eval.NewEvalCmd()
	cmds = append(cmds, commands.NewLxdAudit(ad.Filters, plChan, completedChan, fi, evaluator))
	return cmds
}

//NewArgFunc return args func
func NewArgFunc() SanitizeArgs {
	return ArgsSanitizer
}

//NewCliArgs return cli args
func NewCliArgs(sa SanitizeArgs) ArgsData {
	ad := sa(os.Args[1:])
	return ad
}

//NewCompletionChan return plugin Completion chan
func NewCompletionChan() chan bool {
	completedChan := make(chan bool)
	return completedChan
}

//NewLxdResultChan return plugin test result chan
func NewLxdResultChan() chan models.MeshCheckResults {
	plChan := make(chan models.MeshCheckResults)
	return plChan
}

//createCliBuilderData return cli params and commands
func createCliBuilderData(ca []string, cmd []cli.Command) map[string]cli.CommandFactory {
	// read cli args
	cmdFactory := make(map[string]cli.CommandFactory)
	// build cli commands
	for index, a := range cmd {
		cmdFactory[ca[index]] = func() (cli.Command, error) {
			return a, nil
		}
	}
	return cmdFactory
}

// invokeCommandCli invoke cli command with params
func invokeCommandCli(args []string, commands map[string]cli.CommandFactory) (int, error) {
	app := cli.NewCLI(common.MeshKridikCli, common.MeshKridikVersion)
	app.Args = append(app.Args, args...)
	app.Commands = commands
	app.HelpFunc = LxdProbeHelpFunc(common.MeshKridikCli)
	status, err := app.Run()
	return status, err
}

//ArgsSanitizer sanitize CLI arguments
var ArgsSanitizer SanitizeArgs = func(str []string) ArgsData {
	ad := ArgsData{SpecType: "lxd"}
	args := make([]string, 0)
	if len(str) == 0 {
		args = append(args, "")
	}
	for _, arg := range str {
		arg = strings.Replace(arg, "--", "", -1)
		arg = strings.Replace(arg, "-", "", -1)
		switch {
		case arg == "help", arg == "h":
			ad.Help = true
			args = append(args, arg)
		case strings.HasPrefix(arg, "s="):
			ad.SpecType = arg[len("s="):]
		case strings.HasPrefix(arg, "spec="):
			ad.SpecType = arg[len("spec="):]
		case strings.HasPrefix(arg, "v="):
			ad.SpecVersion = fmt.Sprintf("v%s", arg[len("v="):])
		case strings.HasPrefix(arg, "version="):
			ad.SpecVersion = fmt.Sprintf("v%s", arg[len("version="):])
		default:
			args = append(args, arg)
		}
	}
	if ad.SpecType == "lxd" && len(ad.SpecVersion) == 0 {
		ad.SpecVersion = "v1.0.0"
	}
	ad.Filters = args
	return ad
}

//ArgsData hold cli args data
type ArgsData struct {
	Filters     []string
	Help        bool
	SpecType    string
	SpecVersion string
}

//SanitizeArgs sanitizer func
type SanitizeArgs func(str []string) ArgsData

// LxdProbeHelpFunc mesh-kridik Help function with all supported commands
func LxdProbeHelpFunc(app string) cli.HelpFunc {
	return func(commands map[string]cli.CommandFactory) string {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf(startup.GetHelpSynopsis(), app))
		return buf.String()
	}
}
