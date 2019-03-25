package telemetry

// LocalQuery contains the serviceID for a local query
const LocalQuery string = "weaviate.local.query"

// LocalQueryMeta contains the serviceID for a local meta query
const LocalQueryMeta string = "weaviate.local.query.meta"

// NetworkQuery contains the serviceID for a network query
const NetworkQuery string = "Weaviate.network.query"

// NetworkQueryMeta contains the serviceID for a network meta query
const NetworkQueryMeta string = "weaviate.network.query.meta"

// LocalAdd contains the serviceID for a local add query
const LocalAdd string = "weaviate.local.add"

// LocalAddMeta contains the serviceID for a local add meta query
const LocalAddMeta string = "weaviate.local.add.meta"

// LocalManipulate contains the serviceID for a local manipulate query
const LocalManipulate string = "weaviate.local.manipulate"

// LocalManipulateMeta contains the serviceID for a local manipulate meta query
const LocalManipulateMeta string = "weaviate.local.manipulate.meta"

// LocalToolsMap contains the serviceID for a local tools query
const LocalToolsMap string = "weaviate.local.tools.map"

// NetworkToolsMap contains the serviceID for a network tools query
const NetworkToolsMap string = "weaviate.network.tools.map"

// TypeGQL contains the serviceID for a local query
const TypeGQL string = "GQL"

// TypeREST contains the serviceID for a local query
const TypeREST string = "REST"

// ReportPostFail contains the value used when a report is unable to be sent to the specified URL
const ReportPostFail string = "[POST failed]"

// ReportCBORFail contains the value used when a report fails to be converted to CBOR
const ReportCBORFail string = "[CBOR conversion failed]"
