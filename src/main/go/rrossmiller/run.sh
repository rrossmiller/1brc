clear
rm 1brc
time go build # -race
# go run . weather_stations.csv
# ./1brc weather_stations.csv

./1brc measurements.txt 
# ./1brc measure_100M.txt 1 $@
# ./1brc measurements.txt 50
#10,000,000
