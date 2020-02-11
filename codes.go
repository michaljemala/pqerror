// pqerror package exposes PostgreSQL error codes as constants and provides
// helpers to inspect PostgreSQL errors.
package pqerror

import "github.com/lib/pq"

// See https://www.postgresql.org/docs/11/static/errcodes-appendix.html
// and https://github.com/postgres/postgres/blob/REL_12_STABLE/src/backend/utils/errcodes.txt.
const PqVersion = 12

// IsClass reports whether an error is of type *pq.Error and has a given class.
func IsClass(err error, class pq.ErrorClass) bool {
	pqerr, ok := err.(*pq.Error)
	return ok && pqerr.Code.Class() == class
}

// IsCode reports whether an error is of type *pq.Error and has a given code.
func IsCode(err error, code pq.ErrorCode) bool {
	pqerr, ok := err.(*pq.Error)
	return ok && pqerr.Code == code
}

const (
	ClassSuccessfulCompletion                    = pq.ErrorClass("00")
	ClassWarning                                 = pq.ErrorClass("01")
	ClassNoData                                  = pq.ErrorClass("02")
	ClassSQLStatementNotYetComplete              = pq.ErrorClass("03")
	ClassConnectionException                     = pq.ErrorClass("08")
	ClassTriggeredActionException                = pq.ErrorClass("09")
	ClassFeatureNotSupported                     = pq.ErrorClass("0A")
	ClassInvalidTransactionInitiation            = pq.ErrorClass("0B")
	ClassLocatorException                        = pq.ErrorClass("0F")
	ClassInvalidGrantor                          = pq.ErrorClass("0L")
	ClassInvalidRoleSpecification                = pq.ErrorClass("0P")
	ClassDiagnosticsException                    = pq.ErrorClass("0Z")
	ClassCaseNotFound                            = pq.ErrorClass("20")
	ClassCardinalityViolation                    = pq.ErrorClass("21")
	ClassDataException                           = pq.ErrorClass("22")
	ClassIntegrityConstraintViolation            = pq.ErrorClass("23")
	ClassInvalidCursorState                      = pq.ErrorClass("24")
	ClassInvalidTransactionState                 = pq.ErrorClass("25")
	ClassInvalidSQLStatementName                 = pq.ErrorClass("26")
	ClassTriggeredDataChangeViolation            = pq.ErrorClass("27")
	ClassInvalidAuthorizationSpecification       = pq.ErrorClass("28")
	ClassDependentPrivilegeDescriptorsStillExist = pq.ErrorClass("2B")
	ClassInvalidTransactionTermination           = pq.ErrorClass("2D")
	ClassSQLRoutineException                     = pq.ErrorClass("2F")
	ClassInvalidCursorName                       = pq.ErrorClass("34")
	ClassExternalRoutineException                = pq.ErrorClass("38")
	ClassExternalRoutineInvocationException      = pq.ErrorClass("39")
	ClassSavepointException                      = pq.ErrorClass("3B")
	ClassInvalidCatalogName                      = pq.ErrorClass("3D")
	ClassInvalidSchemaName                       = pq.ErrorClass("3F")
	ClassTransactionRollback                     = pq.ErrorClass("40")
	ClassSyntaxErrorOrAccessRuleViolation        = pq.ErrorClass("42")
	ClassWithCheckOptionViolation                = pq.ErrorClass("44")
	ClassInsufficientResources                   = pq.ErrorClass("53")
	ClassProgramLimitExceeded                    = pq.ErrorClass("54")
	ClassObjectNotInPrerequisiteState            = pq.ErrorClass("55")
	ClassOperatorIntervention                    = pq.ErrorClass("57")
	ClassSystemError                             = pq.ErrorClass("58")
	ClassSnapshotTooOld                          = pq.ErrorClass("72")
	ClassConfigFileError                         = pq.ErrorClass("F0")
	ClassFdwError                                = pq.ErrorClass("HV")
	ClassPlpgsqlError                            = pq.ErrorClass("P0")
	ClassInternalError                           = pq.ErrorClass("XX")
)

