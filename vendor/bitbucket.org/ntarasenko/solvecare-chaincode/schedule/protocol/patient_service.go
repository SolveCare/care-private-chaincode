package protocol

import (
	careproto "bitbucket.org/ntarasenko/solvecare-chaincode/schedule/protocol/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type PatientService interface {

	DecodeProtoByteString(encodedPatientByteString string) (*careproto.Patient, error)
	SavePatient(stub shim.ChaincodeStubInterface, patient careproto.Patient) (*careproto.Patient, error)
	GetPatientById(stub shim.ChaincodeStubInterface, patientId string) (*careproto.Patient, error)

}