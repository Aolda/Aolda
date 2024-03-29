// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract AoldaClient {
  event CallAolda(string fileHash, string funtionName, string[] arguments);
  mapping(bytes32 => string) private result;

  function callAolda(string memory fileHash, string memory funtionName, string[] memory arguments) public {
    emit CallAolda(fileHash, funtionName, arguments);
  }

  function setValue(bytes32 signature ,string memory value) public{
    result[signature]=value;
  }

  function getValue(bytes32 signature) public view returns(string memory){
    return result[signature];
  }

  function makeSignature(string memory fileHash, string memory funtionName, string[] memory arguments) public pure returns(bytes32) {
    string memory concated = string.concat(fileHash,funtionName);
    for(uint256 i=0;i<arguments.length;i++){
      concated=string.concat(concated,arguments[i]);
    }
    // return concated;
    return keccak256(abi.encode(concated));
    // return stringToBytes32(concated);
  }

  function stringToBytes32(string memory source) internal pure returns (bytes32 convert) {
        // require(bytes(source).length <= 32); // causes error
        // but string have to be max 32 chars
        // https://ethereum.stackexchange.com/questions/9603/understanding-mload-assembly-function
        // http://solidity.readthedocs.io/en/latest/assembly.html
      assembly {
      convert := mload(add(source, 32))
      }
  }//
}