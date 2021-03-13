
package cmdline
import(
	"fmt"
	"github.com/yuansudong/cobra"
	
)
// GlobalFlag 全局Flag
type GlobalFlag struct {
	
	F1 *string
	
	F2 *int32
	
	F3 *int64
	
	F4 *float32
	
}


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
		fmt.Println("App Version  :   ", "")
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
	mGlobalsFlags :=  new(GlobalFlag)
	mRootCommand := new(cobra.Command)
	mRootCommand.Use = "http_dns" 
	mRootCommand.Long = "这是一个Example.exe的程序"
	
		
		mGlobalsFlags.F1 = mRootCommand.PersistentFlags().String(
			"F1",
			"F1_DEF", 
			"F1=V1",
		)
		
	
		
		mGlobalsFlags.F2 = mRootCommand.PersistentFlags().Int32(
			"F2",
			32, 
			"F2=V2",
		)
		
	
		
		mGlobalsFlags.F3 = mRootCommand.PersistentFlags().Int64(
			"F3",
			64, 
			"F3=V3",
		)
		
	
		
		mGlobalsFlags.F4 = mRootCommand.PersistentFlags().Float32(
			"F4",
			32.00, 
			"F4=V4",
		)
		
	
	
	mRootCommand.AddCommand(_InitVersion())
	return mRootCommand
}

// Execute 执行入口
func Execute() error {
	return _Root.Execute()
}


