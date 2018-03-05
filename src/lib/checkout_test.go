package checkout

import (
 "reflect"
 "runtime"
 "testing"
 "strings"
)

func TestAddClassic(t *testing.T)  {
  for i := 1; i <= 5; i++ {
		Add(Classic)
    if(ClassicNum != i){
      t.Errorf("Wrong number of Classic ads, expected: %d, actual: %d", i, ClassicNum)
    }
	}
}
func TestAddStandout(t *testing.T)  {
  for i := 1; i <= 5; i++ {
		Add(Standout)
    if(StandoutNum != i){
      t.Errorf("Wrong number of Standout ads, expected: %d, actual: %d", i, StandoutNum)
    }
	}
}
func TestAddPremium(t *testing.T)  {
  for i := 1; i <= 5; i++ {
		Add(Premium)
    if(PremiumNum != i){
      t.Errorf("Wrong number of Premium ads, expected: %d, actual: %d", i, PremiumNum)
    }
	}
}

func TestNewCustomerOrdinary(t *testing.T)  {
  NewCustomer("blah")
  if strategies[0] == nil || strategies[1] == nil || strategies[2] == nil {
    t.Errorf("Wrong mapping of strategies: one of the strategies is nil")
  }
  name := GetFunctionName(strategies[0])
  if !strings.Contains(name, "ClassicStandardDeal") {
        t.Errorf("Wrong mapping for Classic Strategy: %s", name)
  }
  name = GetFunctionName(strategies[1])
  if !strings.Contains(name, "StandoutStandardDeal") {
        t.Errorf("Wrong mapping for Standout Strategy: %s", name)
  }
  name = GetFunctionName(strategies[2])
  if !strings.Contains(name, "PremiumStandardDeal") {
        t.Errorf("Wrong mapping for Premium Strategy: %s", name)
  }
}

func TestNewCustomerFord(t *testing.T)  {
  NewCustomer("Ford")
  if strategies[0] == nil || strategies[1] == nil || strategies[2] == nil {
    t.Errorf("Wrong mapping of strategies: one of the strategies is nil")
  }
  name := GetFunctionName(strategies[0])
  if !strings.Contains(name, "ClassicFiveForFourDeal") {
        t.Errorf("Wrong mapping for Classic Strategy: %s", name)
  }
  name = GetFunctionName(strategies[1])
  if !strings.Contains(name, "DiscountStandOutPriceDropsDealB") {
        t.Errorf("Wrong mapping for Standout Strategy: %s", name)
  }
  name = GetFunctionName(strategies[2])
  if !strings.Contains(name, "DiscountThreePremiumPriceDeal") {
        t.Errorf("Wrong mapping for Premium Strategy: %s", name)
  }
}
func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
