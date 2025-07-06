package Contest

import (
	"regexp"
	"sort"
)

type CouponInfo struct {
	code         string
	businessLine string
}

var businessLinePriority = map[string]int{
	"electronics": 0,
	"grocery":     1,
	"pharmacy":    2,
	"restaurant":  3,
}

var validBusinessLines = map[string]bool{
	"electronics": true,
	"grocery":     true,
	"pharmacy":    true,
	"restaurant":  true,
}

var codeRegex = regexp.MustCompile("^[a-zA-Z0-9_]+$")

func isValidCouponCode(code string) bool {
	if code == "" {
		return false
	}
	return codeRegex.MatchString(code)
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var validateCoupons []CouponInfo
	for i := 0; i < len(code); i++ {
		if !isValidCouponCode(code[i]) {
			continue
		}
		if !validBusinessLines[businessLine[i]] {
			continue
		}
		if !isActive[i] {
			continue
		}
		validateCoupons = append(validateCoupons, CouponInfo{
			code:         code[i],
			businessLine: businessLine[i],
		})
	}

	sort.Slice(validateCoupons, func(i, j int) bool {
		priorityI := businessLinePriority[validateCoupons[i].businessLine]
		priorityJ := businessLinePriority[validateCoupons[j].businessLine]

		if priorityI != priorityJ {
			return priorityI < priorityJ
		}
		return validateCoupons[i].code < validateCoupons[j].code
	})

	var res []string
	for _, coupon := range validateCoupons {
		res = append(res, coupon.code)
	}
	return res
}
