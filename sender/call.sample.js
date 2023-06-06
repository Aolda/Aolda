const { ethers } = require("ethers");

const endPoint = "HTTP://127.0.0.1:7545";
const Address = "0x2EB0558f2189bc84cE7C4CB76814335f3566c3B7";
const Abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "string",
        name: "fileHash",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "funtionName",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string[]",
        name: "arguments",
        type: "string[]",
      },
    ],
    name: "CallAolda",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "fileHash",
        type: "string",
      },
      {
        internalType: "string",
        name: "funtionName",
        type: "string",
      },
      {
        internalType: "string[]",
        name: "arguments",
        type: "string[]",
      },
    ],
    name: "callAolda",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "signature",
        type: "bytes32",
      },
    ],
    name: "getValue",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "fileHash",
        type: "string",
      },
      {
        internalType: "string",
        name: "funtionName",
        type: "string",
      },
      {
        internalType: "string[]",
        name: "arguments",
        type: "string[]",
      },
    ],
    name: "makeSignature",
    outputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    stateMutability: "pure",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "signature",
        type: "bytes32",
      },
      {
        internalType: "string",
        name: "value",
        type: "string",
      },
    ],
    name: "setValue",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

async function main() {
  const provider = new ethers.providers.JsonRpcProvider(endPoint);

  const signer = new ethers.Wallet(
    "9842c29eea3ceb3d1ac3ba6b5e75929d9ec8a4d97cb52d12fd0b382b5dd53cb3",
    provider
  );

  const Contract = new ethers.Contract(Address, Abi, signer);
  const estimatedGasLimit = await Contract.estimateGas.callAolda(
    "add.js",
    "add",
    ["1", "2"]
  );
  const approveTxUnsigned = await Contract.populateTransaction.callAolda(
    "add.js",
    "add",
    ["1", "2"]
  );
  approveTxUnsigned.chainId = 1337; // chainId 1 for Ethereum mainnet
  approveTxUnsigned.gasLimit = estimatedGasLimit;
  approveTxUnsigned.gasPrice = await provider.getGasPrice();
  // approveTxUnsigned.value = 1000;
  approveTxUnsigned.nonce = await provider.getTransactionCount(
    "0x2EB0558f2189bc84cE7C4CB76814335f3566c3B7"
  );
  const approveTxSigned = await signer.signTransaction(approveTxUnsigned);
  const submittedTx = await provider.sendTransaction(approveTxSigned);
  const approveReceipt = await submittedTx.wait();
  if (approveReceipt.status === 0)
    throw new Error("Approve transaction failed");
}

main();
