package service

type (
	BackendType string
)

const (
	DefaultContinueAsNewPageSizeInBytes = 1024 * 1024

	TaskQueue = "Interpreter_DEFAULT"

	StateStartApi        = "/api/v1/workflowState/start"
	StateDecideApi       = "/api/v1/workflowState/decide"
	WorkflowWorkerRpcApi = "/api/v1/workflowWorker/rpc"

	GetDataObjectsWorkflowQueryType      = "GetDataObjects"
	GetSearchAttributesWorkflowQueryType = "GetSearchAttributes"
	GetCurrentTimerInfosQueryType        = "GetCurrentTimerInfos"
	ContinueAsNewDumpQueryType           = "ContinueAsNewDump"
	DebugDumpQueryType                   = "DebugNewDump"
	PrepareRpcQueryType                  = "GetDataObjects"

	SearchAttributeGlobalVersion     = "IwfGlobalWorkflowVersion"
	SearchAttributeExecutingStateIds = "IwfExecutingStateIds"
	SearchAttributeIwfWorkflowType   = "IwfWorkflowType"

	BackendTypeCadence  BackendType = "cadence"
	BackendTypeTemporal BackendType = "temporal"

	IwfSystemSignalPrefix         = "__IwfSystem_"
	SkipTimerSignalChannelName    = IwfSystemSignalPrefix + "SkipTimerChannel"
	FailWorkflowSignalChannelName = IwfSystemSignalPrefix + "FailWorkflowChannel"
	UpdateConfigSignalChannelName = IwfSystemSignalPrefix + "UpdateWorkflowConfig"
	ExecuteRpcSignalChannelName   = IwfSystemSignalPrefix + "ExecuteRpc"
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
)

var ValidClosingWorkflowStateId = map[string]bool{
	GracefulCompletingWorkflowStateId: true,
	ForceCompletingWorkflowStateId:    true,
	ForceFailingWorkflowStateId:       true,
}
