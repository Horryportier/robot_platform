#include "HardwareSerial.h"
#include <Arduino.h>

// connect on serial 
// send msg 
// receive response
// validate response

void setup(){
    Serial.begin(115200);
}

void loop(){
    Serial.println("hello from esp");
}

