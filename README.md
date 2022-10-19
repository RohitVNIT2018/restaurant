# restaurant
#Sample output
<pre>
go run task.go
Welcome to Restaurent
Top 3 items consumed are :
1 )  Steam Rice consumed  15  times
2 )  Tava Roti consumed  14  times
3 )  Egg Curry consumed  11  times
Error : duplicate eater_Id 12001 with menuId 304
</pre>

#sample test output
<br/>
<pre>
go test ./analyzeTest -v
=== RUN   TestCheckDuplicates
--- PASS: TestCheckDuplicates (0.00s)
=== RUN   TestTopItemsConsumed
--- PASS: TestTopItemsConsumed (0.00s)
PASS
ok      task/analyzeTest        1.595s
</pre>
