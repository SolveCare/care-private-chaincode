package protocol

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	careproto "bitbucket.org/ntarasenko/solvecare-chaincode/schedule/protocol/proto"
)

type Scheduler interface {

	GetByOwnerId(stub shim.ChaincodeStubInterface, ownerId string) (*careproto.Schedule, error)
	Save(stub shim.ChaincodeStubInterface, schedule careproto.Schedule) (*careproto.Schedule, error)
	ConstructScheduleKey(ownerId string) string
}
