# PROTOCOL.md


#### STRUCTURE
|1b| |1b| |Nb|
---- ---- ----
 |   |    msg
 |   msg length
msg type

##### Messages
- turn_motor
    msg: 3 byte
    motor|direction|speed
