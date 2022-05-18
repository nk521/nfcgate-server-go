package main

// MifareDesfireATQA := [2]byte{0x03, 0x44}
// MifareDesfireSAK := [1]byte{0x20}
// MifareDesfireATS := [5]byte{0x06, 0x75, 0x77, 0x81, 0x02, 0x80} // |06 75 77 81 02| 80(HST BYTE) XX XX (CRC)

// https://www.st.com/resource/en/datasheet/st25ta64k.pdf page 37, http://www.emutag.com/iso/14443-4.pdf page 15,26
// 0x06 - TL (length of ATS response)
// 0x75 - T0 FSCI (he maximum frame size that the PICC is able to recv.) | least significant nibble defines FSD. In this case its 0x5 which is 64 bytes.
// 0x77 - TA1 (data rate value)
// 0x81 - TB1 (b8-b5 == FWI | 0b1000 in this case, b4-b1 == SFGI | 0b0001 in this case) | change to 0xE1 for max FWI ~4949 ms
// 0x02 - TC1 (DID support)
// 0x80 - HIST BYTE

type MifareDesireDF struct {
	AID       uint32
	FID       uint16
	DFName    [16]uint8
	DFNameLen uint64
}

type MifareDesfireFileSettings struct {
	FileType              uint8
	CommunicationSettings uint8
	AccessRights          uint8
	// standard file
	StandardFile struct {
		Size uint32
	}
	// value file
	ValueFile struct {
		LowerLimit           uint32
		UpperLimit           uint32
		LimitedCreditValue   uint32
		LimitedCreditEnabled uint8
	}
	// linear record file
	LinearRecordFile struct {
		RecordSize             uint32
		MaxNumberOfRecords     uint32
		CurrentNumberOfRecords uint32
	}
}

type MifareDesfireVersionInfo struct {
	Hardware struct {
		VendorId     uint8
		Type         uint8
		Subtype      uint8
		VersionMajor uint8
		VersionMinor uint8
		StorageSize  uint8
		Protocol     uint8
	}
	Software struct {
		VendorId     uint8
		Type         uint8
		Subtype      uint8
		VersionMajor uint8
		VersionMinor uint8
		StorageSize  uint8
		Protocol     uint8
	}
	UID         [7]uint8
	BatchNumber [5]uint8
	ProdWeek    uint8
	ProdYear    uint8
}
