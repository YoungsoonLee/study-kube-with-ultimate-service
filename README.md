package oriented design pattern  
immport can doen, not up  
    app  
        business  
            foundation  

expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,mem:memstats.Alloc"

hey -m GET -c 100 -n 10000 "http://localhost:3000/readiness"