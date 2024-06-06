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

#SITE 1
#go run app-base -n A1 -m g -port 4444 < /tmp/in_A1 > /tmp/out_A1 &
go run app-base -n A1 < /tmp/in_A1 > /tmp/out_A1 &
go run app-control -n C1 -nbsites 4 < /tmp/in_C1 > /tmp/out_C1 &
go run app-net -n N1 -r '[3,2]' -nbsites 4 < /tmp/in_N1 > /tmp/out_N1 &
#open -a "Google Chrome" http://localhost:63340/pixel-war/app-base/frontend/index.html

#SITE 2
#go run app-base -n A2 -m g -port 4445 < /tmp/in_A2 > /tmp/out_A2 &
go run app-base -n A2 < /tmp/in_A2 > /tmp/out_A2 &
go run app-control -n C2 -nbsites 4 < /tmp/in_C2 > /tmp/out_C2 &
go run app-net -n N2 -r '[1,3]' -nbsites 4 < /tmp/in_N2 > /tmp/out_N2 &
#open -a "Google Chrome" http://localhost:63340/pixel-war/app-base/frontend/index.html

#SITE 3
#go run app-base -n A3 -m g -port 4446 < /tmp/in_A3 > /tmp/out_A3 &
go run app-base -n A3 < /tmp/in_A3 > /tmp/out_A3 &
go run app-control -n C3 -nbsites 4 < /tmp/in_C3 > /tmp/out_C3 &
go run app-net -n N3 -r '[2,4;4,1]' -nbsites 4 < /tmp/in_N3 > /tmp/out_N3 &
#open -a "Google Chrome" http://localhost:63340/pixel-war/app-base/frontend/index.html

#SITE 4
#go run app-base -n A4 -m g -port 4447 < /tmp/in_A4 > /tmp/out_A4 &
go run app-base -n A4 < /tmp/in_A4 > /tmp/out_A4 &
go run app-control -n C4 -nbsites 4 < /tmp/in_C4 > /tmp/out_C4 &
go run app-net -n N4 -r '[3,3]' -nbsites 4 < /tmp/in_N4 > /tmp/out_N4 &
#open -a "Google Chrome" http://localhost:63340/pixel-war/app-base/frontend/index.html


cat /tmp/out_A1 > /tmp/in_C1 &
cat /tmp/out_C1 | tee /tmp/in_A1 > /tmp/in_N1 &
#CONNEXIONS INTERSITES
cat /tmp/out_N1 | tee /tmp/in_C1 | tee /tmp/in_N2 > /tmp/in_N3 &

cat /tmp/out_A2 > /tmp/in_C2 &
cat /tmp/out_C2 | tee /tmp/in_A2 > /tmp/in_N2 &
#CONNEXIONS INTERSITES
cat /tmp/out_N2 | tee /tmp/in_C2 | tee /tmp/in_N3 > /tmp/in_N1 &

cat /tmp/out_A3 > /tmp/in_C3 &
cat /tmp/out_C3 | tee /tmp/in_A3 > /tmp/in_N3 &
#CONNEXIONS INTERSITES
cat /tmp/out_N3 | tee /tmp/in_C3 | tee /tmp/in_N1 | tee /tmp/in_N2 > /tmp/in_N4 &

cat /tmp/out_A4 > /tmp/in_C4 &
cat /tmp/out_C4 | tee /tmp/in_A4 > /tmp/in_N4 &
#CONNEXIONS INTERSITES
cat /tmp/out_N4 | tee /tmp/in_C4 > /tmp/in_N3 &