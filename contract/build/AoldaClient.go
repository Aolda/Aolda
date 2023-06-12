// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aoldaClient

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AoldaClientMetaData contains all meta data concerning the AoldaClient contract.
var AoldaClientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"funtionName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"arguments\",\"type\":\"string[]\"}],\"name\":\"CallAolda\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funtionName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"arguments\",\"type\":\"string[]\"}],\"name\":\"callAolda\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funtionName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"arguments\",\"type\":\"string[]\"}],\"name\":\"makeSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610773806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80636984394014610051578063916c47591461007a578063c7bfebd61461009b578063cd6b7a26146100b0575b600080fd5b61006461005f366004610277565b6100c3565b60405161007191906102e0565b60405180910390f35b61008d6100883660046103b1565b610165565b604051908152602001610071565b6100ae6100a93660046104bb565b61021a565b005b6100ae6100be3660046103b1565b610237565b60008181526020819052604090208054606091906100e090610502565b80601f016020809104026020016040519081016040528092919081815260200182805461010c90610502565b80156101595780601f1061012e57610100808354040283529160200191610159565b820191906000526020600020905b81548152906001019060200180831161013c57829003601f168201915b50505050509050919050565b600080848460405160200161017b92919061053c565b604051602081830303815290604052905060005b83518110156101e857818482815181106101ab576101ab61056b565b60200260200101516040516020016101c492919061053c565b604051602081830303815290604052915080806101e090610581565b91505061018f565b50806040516020016101fa91906102e0565b604051602081830303815290604052805190602001209150509392505050565b600082815260208190526040902061023282826105f6565b505050565b7f1d37bd08d8b1245ea7f84485bab95b6726cc1fd59b42ebf52530669bb1137f6283838360405161026a939291906106b6565b60405180910390a1505050565b60006020828403121561028957600080fd5b5035919050565b60005b838110156102ab578181015183820152602001610293565b50506000910152565b600081518084526102cc816020860160208601610290565b601f01601f19169290920160200192915050565b6020815260006102f360208301846102b4565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610339576103396102fa565b604052919050565b600082601f83011261035257600080fd5b813567ffffffffffffffff81111561036c5761036c6102fa565b61037f601f8201601f1916602001610310565b81815284602083860101111561039457600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156103c657600080fd5b833567ffffffffffffffff808211156103de57600080fd5b6103ea87838801610341565b945060209150818601358181111561040157600080fd5b61040d88828901610341565b94505060408601358181111561042257600080fd5b8601601f8101881361043357600080fd5b803582811115610445576104456102fa565b8060051b610454858201610310565b918252828101850191858101908b84111561046e57600080fd5b86850192505b838310156104aa5782358681111561048c5760008081fd5b61049a8d8983890101610341565b8352509186019190860190610474565b809750505050505050509250925092565b600080604083850312156104ce57600080fd5b82359150602083013567ffffffffffffffff8111156104ec57600080fd5b6104f885828601610341565b9150509250929050565b600181811c9082168061051657607f821691505b60208210810361053657634e487b7160e01b600052602260045260246000fd5b50919050565b6000835161054e818460208801610290565b835190830190610562818360208801610290565b01949350505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016105a157634e487b7160e01b600052601160045260246000fd5b5060010190565b601f82111561023257600081815260208120601f850160051c810160208610156105cf5750805b601f850160051c820191505b818110156105ee578281556001016105db565b505050505050565b815167ffffffffffffffff811115610610576106106102fa565b6106248161061e8454610502565b846105a8565b602080601f83116001811461065957600084156106415750858301515b600019600386901b1c1916600185901b1785556105ee565b600085815260208120601f198616915b8281101561068857888601518255948401946001909101908401610669565b50858210156106a65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6060815260006106c960608301866102b4565b6020838203818501526106dc82876102b4565b915083820360408501528185518084528284019150828160051b85010183880160005b8381101561072d57601f1987840301855261071b8383516102b4565b948601949250908501906001016106ff565b50909a995050505050505050505056fea26469706673582212206a437bb6f7522d58940df4d52d1f01c2f18dd749f568d814e30934984a7c4f7b64736f6c63430008110033",
}

// AoldaClientABI is the input ABI used to generate the binding from.
// Deprecated: Use AoldaClientMetaData.ABI instead.
var AoldaClientABI = AoldaClientMetaData.ABI

// AoldaClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AoldaClientMetaData.Bin instead.
var AoldaClientBin = AoldaClientMetaData.Bin

// DeployAoldaClient deploys a new Ethereum contract, binding an instance of AoldaClient to it.
func DeployAoldaClient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AoldaClient, error) {
	parsed, err := AoldaClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AoldaClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AoldaClient{AoldaClientCaller: AoldaClientCaller{contract: contract}, AoldaClientTransactor: AoldaClientTransactor{contract: contract}, AoldaClientFilterer: AoldaClientFilterer{contract: contract}}, nil
}

