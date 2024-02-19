#include "HardwareSerial.h"
#include <Arduino.h>

// connect on serial
// wait for msg to arrive
// parse it
// send response

enum Direction { Forward = 1, Backwords = -1, None = 0 };

typedef struct {
  uint8_t a;
  uint8_t b;
} MotorPins;
/// motor struct using MX1503 or any other simple h-bridge
typedef struct {
  MotorPins pins;
  Direction direction;
  uint8_t speed;
} Motor;

typedef struct {
  Motor m1;
  Motor m2;
  Motor m3;
  Motor m4;
} State;

State state{{{2, 3}, None, 0}, {{4, 5}, None, 0}, {{6, 7}, None, 0}, {{8, 9}, None, 0}};

void setup() {
 Serial.begin(115200);
}

void loop() {
}
