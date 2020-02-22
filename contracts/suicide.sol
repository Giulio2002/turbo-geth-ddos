contract Suicide {
    function start() public {
        assembly {
            selfdestruct(0)
        }     
    }
}