#!/bin/bash
header="%-10s %8s %13s %15s %3s\n"
format="%-30s %8s %13s %15s %3s\n"

lvalue1="10000000000"
rvalue1="10000000000"

lvalue2="20000000000"
rvalue2="20000000000"

lvalue3="10000"
rvalue3="10000"

lvalue4="300000000"
rvlaue4="300000000"

lvalue5="2000"
rvalue5="2000"

value6="1500"

value7="2200"

lvalue8="2300"
rvalue8="2300"


formula1="formula1"
formula2="formula2"
formula3="formula3"
formula4="formula4"
formula5="formula5"
formula6="formula6"
formula7="formula7"
formula8="formula8"

check1="Total Received Packets"
check2="Total Transmitted Packets"
check3="Packets Lost"
check4="Media Access Control Address"
check5="IPv4 Default Gateway"
check6="Input Rate (bits/sec)"
check7="Output Rate (bits/sec)"
check8="Ignored Packets"


printf "$header" "CHECK" "FORMULA" "LVALUE" "RVALUE" "PASSED"
printf "$format" \
"$check1" "$formula1" "$lvalue1" "$rvalue1" "y/n" \
"$check2" "$formula2" "$lvalue2" "$rvalue2" "y/n" \
"$check3" "$formula3" "$lvalue3" "$rvalue3" "y/n" \
"$check4" "$formula4" "$lvalue4" "$rvalue4" "y/n" \
"$check5" "$formula5" "$lvalue5" "$rvalue5" "y/n" \
"$check6" "$formula6" "$value6" "              " "y/n" \
"$check7" "$formula7" "$value7" "              " "y/n" \
"$check8" "$formula8" "$lvalue8" "$rvalue8" "y/n"
