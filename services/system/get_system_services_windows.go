package system

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

func platformGetSystemServices(response *pb.GetSystemServicesResponse) (*pb.GetSystemServicesResponse, error) {
	err := ole.CoInitialize(0)
	if err != nil {

		return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("error initializing OLE: %v", err), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	defer ole.CoUninitialize()

	// Create WMI Service object
	locator, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("failed to create WbemScripting.SWbemLocator object: %v", err), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	defer locator.Release()

	wmiService, err := locator.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("failed to query WbemScripting.SWbemLocator interface: %v", err), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	defer wmiService.Release()

	// Connect to the WMI service
	serviceConnection, err := oleutil.CallMethod(wmiService, "ConnectServer")
	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("failed to connect to WMI service: %v", err), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	service := serviceConnection.ToIDispatch()
	defer service.Release()

	// Execute WMI query to get all services
	result, err := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_Service")
	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("failed to execute query on WMI service: %v", err), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	serviceList := result.ToIDispatch()
	defer serviceList.Release()

	// Iterate over the services
	countVar, _ := oleutil.GetProperty(serviceList, "Count")
	count := int(countVar.Val)
	for i := 0; i < count; i++ {
		item, _ := oleutil.CallMethod(serviceList, "ItemIndex", i)
		serviceItem := item.ToIDispatch()
		if serviceItem == nil {
			continue
		}

		name, _ := oleutil.GetProperty(serviceItem, "Name")
		state, _ := oleutil.GetProperty(serviceItem, "State")
		desc, _ := oleutil.GetProperty(serviceItem, "Description")

		response.List = append(response.List, &pb.SystemServiceInfo{
			Name:        name.ToString(),
			State:       state.ToString(),
			Description: desc.ToString(),
		})
		serviceItem.Release()
	}
	utils.LogInfo(fmt.Sprintf("GetSystemServices response len: %d", len(response.List)))

	return response, nil
}
