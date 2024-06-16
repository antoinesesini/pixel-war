mkfifo /tmp/in_A1 /tmp/out_A1
mkfifo /tmp/in_C1 /tmp/out_C1
mkfifo /tmp/in_N1 /tmp/out_N1

mkfifo /tmp/in_A2 /tmp/out_A2
mkfifo /tmp/in_C2 /tmp/out_C2
mkfifo /tmp/in_N2 /tmp/out_N2

mkfifo /tmp/in_A3 /tmp/out_A3
mkfifo /tmp/in_C3 /tmp/out_C3
mkfifo /tmp/in_N3 /tmp/out_N3

mkfifo /tmp/in_A4 /tmp/out_A4
mkfifo /tmp/in_C4 /tmp/out_C4
mkfifo /tmp/in_N4 /tmp/out_N4

mkfifo /tmp/in_A5 /tmp/out_A5
mkfifo /tmp/in_C5 /tmp/out_C5
mkfifo /tmp/in_N5 /tmp/out_N5

mkfifo /tmp/in_A6 /tmp/out_A6
mkfifo /tmp/in_C6 /tmp/out_C6
mkfifo /tmp/in_N6 /tmp/out_N6

#SITE 1
portB1=4444
portN1=4445
go run app-base -n A1 -m g -port $portB1 < /tmp/in_A1 > /tmp/out_A1 &
#go run app-base -n A1 < /tmp/in_A1 > /tmp/out_A1 &
go run app-control -n C1 -nbsites 4 < /tmp/in_C1 > /tmp/out_C1 &
go run app-net -n N1 -r '[3,2]' -nbsites 4 -port $portN1 -v 2 < /tmp/in_N1 > /tmp/out_N1 &
open -a "Google Chrome" "http://localhost:63340/pixel-war/app-base/frontend/index.html?portB=$portB1&portN=$portN1"

#SITE 2
portB2=4446
portN2=4447
go run app-base -n A2 -m g -port $portB2 < /tmp/in_A2 > /tmp/out_A2 &
#go run app-base -n A2 < /tmp/in_A2 > /tmp/out_A2 &
go run app-control -n C2 -nbsites 4 < /tmp/in_C2 > /tmp/out_C2 &
go run app-net -n N2 -r '[1,3]' -nbsites 4 -port $portN2 -v 2 < /tmp/in_N2 > /tmp/out_N2 &
open -a "Google Chrome" "http://localhost:63340/pixel-war/app-base/frontend/index.html?portB=$portB2&portN=$portN2"

#SITE 3
portB3=4448
portN3=4449
go run app-base -n A3 -m g -port $portB3 < /tmp/in_A3 > /tmp/out_A3 &
#go run app-base -n A3 < /tmp/in_A3 > /tmp/out_A3 &
go run app-control -n C3 -nbsites 4 < /tmp/in_C3 > /tmp/out_C3 &
go run app-net -n N3 -r '[2,4;4,1]' -nbsites 4 -port $portN3 -v 3 < /tmp/in_N3 > /tmp/out_N3 &
open -a "Google Chrome" "http://localhost:63340/pixel-war/app-base/frontend/index.html?portB=$portB3&portN=$portN3"

#SITE 4
portB4=4450
portN4=4451
go run app-base -n A4 -m g -port $portB4 < /tmp/in_A4 > /tmp/out_A4 &
#go run app-base -n A4 < /tmp/in_A4 > /tmp/out_A4 &
go run app-control -n C4 -nbsites 4 < /tmp/in_C4 > /tmp/out_C4 &
go run app-net -n N4 -r '[3,3]' -nbsites 4 -port $portN4 -v 1 < /tmp/in_N4 > /tmp/out_N4 &
open -a "Google Chrome" "http://localhost:63340/pixel-war/app-base/frontend/index.html?portB=$portB4&portN=$portN4"

#SITE 5
portB5=4452
portN5=4453
go run app-base -n A5 -m g -port $portB5 < /tmp/in_A5 > /tmp/out_A5 &
#go run app-base -n A5 < /tmp/in_A5 > /tmp/out_A5 &
go run app-control -n C5 < /tmp/in_C5 > /tmp/out_C5 &
go run app-net -n N5 -r '[4,4]' -port $portN5 -e inactif -t 30 -c 4 < /tmp/in_N5 > /tmp/out_N5 &
open -a "Google Chrome" "http://localhost:63340/pixel-war/app-base/frontend/index.html?portB=$portB5&portN=$portN5"


#SITE 6
portB6=4454
portN6=4455
go run app-base -n A6 -m g -port $portB6 < /tmp/in_A6 > /tmp/out_A6 &
#go run app-base -n A6 < /tmp/in_A6 > /tmp/out_A6 &
go run app-control -n C6 < /tmp/in_C6 > /tmp/out_C6 &
go run app-net -n N6 -r '[2,2]' -port $portN6 -e inactif -t 30 -c 2 < /tmp/in_N6 > /tmp/out_N6 &
open -a "Google Chrome" "http://localhost:63340/pixel-war/app-base/frontend/index.html?portB=$portB6&portN=$portN6"


cat /tmp/out_A1 > /tmp/in_C1 &
cat /tmp/out_C1 | tee /tmp/in_A1 > /tmp/in_N1 &
#CONNEXIONS INTERSITES
cat /tmp/out_N1 | tee /tmp/in_C1 | tee /tmp/in_N2 > /tmp/in_N3 &

cat /tmp/out_A2 > /tmp/in_C2 &
cat /tmp/out_C2 | tee /tmp/in_A2 > /tmp/in_N2 &
#CONNEXIONS INTERSITES
cat /tmp/out_N2 | tee /tmp/in_C2 | tee /tmp/in_N3 | tee /tmp/in_N6 > /tmp/in_N1 &

cat /tmp/out_A3 > /tmp/in_C3 &
cat /tmp/out_C3 | tee /tmp/in_A3 > /tmp/in_N3 &
#CONNEXIONS INTERSITES
cat /tmp/out_N3 | tee /tmp/in_C3 | tee /tmp/in_N1 | tee /tmp/in_N2 > /tmp/in_N4 &

cat /tmp/out_A4 > /tmp/in_C4 &
cat /tmp/out_C4 | tee /tmp/in_A4 > /tmp/in_N4 &
#CONNEXIONS INTERSITES
cat /tmp/out_N4 | tee /tmp/in_C4 | tee /tmp/in_N3 > /tmp/in_N5 &

cat /tmp/out_A5 > /tmp/in_C5 &
cat /tmp/out_C5 | tee /tmp/in_A5 > /tmp/in_N5 &
#CONNEXIONS INTERSITES
cat /tmp/out_N5 | tee /tmp/in_C5 | tee /tmp/in_N6 > /tmp/in_N4 &

cat /tmp/out_A6 > /tmp/in_C6 &
cat /tmp/out_C6 | tee /tmp/in_A6 > /tmp/in_N6 &
#CONNEXIONS INTERSITES
cat /tmp/out_N6 | tee /tmp/in_C6 | tee /tmp/in_N2 > /tmp/in_N5 &