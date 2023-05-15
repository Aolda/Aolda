const { ethers } = require("ethers");

const endPoint= "HTTP://127.0.0.1:7545";
const depositAddress = "0xBFaC0a7C8FFEbbaa4f6491dED70e786b5524a9f6";
const depositAbi = [
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "string",
          "name": "fileHash",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "funtionName",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "string[]",
          "name": "arguments",
          "type": "string[]"
        }
      ],
      "name": "CallAolda",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "fileHash",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "funtionName",
          "type": "string"
        },
        {
          "internalType": "string[]",
          "name": "arguments",
          "type": "string[]"
        }
      ],
      "name": "callAolda",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "signature",
          "type": "bytes32"
        }
      ],
      "name": "getValue",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "fileHash",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "funtionName",
          "type": "string"
        },
        {
          "internalType": "string[]",
          "name": "arguments",
          "type": "string[]"
        }
      ],
      "name": "makeSignature",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "signature",
          "type": "bytes32"
        },
        {
          "internalType": "string",
          "name": "value",
          "type": "string"
        }
      ],
      "name": "setValue",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]

async function main(){
  const provider = new ethers.providers.JsonRpcProvider(endPoint);

  const signer = new ethers.Wallet("6ade8a9ba4d89d8074899b71c98dbd92d2adc31a7b7b9014124482e4eabc35c6", provider);

  const depositContract = new ethers.Contract(depositAddress, depositAbi,signer);
  const estimatedGasLimit = await depositContract.estimateGas.callAolda("add.js","add",["1","2"]);
  const approveTxUnsigned = await depositContract.populateTransaction.callAolda("add.js","add",["1","2"]);
  approveTxUnsigned.chainId = 1337; // chainId 1 for Ethereum mainnet
  approveTxUnsigned.gasLimit = estimatedGasLimit;
  approveTxUnsigned.gasPrice = await provider.getGasPrice();
  approveTxUnsigned.value = 1000;
  approveTxUnsigned.nonce = await provider.getTransactionCount("0x3ED40b5CD09dF63F7cf5b7859Efee85a5BCBFCaf");
  const approveTxSigned = await signer.signTransaction(approveTxUnsigned);
  const submittedTx = await provider.sendTransaction(approveTxSigned);
  const approveReceipt = await submittedTx.wait();
  if (approveReceipt.status === 0)
      throw new Error("Approve transaction failed");
}

main();