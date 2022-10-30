// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./interfaces/ISlashCustomPlugin.sol";
import "../libraries/UniversalERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ReceivePayment is ISlashCustomPlugin, Ownable {
    using UniversalERC20 for IERC20;

    address public reserveContract;

    function updateReservationContractAddress(address _newAddress) external onlyOwner {
        _reservationAddress(_newAddress);
    }

    function _reservationAddress(address _newAddress) internal {
        reserveContract = _newAddress;
    }

    function supportSlashExtensionInterface() external returns (bool) {
        return true;
    }

    function receivePayment(
        address receiveToken,
        uint256 amount,
        string memory paymentId,
        string memory optional
    ) external payable override {
        require(amount >= 0, "invalid amount");
        require(receiveToken != address(0), "invalid token");

        address merchantAddress = getMerchant(paymentId);
        IERC20(receiveToken).universalTransferFrom(
            msg.sender,
            merchantAddress,
            amount
        );
    }

    function getMerchant(string memory paymentId) public returns (address) {
        bytes memory payload = abi.encodeWithSignature("getMerchant(string)", paymentId);
        (bool success, bytes memory returnData) = address(reserveContract).call(payload);
        return bytesToAddress(returnData);
    }

    function bytesToAddress(bytes memory bys) private pure returns (address addr) {
        assembly {
            addr := mload(add(bys, 32))
        }
    }
}