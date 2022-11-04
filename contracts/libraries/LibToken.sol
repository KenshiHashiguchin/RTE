// SPDX-License-Identifier: CC0

pragma solidity ^0.8.0;

import {LibDiamond} from "./LibDiamond.sol";

library LibToken {
    bytes32 constant TOKEN_STORAGE_POSITION = keccak256("diamond.standard.token.storage");

    struct TokenStorage {
        uint256 mintAmount;
        address utilityToken;
    }

    function tokenStorage() internal pure returns (TokenStorage storage rs) {
        bytes32 position = TOKEN_STORAGE_POSITION;
        assembly {
            rs.slot := position
        }
    }
}