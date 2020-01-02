package define

// Admin or Normal User
type Role uint8

const (
	NormalUser Role = iota + 1 // 普通用户 只对自已创建的主机或者主机组具有操作权限
	AdminUser                  // 管理员 具有所有操作
)

// task type
type TaskType uint8

const (
	Shell TaskType = iota + 1
	Api
)

// run crocodile as server or client
type RunMode uint8

const (
	Server RunMode = iota + 1
	Client
)

// task type (parent task,master task, child task)
type TaskRespType uint8

const (
	MasterTask = iota + 1
	ParentTask
	ChildTask
)

func GetUserByRole(r Role) string {
	switch r {
	case AdminUser:
		return "Admin"
	case NormalUser:
		return "Normal"
	default:
		return ""
	}
}

// 定义结构体
type common struct {
	Id         string `json:"id"`
	Name       string `json:"name,omitempty"`
	CreateTime string `json:"create_time,omitempty"` // 创建时间
	UpdateTime string `json:"update_time,omitempty"` // 最后一次更新时间
	Remark     string `json:"remark"`                // 备注
}

// 用户
type User struct {
	Role     Role   `json:"role" binding:"gte=1,lte=2"`     // 用户类型: 2 管理员 1 普通用户
	Forbid   int    `json:"forbid" binding:"gte=0,lte=1"`   // 禁止用户: 0 未禁止 1 已禁止
	Email    string `json:"email" binding:"required,email"` // 用户邮箱 日后任务的通知信息会发送给此邮件
	Password string `json:"password,omitempty"`
	common
}

// 主机组
type HostGroup struct {
	HostsID []string `json:"addrs"`
	//Hosts       []*Host `json:"hosts,omitempty"`       // WorkerId
	CreateByUid string `json:"create_byuid"` // 创建人ID
	CreateBy    string `json:"create_by"`    // 创建人ID
	common
}

// 主机信息
// 定时更新主机信息
type Host struct {
	common
	Addr               string `json:"addr"`     // 主机IP
	HostName           string `json:"hostname"` // 主机名
	Online             int    `json:"online"`   // 主机是否在线 0 not online,1 online
	Version            string `json:"version"`  // 版本号
	LastUpdateTimeUnix int64  `json:"last_updatetimeunix"`
	LastUpdateTime     string `json:"last_updatetime"`
}

// 任务
type Task struct {
	TaskType          TaskType    `json:"task_type"`                       // 任务类型
	TaskData          interface{} `json:"taskData"`                        // 任务数据
	Run               int         `json:"run"`                             // 0 为不能运行 1 为可以运行
	ParentTaskIds     []string    `json:"parent_taskids"`                  // 父任务 运行任务前先运行父任务 以父或子任务运行时 任务不会执行自已的父子任务，防止循环依赖
	ParentRunParallel int         `json:"parent_runparallel"`              // 是否以并行运行父任务 0否 1是
	ChildTaskIds      []string    `json:"child_taskids"`                   // 子任务 运行结束后运行子任务
	ChildRunParallel  int         `json:"child_runparallel"`               // 是否以并行运行子任务 0否 1是
	CreateBy          string      `json:"create_by"`                       // 创建人
	CreateByUid       string      `json:"create_byuid"`                    // 创建人ID
	HostGroup         string      `json:"host_group"`                      // 执行计划
	HostGroupId       string      `json:"host_groupid" binding:"required"` // 主机组ID

	Cronexpr     string   `json:"cronexpr" binding:"required"` // 执行任务表达式
	Timeout      int      `json:"timeout"`                     // 任务超时时间
	AlarmUserIds []string `json:"alarm_userids"`               // 报警用户 多个用户
	AutoSwitch   int      `json:"auto_switch"`                 // 运行失败自动切换到其他主机上
	common
}

// 返回值
type TaskResp struct {
	TaskType   TaskRespType `json:"task_type"` // 1 主任务 2 父任务 3 子任务
	TaskId     string       `json:"task_td"`
	Code       int32        `json:"code"`
	ErrMsg     string       `json:"err_msg"`
	RespData   string       `json:"resp_data"`
	WorkerHost string       `json:"worker_host"`
}

// 运行的任务
type RunTask struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	StartTime     string `json:"start_time"`
	StartTimeUnix int64  `json:"start_timeunix"`
	RunTime       int    `json:"run_time"`
}

// 日志
type Log struct {
	RunByTaskId  string      `json:"run_bytaskId"`
	TaskResps    []*TaskResp `json:"task_resps"`
	StartTime    int64       `json:"start_time"`    // ms
	EndTime      int64       `json:"end_time"`      // ms
	TotalRunTime int         `json:"total_runtime"` // ms
}