// AoldaClient is an auto generated Go binding around an Ethereum contract.
type AoldaClient struct {
	AoldaClientCaller     // Read-only binding to the contract
	AoldaClientTransactor // Write-only binding to the contract
	AoldaClientFilterer   // Log filterer for contract events
}

// AoldaClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type AoldaClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoldaClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AoldaClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoldaClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AoldaClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoldaClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AoldaClientSession struct {
	Contract     *AoldaClient      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AoldaClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AoldaClientCallerSession struct {
	Contract *AoldaClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AoldaClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AoldaClientTransactorSession struct {
	Contract     *AoldaClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AoldaClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type AoldaClientRaw struct {
	Contract *AoldaClient // Generic contract binding to access the raw methods on
}

// AoldaClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AoldaClientCallerRaw struct {
	Contract *AoldaClientCaller // Generic read-only contract binding to access the raw methods on
}

// AoldaClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AoldaClientTransactorRaw struct {
	Contract *AoldaClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAoldaClient creates a new instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClient(address common.Address, backend bind.ContractBackend) (*AoldaClient, error) {
	contract, err := bindAoldaClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AoldaClient{AoldaClientCaller: AoldaClientCaller{contract: contract}, AoldaClientTransactor: AoldaClientTransactor{contract: contract}, AoldaClientFilterer: AoldaClientFilterer{contract: contract}}, nil
}

// NewAoldaClientCaller creates a new read-only instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClientCaller(address common.Address, caller bind.ContractCaller) (*AoldaClientCaller, error) {
	contract, err := bindAoldaClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AoldaClientCaller{contract: contract}, nil
}

// NewAoldaClientTransactor creates a new write-only instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClientTransactor(address common.Address, transactor bind.ContractTransactor) (*AoldaClientTransactor, error) {
	contract, err := bindAoldaClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AoldaClientTransactor{contract: contract}, nil
}

// NewAoldaClientFilterer creates a new log filterer instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClientFilterer(address common.Address, filterer bind.ContractFilterer) (*AoldaClientFilterer, error) {
	contract, err := bindAoldaClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AoldaClientFilterer{contract: contract}, nil
}

