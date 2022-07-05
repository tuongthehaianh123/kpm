package control

const MAX_SUBSCRIPTION_ATTEMPTS = 100

type DecodedIndicationMessage struct {
	RequestID             int32
	RequestSequenceNumber int32
	FuncID                int32
	ActionID              int32
	IndSN                 int32
	IndType               int32
	IndHeader             []byte
	IndHeaderLength       int32
	IndMessage            []byte
	IndMessageLength      int32
	CallProcessID         []byte
	CallProcessIDLength   int32
}

type DecodedSubscriptionResponseMessage struct {
	RequestID             int32
	RequestSequenceNumber int32
	FuncID                int32
	ActionAdmittedList    ActionAdmittedListType
	ActionNotAdmittedList ActionNotAdmittedListType
}

type IndicationHeaderFormat1 struct {
	colletStartTime *OctetString
	fileFormatversion	*OctetString
	senderName	*OctetString
	senderType	*OctetString
	vendorName	*OctetString
}

type IndicationHeader struct {
	IndHdrType int32
	IndHdr     interface{}
}

type IndicationMessageFormat2 struct {
	measData MeasurementData
	measCondUEidList MeasurementCondUEidListType
	granulPeriod *int32
}
type MeasurementData struct {
	measDataItem     [8]MeasurementDataItem
	measDataItemCount int
}

type MeasurementDataItem struct {
	measRecord MeasurementRecord
	incompleteFlag *int64
} 

type MeasurementRecord struct {
	measRecordItem     [8]MeasurementRecordItem
	measRecordItemCount int
}

type MeasurementRecordItem struct {
	integer uint32
	real int64
	noValue	int
} 

type MeasurementCondUEidListType struct {
	measCondUEidList     [8]MeasurementCondUEidItem
	measCondUEidListCount int
}

type MeasurementCondUEidItem struct {
	measType MeasurementType
	matchingCond MatchingCondListType
	matchingUEidList *MatchingUEidListType
}
type MatchingUEidListType struct {
	matchingUEidItem     [8]MatchingUEidItem
	matchingUEidItemCount int
}
type MatchingUEidItem struct {
	ueID UEID
}
type UEID struct {
	gNBUEID *UEIDGNB 
	gNBDUUEID *UEIDGNBDU
	gNBCUUPUEID *UEIDGNBCUUP
	ngeNBUEID *UEIDNGENB 
	ngeNBDUUEID *UEIDNGENBDU
	engNBUEID *UEIDENGNB
	eNBUEID *UEIDENB
}

type UEIDENB struct {
	mMEUES1APID int32
	gUMMEI GUMMEI
	meNBUEX2APID *int32
	meNBUEX2APIDExtension *int32
	globalENBID *GlobalENBID	
}

type GlobalENBID struct {
	pLMNIdentity OctetString
	eNBID ENBID
}
 type GUMMEI struct {
	pLMNIdentity OctetString
	mMEGroupID OctetString
	mMECode OctetString
}

type UEIDENGNB struct {
	meNBUEX2APID int32
	meNBUEX2APIDExtension *int32
	globalENBID GlobalENBID
	gNBCUUEF1APID *int32	
	gNBCUCPUEE1APIDList *UEIDGNBCUCPE1APIDListType	
	ranUEID	*OctetString
}

type GLonalENBID struct {
	pLMNIdentity OctetString
	eNBID ENBID
}

type ENBID struct {
	macroeNBID BitString
	homeeNBID BitString
	shortMacroeNBID BitString
	longMacroeNBID BitString
}
type UEIDGNBDU struct {
	gNBCUUEF1APID int32
	ranUEID *OctetString
}

type UEIDGNBCUUP struct {
	gNBCUCPUEE1APID int32
	ran_UEID *OctetString
}

type UEIDNGENBDU struct {
	ngeNBCUUEW1APID int32
}

type UEIDNGENB struct {
	amfUENGAPID int
	guami GUAMI
	ngeNBCUUEW1APID	*int32
	mNGRANUEXnAPID *int32	
	globalNgENBID *GlobalNgENBID	
	globalNGRANNodeID *GlobalNGRANNodeID
}
type GlobalNgENBID struct {
	pLMNIdentity OctetString
	ngENBID NgENBID
}

type NgENBID struct {
	macroNgENBID BitString
	shortMacroNgENBID BitString
	longMacroNgENBID BitString
} 
type UEIDGNB struct {
	amfUENGAPID int
	guami GUAMI
	gNBCUUEF1APIDList	*UEIDGNBCUF1APIDListType
	gNBCUCPUEE1APIDList	*UEIDGNBCUCPE1APIDListType
	ranUEID OctetString
	mNGRANUEXnAPID	int32
	globalGNBID *GlobalGNBID	
	globalNGRANNodeID	*GlobalNGRANNodeID
}
type GUAMI struct {
	pLMNIdentity OctetString
	aMFRegionID BitString
	aMFSetID BitString
	aMFPointer BitString
}

