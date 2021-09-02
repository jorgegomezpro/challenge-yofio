package models

func ToInvestmentResponse(v300, v500, v700 int32) InvestmentResponse {
	return InvestmentResponse{
		CreditType_300: v300,
		CreditType_500: v500,
		CreditType_700: v700,
	}
}
