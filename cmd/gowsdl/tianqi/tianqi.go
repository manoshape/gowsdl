
package tianqi

import (
	"encoding/xml"
	"time"
 	"bytes"
 	"crypto/tls"
 	"io/ioutil"
	"net/http"
	"log"
	"net"

	
		
	
)

// against "unused imports"
var _ time.Time
var _ xml.Name














	

	

	
		
			
			
			
				type GetSupportCity struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getSupportCity"`
					
						
	   ByProvinceName string `xml:"byProvinceName,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetSupportCityResponse struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getSupportCityResponse"`
					
						
	   GetSupportCityResult *ArrayOfString `xml:"getSupportCityResult,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetSupportProvince struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getSupportProvince"`
					
						
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetSupportProvinceResponse struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getSupportProvinceResponse"`
					
						
	   GetSupportProvinceResult *ArrayOfString `xml:"getSupportProvinceResult,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetSupportDataSet struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getSupportDataSet"`
					
						
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetSupportDataSetResponse struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getSupportDataSetResponse"`
					
						
	  
	GetSupportDataSetResult struct {
	
		
			
	  

 
	

			
	

			
	

			
	

			
	

		
	
	} `xml:"getSupportDataSetResult,omitempty"`
 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetWeatherbyCityName struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getWeatherbyCityName"`
					
						
	   TheCityName string `xml:"theCityName,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetWeatherbyCityNameResponse struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getWeatherbyCityNameResponse"`
					
						
	   GetWeatherbyCityNameResult *ArrayOfString `xml:"getWeatherbyCityNameResult,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetWeatherbyCityNamePro struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getWeatherbyCityNamePro"`
					
						
	   TheCityName string `xml:"theCityName,omitempty"` 
	   TheUserID string `xml:"theUserID,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
			
			
			
				type GetWeatherbyCityNameProResponse struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ getWeatherbyCityNameProResponse"`
					
						
	   GetWeatherbyCityNameProResult *ArrayOfString `xml:"getWeatherbyCityNameProResult,omitempty"` 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	
		
	
		
			
			
			
				type DataSet struct {
					XMLName xml.Name `xml:"http://WebXml.com.cn/ DataSet"`
					
						
	  

 
	

						
	

						
	

						
	

						
	

					
				}
			
		
	

	
		
		
		type ArrayOfString struct {
			XMLName xml.Name `xml:"http://WebXml.com.cn/ ArrayOfString"`
			
				
	   String []string `xml:"string,omitempty"` 
	

				
	

				
	

				
	

				
	

			
		}
	



	
	type WeatherWebServiceSoap struct {
		client *SOAPClient
	}

	func NewWeatherWebServiceSoap(url string, tls bool, auth *BasicAuth) *WeatherWebServiceSoap {
		if url == "" {
			url = "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx"
		}
		client := NewSOAPClient(url, tls, auth)

		return &WeatherWebServiceSoap{
			client: client,
		}
	}

	
		
		
		
		

		
		
		/* <br /><h3>查询本天气预报Web Services支持的国内外城市或地区信息</h3><p>输入参数：byProvinceName = 指定的洲或国内的省份，若为ALL或空则表示返回全部城市；返回数据：一个一维字符串数组 String()，结构为：城市名称(城市代码)。</p><br /> */
		func (service *WeatherWebServiceSoap) GetSupportCity (request *GetSupportCity) (*GetSupportCityResponse, error) {
			response := new(GetSupportCityResponse)
			err := service.client.Call("http://WebXml.com.cn/getSupportCity", request, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br /><h3>获得本天气预报Web Services支持的洲、国内外省份和城市信息</h3><p>输入参数：无； 返回数据：一个一维字符串数组 String()，内容为洲或国内省份的名称。</p><br /> */
		func (service *WeatherWebServiceSoap) GetSupportProvince (request *GetSupportProvince) (*GetSupportProvinceResponse, error) {
			response := new(GetSupportProvinceResponse)
			err := service.client.Call("http://WebXml.com.cn/getSupportProvince", request, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>获得本天气预报Web Services支持的洲、国内外省份和城市信息</h3><p>输入参数：无；返回：DataSet 。DataSet.Tables(0) 为支持的洲和国内省份数据，DataSet.Tables(1) 为支持的国内外城市或地区数据。DataSet.Tables(0).Rows(i).Item("ID") 主键对应 DataSet.Tables(1).Rows(i).Item("ZoneID") 外键。<br />Tables(0)：ID = ID主键，Zone = 支持的洲、省份；Tables(1)：ID 主键，ZoneID = 对应Tables(0)ID的外键，Area = 城市或地区，AreaCode = 城市或地区代码。</p><br /> */
		func (service *WeatherWebServiceSoap) GetSupportDataSet (request *GetSupportDataSet) (*GetSupportDataSetResponse, error) {
			response := new(GetSupportDataSetResponse)
			err := service.client.Call("http://WebXml.com.cn/getSupportDataSet", request, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>根据城市或地区名称查询获得未来三天内天气情况、现在的天气实况、天气和生活指数</h3><p>调用方法如下：输入参数：theCityName = 城市中文名称(国外城市可用英文)或城市代码(不输入默认为上海市)，如：上海 或 58367，如有城市名称重复请使用城市代码查询(可通过 getSupportCity 或 getSupportDataSet 获得)；返回数据： 一个一维数组 String(22)，共有23个元素。<br />String(0) 到 String(4)：省份，城市，城市代码，城市图片名称，最后更新时间。String(5) 到 String(11)：当天的 气温，概况，风向和风力，天气趋势开始图片名称(以下称：图标一)，天气趋势结束图片名称(以下称：图标二)，现在的天气实况，天气和生活指数。String(12) 到 String(16)：第二天的 气温，概况，风向和风力，图标一，图标二。String(17) 到 String(21)：第三天的 气温，概况，风向和风力，图标一，图标二。String(22) 被查询的城市或地区的介绍 <br /><a href="http://www.webxml.com.cn/images/weather.zip">下载天气图标<img src="http://www.webxml.com.cn/images/download_w.gif" border="0" align="absbottom" /></a>(包含大、中、小尺寸) <a href="http://www.webxml.com.cn/zh_cn/weather_icon.aspx" target="_blank">天气图例说明</a> <a href="http://www.webxml.com.cn/files/weather_eg.zip">调用此天气预报Web Services实例下载</a> (VB ASP.net 2.0)</p><br /> */
		func (service *WeatherWebServiceSoap) GetWeatherbyCityName (request *GetWeatherbyCityName) (*GetWeatherbyCityNameResponse, error) {
			response := new(GetWeatherbyCityNameResponse)
			err := service.client.Call("http://WebXml.com.cn/getWeatherbyCityName", request, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>根据城市或地区名称查询获得未来三天内天气情况、现在的天气实况、天气和生活指数（For商业用户）</h3><p>调用方法同 getWeatherbyCityName，输入参数：theUserID = 商业用户ID</p><br /> */
		func (service *WeatherWebServiceSoap) GetWeatherbyCityNamePro (request *GetWeatherbyCityNamePro) (*GetWeatherbyCityNameProResponse, error) {
			response := new(GetWeatherbyCityNameProResponse)
			err := service.client.Call("http://WebXml.com.cn/getWeatherbyCityNamePro", request, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	

	
	type WeatherWebServiceHttpGet struct {
		client *SOAPClient
	}

	func NewWeatherWebServiceHttpGet(url string, tls bool, auth *BasicAuth) *WeatherWebServiceHttpGet {
		if url == "" {
			url = ""
		}
		client := NewSOAPClient(url, tls, auth)

		return &WeatherWebServiceHttpGet{
			client: client,
		}
	}

	
		
		
		
		

		
		
		/* <br /><h3>查询本天气预报Web Services支持的国内外城市或地区信息</h3><p>输入参数：byProvinceName = 指定的洲或国内的省份，若为ALL或空则表示返回全部城市；返回数据：一个一维字符串数组 String()，结构为：城市名称(城市代码)。</p><br /> */

		
	
		
		
		
		

		
		
		/* <br /><h3>获得本天气预报Web Services支持的洲、国内外省份和城市信息</h3><p>输入参数：无； 返回数据：一个一维字符串数组 String()，内容为洲或国内省份的名称。</p><br /> */
		func (service *WeatherWebServiceHttpGet) GetSupportProvince () (*ArrayOfString, error) {
			response := new(ArrayOfString)
			err := service.client.Call("", nil, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>获得本天气预报Web Services支持的洲、国内外省份和城市信息</h3><p>输入参数：无；返回：DataSet 。DataSet.Tables(0) 为支持的洲和国内省份数据，DataSet.Tables(1) 为支持的国内外城市或地区数据。DataSet.Tables(0).Rows(i).Item("ID") 主键对应 DataSet.Tables(1).Rows(i).Item("ZoneID") 外键。<br />Tables(0)：ID = ID主键，Zone = 支持的洲、省份；Tables(1)：ID 主键，ZoneID = 对应Tables(0)ID的外键，Area = 城市或地区，AreaCode = 城市或地区代码。</p><br /> */
		func (service *WeatherWebServiceHttpGet) GetSupportDataSet () (*DataSet, error) {
			response := new(DataSet)
			err := service.client.Call("", nil, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>根据城市或地区名称查询获得未来三天内天气情况、现在的天气实况、天气和生活指数</h3><p>调用方法如下：输入参数：theCityName = 城市中文名称(国外城市可用英文)或城市代码(不输入默认为上海市)，如：上海 或 58367，如有城市名称重复请使用城市代码查询(可通过 getSupportCity 或 getSupportDataSet 获得)；返回数据： 一个一维数组 String(22)，共有23个元素。<br />String(0) 到 String(4)：省份，城市，城市代码，城市图片名称，最后更新时间。String(5) 到 String(11)：当天的 气温，概况，风向和风力，天气趋势开始图片名称(以下称：图标一)，天气趋势结束图片名称(以下称：图标二)，现在的天气实况，天气和生活指数。String(12) 到 String(16)：第二天的 气温，概况，风向和风力，图标一，图标二。String(17) 到 String(21)：第三天的 气温，概况，风向和风力，图标一，图标二。String(22) 被查询的城市或地区的介绍 <br /><a href="http://www.webxml.com.cn/images/weather.zip">下载天气图标<img src="http://www.webxml.com.cn/images/download_w.gif" border="0" align="absbottom" /></a>(包含大、中、小尺寸) <a href="http://www.webxml.com.cn/zh_cn/weather_icon.aspx" target="_blank">天气图例说明</a> <a href="http://www.webxml.com.cn/files/weather_eg.zip">调用此天气预报Web Services实例下载</a> (VB ASP.net 2.0)</p><br /> */

	
		
		
		
		

		
		
		/* <br><h3>根据城市或地区名称查询获得未来三天内天气情况、现在的天气实况、天气和生活指数（For商业用户）</h3><p>调用方法同 getWeatherbyCityName，输入参数：theUserID = 商业用户ID</p><br /> */

		
	

	
	type WeatherWebServiceHttpPost struct {
		client *SOAPClient
	}

	func NewWeatherWebServiceHttpPost(url string, tls bool, auth *BasicAuth) *WeatherWebServiceHttpPost {
		if url == "" {
			url = ""
		}
		client := NewSOAPClient(url, tls, auth)

		return &WeatherWebServiceHttpPost{
			client: client,
		}
	}

	
		
		
		
		

		
		
		/* <br /><h3>查询本天气预报Web Services支持的国内外城市或地区信息</h3><p>输入参数：byProvinceName = 指定的洲或国内的省份，若为ALL或空则表示返回全部城市；返回数据：一个一维字符串数组 String()，结构为：城市名称(城市代码)。</p><br /> */

		
	
		
		
		
		

		
		
		/* <br /><h3>获得本天气预报Web Services支持的洲、国内外省份和城市信息</h3><p>输入参数：无； 返回数据：一个一维字符串数组 String()，内容为洲或国内省份的名称。</p><br /> */
		func (service *WeatherWebServiceHttpPost) GetSupportProvince () (*ArrayOfString, error) {
			response := new(ArrayOfString)
			err := service.client.Call("", nil, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>获得本天气预报Web Services支持的洲、国内外省份和城市信息</h3><p>输入参数：无；返回：DataSet 。DataSet.Tables(0) 为支持的洲和国内省份数据，DataSet.Tables(1) 为支持的国内外城市或地区数据。DataSet.Tables(0).Rows(i).Item("ID") 主键对应 DataSet.Tables(1).Rows(i).Item("ZoneID") 外键。<br />Tables(0)：ID = ID主键，Zone = 支持的洲、省份；Tables(1)：ID 主键，ZoneID = 对应Tables(0)ID的外键，Area = 城市或地区，AreaCode = 城市或地区代码。</p><br /> */
		func (service *WeatherWebServiceHttpPost) GetSupportDataSet () (*DataSet, error) {
			response := new(DataSet)
			err := service.client.Call("", nil, response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}
		
	
		
		
		
		

		
		
		/* <br><h3>根据城市或地区名称查询获得未来三天内天气情况、现在的天气实况、天气和生活指数</h3><p>调用方法如下：输入参数：theCityName = 城市中文名称(国外城市可用英文)或城市代码(不输入默认为上海市)，如：上海 或 58367，如有城市名称重复请使用城市代码查询(可通过 getSupportCity 或 getSupportDataSet 获得)；返回数据： 一个一维数组 String(22)，共有23个元素。<br />String(0) 到 String(4)：省份，城市，城市代码，城市图片名称，最后更新时间。String(5) 到 String(11)：当天的 气温，概况，风向和风力，天气趋势开始图片名称(以下称：图标一)，天气趋势结束图片名称(以下称：图标二)，现在的天气实况，天气和生活指数。String(12) 到 String(16)：第二天的 气温，概况，风向和风力，图标一，图标二。String(17) 到 String(21)：第三天的 气温，概况，风向和风力，图标一，图标二。String(22) 被查询的城市或地区的介绍 <br /><a href="http://www.webxml.com.cn/images/weather.zip">下载天气图标<img src="http://www.webxml.com.cn/images/download_w.gif" border="0" align="absbottom" /></a>(包含大、中、小尺寸) <a href="http://www.webxml.com.cn/zh_cn/weather_icon.aspx" target="_blank">天气图例说明</a> <a href="http://www.webxml.com.cn/files/weather_eg.zip">调用此天气预报Web Services实例下载</a> (VB ASP.net 2.0)</p><br /> */

		
	
		
		
		
		

		
		
		/* <br><h3>根据城市或地区名称查询获得未来三天内天气情况、现在的天气实况、天气和生活指数（For商业用户）</h3><p>调用方法同 getWeatherbyCityName，输入参数：theUserID = 商业用户ID</p><br /> */

		
	


var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login string
	Password string
}

type SOAPClient struct {
	url string
	tls bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url: url,
		tls: tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	err := encoder.Encode(envelope)
	if err == nil {
		err = encoder.Flush()
	}

	log.Println(buffer.String())
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
