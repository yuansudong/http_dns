
package cmdline
import(
	"fmt"
	"github.com/yuansudong/cobra"
	"github.com/yuansudong/http_dns/flags"
	
	service "github.com/yuansudong/http_dns/exec/service"
	
)

var (
	// _Branch 分支名称
	_GitBranch string = "UNKNOWN"
	// _GitCommitID 最近一次的提交ID
	_GitCommit string = "UNKNOWN"
	// _GitAccount 提交人的名字
	_GitAccount string = "UNKNOWN"
	// _DateTime 编译的时间
	_DateTime string = "UNKNOWN"
	// _GoVersion GO的编译版本
	_GoVersion string = "UNKNOWN"
	// _OS 编译时的操作系统
	_OS string = "UNKNOWN"
	// _CPU类型
	_Arch string = "UNKNOWN"
)

// _InitVersion 用于初始化版本
func _InitVersion() *cobra.Command {
	mCommand := new(cobra.Command)
	mCommand.Use = "version"
	mCommand.Long = "查看编译以及版本信息"
	mCommand.Short = "查看编译以及版本信息"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("App Name     :   ", "http_dns")
		fmt.Println("App Version  :   ", "v0.0.1")
		fmt.Println("Git Branch   :   ", _GitBranch)
		fmt.Println("Git Commit   :   ", _GitCommit)
		fmt.Println("Git Account  :   ", _GitAccount)
		fmt.Println("Go Version   :   ", _GoVersion)
		fmt.Println("Build System :   ", _OS)
		fmt.Println("Build Time   :   ", _DateTime)
		fmt.Println("Build Arch   :   ", _Arch)
	}
	return mCommand
}

var _Root *cobra.Command = _InitRoot()
func _InitRoot() *cobra.Command {
	mGlobalsFlags :=  new(flags.GlobalFlag)
	mRootCommand := new(cobra.Command)
	mRootCommand.Use = "http_dns" 
	mRootCommand.Long = "所有的域名查询服务都要走这里"
	
		
		mGlobalsFlags.F1 = mRootCommand.PersistentFlags().String(
			"F1",
			"F1_DEF", 
			"F1=V1",
		)
		
	
	
    mRootCommand.AddCommand(_InitSubService(mGlobalsFlags))
	
	mRootCommand.AddCommand(_InitVersion())
	return mRootCommand
}

// Execute 执行入口
func Execute() error {
	return _Root.Execute()
}

func _InitSubService(mGlobal *flags.GlobalFlag) *cobra.Command {
	mLocal := new(flags.LocalServiceFlag)
	mCommand :=  new(cobra.Command)
	mCommand.Use = "service"
	mCommand.Long = "开启一个http_dns服务"
	mCommand.Short = "开启一个http_dns服务"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		service.Startup(mGlobal,mLocal)
	} 
	
		
		mLocal.DbDns = mCommand.PersistentFlags().String(
			"db_dns",
			"loadlhost", 
			"A1=V1",
		)
		
	
	return mCommand
}


