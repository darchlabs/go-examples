// SPDX-Licence-Identifier: MIT

pragma solidity ^0.8.10;

contract ExampleContract {
    event Example(
        uint256 amount,
        address indexed userAddr,
        uint256 amount1,
        uint256 indexed num,
        string indexed lala,
        string lala2
    );

    uint256 private _amount;

    constructor() {
        _amount = 0;
    }

    function example(uint256 amount) external {
        _amount = amount;
        emit Example(amount, msg.sender, amount + 1, 5, "nana", "dada");
    }

    function getExample() external view returns (uint256) {
        return _amount;
    }

    function checkUpkeep() {} // => view

    function performUpkeep() {
        // check

         getExample()
    }
}
