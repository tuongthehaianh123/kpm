/*
==================================================================================
  Copyright (c) 2019 AT&T Intellectual Property.
  Copyright (c) 2019 Nokia

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
==================================================================================
*/

package control

/*
#include <e2sm/wrapper.h>
#cgo LDFLAGS: -le2smwrapper
#cgo CFLAGS: -I/usr/local/include/e2sm
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"unsafe"
)

type E2sm struct {
}

func (c *E2sm) SetEventTriggerDefinition(buffer []byte, eventTriggerCount int, RTPeriods []int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	periods := unsafe.Pointer(&RTPeriods[0])
	size := C.e2sm_encode_ric_event_trigger_definition(cptr, C.size_t(len(buffer)), C.size_t(eventTriggerCount), (*C.long)(periods))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2sm) SetActionDefinition(buffer []byte, ricStyleType int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	size := C.e2sm_encode_ric_action_definition(cptr, C.size_t(len(buffer)), C.long(ricStyleType))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set ActionDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2sm) GetIndicationHeader(buffer []byte) (indHdr *IndicationHeader, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	indHdr = &IndicationHeader{}
	decodedHdr := C.e2sm_decode_ric_indication_header(cptr, C.size_t(len(buffer)))
	if decodedHdr == nil {
		return indHdr, errors.New("e2sm wrapper is unable to get IndicationHeader due to wrong or invalid input")
	}
	defer C.e2sm_free_ric_indication_header(decodedHdr)

	indHdr.IndHdrType = int32(decodedHdr.present)
	if indHdr.IndHdrType == 1 {
		indHdrFormat1 := &IndicationHeaderFormat1{}
		indHdrFormat1_C := *(**C.E2SM_KPM_IndicationHeader_Format1_t)(unsafe.Pointer(&decodedHdr.choice[0]))
		
		colletStartTime_C := (*C.TimeStamp_t)(indHdrFormat1_C.colletStartTime)
		colletStartTime_C := indHdrFormat1_C.colletStartTime
		indHdrFormat1.colletStartTime.Buf = C.GoBytes(unsafe.Pointer(colletStartTime_C.buf), C.int(colletStartTime_C.size))
		indHdrFormat1.colletStartTime.Size = int(colletStartTime_C.size)

		fileFormatversion_C := (*C.PrintableString_t)(indHdrFormat1_C.fileFormatversion)
		indHdrFormat1.fileFormatversion.Buf = C.GoBytes(unsafe.Pointer(fileFormatversion_C.buf), C.int(fileFormatversion_C.size))
		indHdrFormat1.fileFormatversion.Size = int(fileFormatversion_C.size)

		senderType_C := (*C.PrintableString_t)(indHdrFormat1_C.senderType)
		indHdrFormat1.senderType.Buf = C.GoBytes(unsafe.Pointer(senderType_C.buf), C.int(senderType_C.size))
		indHdrFormat1.senderType.Size = int(senderType_C.size)

		senderType_C := (*C.PrintableString_t)(indHdrFormat1_C.senderType)
		indHdrFormat1.senderType.Buf = C.GoBytes(unsafe.Pointer(senderType_C.buf), C.int(senderType_C.size))
		indHdrFormat1.senderType.Size = int(senderType_C.size)

		vendorName_C := (*C.PrintableString_t)(indHdrFormat1_C.vendorName)
		indHdrFormat1.vendorName.Buf = C.GoBytes(unsafe.Pointer(vendorName_C.buf), C.int(vendorName_C.size))
		indHdrFormat1.vendorName.Size = int(vendorName_C.size)		

		indHdr.IndHdr = indHdrFormat1
	} else {
		return indHdr, errors.New("Unknown RIC Indication Header type")
	}

	return
}

func (c *E2sm) GetIndicationMessage(buffer []byte) (indMsg *IndicationMessage, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	indMsg = &IndicationMessage{}
	decodedMsg := C.e2sm_decode_ric_indication_message(cptr, C.size_t(len(buffer)))
	if decodedMsg == nil {
		return indMsg, errors.New("e2sm wrapper is unable to get IndicationMessage due to wrong or invalid input")
	}
	defer C.e2sm_free_ric_indication_message(decodedMsg)

	indMsg.StyleType = int64(decodedMsg.ric_Style_Type)

	indMsg.IndMsgType = int32(decodedMsg.indicationMessage.present)

	if indMsg.IndMsgType == 2 {
		indMsgFormat2 := &IndicationMessageFormat2{}
		indMsgFormat2_C := *(**C.E2SM_KPM_IndicationMessage_Format2_t)(unsafe.Pointer(&decodedMsg.indicationMessage.choice[0]))

		indMsgFormat2.measData = int(indMsgFormat2_C.measData)
		indMsgFormat2.granulPeriod = int(indMsgFormat2_C.granulPeriod)
		indMsgFormat2.measCondUEidListCount = int(indMsgFormat2_C.measCondUEidList.list.count)
		for i := 0; i < indMsgFormat2.measCondUEidListCount; i++ {
			measCondUEidListItem := &indMsgFormat2.measCondUEidList[i]
			var sizeof_MeasurementCondUEidList_t *C.MeasurementCondUEidList_t
			MeasurementCondUEidItem_C := *(**C.MeasurementCondUEidList_t)(unsafe.Pointer(uintptr(unsafe.Pointer(indMsgFormat2_C.MeasurementCondUEidList.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_MeasurementCondUEidList_t)))

			if MeasurementCondUEidItem_C.measType != nil {
				measType := &MeasurementType{}

				measType.measID = int32(MeasurementCondUEidItem_C.measType.measID)

				// if pfContainer.ContainerType == 1 {
				// 	oDU_PF := &ODUPFContainerType{}
				// 	oDU_PF_C := *(**C.ODU_PF_Container_t)(unsafe.Pointer(&pmContainer_C.performanceContainer.choice[0]))

				// 	oDU_PF.CellResourceReportCount = int(oDU_PF_C.cellResourceReportList.list.count)
				// 	for j := 0; j < oDU_PF.CellResourceReportCount; j++ {
				// 		cellResourceReport := &oDU_PF.CellResourceReports[j]
				// 		var sizeof_CellResourceReportListItem_t *C.CellResourceReportListItem_t
				// 		cellResourceReport_C := *(**C.CellResourceReportListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oDU_PF_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_CellResourceReportListItem_t)))

				// 		cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
				// 		cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

				// 		cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
				// 		cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
				// 		cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

				// 		if cellResourceReport_C.dl_TotalofAvailablePRBs != nil {
				// 			cellResourceReport.TotalofAvailablePRBs.DL = int64(*cellResourceReport_C.dl_TotalofAvailablePRBs)
				// 		} else {
				// 			cellResourceReport.TotalofAvailablePRBs.DL = -1
				// 		}

				// 		if cellResourceReport_C.ul_TotalofAvailablePRBs != nil {
				// 			cellResourceReport.TotalofAvailablePRBs.UL = int64(*cellResourceReport_C.ul_TotalofAvailablePRBs)
				// 		} else {
				// 			cellResourceReport.TotalofAvailablePRBs.UL = -1
				// 		}

				// 		cellResourceReport.ServedPlmnPerCellCount = int(cellResourceReport_C.servedPlmnPerCellList.list.count)
				// 		for k := 0; k < cellResourceReport.ServedPlmnPerCellCount; k++ {
				// 			servedPlmnPerCell := cellResourceReport.ServedPlmnPerCells[k]
				// 			var sizeof_ServedPlmnPerCellListItem_t *C.ServedPlmnPerCellListItem_t
				// 			servedPlmnPerCell_C := *(**C.ServedPlmnPerCellListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.servedPlmnPerCellList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_ServedPlmnPerCellListItem_t)))

				// 			servedPlmnPerCell.PlmnID.Buf = C.GoBytes(unsafe.Pointer(servedPlmnPerCell_C.pLMN_Identity.buf), C.int(servedPlmnPerCell_C.pLMN_Identity.size))
				// 			servedPlmnPerCell.PlmnID.Size = int(servedPlmnPerCell_C.pLMN_Identity.size)

				// 			if servedPlmnPerCell_C.du_PM_5GC != nil {
				// 				duPM5GC := &DUPM5GCContainerType{}
				// 				duPM5GC_C := (*C.FGC_DU_PM_Container_t)(servedPlmnPerCell_C.du_PM_5GC)

				// 				duPM5GC.SlicePerPlmnPerCellCount = int(duPM5GC_C.slicePerPlmnPerCellList.list.count)
				// 				for l := 0; l < duPM5GC.SlicePerPlmnPerCellCount; l++ {
				// 					slicePerPlmnPerCell := &duPM5GC.SlicePerPlmnPerCells[l]
				// 					var sizeof_SlicePerPlmnPerCellListItem_t *C.SlicePerPlmnPerCellListItem_t
				// 					slicePerPlmnPerCell_C := *(**C.SlicePerPlmnPerCellListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(duPM5GC_C.slicePerPlmnPerCellList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_SlicePerPlmnPerCellListItem_t)))

				// 					slicePerPlmnPerCell.SliceID.SST.Buf = C.GoBytes(unsafe.Pointer(slicePerPlmnPerCell_C.sliceID.sST.buf), C.int(slicePerPlmnPerCell_C.sliceID.sST.size))
				// 					slicePerPlmnPerCell.SliceID.SST.Size = int(slicePerPlmnPerCell_C.sliceID.sST.size)

				// 					if slicePerPlmnPerCell_C.sliceID.sD != nil {
				// 						slicePerPlmnPerCell.SliceID.SD = &OctetString{}
				// 						slicePerPlmnPerCell.SliceID.SD.Buf = C.GoBytes(unsafe.Pointer(slicePerPlmnPerCell_C.sliceID.sD.buf), C.int(slicePerPlmnPerCell_C.sliceID.sD.size))
				// 						slicePerPlmnPerCell.SliceID.SD.Size = int(slicePerPlmnPerCell_C.sliceID.sD.size)
				// 					}

				// 					slicePerPlmnPerCell.FQIPERSlicesPerPlmnPerCellCount = int(slicePerPlmnPerCell_C.fQIPERSlicesPerPlmnPerCellList.list.count)
				// 					for m := 0; m < slicePerPlmnPerCell.FQIPERSlicesPerPlmnPerCellCount; m++ {
				// 						fQIPerSlicesPerPlmnPerCell := &slicePerPlmnPerCell.FQIPERSlicesPerPlmnPerCells[m]
				// 						var sizeof_FQIPERSlicesPerPlmnPerCellListItem_t *C.FQIPERSlicesPerPlmnPerCellListItem_t
				// 						fQIPerSlicesPerPlmnPerCell_C := *(**C.FQIPERSlicesPerPlmnPerCellListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(slicePerPlmnPerCell_C.fQIPERSlicesPerPlmnPerCellList.list.array)) + (uintptr)(m)*unsafe.Sizeof(sizeof_FQIPERSlicesPerPlmnPerCellListItem_t)))

				// 						fQIPerSlicesPerPlmnPerCell.FiveQI = int64(fQIPerSlicesPerPlmnPerCell_C.fiveQI)

				// 						if fQIPerSlicesPerPlmnPerCell_C.dl_PRBUsage != nil {
				// 							fQIPerSlicesPerPlmnPerCell.PrbUsage.DL = int64(*fQIPerSlicesPerPlmnPerCell_C.dl_PRBUsage)
				// 						} else {
				// 							fQIPerSlicesPerPlmnPerCell.PrbUsage.DL = -1
				// 						}

				// 						if fQIPerSlicesPerPlmnPerCell_C.ul_PRBUsage != nil {
				// 							fQIPerSlicesPerPlmnPerCell.PrbUsage.UL = int64(*fQIPerSlicesPerPlmnPerCell_C.ul_PRBUsage)
				// 						} else {
				// 							fQIPerSlicesPerPlmnPerCell.PrbUsage.UL = -1
				// 						}
				// 					}
				// 				}

				// 				servedPlmnPerCell.DUPM5GC = duPM5GC
				// 			}

				// 			if servedPlmnPerCell_C.du_PM_EPC != nil {
				// 				duPMEPC := &DUPMEPCContainerType{}
				// 				duPMEPC_C := (*C.EPC_DU_PM_Container_t)(servedPlmnPerCell_C.du_PM_EPC)

				// 				duPMEPC.PerQCIReportCount = int(duPMEPC_C.perQCIReportList.list.count)
				// 				for l := 0; l < duPMEPC.PerQCIReportCount; l++ {
				// 					perQCIReport := &duPMEPC.PerQCIReports[l]
				// 					var sizeof_PerQCIReportListItem_t *C.PerQCIReportListItem_t
				// 					perQCIReport_C := *(**C.PerQCIReportListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(duPMEPC_C.perQCIReportList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_PerQCIReportListItem_t)))

				// 					perQCIReport.QCI = int64(perQCIReport_C.qci)

				// 					if perQCIReport_C.dl_PRBUsage != nil {
				// 						perQCIReport.PrbUsage.DL = int64(*perQCIReport_C.dl_PRBUsage)
				// 					} else {
				// 						perQCIReport.PrbUsage.DL = -1
				// 					}

				// 					if perQCIReport_C.ul_PRBUsage != nil {
				// 						perQCIReport.PrbUsage.UL = int64(*perQCIReport_C.ul_PRBUsage)
				// 					} else {
				// 						perQCIReport.PrbUsage.UL = -1
				// 					}
				// 				}

				// 				servedPlmnPerCell.DUPMEPC = duPMEPC
				// 			}
				// 		}
				// 	}

				// 	pfContainer.Container = oDU_PF
				// } else if pfContainer.ContainerType == 2 {
				// 	oCU_CP_PF := &OCUCPPFContainerType{}
				// 	oCU_CP_PF_C := *(**C.OCUCP_PF_Container_t)(unsafe.Pointer(&pmContainer_C.performanceContainer.choice[0]))

				// 	if oCU_CP_PF_C.gNB_CU_CP_Name != nil {
				// 		oCU_CP_PF.GNBCUCPName = &PrintableString{}
				// 		oCU_CP_PF.GNBCUCPName.Buf = C.GoBytes(unsafe.Pointer(oCU_CP_PF_C.gNB_CU_CP_Name.buf), C.int(oCU_CP_PF_C.gNB_CU_CP_Name.size))
				// 		oCU_CP_PF.GNBCUCPName.Size = int(oCU_CP_PF_C.gNB_CU_CP_Name.size)
				// 	}

				// 	if oCU_CP_PF_C.cu_CP_Resource_Status.numberOfActive_UEs != nil {
				// 		oCU_CP_PF.CUCPResourceStatus.NumberOfActiveUEs = int64(*oCU_CP_PF_C.cu_CP_Resource_Status.numberOfActive_UEs)
				// 	}

				// 	pfContainer.Container = oCU_CP_PF
				// } else if pfContainer.ContainerType == 3 {
				// 	oCU_UP_PF := &OCUUPPFContainerType{}
				// 	oCU_UP_PF_C := *(**C.OCUUP_PF_Container_t)(unsafe.Pointer(&pmContainer_C.performanceContainer.choice[0]))

				// 	if oCU_UP_PF_C.gNB_CU_UP_Name != nil {
				// 		oCU_UP_PF.GNBCUUPName = &PrintableString{}
				// 		oCU_UP_PF.GNBCUUPName.Buf = C.GoBytes(unsafe.Pointer(oCU_UP_PF_C.gNB_CU_UP_Name.buf), C.int(oCU_UP_PF_C.gNB_CU_UP_Name.size))
				// 		oCU_UP_PF.GNBCUUPName.Size = int(oCU_UP_PF_C.gNB_CU_UP_Name.size)
				// 	}

				// 	oCU_UP_PF.CUUPPFContainerItemCount = int(oCU_UP_PF_C.pf_ContainerList.list.count)
				// 	for j := 0; j < oCU_UP_PF.CUUPPFContainerItemCount; j++ {
				// 		cuUPPFContainer := &oCU_UP_PF.CUUPPFContainerItems[j]
				// 		var sizeof_PF_ContainerListItem_t *C.PF_ContainerListItem_t
				// 		cuUPPFContainer_C := *(**C.PF_ContainerListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oCU_UP_PF_C.pf_ContainerList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_PF_ContainerListItem_t)))

				// 		cuUPPFContainer.InterfaceType = int64(cuUPPFContainer_C.interface_type)

				// 		cuUPPFContainer.OCUUPPMContainer.CUUPPlmnCount = int(cuUPPFContainer_C.o_CU_UP_PM_Container.plmnList.list.count)
				// 		for k := 0; k < cuUPPFContainer.OCUUPPMContainer.CUUPPlmnCount; k++ {
				// 			cuUPPlmn := &cuUPPFContainer.OCUUPPMContainer.CUUPPlmns[k]
				// 			var sizeof_PlmnID_List_t *C.PlmnID_List_t
				// 			cuUPPlmn_C := *(**C.PlmnID_List_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cuUPPFContainer_C.o_CU_UP_PM_Container.plmnList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_PlmnID_List_t)))

				// 			cuUPPlmn.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cuUPPlmn_C.pLMN_Identity.buf), C.int(cuUPPlmn_C.pLMN_Identity.size))
				// 			cuUPPlmn.PlmnID.Size = int(cuUPPlmn_C.pLMN_Identity.size)

				// 			if cuUPPlmn_C.cu_UP_PM_5GC != nil {
				// 				cuUPPM5GC := &CUUPPM5GCType{}
				// 				cuUPPM5GC_C := (*C.FGC_CUUP_PM_Format_t)(cuUPPlmn_C.cu_UP_PM_5GC)

				// 				cuUPPM5GC.SliceToReportCount = int(cuUPPM5GC_C.sliceToReportList.list.count)
				// 				for l := 0; l < cuUPPM5GC.SliceToReportCount; l++ {
				// 					sliceToReport := &cuUPPM5GC.SliceToReports[l]
				// 					var sizeof_SliceToReportListItem_t *C.SliceToReportListItem_t
				// 					sliceToReport_C := *(**C.SliceToReportListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cuUPPM5GC_C.sliceToReportList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_SliceToReportListItem_t)))

				// 					sliceToReport.SliceID.SST.Buf = C.GoBytes(unsafe.Pointer(sliceToReport_C.sliceID.sST.buf), C.int(sliceToReport_C.sliceID.sST.size))
				// 					sliceToReport.SliceID.SST.Size = int(sliceToReport_C.sliceID.sST.size)

				// 					if sliceToReport_C.sliceID.sD != nil {
				// 						sliceToReport.SliceID.SD = &OctetString{}
				// 						sliceToReport.SliceID.SD.Buf = C.GoBytes(unsafe.Pointer(sliceToReport_C.sliceID.sD.buf), C.int(sliceToReport_C.sliceID.sD.size))
				// 						sliceToReport.SliceID.SD.Size = int(sliceToReport_C.sliceID.sD.size)
				// 					}

				// 					sliceToReport.FQIPERSlicesPerPlmnCount = int(sliceToReport_C.fQIPERSlicesPerPlmnList.list.count)
				// 					for m := 0; m < sliceToReport.FQIPERSlicesPerPlmnCount; m++ {
				// 						fQIPerSlicesPerPlmn := &sliceToReport.FQIPERSlicesPerPlmns[m]
				// 						var sizeof_FQIPERSlicesPerPlmnListItem_t *C.FQIPERSlicesPerPlmnListItem_t
				// 						fQIPerSlicesPerPlmn_C := *(**C.FQIPERSlicesPerPlmnListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(sliceToReport_C.fQIPERSlicesPerPlmnList.list.array)) + (uintptr)(m)*unsafe.Sizeof(sizeof_FQIPERSlicesPerPlmnListItem_t)))

				// 						fQIPerSlicesPerPlmn.FiveQI = int64(fQIPerSlicesPerPlmn_C.fiveQI)

				// 						if fQIPerSlicesPerPlmn_C.pDCPBytesDL != nil {
				// 							fQIPerSlicesPerPlmn.PDCPBytesDL = &Integer{}
				// 							fQIPerSlicesPerPlmn.PDCPBytesDL.Buf = C.GoBytes(unsafe.Pointer(fQIPerSlicesPerPlmn_C.pDCPBytesDL.buf), C.int(fQIPerSlicesPerPlmn_C.pDCPBytesDL.size))
				// 							fQIPerSlicesPerPlmn.PDCPBytesDL.Size = int(fQIPerSlicesPerPlmn_C.pDCPBytesDL.size)
				// 						}

				// 						if fQIPerSlicesPerPlmn_C.pDCPBytesUL != nil {
				// 							fQIPerSlicesPerPlmn.PDCPBytesUL = &Integer{}
				// 							fQIPerSlicesPerPlmn.PDCPBytesUL.Buf = C.GoBytes(unsafe.Pointer(fQIPerSlicesPerPlmn_C.pDCPBytesUL.buf), C.int(fQIPerSlicesPerPlmn_C.pDCPBytesUL.size))
				// 							fQIPerSlicesPerPlmn.PDCPBytesUL.Size = int(fQIPerSlicesPerPlmn_C.pDCPBytesUL.size)
				// 						}
				// 					}
				// 				}

				// 				cuUPPlmn.CUUPPM5GC = cuUPPM5GC
				// 			}

				// 			if cuUPPlmn_C.cu_UP_PM_EPC != nil {
				// 				cuUPPMEPC := &CUUPPMEPCType{}
				// 				cuUPPMEPC_C := (*C.EPC_CUUP_PM_Format_t)(cuUPPlmn_C.cu_UP_PM_EPC)

				// 				cuUPPMEPC.CUUPPMEPCPerQCIReportCount = int(cuUPPMEPC_C.perQCIReportList.list.count)
				// 				for l := 0; l < cuUPPMEPC.CUUPPMEPCPerQCIReportCount; l++ {
				// 					perQCIReport := &cuUPPMEPC.CUUPPMEPCPerQCIReports[l]
				// 					var sizeof_PerQCIReportListItemFormat_t *C.PerQCIReportListItemFormat_t
				// 					perQCIReport_C := *(**C.PerQCIReportListItemFormat_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cuUPPMEPC_C.perQCIReportList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_PerQCIReportListItemFormat_t)))

				// 					perQCIReport.QCI = int64(perQCIReport_C.qci)

				// 					if perQCIReport_C.pDCPBytesDL != nil {
				// 						perQCIReport.PDCPBytesDL = &Integer{}
				// 						perQCIReport.PDCPBytesDL.Buf = C.GoBytes(unsafe.Pointer(perQCIReport_C.pDCPBytesDL.buf), C.int(perQCIReport_C.pDCPBytesDL.size))
				// 						perQCIReport.PDCPBytesDL.Size = int(perQCIReport_C.pDCPBytesDL.size)
				// 					}

				// 					if perQCIReport_C.pDCPBytesUL != nil {
				// 						perQCIReport.PDCPBytesUL = &Integer{}
				// 						perQCIReport.PDCPBytesUL.Buf = C.GoBytes(unsafe.Pointer(perQCIReport_C.pDCPBytesUL.buf), C.int(perQCIReport_C.pDCPBytesUL.size))
				// 						perQCIReport.PDCPBytesUL.Size = int(perQCIReport_C.pDCPBytesUL.size)
				// 					}
				// 				}

				// 				cuUPPlmn.CUUPPMEPC = cuUPPMEPC
				// 			}
				// 		}
				// 	}

				// 	pfContainer.Container = oCU_UP_PF
				// } else {
				// 	return indMsg, errors.New("Unknown PF Container type")
				// }

				measCondUEidListItem.measType = measType
			}

			// if pmContainer_C.theRANContainer != nil {
			// 	ranContainer := &RANContainerType{}

			// 	ranContainer.Timestamp.Buf = C.GoBytes(unsafe.Pointer(pmContainer_C.theRANContainer.timestamp.buf), C.int(pmContainer_C.theRANContainer.timestamp.size))
			// 	ranContainer.Timestamp.Size = int(pmContainer_C.theRANContainer.timestamp.size)

			// 	ranContainer.ContainerType = int32(pmContainer_C.theRANContainer.reportContainer.present)

			// 	if ranContainer.ContainerType == 1 {
			// 		oDU_UE := &DUUsageReportType{}
			// 		oDU_UE_C := *(**C.DU_Usage_Report_Per_UE_t)(unsafe.Pointer(&pmContainer_C.theRANContainer.reportContainer.choice[0]))

			// 		oDU_UE.CellResourceReportItemCount = int(oDU_UE_C.cellResourceReportList.list.count)
			// 		for j := 0; j < oDU_UE.CellResourceReportItemCount; j++ {
			// 			cellResourceReport := &oDU_UE.CellResourceReportItems[j]
			// 			var sizeof_DU_Usage_Report_CellResourceReportItem_t *C.DU_Usage_Report_CellResourceReportItem_t
			// 			cellResourceReport_C := *(**C.DU_Usage_Report_CellResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oDU_UE_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_DU_Usage_Report_CellResourceReportItem_t)))

			// 			cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
			// 			cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

			// 			cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
			// 			cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
			// 			cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

			// 			cellResourceReport.UeResourceReportItemCount = int(cellResourceReport_C.ueResourceReportList.list.count)
			// 			for k := 0; k < cellResourceReport.UeResourceReportItemCount; k++ {
			// 				ueResourceReport := &cellResourceReport.UeResourceReportItems[k]
			// 				var sizeof_DU_Usage_Report_UeResourceReportItem_t *C.DU_Usage_Report_UeResourceReportItem_t
			// 				ueResourceReport_C := *(**C.DU_Usage_Report_UeResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.ueResourceReportList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_DU_Usage_Report_UeResourceReportItem_t)))

			// 				ueResourceReport.CRNTI.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.c_RNTI.buf), C.int(ueResourceReport_C.c_RNTI.size))
			// 				ueResourceReport.CRNTI.Size = int(ueResourceReport_C.c_RNTI.size)

			// 				if ueResourceReport_C.dl_PRBUsage != nil {
			// 					ueResourceReport.PRBUsageDL = int64(*ueResourceReport_C.dl_PRBUsage)
			// 				} else {
			// 					ueResourceReport.PRBUsageDL = -1
			// 				}

			// 				if ueResourceReport_C.ul_PRBUsage != nil {
			// 					ueResourceReport.PRBUsageUL = int64(*ueResourceReport_C.ul_PRBUsage)
			// 				} else {
			// 					ueResourceReport.PRBUsageUL = -1
			// 				}
			// 			}
			// 		}

			// 		ranContainer.Container = oDU_UE
			// 	} else if ranContainer.ContainerType == 2 {
			// 		oCU_CP_UE := &CUCPUsageReportType{}
			// 		oCU_CP_UE_C := *(**C.CU_CP_Usage_Report_Per_UE_t)(unsafe.Pointer(&pmContainer_C.theRANContainer.reportContainer.choice[0]))

			// 		oCU_CP_UE.CellResourceReportItemCount = int(oCU_CP_UE_C.cellResourceReportList.list.count)
			// 		for j := 0; j < oCU_CP_UE.CellResourceReportItemCount; j++ {
			// 			cellResourceReport := &oCU_CP_UE.CellResourceReportItems[j]
			// 			var sizeof_CU_CP_Usage_Report_CellResourceReportItem_t *C.CU_CP_Usage_Report_CellResourceReportItem_t
			// 			cellResourceReport_C := *(**C.CU_CP_Usage_Report_CellResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oCU_CP_UE_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_CU_CP_Usage_Report_CellResourceReportItem_t)))

			// 			cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
			// 			cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

			// 			cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
			// 			cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
			// 			cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

			// 			cellResourceReport.UeResourceReportItemCount = int(cellResourceReport_C.ueResourceReportList.list.count)
			// 			for k := 0; k < cellResourceReport.UeResourceReportItemCount; k++ {
			// 				ueResourceReport := &cellResourceReport.UeResourceReportItems[k]
			// 				var sizeof_CU_CP_Usage_Report_UeResourceReportItem_t *C.CU_CP_Usage_Report_UeResourceReportItem_t
			// 				ueResourceReport_C := *(**C.CU_CP_Usage_Report_UeResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.ueResourceReportList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_CU_CP_Usage_Report_UeResourceReportItem_t)))

			// 				ueResourceReport.CRNTI.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.c_RNTI.buf), C.int(ueResourceReport_C.c_RNTI.size))
			// 				ueResourceReport.CRNTI.Size = int(ueResourceReport_C.c_RNTI.size)

			// 				if ueResourceReport_C.serving_Cell_RF_Type != nil {
			// 					ueResourceReport.ServingCellRF = &OctetString{}
			// 					ueResourceReport.ServingCellRF.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.serving_Cell_RF_Type.buf), C.int(ueResourceReport_C.serving_Cell_RF_Type.size))
			// 					ueResourceReport.ServingCellRF.Size = int(ueResourceReport_C.serving_Cell_RF_Type.size)
			// 				}

			// 				if ueResourceReport_C.neighbor_Cell_RF != nil {
			// 					ueResourceReport.NeighborCellRF = &OctetString{}
			// 					ueResourceReport.NeighborCellRF.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.neighbor_Cell_RF.buf), C.int(ueResourceReport_C.neighbor_Cell_RF.size))
			// 					ueResourceReport.NeighborCellRF.Size = int(ueResourceReport_C.neighbor_Cell_RF.size)
			// 				}
			// 			}
			// 		}

			// 		ranContainer.Container = oCU_CP_UE
			// 	} else if ranContainer.ContainerType == 3 {
			// 		oCU_UP_UE := &CUUPUsageReportType{}
			// 		oCU_UP_UE_C := *(**C.CU_UP_Usage_Report_Per_UE_t)(unsafe.Pointer(&pmContainer_C.theRANContainer.reportContainer.choice[0]))

			// 		oCU_UP_UE.CellResourceReportItemCount = int(oCU_UP_UE_C.cellResourceReportList.list.count)
			// 		for j := 0; j < oCU_UP_UE.CellResourceReportItemCount; j++ {
			// 			cellResourceReport := &oCU_UP_UE.CellResourceReportItems[j]
			// 			var sizeof_CU_UP_Usage_Report_CellResourceReportItem_t *C.CU_UP_Usage_Report_CellResourceReportItem_t
			// 			cellResourceReport_C := *(**C.CU_UP_Usage_Report_CellResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oCU_UP_UE_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_CU_UP_Usage_Report_CellResourceReportItem_t)))

			// 			cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
			// 			cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

			// 			cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
			// 			cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
			// 			cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

			// 			cellResourceReport.UeResourceReportItemCount = int(cellResourceReport_C.ueResourceReportList.list.count)
			// 			for k := 0; k < cellResourceReport.UeResourceReportItemCount; k++ {
			// 				ueResourceReport := &cellResourceReport.UeResourceReportItems[k]
			// 				var sizeof_CU_UP_Usage_Report_UeResourceReportItem_t *C.CU_UP_Usage_Report_UeResourceReportItem_t
			// 				ueResourceReport_C := *(**C.CU_UP_Usage_Report_UeResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.ueResourceReportList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_CU_UP_Usage_Report_UeResourceReportItem_t)))

			// 				ueResourceReport.CRNTI.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.c_RNTI.buf), C.int(ueResourceReport_C.c_RNTI.size))
			// 				ueResourceReport.CRNTI.Size = int(ueResourceReport_C.c_RNTI.size)

			// 				if ueResourceReport_C.pDCPBytesDL != nil {
			// 					ueResourceReport.PDCPBytesDL = &Integer{}
			// 					ueResourceReport.PDCPBytesDL.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.pDCPBytesDL.buf), C.int(ueResourceReport_C.pDCPBytesDL.size))
			// 					ueResourceReport.PDCPBytesDL.Size = int(ueResourceReport_C.pDCPBytesDL.size)
			// 				}

			// 				if ueResourceReport_C.pDCPBytesUL != nil {
			// 					ueResourceReport.PDCPBytesUL = &Integer{}
			// 					ueResourceReport.PDCPBytesUL.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.pDCPBytesUL.buf), C.int(ueResourceReport_C.pDCPBytesUL.size))
			// 					ueResourceReport.PDCPBytesUL.Size = int(ueResourceReport_C.pDCPBytesUL.size)
			// 				}
			// 			}
			// 		}

			// 		ranContainer.Container = oCU_UP_UE
			// 	} else {
			// 		return indMsg, errors.New("Unknown RAN Container type")
			// 	}

			// 	pmContainer.RANContainer = ranContainer
			// }
		}

		indMsg.IndMsg = indMsgFormat2
	} else {
		return indMsg, errors.New("Unknown RIC Indication Message Format")
	}

	return
}
