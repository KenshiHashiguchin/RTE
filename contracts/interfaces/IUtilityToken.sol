pragma solidity ^0.8.17;

interface IUtilityToken {
    function mint(address to, uint256 amount) external returns (bool);
}