const (
	// Class 00 - Successful Completion
	SuccessfulCompletion = pq.ErrorCode("00000")

	// Class 01 - Warning
	Warning                                 = pq.ErrorCode("01000")
	WarningDynamicResultSetsReturned        = pq.ErrorCode("0100C")
	WarningImplicitZeroBitPadding           = pq.ErrorCode("01008")
	WarningNullValueEliminatedInSetFunction = pq.ErrorCode("01003")
	WarningPrivilegeNotGranted              = pq.ErrorCode("01007")
	WarningPrivilegeNotRevoked              = pq.ErrorCode("01006")
	WarningStringDataRightTruncation        = pq.ErrorCode("01004")
	WarningDeprecatedFeature                = pq.ErrorCode("01P01")

	// Class 02 - No Data (this is also a warning class per the SQL standard)
	NoData                                = pq.ErrorCode("02000")
	NoAdditionalDynamicResultSetsReturned = pq.ErrorCode("02001")

	// Class 03 - SQL Statement Not Yet Complete
	SQLStatementNotYetComplete = pq.ErrorCode("03000")

	// Class 08 - Connection Exception
	ConnectionException                           = pq.ErrorCode("08000")
	ConnectionDoesNotExist                        = pq.ErrorCode("08003")
	ConnectionFailure                             = pq.ErrorCode("08006")
	SQLClientUnableToEstablishSQLConnection       = pq.ErrorCode("08001")
	SQLServerRejectedEstablishmentOfSQLConnection = pq.ErrorCode("08004")
	TransactionResolutionUnknown                  = pq.ErrorCode("08007")
	ProtocolViolation                             = pq.ErrorCode("08P01")

	// Class 09 - Triggered Action Exception
	TriggeredActionException = pq.ErrorCode("09000")

	// Class 0A - Feature Not Supported
	FeatureNotSupported = pq.ErrorCode("0A000")

	// Class 0B - Invalid Transaction Initiation
	InvalidTransactionInitiation = pq.ErrorCode("0B000")

	// Class 0F - Locator Exception
	LocatorException            = pq.ErrorCode("0F000")
	InvalidLocatorSpecification = pq.ErrorCode("0F001")

	// Class 0L - Invalid Grantor
	InvalidGrantor        = pq.ErrorCode("0L000")
	InvalidGrantOperation = pq.ErrorCode("0LP01")

	// Class 0P - Invalid Role Specification
	InvalidRoleSpecification = pq.ErrorCode("0P000")

	// Class 0Z - Diagnostics Exception
	DiagnosticsException                           = pq.ErrorCode("0Z000")
	StackedDiagnosticsAccessedWithoutActiveHandler = pq.ErrorCode("0Z002")

	// Class 20 - Case Not Found
	CaseNotFound = pq.ErrorCode("20000")

	// Class 21 - Cardinality Violation
	CardinalityViolation = pq.ErrorCode("21000")

	// Class 22 - Data Exception
	DataException                         = pq.ErrorCode("22000")
	ArraySubscriptError                   = pq.ErrorCode("2202E")
	ArrayElementError                     = pq.ErrorCode("2202E")
	CharacterNotInRepertoire              = pq.ErrorCode("22021")
	DatetimeFieldOverflow                 = pq.ErrorCode("22008")
	DatetimeValueOutOfRange               = pq.ErrorCode("22008")
	DivisionByZero                        = pq.ErrorCode("22012")
	ErrorInAssignment                     = pq.ErrorCode("22005")
	EscapeCharacterConflict               = pq.ErrorCode("2200B")
	IndicatorOverflow                     = pq.ErrorCode("22022")
	IntervalFieldOverflow                 = pq.ErrorCode("22015")
	InvalidArgumentForLogarithm           = pq.ErrorCode("2201E")
	InvalidArgumentForNtileFunction       = pq.ErrorCode("22014")
	InvalidArgumentForNthValueFunction    = pq.ErrorCode("22016")
	InvalidArgumentForPowerFunction       = pq.ErrorCode("2201F")
	InvalidArgumentForWidthBucketFunction = pq.ErrorCode("2201G")
	InvalidCharacterValueForCast          = pq.ErrorCode("22018")
	InvalidDatetimeFormat                 = pq.ErrorCode("22007")
	InvalidEscapeCharacter                = pq.ErrorCode("22019")
	InvalidEscapeOctet                    = pq.ErrorCode("2200D")
	InvalidEscapeSequence                 = pq.ErrorCode("22025")
	NonstandardUseOfEscapeCharacter       = pq.ErrorCode("22P06")
	InvalidIndicatorParameterValue        = pq.ErrorCode("22010")
	InvalidParameterValue                 = pq.ErrorCode("22023")
	InvalidRegularExpression              = pq.ErrorCode("2201B")
	InvalidRowCountInLimitClause          = pq.ErrorCode("2201W")
	InvalidRowCountInResultOffsetClause   = pq.ErrorCode("2201X")
	InvalidTablesampleArgument            = pq.ErrorCode("2202H")
	InvalidTablesampleRepeat              = pq.ErrorCode("2202G")
	InvalidTimeZoneDisplacementValue      = pq.ErrorCode("22009")
	InvalidUseOfEscapeCharacter           = pq.ErrorCode("2200C")
	MostSpecificTypeMismatch              = pq.ErrorCode("2200G")
	NullValueNotAllowed                   = pq.ErrorCode("22004")
	NullValueNoIndicatorParameter         = pq.ErrorCode("22002")
	NumericValueOutOfRange                = pq.ErrorCode("22003")
	StringDataLengthMismatch              = pq.ErrorCode("22026")
	StringDataRightTruncation             = pq.ErrorCode("22001")
	SubstringError                        = pq.ErrorCode("22011")
	TrimError                             = pq.ErrorCode("22027")
	UnterminatedCString                   = pq.ErrorCode("22024")
	ZeroLengthCharacterString             = pq.ErrorCode("2200F")
	FloatingPointException                = pq.ErrorCode("22P01")
	InvalidTextRepresentation             = pq.ErrorCode("22P02")
	InvalidBinaryRepresentation           = pq.ErrorCode("22P03")
	BadCopyFileFormat                     = pq.ErrorCode("22P04")
	UntranslatableCharacter               = pq.ErrorCode("22P05")
	NotAnXmlDocument                      = pq.ErrorCode("2200L")
	InvalidXmlDocument                    = pq.ErrorCode("2200M")
	InvalidXmlContent                     = pq.ErrorCode("2200N")
	InvalidXmlComment                     = pq.ErrorCode("2200S")
	InvalidXmlProcessingInstruction       = pq.ErrorCode("2200T")
	DuplicateJsonObjectKeyValue           = pq.ErrorCode("22030")
	InvalidJsonText                       = pq.ErrorCode("22032")
	InvalidJsonSubscript                  = pq.ErrorCode("22033")
	MoreThanOneJsonItem                   = pq.ErrorCode("22034")
	NoJsonItem                            = pq.ErrorCode("22035")
	NonNumericJsonItem                    = pq.ErrorCode("22036")
	NonUniqueKeysInJsonObject             = pq.ErrorCode("22037")
	SingletonJsonItemRequired             = pq.ErrorCode("22038")
	JsonArrayNotFound                     = pq.ErrorCode("22039")
	JsonMemberNotFound                    = pq.ErrorCode("2203A")
	JsonNumberNotFound                    = pq.ErrorCode("2203B")
	JsonObjectNotFound                    = pq.ErrorCode("2203C")
	JsonScalarRequired                    = pq.ErrorCode("2203F")
	TooManyJsonArrayElements              = pq.ErrorCode("2203D")
	TooManyJsonObjectMembers              = pq.ErrorCode("2203E")

	// Class 23 - Integrity Constraint Violation
	IntegrityConstraintViolation = pq.ErrorCode("23000")
	RestrictViolation            = pq.ErrorCode("23001")
	NotNullViolation             = pq.ErrorCode("23502")
	ForeignKeyViolation          = pq.ErrorCode("23503")
	UniqueViolation              = pq.ErrorCode("23505")
	CheckViolation               = pq.ErrorCode("23514")
	ExclusionViolation           = pq.ErrorCode("23P01")

	// Class 24 - Invalid Cursor State
	InvalidCursorState = pq.ErrorCode("24000")

	// Class 25 - Invalid Transaction State
	InvalidTransactionState                         = pq.ErrorCode("25000")
	ActiveSQLTransaction                            = pq.ErrorCode("25001")
	BranchTransactionAlreadyActive                  = pq.ErrorCode("25002")
	HeldCursorRequiresSameIsolationLevel            = pq.ErrorCode("25008")
	InappropriateAccessModeForBranchTransaction     = pq.ErrorCode("25003")
	InappropriateIsolationLevelForBranchTransaction = pq.ErrorCode("25004")
	NoActiveSQLTransactionForBranchTransaction      = pq.ErrorCode("25005")
	ReadOnlySQLTransaction                          = pq.ErrorCode("25006")
	SchemaAndDataStatementMixingNotSupported        = pq.ErrorCode("25007")
	NoActiveSQLTransaction                          = pq.ErrorCode("25P01")
	InFailedSQLTransaction                          = pq.ErrorCode("25P02")
	IdleInTransactionSessionTimeout                 = pq.ErrorCode("25P03")

	// Class 26 - Invalid SQL Statement Name
	InvalidSQLStatementName = pq.ErrorCode("26000")

	// Class 27 - Triggered Data Change Violation
	TriggeredDataChangeViolation = pq.ErrorCode("27000")

	// Class 28 - Invalid Authorization Specification
	InvalidAuthorizationSpecification = pq.ErrorCode("28000")
	InvalidPassword                   = pq.ErrorCode("28P01")

	// Class 2B - Dependent Privilege Descriptors Still Exist
	DependentPrivilegeDescriptorsStillExist = pq.ErrorCode("2B000")
	DependentObjectsStillExist              = pq.ErrorCode("2BP01")

	// Class 2D - Invalid Transaction Termination
	InvalidTransactionTermination = pq.ErrorCode("2D000")

	// Class 2F - SQL Routine Exception
	SQLRoutineException               = pq.ErrorCode("2F000")
	FunctionExecutedNoReturnStatement = pq.ErrorCode("2F005")
	ModifyingSQLDataNotPermitted      = pq.ErrorCode("2F002")
	ProhibitedSQLStatementAttempted   = pq.ErrorCode("2F003")
	ReadingSQLDataNotPermitted        = pq.ErrorCode("2F004")

	// Class 34 - Invalid Cursor Name
	InvalidCursorName = pq.ErrorCode("34000")

	// Class 38 - External Routine Exception
	ExternalRoutineException                                 = pq.ErrorCode("38000")
	ExternalRoutineException_ContainingSQLNotPermitted       = pq.ErrorCode("38001")
	ExternalRoutineException_ModifyingSQLDataNotPermitted    = pq.ErrorCode("38002")
	ExternalRoutineException_ProhibitedSQLStatementAttempted = pq.ErrorCode("38003")
	ExternalRoutineException_ReadingSQLDataNotPermitted      = pq.ErrorCode("38004")

	// Class 39 - External Routine Invocation Exception
	ExternalRoutineInvocationException                              = pq.ErrorCode("39000")
	ExternalRoutineInvocationException_InvalidSQLstateReturned      = pq.ErrorCode("39001")
	ExternalRoutineInvocationException_NullValueNotAllowed          = pq.ErrorCode("39004")
	ExternalRoutineInvocationException_TriggerProtocolViolated      = pq.ErrorCode("39P01")
	ExternalRoutineInvocationException_SrfProtocolViolated          = pq.ErrorCode("39P02")
	ExternalRoutineInvocationException_EventTriggerProtocolViolated = pq.ErrorCode("39P03")

	// Class 3B - Savepoint Exception
	SavepointException            = pq.ErrorCode("3B000")
	InvalidSavepointSpecification = pq.ErrorCode("3B001")

	// Class 3D - Invalid Catalog Name
	InvalidCatalogName = pq.ErrorCode("3D000")

	// Class 3F - Invalid Schema Name
	InvalidSchemaName = pq.ErrorCode("3F000")

	// Class 40 - Transaction Rollback
	TransactionRollback                     = pq.ErrorCode("40000")
	TransactionIntegrityConstraintViolation = pq.ErrorCode("40002")
	SerializationFailure                    = pq.ErrorCode("40001")
	StatementCompletionUnknown              = pq.ErrorCode("40003")
	DeadlockDetected                        = pq.ErrorCode("40P01")

	// Class 42 - Syntax Error or Access Rule Violation
	SyntaxErrorOrAccessRuleViolation   = pq.ErrorCode("42000")
	SyntaxError                        = pq.ErrorCode("42601")
	InsufficientPrivilege              = pq.ErrorCode("42501")
	CannotCoerce                       = pq.ErrorCode("42846")
	GroupingError                      = pq.ErrorCode("42803")
	WindowingError                     = pq.ErrorCode("42P20")
	InvalidRecursion                   = pq.ErrorCode("42P19")
	InvalidForeignKey                  = pq.ErrorCode("42830")
	InvalidName                        = pq.ErrorCode("42602")
	NameTooLong                        = pq.ErrorCode("42622")
	ReservedName                       = pq.ErrorCode("42939")
	DatatypeMismatch                   = pq.ErrorCode("42804")
	IndeterminateDatatype              = pq.ErrorCode("42P18")
	CollationMismatch                  = pq.ErrorCode("42P21")
	IndeterminateCollation             = pq.ErrorCode("42P22")
	WrongObjectType                    = pq.ErrorCode("42809")
	UndefinedColumn                    = pq.ErrorCode("42703")
	UndefinedFunction                  = pq.ErrorCode("42883")
	UndefinedTable                     = pq.ErrorCode("42P01")
	UndefinedParameter                 = pq.ErrorCode("42P02")
	UndefinedObject                    = pq.ErrorCode("42704")
	DuplicateColumn                    = pq.ErrorCode("42701")
	DuplicateCursor                    = pq.ErrorCode("42P03")
	DuplicateDatabase                  = pq.ErrorCode("42P04")
	DuplicateFunction                  = pq.ErrorCode("42723")
	DuplicatePreparedStatement         = pq.ErrorCode("42P05")
	DuplicateSchema                    = pq.ErrorCode("42P06")
	DuplicateTable                     = pq.ErrorCode("42P07")
	DuplicateAlias                     = pq.ErrorCode("42712")
	DuplicateObject                    = pq.ErrorCode("42710")
	AmbiguousColumn                    = pq.ErrorCode("42702")
	AmbiguousFunction                  = pq.ErrorCode("42725")
	AmbiguousParameter                 = pq.ErrorCode("42P08")
	AmbiguousAlias                     = pq.ErrorCode("42P09")
	InvalidColumnReference             = pq.ErrorCode("42P10")
	InvalidColumnDefinition            = pq.ErrorCode("42611")
	InvalidCursorDefinition            = pq.ErrorCode("42P11")
	InvalidDatabaseDefinition          = pq.ErrorCode("42P12")
	InvalidFunctionDefinition          = pq.ErrorCode("42P13")
	InvalidPreparedStatementDefinition = pq.ErrorCode("42P14")
	InvalidSchemaDefinition            = pq.ErrorCode("42P15")
	InvalidTableDefinition             = pq.ErrorCode("42P16")
	InvalidObjectDefinition            = pq.ErrorCode("42P17")

	// Class 44 - WITH CHECK OPTION Violation
	WithCheckOptionViolation = pq.ErrorCode("44000")

	// Class 53 - Insufficient Resources
	InsufficientResources      = pq.ErrorCode("53000")
	DiskFull                   = pq.ErrorCode("53100")
	OutOfMemory                = pq.ErrorCode("53200")
	TooManyConnections         = pq.ErrorCode("53300")
	ConfigurationLimitExceeded = pq.ErrorCode("53400")

	// Class 54 - Program Limit Exceeded
	ProgramLimitExceeded = pq.ErrorCode("54000")
	StatementTooComplex  = pq.ErrorCode("54001")
	TooManyColumns       = pq.ErrorCode("54011")
	TooManyArguments     = pq.ErrorCode("54023")

	// Class 55 - Object Not In Prerequisite State
	ObjectNotInPrerequisiteState = pq.ErrorCode("55000")
	ObjectInUse                  = pq.ErrorCode("55006")
	CantChangeRuntimeParam       = pq.ErrorCode("55P02")
	LockNotAvailable             = pq.ErrorCode("55P03")
	UnsafeNewEnumValueUsage      = pq.ErrorCode("55P04")

	// Class 57 - Operator Intervention
	OperatorIntervention = pq.ErrorCode("57000")
	QueryCanceled        = pq.ErrorCode("57014")
	AdminShutdown        = pq.ErrorCode("57P01")
	CrashShutdown        = pq.ErrorCode("57P02")
	CannotConnectNow     = pq.ErrorCode("57P03")
	DatabaseDropped      = pq.ErrorCode("57P04")

	// Class 58 - System Error (errors external to PostgreSQL itself)
	SystemError   = pq.ErrorCode("58000")
	IoError       = pq.ErrorCode("58030")
	UndefinedFile = pq.ErrorCode("58P01")
	DuplicateFile = pq.ErrorCode("58P02")

	// Class 72 — Snapshot Failure
	SnapshotTooOld = pq.ErrorCode("72000")

	// Class F0 - Configuration File Error
	ConfigFileError = pq.ErrorCode("F0000")
	LockFileExists  = pq.ErrorCode("F0001")

	// Class HV - Foreign Data Wrapper Error (SQL/MED)
	FdwError                             = pq.ErrorCode("HV000")
	FdwColumnNameNotFound                = pq.ErrorCode("HV005")
	FdwDynamicParameterValueNeeded       = pq.ErrorCode("HV002")
	FdwFunctionSequenceError             = pq.ErrorCode("HV010")
	FdwInconsistentDescriptorInformation = pq.ErrorCode("HV021")
	FdwInvalidAttributeValue             = pq.ErrorCode("HV024")
	FdwInvalidColumnName                 = pq.ErrorCode("HV007")
	FdwInvalidColumnNumber               = pq.ErrorCode("HV008")
	FdwInvalidDataType                   = pq.ErrorCode("HV004")
	FdwInvalidDataTypeDescriptors        = pq.ErrorCode("HV006")
	FdwInvalidDescriptorFieldIdentifier  = pq.ErrorCode("HV091")
	FdwInvalidHandle                     = pq.ErrorCode("HV00B")
	FdwInvalidOptionIndex                = pq.ErrorCode("HV00C")
	FdwInvalidOptionName                 = pq.ErrorCode("HV00D")
	FdwInvalidStringLengthOrBufferLength = pq.ErrorCode("HV090")
	FdwInvalidStringFormat               = pq.ErrorCode("HV00A")
	FdwInvalidUseOfNullPointer           = pq.ErrorCode("HV009")
	FdwTooManyHandles                    = pq.ErrorCode("HV014")
	FdwOutOfMemory                       = pq.ErrorCode("HV001")
	FdwNoSchemas                         = pq.ErrorCode("HV00P")
	FdwOptionNameNotFound                = pq.ErrorCode("HV00J")
	FdwReplyHandle                       = pq.ErrorCode("HV00K")
	FdwSchemaNotFound                    = pq.ErrorCode("HV00Q")
	FdwTableNotFound                     = pq.ErrorCode("HV00R")
	FdwUnableToCreateExecution           = pq.ErrorCode("HV00L")
	FdwUnableToCreateReply               = pq.ErrorCode("HV00M")
	FdwUnableToEstablishConnection       = pq.ErrorCode("HV00N")

	// Class P0 - PL/pgSQL Error
	PLpgSQLError   = pq.ErrorCode("P0000")
	RaiseException = pq.ErrorCode("P0001")
	NoDataFound    = pq.ErrorCode("P0002")
	TooManyRows    = pq.ErrorCode("P0003")
	AssertFailure  = pq.ErrorCode("P0004")

	// Class XX - Internal Error
	InternalError  = pq.ErrorCode("XX000")
	DataCorrupted  = pq.ErrorCode("XX001")
	IndexCorrupted = pq.ErrorCode("XX002")
)
