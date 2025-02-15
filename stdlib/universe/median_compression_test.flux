package universe_test


import "testing"
import "csv"

option now = () => 2030-01-01T00:00:00Z

inData =
    "
#datatype,string,long,string,string,dateTime:RFC3339,double
#group,false,false,true,true,false,false
#default,_result,,,,,
,result,table,_measurement,_field,_time,_value
,,0,SOYcRk,NC7N,2018-12-18T21:12:45Z,55
,,0,SOYcRk,NC7N,2018-12-18T21:12:55Z,15
,,0,SOYcRk,NC7N,2018-12-18T21:13:05Z,25
,,0,SOYcRk,NC7N,2018-12-18T21:13:15Z,5
,,0,SOYcRk,NC7N,2018-12-18T21:13:25Z,105
,,0,SOYcRk,NC7N,2018-12-18T21:13:35Z,45
"
outData =
    "
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,string,double
#group,false,false,true,true,true,true,false
#default,_result,,,,,,
,result,table,_start,_stop,_measurement,_field,_value
,,0,2018-12-01T00:00:00Z,2030-01-01T00:00:00Z,SOYcRk,NC7N,41.666666666666664
"

testcase median {
    got =
        csv.from(csv: inData)
            |> testing.load()
            |> range(start: 2018-12-01T00:00:00Z)
            |> median(compression: 1.0)
    want = csv.from(csv: outData)

    testing.diff(got, want)
}
