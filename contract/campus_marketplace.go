// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// CampusMarketplaceResponse is an auto generated low-level Go binding around an user-defined struct.
type CampusMarketplaceResponse struct {
	Success bool
	Code    *big.Int
	Message string
	Data    *big.Int
}

// CampusMarketplaceMetaData contains all meta data concerning the CampusMarketplace contract.
var CampusMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"emailHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"bindEmail\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structCampusMarketplace.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"}],\"name\":\"cancelTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structCampusMarketplace.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_imageHash\",\"type\":\"string\"}],\"name\":\"createListing\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structCampusMarketplace.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structCampusMarketplace.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"forceDelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"}],\"name\":\"DisputeRaised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellerWins\",\"type\":\"bool\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"emailHash\",\"type\":\"bytes32\"}],\"name\":\"EmailBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ListingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"ListingForceDelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"raiseDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"}],\"name\":\"releaseFunds\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structCampusMarketplace.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"releaseToSeller\",\"type\":\"bool\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setDebugMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"TransactionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransactionCreated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToEmail\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debugMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"emailToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"imageHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumCampusMarketplace.TransactionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"disputeInitiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"disputeReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"verifyLogin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CampusMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use CampusMarketplaceMetaData.ABI instead.
var CampusMarketplaceABI = CampusMarketplaceMetaData.ABI

// CampusMarketplace is an auto generated Go binding around an Ethereum contract.
type CampusMarketplace struct {
	CampusMarketplaceCaller     // Read-only binding to the contract
	CampusMarketplaceTransactor // Write-only binding to the contract
	CampusMarketplaceFilterer   // Log filterer for contract events
}

// CampusMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampusMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampusMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampusMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampusMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampusMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampusMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampusMarketplaceSession struct {
	Contract     *CampusMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CampusMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampusMarketplaceCallerSession struct {
	Contract *CampusMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CampusMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampusMarketplaceTransactorSession struct {
	Contract     *CampusMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CampusMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampusMarketplaceRaw struct {
	Contract *CampusMarketplace // Generic contract binding to access the raw methods on
}

// CampusMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampusMarketplaceCallerRaw struct {
	Contract *CampusMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// CampusMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampusMarketplaceTransactorRaw struct {
	Contract *CampusMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampusMarketplace creates a new instance of CampusMarketplace, bound to a specific deployed contract.
func NewCampusMarketplace(address common.Address, backend bind.ContractBackend) (*CampusMarketplace, error) {
	contract, err := bindCampusMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplace{CampusMarketplaceCaller: CampusMarketplaceCaller{contract: contract}, CampusMarketplaceTransactor: CampusMarketplaceTransactor{contract: contract}, CampusMarketplaceFilterer: CampusMarketplaceFilterer{contract: contract}}, nil
}

// NewCampusMarketplaceCaller creates a new read-only instance of CampusMarketplace, bound to a specific deployed contract.
func NewCampusMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*CampusMarketplaceCaller, error) {
	contract, err := bindCampusMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceCaller{contract: contract}, nil
}

// NewCampusMarketplaceTransactor creates a new write-only instance of CampusMarketplace, bound to a specific deployed contract.
func NewCampusMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*CampusMarketplaceTransactor, error) {
	contract, err := bindCampusMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceTransactor{contract: contract}, nil
}

// NewCampusMarketplaceFilterer creates a new log filterer instance of CampusMarketplace, bound to a specific deployed contract.
func NewCampusMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*CampusMarketplaceFilterer, error) {
	contract, err := bindCampusMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceFilterer{contract: contract}, nil
}

// bindCampusMarketplace binds a generic wrapper to an already deployed contract.
func bindCampusMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CampusMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampusMarketplace *CampusMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampusMarketplace.Contract.CampusMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampusMarketplace *CampusMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CampusMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampusMarketplace *CampusMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CampusMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampusMarketplace *CampusMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampusMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampusMarketplace *CampusMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampusMarketplace *CampusMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.contract.Transact(opts, method, params...)
}

