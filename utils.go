package main

import (
	nfc_pb "github.com/nk521/nfcgate-server-go/nfcgate_proto"
	"google.golang.org/protobuf/proto"
)

type UnmarshalledStruct struct {
	ServerData *nfc_pb.ServerData
	NfcData    *nfc_pb.NFCData
}

func unmarshallServer(marshalledBytes []byte) *nfc_pb.ServerData {
	server := nfc_pb.ServerData{}
	proto.Unmarshal(marshalledBytes, &server)
	return &server
}

func unmarshallNFCData(marshalledBytes []byte) *nfc_pb.NFCData {
	nfcData := nfc_pb.NFCData{}
	proto.Unmarshal(marshalledBytes, &nfcData)
	return &nfcData
}

func unmarshallServerAndNFCData(marshalledBytes []byte) *nfc_pb.NFCData {
	server := *unmarshallServer(marshalledBytes)
	nfcData := *unmarshallNFCData(server.GetData())
	return &nfcData
}

func unmarshallServerAndNFCDataStruct(marshalledBytes []byte) UnmarshalledStruct {
	server := *unmarshallServer(marshalledBytes)
	nfcData := *unmarshallNFCData(server.GetData())
	return UnmarshalledStruct{
		ServerData: &server,
		NfcData:    &nfcData,
	}
}

func marshallServer(opcode nfc_pb.ServerData_Opcode, data []byte) []byte {
	server := nfc_pb.ServerData{
		Opcode: opcode,
		Data:   data,
	}
	marshalled_data, err := proto.Marshal(&server)
	checkError(err)
	return marshalled_data
}

func marshallNFCData(data_source nfc_pb.NFCData_DataSource, data_type nfc_pb.NFCData_DataType, timestamp int64, data []byte) []byte {
	nfcData := nfc_pb.NFCData{
		DataSource: data_source,
		DataType:   data_type,
		Timestamp:  timestamp,
		Data:       data,
	}
	marshalled_data, err := proto.Marshal(&nfcData)
	checkError(err)
	return marshalled_data
}

func marshallNFCDataAndServer(opcode nfc_pb.ServerData_Opcode, data_source nfc_pb.NFCData_DataSource, data_type nfc_pb.NFCData_DataType, timestamp int64, data []byte) []byte {
	server := marshallNFCData(data_source, data_type, timestamp, data)
	nfcData := marshallServer(opcode, server)
	return nfcData
}
