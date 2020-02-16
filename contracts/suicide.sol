pragma solidity ^0.5.0;

contract SuicideAttack {
    function loop() public {
        uint g = gasleft();
        Suicide s1 = new Suicide();
        while(g > 40000) {
            Suicide s2 = new Suicide();
            s1.kill(address(s2));
            s1 = s2;
            g = gasleft();
        }
    }
}

contract Suicide {
    /* Test How Much Gas is Used here*/
    function kill(address guy) public {
        assembly{
            selfdestruct(guy)
        }
    }
}