// AddressToEmail is a free data retrieval call binding the contract method 0x31268449.
//
// Solidity: function addressToEmail(address ) view returns(bytes32)
func (_CampusMarketplace *CampusMarketplaceCaller) AddressToEmail(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "addressToEmail", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToEmail is a free data retrieval call binding the contract method 0x31268449.
//
// Solidity: function addressToEmail(address ) view returns(bytes32)
func (_CampusMarketplace *CampusMarketplaceSession) AddressToEmail(arg0 common.Address) ([32]byte, error) {
	return _CampusMarketplace.Contract.AddressToEmail(&_CampusMarketplace.CallOpts, arg0)
}

// AddressToEmail is a free data retrieval call binding the contract method 0x31268449.
//
// Solidity: function addressToEmail(address ) view returns(bytes32)
func (_CampusMarketplace *CampusMarketplaceCallerSession) AddressToEmail(arg0 common.Address) ([32]byte, error) {
	return _CampusMarketplace.Contract.AddressToEmail(&_CampusMarketplace.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CampusMarketplace *CampusMarketplaceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CampusMarketplace *CampusMarketplaceSession) Admin() (common.Address, error) {
	return _CampusMarketplace.Contract.Admin(&_CampusMarketplace.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CampusMarketplace *CampusMarketplaceCallerSession) Admin() (common.Address, error) {
	return _CampusMarketplace.Contract.Admin(&_CampusMarketplace.CallOpts)
}

// DebugMode is a free data retrieval call binding the contract method 0x1c9568be.
//
// Solidity: function debugMode() view returns(bool)
func (_CampusMarketplace *CampusMarketplaceCaller) DebugMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "debugMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DebugMode is a free data retrieval call binding the contract method 0x1c9568be.
//
// Solidity: function debugMode() view returns(bool)
func (_CampusMarketplace *CampusMarketplaceSession) DebugMode() (bool, error) {
	return _CampusMarketplace.Contract.DebugMode(&_CampusMarketplace.CallOpts)
}

// DebugMode is a free data retrieval call binding the contract method 0x1c9568be.
//
// Solidity: function debugMode() view returns(bool)
func (_CampusMarketplace *CampusMarketplaceCallerSession) DebugMode() (bool, error) {
	return _CampusMarketplace.Contract.DebugMode(&_CampusMarketplace.CallOpts)
}

// EmailToAddress is a free data retrieval call binding the contract method 0x54817832.
//
// Solidity: function emailToAddress(bytes32 ) view returns(address)
func (_CampusMarketplace *CampusMarketplaceCaller) EmailToAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "emailToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmailToAddress is a free data retrieval call binding the contract method 0x54817832.
//
// Solidity: function emailToAddress(bytes32 ) view returns(address)
func (_CampusMarketplace *CampusMarketplaceSession) EmailToAddress(arg0 [32]byte) (common.Address, error) {
	return _CampusMarketplace.Contract.EmailToAddress(&_CampusMarketplace.CallOpts, arg0)
}

// EmailToAddress is a free data retrieval call binding the contract method 0x54817832.
//
// Solidity: function emailToAddress(bytes32 ) view returns(address)
func (_CampusMarketplace *CampusMarketplaceCallerSession) EmailToAddress(arg0 [32]byte) (common.Address, error) {
	return _CampusMarketplace.Contract.EmailToAddress(&_CampusMarketplace.CallOpts, arg0)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_CampusMarketplace *CampusMarketplaceCaller) ListingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "listingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_CampusMarketplace *CampusMarketplaceSession) ListingCount() (*big.Int, error) {
	return _CampusMarketplace.Contract.ListingCount(&_CampusMarketplace.CallOpts)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_CampusMarketplace *CampusMarketplaceCallerSession) ListingCount() (*big.Int, error) {
	return _CampusMarketplace.Contract.ListingCount(&_CampusMarketplace.CallOpts)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 id, string title, string description, uint256 price, address seller, bool isActive, string imageHash)
func (_CampusMarketplace *CampusMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Title       string
	Description string
	Price       *big.Int
	Seller      common.Address
	IsActive    bool
	ImageHash   string
}, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Title       string
		Description string
		Price       *big.Int
		Seller      common.Address
		IsActive    bool
		ImageHash   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.ImageHash = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 id, string title, string description, uint256 price, address seller, bool isActive, string imageHash)
func (_CampusMarketplace *CampusMarketplaceSession) Listings(arg0 *big.Int) (struct {
	Id          *big.Int
	Title       string
	Description string
	Price       *big.Int
	Seller      common.Address
	IsActive    bool
	ImageHash   string
}, error) {
	return _CampusMarketplace.Contract.Listings(&_CampusMarketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 id, string title, string description, uint256 price, address seller, bool isActive, string imageHash)
func (_CampusMarketplace *CampusMarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	Id          *big.Int
	Title       string
	Description string
	Price       *big.Int
	Seller      common.Address
	IsActive    bool
	ImageHash   string
}, error) {
	return _CampusMarketplace.Contract.Listings(&_CampusMarketplace.CallOpts, arg0)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_CampusMarketplace *CampusMarketplaceCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_CampusMarketplace *CampusMarketplaceSession) TransactionCount() (*big.Int, error) {
	return _CampusMarketplace.Contract.TransactionCount(&_CampusMarketplace.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_CampusMarketplace *CampusMarketplaceCallerSession) TransactionCount() (*big.Int, error) {
	return _CampusMarketplace.Contract.TransactionCount(&_CampusMarketplace.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(uint256 id, uint256 listingId, address buyer, address seller, uint256 amount, uint8 status, address disputeInitiator, string disputeReason, uint256 createdAt)
func (_CampusMarketplace *CampusMarketplaceCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id               *big.Int
	ListingId        *big.Int
	Buyer            common.Address
	Seller           common.Address
	Amount           *big.Int
	Status           uint8
	DisputeInitiator common.Address
	DisputeReason    string
	CreatedAt        *big.Int
}, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Id               *big.Int
		ListingId        *big.Int
		Buyer            common.Address
		Seller           common.Address
		Amount           *big.Int
		Status           uint8
		DisputeInitiator common.Address
		DisputeReason    string
		CreatedAt        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ListingId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Seller = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.DisputeInitiator = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.DisputeReason = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(uint256 id, uint256 listingId, address buyer, address seller, uint256 amount, uint8 status, address disputeInitiator, string disputeReason, uint256 createdAt)
func (_CampusMarketplace *CampusMarketplaceSession) Transactions(arg0 *big.Int) (struct {
	Id               *big.Int
	ListingId        *big.Int
	Buyer            common.Address
	Seller           common.Address
	Amount           *big.Int
	Status           uint8
	DisputeInitiator common.Address
	DisputeReason    string
	CreatedAt        *big.Int
}, error) {
	return _CampusMarketplace.Contract.Transactions(&_CampusMarketplace.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(uint256 id, uint256 listingId, address buyer, address seller, uint256 amount, uint8 status, address disputeInitiator, string disputeReason, uint256 createdAt)
func (_CampusMarketplace *CampusMarketplaceCallerSession) Transactions(arg0 *big.Int) (struct {
	Id               *big.Int
	ListingId        *big.Int
	Buyer            common.Address
	Seller           common.Address
	Amount           *big.Int
	Status           uint8
	DisputeInitiator common.Address
	DisputeReason    string
	CreatedAt        *big.Int
}, error) {
	return _CampusMarketplace.Contract.Transactions(&_CampusMarketplace.CallOpts, arg0)
}

// VerifyLogin is a free data retrieval call binding the contract method 0xd53dbb88.
//
// Solidity: function verifyLogin(address user) view returns(bool)
func (_CampusMarketplace *CampusMarketplaceCaller) VerifyLogin(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _CampusMarketplace.contract.Call(opts, &out, "verifyLogin", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyLogin is a free data retrieval call binding the contract method 0xd53dbb88.
//
// Solidity: function verifyLogin(address user) view returns(bool)
func (_CampusMarketplace *CampusMarketplaceSession) VerifyLogin(user common.Address) (bool, error) {
	return _CampusMarketplace.Contract.VerifyLogin(&_CampusMarketplace.CallOpts, user)
}

// VerifyLogin is a free data retrieval call binding the contract method 0xd53dbb88.
//
// Solidity: function verifyLogin(address user) view returns(bool)
func (_CampusMarketplace *CampusMarketplaceCallerSession) VerifyLogin(user common.Address) (bool, error) {
	return _CampusMarketplace.Contract.VerifyLogin(&_CampusMarketplace.CallOpts, user)
}

// BindEmail is a paid mutator transaction binding the contract method 0x9f5e5062.
//
// Solidity: function bindEmail(bytes32 emailHash, bytes signature) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactor) BindEmail(opts *bind.TransactOpts, emailHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "bindEmail", emailHash, signature)
}

// BindEmail is a paid mutator transaction binding the contract method 0x9f5e5062.
//
// Solidity: function bindEmail(bytes32 emailHash, bytes signature) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceSession) BindEmail(emailHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.BindEmail(&_CampusMarketplace.TransactOpts, emailHash, signature)
}

// BindEmail is a paid mutator transaction binding the contract method 0x9f5e5062.
//
// Solidity: function bindEmail(bytes32 emailHash, bytes signature) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactorSession) BindEmail(emailHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.BindEmail(&_CampusMarketplace.TransactOpts, emailHash, signature)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x3380c0d8.
//
// Solidity: function cancelTransaction(uint256 _txId) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactor) CancelTransaction(opts *bind.TransactOpts, _txId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "cancelTransaction", _txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x3380c0d8.
//
// Solidity: function cancelTransaction(uint256 _txId) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceSession) CancelTransaction(_txId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CancelTransaction(&_CampusMarketplace.TransactOpts, _txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x3380c0d8.
//
// Solidity: function cancelTransaction(uint256 _txId) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactorSession) CancelTransaction(_txId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CancelTransaction(&_CampusMarketplace.TransactOpts, _txId)
}

// CreateListing is a paid mutator transaction binding the contract method 0x744409c6.
//
// Solidity: function createListing(string _title, string _description, uint256 _price, string _imageHash) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactor) CreateListing(opts *bind.TransactOpts, _title string, _description string, _price *big.Int, _imageHash string) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "createListing", _title, _description, _price, _imageHash)
}

// CreateListing is a paid mutator transaction binding the contract method 0x744409c6.
//
// Solidity: function createListing(string _title, string _description, uint256 _price, string _imageHash) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceSession) CreateListing(_title string, _description string, _price *big.Int, _imageHash string) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CreateListing(&_CampusMarketplace.TransactOpts, _title, _description, _price, _imageHash)
}

// CreateListing is a paid mutator transaction binding the contract method 0x744409c6.
//
// Solidity: function createListing(string _title, string _description, uint256 _price, string _imageHash) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactorSession) CreateListing(_title string, _description string, _price *big.Int, _imageHash string) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CreateListing(&_CampusMarketplace.TransactOpts, _title, _description, _price, _imageHash)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0x9be27a45.
//
// Solidity: function createTransaction(uint256 _listingId, bytes _signature) payable returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactor) CreateTransaction(opts *bind.TransactOpts, _listingId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "createTransaction", _listingId, _signature)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0x9be27a45.
//
// Solidity: function createTransaction(uint256 _listingId, bytes _signature) payable returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceSession) CreateTransaction(_listingId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CreateTransaction(&_CampusMarketplace.TransactOpts, _listingId, _signature)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0x9be27a45.
//
// Solidity: function createTransaction(uint256 _listingId, bytes _signature) payable returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactorSession) CreateTransaction(_listingId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.CreateTransaction(&_CampusMarketplace.TransactOpts, _listingId, _signature)
}

// ForceDelist is a paid mutator transaction binding the contract method 0xdc0a927c.
//
// Solidity: function forceDelist(uint256 listingId) returns()
func (_CampusMarketplace *CampusMarketplaceTransactor) ForceDelist(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "forceDelist", listingId)
}

// ForceDelist is a paid mutator transaction binding the contract method 0xdc0a927c.
//
// Solidity: function forceDelist(uint256 listingId) returns()
func (_CampusMarketplace *CampusMarketplaceSession) ForceDelist(listingId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.ForceDelist(&_CampusMarketplace.TransactOpts, listingId)
}

// ForceDelist is a paid mutator transaction binding the contract method 0xdc0a927c.
//
// Solidity: function forceDelist(uint256 listingId) returns()
func (_CampusMarketplace *CampusMarketplaceTransactorSession) ForceDelist(listingId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.ForceDelist(&_CampusMarketplace.TransactOpts, listingId)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x5aef573c.
//
// Solidity: function raiseDispute(uint256 transactionId, string reason) returns()
func (_CampusMarketplace *CampusMarketplaceTransactor) RaiseDispute(opts *bind.TransactOpts, transactionId *big.Int, reason string) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "raiseDispute", transactionId, reason)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x5aef573c.
//
// Solidity: function raiseDispute(uint256 transactionId, string reason) returns()
func (_CampusMarketplace *CampusMarketplaceSession) RaiseDispute(transactionId *big.Int, reason string) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.RaiseDispute(&_CampusMarketplace.TransactOpts, transactionId, reason)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x5aef573c.
//
// Solidity: function raiseDispute(uint256 transactionId, string reason) returns()
func (_CampusMarketplace *CampusMarketplaceTransactorSession) RaiseDispute(transactionId *big.Int, reason string) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.RaiseDispute(&_CampusMarketplace.TransactOpts, transactionId, reason)
}

