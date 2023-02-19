package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
*/ 
 type FunFacts struct {
      Sun []string 
      Luna []string
      Terra []string
  SunTempCore float64
  SunTempSurf float64 
 
  }
  var fun FunFacts 
  func setFunFacts() {
    fun = FunFacts {}
    fun.SunTempCore = 15000000
    fun.SunTempSurf = 5506.85
    fun.Sun = []string {
      "Temperatur på Solens ytre lag ",
      "Temperatur i Solens kjerne ",
    }
    fun.Luna = []string {
      "Temperatur på Månens overflate på natten -183°C",
      "Temperatur på Månens overflate på dagen 106°C",
    }
    fun.Terra = []string {
      "Høyeste temperatur målt på Jordens overflate 56.7°C",
      "Laveste temperatur målt på Jordens overflate -89.4°C",
    }
  }
  func GetTempFacts(about string) []float64 {
   
    if(about == "sun"){
      return []float64 { fun.SunTempSurf, fun.SunTempCore}
    }
    if(about == "luna"){
      return  []float64{ 0.0, 0.0 }
    }
    if(about == "terra"){
      return  []float64{ 0.0, 0.0 }
    }
    return  []float64{ 0.0, 0.0 }
  
  }

func GetFunFacts(about string) []string {
 setFunFacts()
 
  if(about == "sun"){
    return fun.Sun
  }
  if(about == "luna"){
    return fun.Luna
  }
  if(about == "terra"){
    return fun.Terra
  }
  return fun.Sun
}
