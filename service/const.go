package service

type (
	BackendType string
)

const (
	EnvNameDebugMode = "DEBUG_MODE"

	DefaultContinueAsNewPageSizeInBytes = 1024 * 1024

	// HttpStatusCodeSpecial4xxError is a special deprecated code for this because I can't find an official one for this case
	HttpStatusCodeSpecial4xxError = 420

	TaskQueue = "Interpreter_DEFAULT"

	StateStartApi        = "/api/v1/workflowState/start"
	StateDecideApi       = "/api/v1/workflowState/decide"
	WorkflowWorkerRpcApi = "/api/v1/workflowWorker/rpc"

	GetDataObjectsWorkflowQueryType      = "GetDataObjects"
	GetSearchAttributesWorkflowQueryType = "GetSearchAttributes"
	GetCurrentTimerInfosQueryType        = "GetCurrentTimerInfos"
	ContinueAsNewDumpQueryType           = "ContinueAsNewDump"
	DebugDumpQueryType                   = "DebugNewDump"
	PrepareRpcQueryType                  = "PrepareRpcQueryType"

	SearchAttributeGlobalVersion     = "IwfGlobalWorkflowVersion"
	SearchAttributeExecutingStateIds = "IwfExecutingStateIds"
	SearchAttributeIwfWorkflowType   = "IwfWorkflowType"

	BackendTypeCadence  BackendType = "cadence"
	BackendTypeTemporal BackendType = "temporal"

	IwfSystemConstPrefix = "__IwfSystem_"

	SkipTimerSignalChannelName    = IwfSystemConstPrefix + "SkipTimerChannel"
	FailWorkflowSignalChannelName = IwfSystemConstPrefix + "FailWorkflowChannel"
	UpdateConfigSignalChannelName = IwfSystemConstPrefix + "UpdateWorkflowConfig"
	ExecuteRpcSignalChannelName   = IwfSystemConstPrefix + "ExecuteRpc"

	WorkerUrlMemoKey = IwfSystemConstPrefix + "WorkerUrl"
)

var ValidIwfSystemSignalNames = map[string]bool{
	SkipTimerSignalChannelName:    true,
	FailWorkflowSignalChannelName: true,
	UpdateConfigSignalChannelName: true,
	ExecuteRpcSignalChannelName:   true,
}

const (
	GracefulCompletingWorkflowStateId = "_SYS_GRACEFUL_COMPLETING_WORKFLOW"
	ForceCompletingWorkflowStateId    = "_SYS_FORCE_COMPLETING_WORKFLOW"
	ForceFailingWorkflowStateId       = "_SYS_FORCE_FAILING_WORKFLOW"
	DeadEndWorkflowStateId            = "_SYS_DEAD_END"
)

var ValidClosingWorkflowStateId = map[string]bool{
	GracefulCompletingWorkflowStateId: true,
	ForceCompletingWorkflowStateId:    true,
	ForceFailingWorkflowStateId:       true,
	DeadEndWorkflowStateId:            true,
}