// ReleaseFunds is a paid mutator transaction binding the contract method 0x4d68282f.
//
// Solidity: function releaseFunds(uint256 _txId) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactor) ReleaseFunds(opts *bind.TransactOpts, _txId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "releaseFunds", _txId)
}

// ReleaseFunds is a paid mutator transaction binding the contract method 0x4d68282f.
//
// Solidity: function releaseFunds(uint256 _txId) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceSession) ReleaseFunds(_txId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.ReleaseFunds(&_CampusMarketplace.TransactOpts, _txId)
}

// ReleaseFunds is a paid mutator transaction binding the contract method 0x4d68282f.
//
// Solidity: function releaseFunds(uint256 _txId) returns((bool,uint256,string,uint256))
func (_CampusMarketplace *CampusMarketplaceTransactorSession) ReleaseFunds(_txId *big.Int) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.ReleaseFunds(&_CampusMarketplace.TransactOpts, _txId)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x34b25ee2.
//
// Solidity: function resolveDispute(uint256 transactionId, bool releaseToSeller) returns()
func (_CampusMarketplace *CampusMarketplaceTransactor) ResolveDispute(opts *bind.TransactOpts, transactionId *big.Int, releaseToSeller bool) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "resolveDispute", transactionId, releaseToSeller)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x34b25ee2.
//
// Solidity: function resolveDispute(uint256 transactionId, bool releaseToSeller) returns()
func (_CampusMarketplace *CampusMarketplaceSession) ResolveDispute(transactionId *big.Int, releaseToSeller bool) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.ResolveDispute(&_CampusMarketplace.TransactOpts, transactionId, releaseToSeller)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x34b25ee2.
//
// Solidity: function resolveDispute(uint256 transactionId, bool releaseToSeller) returns()
func (_CampusMarketplace *CampusMarketplaceTransactorSession) ResolveDispute(transactionId *big.Int, releaseToSeller bool) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.ResolveDispute(&_CampusMarketplace.TransactOpts, transactionId, releaseToSeller)
}

