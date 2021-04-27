package goseelog

import (
	"fmt"
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func init() {
	DisableLog()
	LoadAppConfig()
}

// 根据配置文件初始化seelog的配置信息，这里我们把配置文件通过字符串读取设置好了，当然也可以通过
// 读取XML文件
func LoadAppConfig() {
	// seelog: minlevel参数可选，如果被配置，高于或等于此级别的日志会被记录，同理maxlevel
	// outputs: 输出信息的目的地，这里分成了两份数据，一份记录到 log rotate文件里面，另一份设置了 filter，如果
	// 这个错误级别是 critical，那么将发送报警邮件
	// formats: 定义了各种日志格式
	// UseLogger: 设置当前的日志器为相应的日志处理
	appConfig := `
	<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="/data/logs/roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
            <file path="/data/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="astaxie@gmail.com" sendername="ShortUrl API" hostname="smtp.gmail.com" hostport="587" username="mailusername" password="mailpassword">
                <recipient address="xiemengjun@gmail.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
	    <format id="critical" format="%File %FullPath %Func %Msg%n" />
	    <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
	`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UserLogger(logger)

}

// 初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
func DisableLog() {
	Logger = seelog.Disabled
}

func UserLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

func SeelogMain() {
	defer seelog.Flush()
	seelog.Info("Hello from Seelog!")
}

