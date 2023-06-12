// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.7;

library MerkleProof {
    function getMerkleRoot(bytes[][] memory elements) internal pure returns (bytes32) {
        uint256 len = elements.length;

        if (len == 0) {
            return bytes32(0);
        }

        bytes32[] memory hashes = new bytes32[](len);
        for (uint256 i = 0; i < len; i++) {
            hashes[i] = keccak256(abi.encode(elements[i]));
        }

        uint256 j = 0;
        for (uint256 step = 1; step < len; step *= 2) {
            for (uint256 i = 0; i < len - step; i += step * 2) {
                bytes32 left = hashes[i + j];
                bytes32 right = hashes[i + j + step];
                hashes[i + j] = keccak256(abi.encodePacked(left, right));
            }
            j += step;
        }

        return hashes[0];
    }

    function verify(bytes32[] memory proof, bytes32 root, bytes32 leaf) internal pure returns (bool) {
        bytes32 el;
        bytes32 h = leaf;

        for (uint256 i = 0; i < proof.length; i += 1) {
            el = proof[i];

            if (h < el) {
                h = keccak256(abi.encodePacked(h, el));
            } else {
                h = keccak256(abi.encodePacked(el, h));
            }
        }

        return h == root;
    }
}
