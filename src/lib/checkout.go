package checkout

import (
	//"fmt"
	"math"
)


type JobType string
type JobPrice float32

const (
	Classic  JobType = "Classic"
	Standout JobType = "Standout"
	Premium  JobType = "Premium"
)

const (
	ClassicPrice  float64 = 269.99
	StandoutPrice float64 = 322.99
	PremiumPrice  float64 = 394.99
)

const (
	ClassicId  int = 0
	StandOutId int = 1
	PremiumId  int = 2
)

type CheckoutStrategy func(int) float64

//Should be float32?
var ClassicNum, StandoutNum, PremiumNum int

var strategies [3]CheckoutStrategy

func NewCustomer(name string) {
	ClassicNum = 0
	StandoutNum = 0
	PremiumNum = 0
switch name {
case "Uniliver":
	strategies[0] = ClassicThreeForTwoDeal
	strategies[1] = StandoutStandardDeal
	strategies[2] = PremiumStandardDeal
case "Apple":
	strategies[0] = ClassicStandardDeal
	strategies[1] = DiscountStandOutPriceDropsDealA
	strategies[2] = PremiumStandardDeal
case "Nike":
	strategies[0] = ClassicStandardDeal
	strategies[1] = StandoutStandardDeal
	strategies[2] = DiscountFourPremiumPriceDeal
case "Ford":
	strategies[0] = ClassicFiveForFourDeal
	strategies[1] = DiscountStandOutPriceDropsDealB
	strategies[2] = DiscountThreePremiumPriceDeal
default:
	strategies[0] = ClassicStandardDeal
	strategies[1] = StandoutStandardDeal
	strategies[2] = PremiumStandardDeal
	}
}

func Add(jobType JobType) {
	switch jobType {
	case Classic:
		ClassicNum++
	case Standout:
		StandoutNum++
	case Premium:
		PremiumNum++
	}
}

func Total() float64 {

	var ClassicTotal, StandoutTotal, PremiumTotal float64 = 0, 0, 0
	ClassicTotal = strategies[ClassicId](ClassicNum)
	StandoutTotal = strategies[StandOutId](StandoutNum)
	PremiumTotal = strategies[PremiumId](PremiumNum)
	return ClassicTotal + StandoutTotal + PremiumTotal
}

//TODO: use the fucntion type instead
//func ClassicThreeForTwoDeal := CheckoutStrategy(func(int) float32 {
func ClassicThreeForTwoDeal(numJobs int) float64 {
  floorPart := math.Floor(float64(numJobs) / 3)
	modPart := math.Mod(float64(numJobs), 3)
	return (floorPart * 2 + modPart) * ClassicPrice
}

func ClassicFiveForFourDeal(numJobs int) float64 {
  floorPart := math.Floor(float64(numJobs) / 5)
	modPart := math.Mod(float64(numJobs), 5)
	return (floorPart * 4 + modPart) * ClassicPrice
}
func DiscountStandOutPriceDropsDealA(numJobs int) float64 {
	discountPrice := 322.99
	return float64(numJobs) * discountPrice
}
func DiscountStandOutPriceDropsDealB(numJobs int) float64 {
	discountPrice := 309.99
	return float64(numJobs) * discountPrice
}

func DiscountFourPremiumPriceDeal(numJobs int) float64{
	discountPrice := 379.99
	currPrice := PremiumPrice
	if numJobs >= 4 {
		currPrice = discountPrice
	}
	return float64(numJobs) * currPrice
}
func DiscountThreePremiumPriceDeal(numJobs int) float64{
	discountPrice := 389.99
	currPrice := PremiumPrice
	if numJobs >= 3 {
		currPrice = discountPrice
	}
	return float64(numJobs) * currPrice
}
func ClassicStandardDeal(numJobs int) float64 {
	return float64(numJobs) * ClassicPrice
}
func  StandoutStandardDeal(numJobs int) float64 {
	return float64(numJobs) * StandoutPrice
}
func  PremiumStandardDeal(numJobs int) float64 {
	return float64(numJobs) * PremiumPrice
}