type UEIDGNBCUF1APIDListType struct {
	ueIDGNBCUF1APIDItem     [8]UEIDGNBCUF1APIDItem
	ueIDGNBCUF1APIDItemCount int
}

type UEIDGNBCUF1APIDItem struct {
	gNBCUUEF1APID int32
}

type UEIDGNBCUCPE1APIDListType struct {
	ueIDGNBCUCPE1APIDItem    [8]UEIDGNBCUCPE1APIDItem
	ueIDGNBCUCPE1APIDItemCount int
}

type UEIDGNBCUCPE1APIDItem struct {
	gNBCUCPUEE1APID int32
}


type GlobalNGRANNodeID struct {
	gNB *GlobalGNBID
	ngeNB *GlobalNgENBID
}

type MeasurementType struct {
	measName PrintableString
	measID	int32
}	
type GlobalGNBID struct {
	pLMNIdentity OctetString
	gNBID BitString
}

type MatchingCondListType struct {
	matchingCondList   [8]MatchingCondItem
	matchCondListCount int
}
type MatchingCondItem struct {
	measLabel *MeasurementLabel
	testCondInfo *TestCondInfo
}
type  MeasurementLabel struct {
	noLabel *int32	
	plmnID	*OctetString
	sliceID *S_NSSAI	
	fiveQI	*int32
	qFI 	*int32
	qCI 	*int32	
	qCImax	*int32
	qCImin	*int32
	aRPmax	*int32
	aRPmin	*int32
	bitrateRange *int32	
	layerMU_MIMO	*int32
	sUM	*int32
	distBinX	*int32
	distBinY	*int32
	distBinZ	*int32
	preLabelOverride *int32	
	startEndInd	*int32	
	min	*int32
	max	*int32
	avg	*int32
}

type S_NSSAI struct {
	sST OctetString
	sD *OctetString
}
type TestCondInfo struct {
	testType TestCondType
	testExpr int32
	testValue TestCond_Value
}
type TestCondType struct {
	gBR int32;
	aMBR int32
	isStat int32
	isCatM int32
	rSRP int32
	rSRQ int32
}

type TestCond_Value struct {
	valueInt int32
	valueEnum int32
	valueBool int
	valueBitS BitString
	valueOctS OctetString
	valuePrtS PrintableString
}

type IndicationMessage struct {
	StyleType  int64
	IndMsgType int32
	IndMsg     interface{}
}

// type CellMetricsEntry struct {
// 	MeasTimestampPDCPBytes Timestamp `json:"MeasTimestampPDCPBytes"`
// 	CellID 		       string 	 `json:"CellID"`
// 	PDCPBytesDL            int64     `json:"PDCPBytesDL"`
// 	PDCPBytesUL            int64     `json:"PDCPBytesUL"`
// 	MeasTimestampPRB       Timestamp `json:"MeasTimestampAvailPRB"`
// 	AvailPRBDL             int64     `json:"AvailPRBDL"`
// 	AvailPRBUL             int64     `json:"AvailPRBUL"`
// 	MeasPeriodPDCP	       int64	 `json:"MeasPeriodPDCPBytes"`
// 	MeasPeriodPRB	       int64	 `json:"MeasPeriodAvailPRB"`
// }

// type UeMetricsEntry struct {
// 	UeID                   int64     `json:"UEID"`
// 	ServingCellID          string    `json:"ServingCellID"`
// 	MeasTimestampPDCPBytes Timestamp `json:"MeasTimestampUEPDCPBytes"`
// 	PDCPBytesDL            int64     `json:"UEPDCPBytesDL"`
// 	PDCPBytesUL            int64     `json:"UEPDCPBytesUL"`
// 	MeasTimestampPRB       Timestamp `json:"MeasTimestampUEPRBUsage"`
// 	PRBUsageDL             int64     `json:"UEPRBUsageDL"`
// 	PRBUsageUL             int64     `json:"UEPRBUsageUL"`
// 	MeasTimeRF             Timestamp `json:"MeasTimestampRF"`
// 	MeasPeriodRF	       int64	 `json:"MeasPeriodRF"`
// 	MeasPeriodPDCP	       int64	 `json:"MeasPeriodUEPDCPBytes"`
// 	MeasPeriodPRB	       int64	 `json:"MeasPeriodUEPRBUsage"`
// 	ServingCellRF   CellRFType           `json:"ServingCellRF"`
// 	NeighborCellsRF []NeighborCellRFType `json:"NeighborCellRF"`
// }

type Timestamp struct {
	TVsec  int64 `json:"tv_sec"`
	TVnsec int64 `json:"tv_nsec"`
}
 
type OctetString struct {
	Buf []byte
	Size int
}
		
type PrintableString OctetString
	
type BitString struct {
	Buf []byte
	Size int
	BitsUnused int
}