// bindAoldaClient binds a generic wrapper to an already deployed contract.
func bindAoldaClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AoldaClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AoldaClient *AoldaClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AoldaClient.Contract.AoldaClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AoldaClient *AoldaClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AoldaClient.Contract.AoldaClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AoldaClient *AoldaClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AoldaClient.Contract.AoldaClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AoldaClient *AoldaClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AoldaClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AoldaClient *AoldaClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AoldaClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AoldaClient *AoldaClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AoldaClient.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x69843940.
//
// Solidity: function getValue(bytes32 signature) view returns(string)
func (_AoldaClient *AoldaClientCaller) GetValue(opts *bind.CallOpts, signature [32]byte) (string, error) {
	var out []interface{}
	err := _AoldaClient.contract.Call(opts, &out, "getValue", signature)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x69843940.
//
// Solidity: function getValue(bytes32 signature) view returns(string)
func (_AoldaClient *AoldaClientSession) GetValue(signature [32]byte) (string, error) {
	return _AoldaClient.Contract.GetValue(&_AoldaClient.CallOpts, signature)
}

// GetValue is a free data retrieval call binding the contract method 0x69843940.
//
// Solidity: function getValue(bytes32 signature) view returns(string)
func (_AoldaClient *AoldaClientCallerSession) GetValue(signature [32]byte) (string, error) {
	return _AoldaClient.Contract.GetValue(&_AoldaClient.CallOpts, signature)
}

// MakeSignature is a free data retrieval call binding the contract method 0x916c4759.
//
// Solidity: function makeSignature(string fileHash, string funtionName, string[] arguments) pure returns(bytes32)
func (_AoldaClient *AoldaClientCaller) MakeSignature(opts *bind.CallOpts, fileHash string, funtionName string, arguments []string) ([32]byte, error) {
	var out []interface{}
	err := _AoldaClient.contract.Call(opts, &out, "makeSignature", fileHash, funtionName, arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeSignature is a free data retrieval call binding the contract method 0x916c4759.
//
// Solidity: function makeSignature(string fileHash, string funtionName, string[] arguments) pure returns(bytes32)
func (_AoldaClient *AoldaClientSession) MakeSignature(fileHash string, funtionName string, arguments []string) ([32]byte, error) {
	return _AoldaClient.Contract.MakeSignature(&_AoldaClient.CallOpts, fileHash, funtionName, arguments)
}

// MakeSignature is a free data retrieval call binding the contract method 0x916c4759.
//
// Solidity: function makeSignature(string fileHash, string funtionName, string[] arguments) pure returns(bytes32)
func (_AoldaClient *AoldaClientCallerSession) MakeSignature(fileHash string, funtionName string, arguments []string) ([32]byte, error) {
	return _AoldaClient.Contract.MakeSignature(&_AoldaClient.CallOpts, fileHash, funtionName, arguments)
}

// CallAolda is a paid mutator transaction binding the contract method 0xcd6b7a26.
//
// Solidity: function callAolda(string fileHash, string funtionName, string[] arguments) returns()
func (_AoldaClient *AoldaClientTransactor) CallAolda(opts *bind.TransactOpts, fileHash string, funtionName string, arguments []string) (*types.Transaction, error) {
	return _AoldaClient.contract.Transact(opts, "callAolda", fileHash, funtionName, arguments)
}

// CallAolda is a paid mutator transaction binding the contract method 0xcd6b7a26.
//
// Solidity: function callAolda(string fileHash, string funtionName, string[] arguments) returns()
func (_AoldaClient *AoldaClientSession) CallAolda(fileHash string, funtionName string, arguments []string) (*types.Transaction, error) {
	return _AoldaClient.Contract.CallAolda(&_AoldaClient.TransactOpts, fileHash, funtionName, arguments)
}

// CallAolda is a paid mutator transaction binding the contract method 0xcd6b7a26.
//
// Solidity: function callAolda(string fileHash, string funtionName, string[] arguments) returns()
func (_AoldaClient *AoldaClientTransactorSession) CallAolda(fileHash string, funtionName string, arguments []string) (*types.Transaction, error) {
	return _AoldaClient.Contract.CallAolda(&_AoldaClient.TransactOpts, fileHash, funtionName, arguments)
}

// SetValue is a paid mutator transaction binding the contract method 0xc7bfebd6.
//
// Solidity: function setValue(bytes32 signature, string value) returns()
func (_AoldaClient *AoldaClientTransactor) SetValue(opts *bind.TransactOpts, signature [32]byte, value string) (*types.Transaction, error) {
	return _AoldaClient.contract.Transact(opts, "setValue", signature, value)
}

// SetValue is a paid mutator transaction binding the contract method 0xc7bfebd6.
//
// Solidity: function setValue(bytes32 signature, string value) returns()
func (_AoldaClient *AoldaClientSession) SetValue(signature [32]byte, value string) (*types.Transaction, error) {
	return _AoldaClient.Contract.SetValue(&_AoldaClient.TransactOpts, signature, value)
}

// SetValue is a paid mutator transaction binding the contract method 0xc7bfebd6.
//
// Solidity: function setValue(bytes32 signature, string value) returns()
func (_AoldaClient *AoldaClientTransactorSession) SetValue(signature [32]byte, value string) (*types.Transaction, error) {
	return _AoldaClient.Contract.SetValue(&_AoldaClient.TransactOpts, signature, value)
}

// AoldaClientCallAoldaIterator is returned from FilterCallAolda and is used to iterate over the raw logs and unpacked data for CallAolda events raised by the AoldaClient contract.
type AoldaClientCallAoldaIterator struct {
	Event *AoldaClientCallAolda // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AoldaClientCallAoldaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoldaClientCallAolda)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AoldaClientCallAolda)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AoldaClientCallAoldaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoldaClientCallAoldaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoldaClientCallAolda represents a CallAolda event raised by the AoldaClient contract.
type AoldaClientCallAolda struct {
	FileHash    string
	FuntionName string
	Arguments   []string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallAolda is a free log retrieval operation binding the contract event 0x1d37bd08d8b1245ea7f84485bab95b6726cc1fd59b42ebf52530669bb1137f62.
//
// Solidity: event CallAolda(string fileHash, string funtionName, string[] arguments)
func (_AoldaClient *AoldaClientFilterer) FilterCallAolda(opts *bind.FilterOpts) (*AoldaClientCallAoldaIterator, error) {

	logs, sub, err := _AoldaClient.contract.FilterLogs(opts, "CallAolda")
	if err != nil {
		return nil, err
	}
	return &AoldaClientCallAoldaIterator{contract: _AoldaClient.contract, event: "CallAolda", logs: logs, sub: sub}, nil
}

// WatchCallAolda is a free log subscription operation binding the contract event 0x1d37bd08d8b1245ea7f84485bab95b6726cc1fd59b42ebf52530669bb1137f62.
//
// Solidity: event CallAolda(string fileHash, string funtionName, string[] arguments)
func (_AoldaClient *AoldaClientFilterer) WatchCallAolda(opts *bind.WatchOpts, sink chan<- *AoldaClientCallAolda) (event.Subscription, error) {

	logs, sub, err := _AoldaClient.contract.WatchLogs(opts, "CallAolda")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoldaClientCallAolda)
				if err := _AoldaClient.contract.UnpackLog(event, "CallAolda", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallAolda is a log parse operation binding the contract event 0x1d37bd08d8b1245ea7f84485bab95b6726cc1fd59b42ebf52530669bb1137f62.
//
// Solidity: event CallAolda(string fileHash, string funtionName, string[] arguments)
func (_AoldaClient *AoldaClientFilterer) ParseCallAolda(log types.Log) (*AoldaClientCallAolda, error) {
	event := new(AoldaClientCallAolda)
	if err := _AoldaClient.contract.UnpackLog(event, "CallAolda", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
