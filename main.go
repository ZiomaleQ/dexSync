package main

import (
	"errors"
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	path, err := getKindleVolume()

	if err != nil {
		fmt.Println(err)
		return
	}

	pathString := *path
	fmt.Println(pathString)

	//dir, err := os.ReadDir(pathString + "/dexSync")

	//if err != nil {
	//fmt.Println(err)
	//return
	//}
}

func getKindleVolume() (*string, error) {
	//get all devices
	devices, err := windows.GetLogicalDrives()

	if err != nil {
		return nil, err
	}

	for i := 0; i < 26; i++ {
		if devices&(1<<uint32(i)) != 0 {
			path := fmt.Sprintf("%c:\\", 'A'+i)
			temp, err := windows.UTF16PtrFromString(path)

			if err != nil {
				return nil, err
			}

			if windows.GetDriveType(temp) != windows.DRIVE_REMOVABLE {
				continue
			}

			buf := make([]uint16, windows.MAX_PATH)

			//read volume information
			err = windows.GetVolumeInformation(temp, &buf[0], windows.MAX_PATH, nil, nil, nil, nil, 0)

			if err != nil {
				return nil, err
			}

			if windows.UTF16ToString(buf) == "Kindle" {
				return &path, nil
			}
		}
	}

	return nil, errors.New("kindle disk not found")
}
