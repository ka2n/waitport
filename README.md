waitport
-----------------

wait until specific port is available.


    waitport -listen :8080
    
    waitport -timout 10s -listen :8080 # w/ timeout


## Install

    go install github.com/ka2n/waitport/cmd/waitport@latest
