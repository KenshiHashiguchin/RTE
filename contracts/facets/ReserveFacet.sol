// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {LibReserve} from "../libraries/LibApp.sol";
import {LibMerchant} from "../libraries/LibMerchant.sol";
import "../libraries/UniversalERC20.sol";
import "../extension/interfaces/ISlashSubmitTransaction.sol";

contract ReserveFacet {
    using UniversalERC20 for IERC20;
    using SafeMath for uint256;

    function getReservation(string memory _paymentId) external view returns (LibReserve.Reserve memory) {
        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        require(reserve.status != LibReserve.Status.None, "Don't exist paymentId");
        return reserve;
    }

    function getMerchant(string memory _paymentId) external view returns (address) {
        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        return reserve.merchant;
    }

    function reserve(string memory _paymentId, address _merchant, address _token, uint _depositAmount, uint _cancelableTime, uint _withdrawableTime) external {
        LibReserve.validateReserve(_paymentId, _merchant, _token, _cancelableTime, _withdrawableTime);
        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        require(reserve.status == LibReserve.Status.None, "Already exist paymentId");

        reserve.token = _token;
        reserve.merchant = _merchant;
        reserve.subscriber = msg.sender;
        reserve.withdrawableTime = _withdrawableTime;
        reserve.cancelableTime = _cancelableTime;
        reserve.depositAmount = _depositAmount;
        reserve.status = LibReserve.Status.Reserved;

        if (_depositAmount > 0) {
            IERC20(_token).universalTransferFrom(
                msg.sender,
                address(this),
                _depositAmount
            );
        }
    }

    function settleReservation(
        uint256 amountIn_,
        uint256 requiredAmountOut_,
        address[] memory path_,
        address[] memory feePath_,
        string memory paymentId_,
        string memory optional_
    ) external payable {
        LibReserve.isSubscriber(paymentId_);
        LibReserve.isStatusReserved(paymentId_);

        LibReserve.ReserveStorage storage reserveStorage = LibReserve.reserveStorage();
        LibReserve.Reserve storage reserve = reserveStorage.reserve[paymentId_];
        require(amountIn_ >= reserve.depositAmount, "Excess payments");

        uint256 amount = amountIn_.sub(reserve.depositAmount);
        IERC20(reserve.token).universalTransfer(address(this), amount);

        IERC20(reserve.token).approve(reserveStorage.swapSubmitAddress, amountIn_);
        ISlashSubmitTransaction(reserveStorage.swapSubmitAddress).submitTransaction(reserve.token, amountIn_, requiredAmountOut_, path_, feePath_, paymentId_, optional_, bytes(""));

        reserve.status = LibReserve.Status.Settled;
        reserve.additionalAmount = amount;

        // TODO SBT
    }

    function cancel(string memory _paymentId) external {
        LibReserve.isSubscriber(_paymentId);
        LibReserve.isCancelable(_paymentId);

        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        reserve.status = LibReserve.Status.Canceled;

        IERC20(reserve.token).universalTransferFrom(
            address(this),
            reserve.subscriber,
            reserve.depositAmount
        );
    }

    function withdrawDeposit(string memory _paymentId) external {
        LibReserve.isWithdrawDeposit(_paymentId);

        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        reserve.status = LibReserve.Status.Canceled;

        IERC20(reserve.token).universalTransferFrom(
            address(this),
            reserve.subscriber,
            reserve.depositAmount
        );
    }

    function transferSwapSubmitAddress(address _newAddress) external {
        LibReserve.isContractOwner();
        LibReserve.ReserveStorage storage reserveStorage = LibReserve.reserveStorage();
        reserveStorage.swapSubmitAddress = _newAddress;
    }

    function swapSubmitAddress() external view returns (address) {
        return LibReserve.reserveStorage().swapSubmitAddress;
    }

}
