// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.7;
import "./MerkleProof.sol";

contract EthereumBlock {
    struct BlockHeader {
        bytes32 BlockHash;
        bytes32 PreviousHash;
        bytes32 Minor;
        uint difficulty;
        uint timestamp;
        uint number;
        bytes32 nonce;
    }

    struct BlockData {
        bytes32 merkleRoot; // 머클 트리 루트 해시
        bytes[][] transactions;
    }

    BlockHeader public header;
    BlockData public data;

    constructor(
        bytes32 _BlockHash,
        bytes32 _PreviousHash,
        bytes32 _Minor,
        uint _difficulty,
        uint _timestamp,
        uint _number,
        bytes32 _nonce,
        bytes[][] memory _transactions
    ) {
        header = BlockHeader({
            BlockHash: _BlockHash,
            PreviousHash: _PreviousHash,
            Minor: _Minor,
            difficulty: _difficulty,
            timestamp: _timestamp,
            number: _number,
            nonce: _nonce
        });

        data = BlockData({
            merkleRoot: bytes32(0),
            transactions: _transactions
        });

        bytes[][] memory transactionHashes = new bytes[][](_transactions.length);
        for (uint i = 0; i < _transactions.length; i++) {
            bytes32[] memory hashes = new bytes32[](_transactions[i].length);
        for (uint j = 0; j < _transactions[i].length; j++) {
                hashes[j] = keccak256(abi.encode(_transactions[i][j]));
            }
        }
        data.merkleRoot = MerkleProof.getMerkleRoot(transactionHashes);
    }

    function mineBlock(bytes32 _nonce) public {
        uint256 nonce = uint256(_nonce);
        while (uint256(keccak256(abi.encodePacked(header.BlockHash, header.PreviousHash, header.Minor, header.difficulty, header.timestamp, header.number, nonce))) >= 2 ** (256 - header.difficulty)) {
            nonce++;
        }

        header.nonce = bytes32(nonce);
        header.BlockHash = keccak256(abi.encodePacked(header.PreviousHash, header.Minor, header.difficulty, header.timestamp, header.number, header.nonce));
    }
}
