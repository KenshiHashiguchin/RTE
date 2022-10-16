// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ISlashSubmitTransaction {

    function submitTransaction(
        address payingToken_,
        uint256 amountIn_,
        uint256 requiredAmountOut_,
        address[] memory path_,
        address[] memory feePath_,
        string memory paymentId_,
        string memory optional_,
        bytes memory reserved_
    ) external payable returns (bytes16 txNumber);

}