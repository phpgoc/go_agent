package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	//这个cmd没有实际意义，只是为了测试
	// Initialize OLE
	err := ole.CoInitialize(0)
	if err != nil {
		fmt.Println("Error initializing OLE:", err)
		return
	}
	defer ole.CoUninitialize()

	// Create WMI Service object
	locator, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		fmt.Println("Failed to create WbemScripting.SWbemLocator object:", err)
		return
	}
	defer locator.Release()

	wmiService, err := locator.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		fmt.Println("Failed to query WbemScripting.SWbemLocator interface:", err)
		return
	}
	defer wmiService.Release()

	// Connect to the WMI service
	serviceConnection, err := oleutil.CallMethod(wmiService, "ConnectServer")
	if err != nil {
		fmt.Println("Failed to connect to WMI service:", err)
		return
	}
	service := serviceConnection.ToIDispatch()
	defer service.Release()

	// Execute WMI query to get all services
	result, err := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_Service")
	if err != nil {
		fmt.Println("Failed to execute query on WMI service:", err)
		return
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

		fmt.Printf("Service Name: %s, State: %s, desc: %s\n", name.ToString(), state.ToString(), desc.ToString())
		serviceItem.Release()
	}
}
