package proxy

const (
	ERC20ProxyBin = `608060405234801561000f575f5ffd5b5060405161122d38038061122d8339818101604052810190610031919061030e565b5f6100418261008160201b60201c565b90507ffcc2869d2ba3df58e629ad93fc7e1e2c05a9f6ce6223e40ff0041b7670ca1a9681604051610072919061036f565b60405180910390a150506104bc565b5f5f61010173ffffffffffffffffffffffffffffffffffffffff16836040516024016100ad91906103da565b6040516020818303038152906040527ff2c298be000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610137919061043e565b5f60405180830381855af49150503d805f811461016f576040519150601f19603f3d011682016040523d82523d5f602084013e610174565b606091505b50509050806101b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101af9061049e565b60405180910390fd5b80915050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610220826101da565b810181811067ffffffffffffffff8211171561023f5761023e6101ea565b5b80604052505050565b5f6102516101c1565b905061025d8282610217565b919050565b5f67ffffffffffffffff82111561027c5761027b6101ea565b5b610285826101da565b9050602081019050919050565b8281835e5f83830152505050565b5f6102b26102ad84610262565b610248565b9050828152602081018484840111156102ce576102cd6101d6565b5b6102d9848285610292565b509392505050565b5f82601f8301126102f5576102f46101d2565b5b81516103058482602086016102a0565b91505092915050565b5f60208284031215610323576103226101ca565b5b5f82015167ffffffffffffffff8111156103405761033f6101ce565b5b61034c848285016102e1565b91505092915050565b5f8115159050919050565b61036981610355565b82525050565b5f6020820190506103825f830184610360565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f6103ac82610388565b6103b68185610392565b93506103c6818560208601610292565b6103cf816101da565b840191505092915050565b5f6020820190508181035f8301526103f281846103a2565b905092915050565b5f81519050919050565b5f81905092915050565b5f610418826103fa565b6104228185610404565b9350610432818560208601610292565b80840191505092915050565b5f610449828461040e565b915081905092915050565b7f45524332303a20726567697374726174696f6e206661696c65640000000000005f82015250565b5f610488601a83610392565b915061049382610454565b602082019050919050565b5f6020820190508181035f8301526104b58161047c565b9050919050565b610d64806104c95f395ff3fe60806040526004361061004d575f3560e01c8063095ea7b31461006157806370a082311461009d578063a9059cbb146100d9578063d3ed45f814610115578063dd62ed3e1461013f57610054565b3661005457005b61005f61010161017b565b005b34801561006c575f5ffd5b50610087600480360381019061008291906108c0565b6101fe565b6040516100949190610918565b60405180910390f35b3480156100a8575f5ffd5b506100c360048036038101906100be9190610931565b6103bc565b6040516100d0919061096b565b60405180910390f35b3480156100e4575f5ffd5b506100ff60048036038101906100fa91906108c0565b610512565b60405161010c9190610918565b60405180910390f35b348015610120575f5ffd5b506101296106d0565b6040516101369190610993565b60405180910390f35b34801561014a575f5ffd5b50610165600480360381019061016091906109ac565b6106d6565b604051610172919061096b565b60405180910390f35b5f5f8273ffffffffffffffffffffffffffffffffffffffff165f366040516101a4929190610a26565b5f60405180830381855af49150503d805f81146101dc576040519150601f19603f3d011682016040523d82523d5f602084013e6101e1565b606091505b5091509150815f81146101f657815160208301f35b815160208301fd5b5f5f5f61010173ffffffffffffffffffffffffffffffffffffffff16858560405160240161022d929190610a3e565b6040516020818303038152906040527f095ea7b3000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102b79190610aad565b5f60405180830381855af49150503d805f81146102ef576040519150601f19603f3d011682016040523d82523d5f602084013e6102f4565b606091505b509150915081610339576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033090610b1d565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92586604051610396919061096b565b60405180910390a3808060200190518101906103b29190610b65565b9250505092915050565b5f5f5f61010173ffffffffffffffffffffffffffffffffffffffff16846040516024016103e99190610993565b6040516020818303038152906040527f70a08231000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516104739190610aad565b5f60405180830381855af49150503d805f81146104ab576040519150601f19603f3d011682016040523d82523d5f602084013e6104b0565b606091505b5091509150816104f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ec90610bda565b60405180910390fd5b808060200190518101906105099190610c0c565b92505050919050565b5f5f5f61010173ffffffffffffffffffffffffffffffffffffffff168585604051602401610541929190610a3e565b6040516020818303038152906040527fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516105cb9190610aad565b5f60405180830381855af49150503d805f8114610603576040519150601f19603f3d011682016040523d82523d5f602084013e610608565b606091505b50915091508161064d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064490610c81565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040516106aa919061096b565b60405180910390a3808060200190518101906106c69190610b65565b9250505092915050565b61010181565b5f5f5f61010173ffffffffffffffffffffffffffffffffffffffff168585604051602401610705929190610c9f565b6040516020818303038152906040527fdd62ed3e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161078f9190610aad565b5f60405180830381855af49150503d805f81146107c7576040519150601f19603f3d011682016040523d82523d5f602084013e6107cc565b606091505b509150915081610811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080890610d10565b60405180910390fd5b808060200190518101906108259190610c0c565b9250505092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61085c82610833565b9050919050565b61086c81610852565b8114610876575f5ffd5b50565b5f8135905061088781610863565b92915050565b5f819050919050565b61089f8161088d565b81146108a9575f5ffd5b50565b5f813590506108ba81610896565b92915050565b5f5f604083850312156108d6576108d561082f565b5b5f6108e385828601610879565b92505060206108f4858286016108ac565b9150509250929050565b5f8115159050919050565b610912816108fe565b82525050565b5f60208201905061092b5f830184610909565b92915050565b5f602082840312156109465761094561082f565b5b5f61095384828501610879565b91505092915050565b6109658161088d565b82525050565b5f60208201905061097e5f83018461095c565b92915050565b61098d81610852565b82525050565b5f6020820190506109a65f830184610984565b92915050565b5f5f604083850312156109c2576109c161082f565b5b5f6109cf85828601610879565b92505060206109e085828601610879565b9150509250929050565b5f81905092915050565b828183375f83830152505050565b5f610a0d83856109ea565b9350610a1a8385846109f4565b82840190509392505050565b5f610a32828486610a02565b91508190509392505050565b5f604082019050610a515f830185610984565b610a5e602083018461095c565b9392505050565b5f81519050919050565b8281835e5f83830152505050565b5f610a8782610a65565b610a9181856109ea565b9350610aa1818560208601610a6f565b80840191505092915050565b5f610ab88284610a7d565b915081905092915050565b5f82825260208201905092915050565b7f45524332303a20617070726f7665206661696c656400000000000000000000005f82015250565b5f610b07601583610ac3565b9150610b1282610ad3565b602082019050919050565b5f6020820190508181035f830152610b3481610afb565b9050919050565b610b44816108fe565b8114610b4e575f5ffd5b50565b5f81519050610b5f81610b3b565b92915050565b5f60208284031215610b7a57610b7961082f565b5b5f610b8784828501610b51565b91505092915050565b7f45524332303a2064656c65676174652063616c6c206661696c656400000000005f82015250565b5f610bc4601b83610ac3565b9150610bcf82610b90565b602082019050919050565b5f6020820190508181035f830152610bf181610bb8565b9050919050565b5f81519050610c0681610896565b92915050565b5f60208284031215610c2157610c2061082f565b5b5f610c2e84828501610bf8565b91505092915050565b7f45524332303a207472616e73666572206661696c6564000000000000000000005f82015250565b5f610c6b601683610ac3565b9150610c7682610c37565b602082019050919050565b5f6020820190508181035f830152610c9881610c5f565b9050919050565b5f604082019050610cb25f830185610984565b610cbf6020830184610984565b9392505050565b7f45524332303a20616c6c6f77616e6365207175657279206661696c65640000005f82015250565b5f610cfa601d83610ac3565b9150610d0582610cc6565b602082019050919050565b5f6020820190508181035f830152610d2781610cee565b905091905056fea2646970667358221220115ee257bb0d95108c8e6dd77167504375849ff19e41c29cb3b7dd604c88f7c964736f6c634300081b0033`
	ERC20ProxyAbi = `[
		{
		  "inputs": [
			{
			  "internalType": "string",
			  "name": "denom",
			  "type": "string"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "constructor"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "owner",
			  "type": "address"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "spender",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "uint256",
			  "name": "value",
			  "type": "uint256"
			}
		  ],
		  "name": "Approval",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bool",
			  "name": "success",
			  "type": "bool"
			}
		  ],
		  "name": "RegistrationSuccess",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "from",
			  "type": "address"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "to",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "uint256",
			  "name": "value",
			  "type": "uint256"
			}
		  ],
		  "name": "Transfer",
		  "type": "event"
		},
		{
		  "stateMutability": "payable",
		  "type": "fallback"
		},
		{
		  "inputs": [],
		  "name": "ERC20_PRECOMPILED_ADDRESS",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "owner",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "spender",
			  "type": "address"
			}
		  ],
		  "name": "allowance",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "spender",
			  "type": "address"
			},
			{
			  "internalType": "uint256",
			  "name": "amount",
			  "type": "uint256"
			}
		  ],
		  "name": "approve",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "account",
			  "type": "address"
			}
		  ],
		  "name": "balanceOf",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "string",
			  "name": "denom",
			  "type": "string"
			}
		  ],
		  "name": "register",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "to",
			  "type": "address"
			},
			{
			  "internalType": "uint256",
			  "name": "amount",
			  "type": "uint256"
			}
		  ],
		  "name": "transfer",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "stateMutability": "payable",
		  "type": "receive"
		}
	  ]`
)
