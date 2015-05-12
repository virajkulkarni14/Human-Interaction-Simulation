#!/bin/sh

CUSTOMER_NAMES=("Ashish" "Krishna" "Akshay" "Shreya" "Viraj")
COFFEE_NAMES=("Ambrosia" "Tesora" "Philharmonic" "Mocha_Tesora" "Ether" "Philtered_Soul" "Silken_Splendor" "Dancing_Water")
COFFEE_SIZES=("small" "medium" "large")
CUSTOMER_RESPONSE=("yes" "no")
CUSTOMER_SWEETNESS_RESPONSE=("low" "high")
declare -i START_PORT=1340

IP="localhost"



#echo "running coffee_shop"
#go run coffee_shop.go 

#sleep 5s

#echo "running barista"
#go run barista.go 

for i in {0..10}
do
	echo $i
	echo ${CUSTOMER_NAMES[$((RANDOM%4+0))]}
	echo ${COFFEE_NAMES[$((RANDOM%7+0))]}
	echo ${COFFEE_SIZES[$((RANDOM%2+0))]}
	echo ${CUSTOMER_RESPONSE[$((RANDOM%1+0))]}
	echo ${CUSTOMER_SWEETNESS_RESPONSE[$((RANDOM%1+0))]}
	echo ${CUSTOMER_RESPONSE[$((RANDOM%1+0))]}
	go run customer.go ${CUSTOMER_NAMES[$((RANDOM%4+0))]} ${COFFEE_NAMES[$((RANDOM%7+0))]} ${COFFEE_SIZES[$((RANDOM%2+0))]} ${CUSTOMER_RESPONSE[$((RANDOM%1+0))]} ${CUSTOMER_SWEETNESS_RESPONSE[$((RANDOM%1+0))]} ${CUSTOMER_RESPONSE[$((RANDOM%1+0))]} $START_PORT localhost & 
	START_PORT=$((START_PORT+1))
	sleep $((RANDOM%10+5))
done

#echo "running customer akshay"
#go run customer.go akshay Ambrosia small yes yes yes 1340 localhost

#echo "running customer krishna"
#go run customer.go krishna Tesora small yes yes yes 1342 localhost

#echo "running customer ashish"
#go run customer.go ashish Philharmonic small yes yes yes 1344 localhost

