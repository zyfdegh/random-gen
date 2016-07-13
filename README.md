# random-gen
An interesting demo showing how to generate random numbers and do a statistic analysis.

# Screenshots

![Screenshot00](https://raw.githubusercontent.com/zyfdegh/random-gen/master/raw/screenshot-00.png)

![Screenshot01](https://raw.githubusercontent.com/zyfdegh/random-gen/master/raw/screenshot-01.png)

# How it works?
This program do the following things:

1. generate **N** random numbers in byte array **arr**, each ranges from 0~255
2. loop **arr** and get maximum(max) and minimum(min) value
3. check if **max==255 && min==0**
4. do above 3 steps 1000 times and count how many times step 3 is statisfied
5. inscrease **N** by 10 and do above steps until **N** is greater than 2016

# Run
Redirect out to file is recommended

```sh
git clone https://github.com/zyfdegh/random-gen.git
go build -o random-gen main.go
./random-gen > statics.txt
```

# Of what use?
The statistics will be used in another project.

# Related
[**go-chart**](https://github.com/zyfdegh/go-chart.git): Draw statistics of this program into a chart.