// SetDebugMode is a paid mutator transaction binding the contract method 0x8811e191.
//
// Solidity: function setDebugMode(bool _enabled) returns()
func (_CampusMarketplace *CampusMarketplaceTransactor) SetDebugMode(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _CampusMarketplace.contract.Transact(opts, "setDebugMode", _enabled)
}

// SetDebugMode is a paid mutator transaction binding the contract method 0x8811e191.
//
// Solidity: function setDebugMode(bool _enabled) returns()
func (_CampusMarketplace *CampusMarketplaceSession) SetDebugMode(_enabled bool) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.SetDebugMode(&_CampusMarketplace.TransactOpts, _enabled)
}

// SetDebugMode is a paid mutator transaction binding the contract method 0x8811e191.
//
// Solidity: function setDebugMode(bool _enabled) returns()
func (_CampusMarketplace *CampusMarketplaceTransactorSession) SetDebugMode(_enabled bool) (*types.Transaction, error) {
	return _CampusMarketplace.Contract.SetDebugMode(&_CampusMarketplace.TransactOpts, _enabled)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CampusMarketplace *CampusMarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampusMarketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CampusMarketplace *CampusMarketplaceSession) Receive() (*types.Transaction, error) {
	return _CampusMarketplace.Contract.Receive(&_CampusMarketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CampusMarketplace *CampusMarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _CampusMarketplace.Contract.Receive(&_CampusMarketplace.TransactOpts)
}

// CampusMarketplaceDisputeRaisedIterator is returned from FilterDisputeRaised and is used to iterate over the raw logs and unpacked data for DisputeRaised events raised by the CampusMarketplace contract.
type CampusMarketplaceDisputeRaisedIterator struct {
	Event *CampusMarketplaceDisputeRaised // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceDisputeRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceDisputeRaised)
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
		it.Event = new(CampusMarketplaceDisputeRaised)
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
func (it *CampusMarketplaceDisputeRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceDisputeRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceDisputeRaised represents a DisputeRaised event raised by the CampusMarketplace contract.
type CampusMarketplaceDisputeRaised struct {
	TransactionId *big.Int
	Initiator     common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDisputeRaised is a free log retrieval operation binding the contract event 0x84a477df8a28a4276ca6dee4458a06c3015f30c477d9c949ede4e13ff8a552b4.
//
// Solidity: event DisputeRaised(uint256 indexed transactionId, address initiator)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterDisputeRaised(opts *bind.FilterOpts, transactionId []*big.Int) (*CampusMarketplaceDisputeRaisedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "DisputeRaised", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceDisputeRaisedIterator{contract: _CampusMarketplace.contract, event: "DisputeRaised", logs: logs, sub: sub}, nil
}

// WatchDisputeRaised is a free log subscription operation binding the contract event 0x84a477df8a28a4276ca6dee4458a06c3015f30c477d9c949ede4e13ff8a552b4.
//
// Solidity: event DisputeRaised(uint256 indexed transactionId, address initiator)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchDisputeRaised(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceDisputeRaised, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "DisputeRaised", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceDisputeRaised)
				if err := _CampusMarketplace.contract.UnpackLog(event, "DisputeRaised", log); err != nil {
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

// ParseDisputeRaised is a log parse operation binding the contract event 0x84a477df8a28a4276ca6dee4458a06c3015f30c477d9c949ede4e13ff8a552b4.
//
// Solidity: event DisputeRaised(uint256 indexed transactionId, address initiator)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseDisputeRaised(log types.Log) (*CampusMarketplaceDisputeRaised, error) {
	event := new(CampusMarketplaceDisputeRaised)
	if err := _CampusMarketplace.contract.UnpackLog(event, "DisputeRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the CampusMarketplace contract.
type CampusMarketplaceDisputeResolvedIterator struct {
	Event *CampusMarketplaceDisputeResolved // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceDisputeResolved)
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
		it.Event = new(CampusMarketplaceDisputeResolved)
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
func (it *CampusMarketplaceDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceDisputeResolved represents a DisputeResolved event raised by the CampusMarketplace contract.
type CampusMarketplaceDisputeResolved struct {
	TransactionId *big.Int
	SellerWins    bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x5a87909bff68caaaaf0b3fd9c74eeccc928832f879315e5c6fb7a73612f26c0c.
//
// Solidity: event DisputeResolved(uint256 indexed transactionId, bool sellerWins)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterDisputeResolved(opts *bind.FilterOpts, transactionId []*big.Int) (*CampusMarketplaceDisputeResolvedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "DisputeResolved", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceDisputeResolvedIterator{contract: _CampusMarketplace.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x5a87909bff68caaaaf0b3fd9c74eeccc928832f879315e5c6fb7a73612f26c0c.
//
// Solidity: event DisputeResolved(uint256 indexed transactionId, bool sellerWins)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceDisputeResolved, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "DisputeResolved", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceDisputeResolved)
				if err := _CampusMarketplace.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x5a87909bff68caaaaf0b3fd9c74eeccc928832f879315e5c6fb7a73612f26c0c.
//
// Solidity: event DisputeResolved(uint256 indexed transactionId, bool sellerWins)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseDisputeResolved(log types.Log) (*CampusMarketplaceDisputeResolved, error) {
	event := new(CampusMarketplaceDisputeResolved)
	if err := _CampusMarketplace.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceEmailBoundIterator is returned from FilterEmailBound and is used to iterate over the raw logs and unpacked data for EmailBound events raised by the CampusMarketplace contract.
type CampusMarketplaceEmailBoundIterator struct {
	Event *CampusMarketplaceEmailBound // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceEmailBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceEmailBound)
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
		it.Event = new(CampusMarketplaceEmailBound)
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
func (it *CampusMarketplaceEmailBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceEmailBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceEmailBound represents a EmailBound event raised by the CampusMarketplace contract.
type CampusMarketplaceEmailBound struct {
	User      common.Address
	EmailHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmailBound is a free log retrieval operation binding the contract event 0x7cc891ce81622bcecb9df2ac19ea3e1524930e4a63ec2393df782907742fd74a.
//
// Solidity: event EmailBound(address indexed user, bytes32 indexed emailHash)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterEmailBound(opts *bind.FilterOpts, user []common.Address, emailHash [][32]byte) (*CampusMarketplaceEmailBoundIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var emailHashRule []interface{}
	for _, emailHashItem := range emailHash {
		emailHashRule = append(emailHashRule, emailHashItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "EmailBound", userRule, emailHashRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceEmailBoundIterator{contract: _CampusMarketplace.contract, event: "EmailBound", logs: logs, sub: sub}, nil
}

// WatchEmailBound is a free log subscription operation binding the contract event 0x7cc891ce81622bcecb9df2ac19ea3e1524930e4a63ec2393df782907742fd74a.
//
// Solidity: event EmailBound(address indexed user, bytes32 indexed emailHash)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchEmailBound(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceEmailBound, user []common.Address, emailHash [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var emailHashRule []interface{}
	for _, emailHashItem := range emailHash {
		emailHashRule = append(emailHashRule, emailHashItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "EmailBound", userRule, emailHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceEmailBound)
				if err := _CampusMarketplace.contract.UnpackLog(event, "EmailBound", log); err != nil {
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

// ParseEmailBound is a log parse operation binding the contract event 0x7cc891ce81622bcecb9df2ac19ea3e1524930e4a63ec2393df782907742fd74a.
//
// Solidity: event EmailBound(address indexed user, bytes32 indexed emailHash)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseEmailBound(log types.Log) (*CampusMarketplaceEmailBound, error) {
	event := new(CampusMarketplaceEmailBound)
	if err := _CampusMarketplace.contract.UnpackLog(event, "EmailBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceFundReleasedIterator is returned from FilterFundReleased and is used to iterate over the raw logs and unpacked data for FundReleased events raised by the CampusMarketplace contract.
type CampusMarketplaceFundReleasedIterator struct {
	Event *CampusMarketplaceFundReleased // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceFundReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceFundReleased)
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
		it.Event = new(CampusMarketplaceFundReleased)
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
func (it *CampusMarketplaceFundReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceFundReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceFundReleased represents a FundReleased event raised by the CampusMarketplace contract.
type CampusMarketplaceFundReleased struct {
	TxId   *big.Int
	Seller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundReleased is a free log retrieval operation binding the contract event 0x182afa9d181d69bec024c0ca773a3858bf5178facff83b16197d0bb1735f543f.
//
// Solidity: event FundReleased(uint256 indexed txId, address seller, uint256 amount)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterFundReleased(opts *bind.FilterOpts, txId []*big.Int) (*CampusMarketplaceFundReleasedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "FundReleased", txIdRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceFundReleasedIterator{contract: _CampusMarketplace.contract, event: "FundReleased", logs: logs, sub: sub}, nil
}

// WatchFundReleased is a free log subscription operation binding the contract event 0x182afa9d181d69bec024c0ca773a3858bf5178facff83b16197d0bb1735f543f.
//
// Solidity: event FundReleased(uint256 indexed txId, address seller, uint256 amount)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchFundReleased(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceFundReleased, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "FundReleased", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceFundReleased)
				if err := _CampusMarketplace.contract.UnpackLog(event, "FundReleased", log); err != nil {
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

// ParseFundReleased is a log parse operation binding the contract event 0x182afa9d181d69bec024c0ca773a3858bf5178facff83b16197d0bb1735f543f.
//
// Solidity: event FundReleased(uint256 indexed txId, address seller, uint256 amount)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseFundReleased(log types.Log) (*CampusMarketplaceFundReleased, error) {
	event := new(CampusMarketplaceFundReleased)
	if err := _CampusMarketplace.contract.UnpackLog(event, "FundReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceListingCreatedIterator is returned from FilterListingCreated and is used to iterate over the raw logs and unpacked data for ListingCreated events raised by the CampusMarketplace contract.
type CampusMarketplaceListingCreatedIterator struct {
	Event *CampusMarketplaceListingCreated // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceListingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceListingCreated)
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
		it.Event = new(CampusMarketplaceListingCreated)
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
func (it *CampusMarketplaceListingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceListingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceListingCreated represents a ListingCreated event raised by the CampusMarketplace contract.
type CampusMarketplaceListingCreated struct {
	Id     *big.Int
	Seller common.Address
	Title  string
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterListingCreated is a free log retrieval operation binding the contract event 0x6b7a3194eaddce2be5200d10785a15a5cd14a78b827cf902f4d23c35e6bd73f8.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, string title, uint256 price)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterListingCreated(opts *bind.FilterOpts, id []*big.Int, seller []common.Address) (*CampusMarketplaceListingCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "ListingCreated", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceListingCreatedIterator{contract: _CampusMarketplace.contract, event: "ListingCreated", logs: logs, sub: sub}, nil
}

// WatchListingCreated is a free log subscription operation binding the contract event 0x6b7a3194eaddce2be5200d10785a15a5cd14a78b827cf902f4d23c35e6bd73f8.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, string title, uint256 price)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchListingCreated(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceListingCreated, id []*big.Int, seller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "ListingCreated", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceListingCreated)
				if err := _CampusMarketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
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

// ParseListingCreated is a log parse operation binding the contract event 0x6b7a3194eaddce2be5200d10785a15a5cd14a78b827cf902f4d23c35e6bd73f8.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, string title, uint256 price)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseListingCreated(log types.Log) (*CampusMarketplaceListingCreated, error) {
	event := new(CampusMarketplaceListingCreated)
	if err := _CampusMarketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceListingForceDelistedIterator is returned from FilterListingForceDelisted and is used to iterate over the raw logs and unpacked data for ListingForceDelisted events raised by the CampusMarketplace contract.
type CampusMarketplaceListingForceDelistedIterator struct {
	Event *CampusMarketplaceListingForceDelisted // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceListingForceDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceListingForceDelisted)
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
		it.Event = new(CampusMarketplaceListingForceDelisted)
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
func (it *CampusMarketplaceListingForceDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceListingForceDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceListingForceDelisted represents a ListingForceDelisted event raised by the CampusMarketplace contract.
type CampusMarketplaceListingForceDelisted struct {
	ListingId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingForceDelisted is a free log retrieval operation binding the contract event 0xaf887738eeadb5276aa770054240678d19beef558a6d1c3bba0dc412f7a86a8f.
//
// Solidity: event ListingForceDelisted(uint256 indexed listingId)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterListingForceDelisted(opts *bind.FilterOpts, listingId []*big.Int) (*CampusMarketplaceListingForceDelistedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "ListingForceDelisted", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceListingForceDelistedIterator{contract: _CampusMarketplace.contract, event: "ListingForceDelisted", logs: logs, sub: sub}, nil
}

// WatchListingForceDelisted is a free log subscription operation binding the contract event 0xaf887738eeadb5276aa770054240678d19beef558a6d1c3bba0dc412f7a86a8f.
//
// Solidity: event ListingForceDelisted(uint256 indexed listingId)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchListingForceDelisted(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceListingForceDelisted, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "ListingForceDelisted", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceListingForceDelisted)
				if err := _CampusMarketplace.contract.UnpackLog(event, "ListingForceDelisted", log); err != nil {
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

// ParseListingForceDelisted is a log parse operation binding the contract event 0xaf887738eeadb5276aa770054240678d19beef558a6d1c3bba0dc412f7a86a8f.
//
// Solidity: event ListingForceDelisted(uint256 indexed listingId)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseListingForceDelisted(log types.Log) (*CampusMarketplaceListingForceDelisted, error) {
	event := new(CampusMarketplaceListingForceDelisted)
	if err := _CampusMarketplace.contract.UnpackLog(event, "ListingForceDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceTransactionCancelledIterator is returned from FilterTransactionCancelled and is used to iterate over the raw logs and unpacked data for TransactionCancelled events raised by the CampusMarketplace contract.
type CampusMarketplaceTransactionCancelledIterator struct {
	Event *CampusMarketplaceTransactionCancelled // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceTransactionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceTransactionCancelled)
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
		it.Event = new(CampusMarketplaceTransactionCancelled)
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
func (it *CampusMarketplaceTransactionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceTransactionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceTransactionCancelled represents a TransactionCancelled event raised by the CampusMarketplace contract.
type CampusMarketplaceTransactionCancelled struct {
	TxId *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransactionCancelled is a free log retrieval operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: event TransactionCancelled(uint256 indexed txId)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterTransactionCancelled(opts *bind.FilterOpts, txId []*big.Int) (*CampusMarketplaceTransactionCancelledIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "TransactionCancelled", txIdRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceTransactionCancelledIterator{contract: _CampusMarketplace.contract, event: "TransactionCancelled", logs: logs, sub: sub}, nil
}

// WatchTransactionCancelled is a free log subscription operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: event TransactionCancelled(uint256 indexed txId)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchTransactionCancelled(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceTransactionCancelled, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "TransactionCancelled", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceTransactionCancelled)
				if err := _CampusMarketplace.contract.UnpackLog(event, "TransactionCancelled", log); err != nil {
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

// ParseTransactionCancelled is a log parse operation binding the contract event 0x956fb32199d8b882b2039a14e1be35ab14f7a80b9089fc223f14b43937173e60.
//
// Solidity: event TransactionCancelled(uint256 indexed txId)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseTransactionCancelled(log types.Log) (*CampusMarketplaceTransactionCancelled, error) {
	event := new(CampusMarketplaceTransactionCancelled)
	if err := _CampusMarketplace.contract.UnpackLog(event, "TransactionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampusMarketplaceTransactionCreatedIterator is returned from FilterTransactionCreated and is used to iterate over the raw logs and unpacked data for TransactionCreated events raised by the CampusMarketplace contract.
type CampusMarketplaceTransactionCreatedIterator struct {
	Event *CampusMarketplaceTransactionCreated // Event containing the contract specifics and raw log

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
func (it *CampusMarketplaceTransactionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampusMarketplaceTransactionCreated)
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
		it.Event = new(CampusMarketplaceTransactionCreated)
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
func (it *CampusMarketplaceTransactionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampusMarketplaceTransactionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampusMarketplaceTransactionCreated represents a TransactionCreated event raised by the CampusMarketplace contract.
type CampusMarketplaceTransactionCreated struct {
	TxId      *big.Int
	ListingId *big.Int
	Buyer     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreated is a free log retrieval operation binding the contract event 0xa8bfac72285d2c7a4231a709d3dee850d0f0d204d4b18a725d6e6b8597ee8ad4.
//
// Solidity: event TransactionCreated(uint256 indexed txId, uint256 indexed listingId, address indexed buyer, uint256 amount)
func (_CampusMarketplace *CampusMarketplaceFilterer) FilterTransactionCreated(opts *bind.FilterOpts, txId []*big.Int, listingId []*big.Int, buyer []common.Address) (*CampusMarketplaceTransactionCreatedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CampusMarketplace.contract.FilterLogs(opts, "TransactionCreated", txIdRule, listingIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &CampusMarketplaceTransactionCreatedIterator{contract: _CampusMarketplace.contract, event: "TransactionCreated", logs: logs, sub: sub}, nil
}

// WatchTransactionCreated is a free log subscription operation binding the contract event 0xa8bfac72285d2c7a4231a709d3dee850d0f0d204d4b18a725d6e6b8597ee8ad4.
//
// Solidity: event TransactionCreated(uint256 indexed txId, uint256 indexed listingId, address indexed buyer, uint256 amount)
func (_CampusMarketplace *CampusMarketplaceFilterer) WatchTransactionCreated(opts *bind.WatchOpts, sink chan<- *CampusMarketplaceTransactionCreated, txId []*big.Int, listingId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CampusMarketplace.contract.WatchLogs(opts, "TransactionCreated", txIdRule, listingIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampusMarketplaceTransactionCreated)
				if err := _CampusMarketplace.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
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

// ParseTransactionCreated is a log parse operation binding the contract event 0xa8bfac72285d2c7a4231a709d3dee850d0f0d204d4b18a725d6e6b8597ee8ad4.
//
// Solidity: event TransactionCreated(uint256 indexed txId, uint256 indexed listingId, address indexed buyer, uint256 amount)
func (_CampusMarketplace *CampusMarketplaceFilterer) ParseTransactionCreated(log types.Log) (*CampusMarketplaceTransactionCreated, error) {
	event := new(CampusMarketplaceTransactionCreated)
	if err := _CampusMarketplace.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
