pragma solidity ^0.5.0;

contract ExhaustionAttack {
    function loop() public view returns(uint) {

        uint g = gasleft();

        while(g > 50000) {
            assembly {
                pop(blockhash(0))
            }
            g = gasleft();
        }
        return gasleft();
    }
}