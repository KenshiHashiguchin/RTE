// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {LibDiamond} from "./LibDiamond.sol";

library LibMerchant {
    bytes32 constant MERCHANT_STORAGE_POSITION = keccak256("diamond.standard.merchant.storage");

    struct MerchantStorage {
        mapping(address => bool) isApprove;
    }

    function merchantStorage() internal pure returns (MerchantStorage storage rs) {
        bytes32 position = MERCHANT_STORAGE_POSITION;
        assembly {
            rs.slot := position
        }
    }
}