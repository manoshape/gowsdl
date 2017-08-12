package main

import(
	"./myservice"
	"fmt"
	"github.com/oshapeman/gowsdl/cmd/gowsdl/ferry"
	"github.com/oshapeman/gowsdl/cmd/gowsdl/tianqi"
)

func main()  {
	auth1 := myservice.BasicAuth{Login:"admin",Password:"1234"}
	client1 := myservice.NewKuaidiSoap("",true, &auth1)
	querycode1 := myservice.KuaidiQuery{Compay:"顺丰",OrderNo:"123123"}
	result1,_ := client1.KuaidiQuery(&querycode1)
	fmt.Println(result1)

	auth2 := ferry.BasicAuth{Login:"admin",Password:"1234"}
	client2 := ferry.NewWSF_x0020_ScheduleSoap("",true, &auth2)
	querycode2 := ferry.GetActiveScheduledSeasons{}
	result2,_ := client2.GetActiveScheduledSeasons(&querycode2)
	fmt.Println(result2)


	auth3 := tianqi.BasicAuth{Login:"admin",Password:"1234"}
	client3 := tianqi.NewWeatherWebServiceSoap("",true, &auth3)
	querycode3 := tianqi.GetWeatherbyCityName{TheCityName:"北京"}
	result3,_ := client3.GetWeatherbyCityName(&querycode3)
	fmt.Println(result3)

}
