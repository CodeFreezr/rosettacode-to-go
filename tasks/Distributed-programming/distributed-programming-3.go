syntax = "proto3";

service TaxComputer {
  rpc Tax(Amount) returns (Amount) {}
}

message Amount {
  int32 cents = 1;
}

//\Distributed-programming\distributed-programming-3.go
