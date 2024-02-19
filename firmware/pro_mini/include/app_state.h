#ifndef __STATE__
#include <Arduino.h>

enum Direction {
    Forward = 1,
    Backwords = -1,
    None = 0
};

/// motor struct using MX1503 or any other simple h-bridge
typedef struct {
  struct {
        uint8_t a;
        uint8_t b; 
  } pins;
  Direction direction;
  uint8_t speed;
} Motor;

typedef struct {
  Motor m1;
  Motor m2;
  Motor m3;
  Motor m4;
} AppState;


#endif // !__STATE
