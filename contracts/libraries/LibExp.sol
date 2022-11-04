// SPDX-License-Identifier: CC0

pragma solidity ^0.8.0;

import {LibDiamond} from "./LibDiamond.sol";

library LibExp {
    bytes32 constant EXP_STORAGE_POSITION = keccak256("diamond.standard.exp.storage");

    struct ExpStorage {
        mapping(address => uint256) executionPoint;
        mapping(address => uint256) failurePoint;
    }

    function expStorage() internal pure returns (ExpStorage storage rs) {
        bytes32 position = EXP_STORAGE_POSITION;
        assembly {
            rs.slot := position
        }
    }
}