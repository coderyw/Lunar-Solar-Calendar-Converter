// code by isee15
// from https://github.com/isee15/Lunar-Solar-Calendar-Converter
//

package lunarsolar

type Lunar struct {
	IsLeap     bool //是否闰月
	LunarDay   int
	LunarMonth int
	LunarYear  int
}

type Solar struct {
	SolarDay   int
	SolarMonth int
	SolarYear  int
}

/**
*农历转公历
**/
func LunarToSolar(lunar Lunar) *Solar {
	var days = lunar_month_days[lunar.LunarYear-lunar_month_days[0]]
	var leap = getBitInt(days, 4, 13)
	var offset = 0
	var loopend = leap
	if !lunar.IsLeap {
		if lunar.LunarMonth <= leap || leap == 0 {
			loopend = lunar.LunarMonth - 1
		} else {
			loopend = lunar.LunarMonth
		}
	}
	for i := 0; i < loopend; i++ {
		offset += getBitInt(days, 1, 12-i) + 29
	}
	offset += lunar.LunarDay

	var solar11 = solar_1_1[lunar.LunarYear-solar_1_1[0]]

	var y = getBitInt(solar11, 12, 9)
	var m = getBitInt(solar11, 4, 5)
	var d = getBitInt(solar11, 5, 0)

	var solar = solarFromInt(solarToInt(y, m, d) + int64(offset-1))
	return solar
}

/**
*公历转农历
**/
func SolarToLunar(solar Solar) *Lunar {
	lunar := &Lunar{}
	var index = solar.SolarYear - solar_1_1[0]
	var data = (solar.SolarYear << 9) | (solar.SolarMonth << 5) | (solar.SolarDay)
	var solar11 = 0
	if solar_1_1[index] > data {
		index--
	}
	solar11 = solar_1_1[index]
	var y = getBitInt(solar11, 12, 9)
	var m = getBitInt(solar11, 4, 5)
	var d = getBitInt(solar11, 5, 0)
	var offset int = int(solarToInt(solar.SolarYear, solar.SolarMonth, solar.SolarDay) - solarToInt(y, m, d))

	var days = lunar_month_days[index]
	var leap = getBitInt(days, 4, 13)

	var lunarY = index + solar_1_1[0]
	var lunarM = 1
	var lunarD = 0
	offset += 1

	for i := 0; i < 13; i++ {
		dm := getBitInt(days, 1, 12-i) + 29
		if offset > dm {
			lunarM++
			offset -= dm
		} else {
			break
		}
	}
	lunarD = (int)(offset)
	lunar.LunarYear = lunarY
	lunar.LunarMonth = lunarM
	lunar.IsLeap = false
	if leap != 0 && lunarM > leap {
		lunar.LunarMonth = lunarM - 1
		if lunarM == leap+1 {
			lunar.IsLeap = true
		}
	}

	lunar.LunarDay = lunarD
	return lunar
}

func getBitInt(data int, length int, shift int) int {
	return (data & (((1 << uint32(length)) - 1) << uint32(shift))) >> uint32(shift)
}

// WARNING: Dates before Oct. 1582 are inaccurate
func solarToInt(y int, m int, d int) int64 {
	m = (m + 9) % 12
	y = y - m/10
	return int64(365*y + y/4 - y/100 + y/400 + (m*306+5)/10 + (d - 1))
}

func solarFromInt(g int64) *Solar {
	var y = (10000*g + 14780) / 3652425
	var ddd = g - (365*y + y/4 - y/100 + y/400)
	if ddd < 0 {
		y--
		ddd = g - (365*y + y/4 - y/100 + y/400)
	}
	var mi = (100*ddd + 52) / 3060
	var mm = (mi+2)%12 + 1
	y = y + (mi+2)/12
	dd := ddd - (mi*306+5)/10 + 1
	var solar = &Solar{}
	solar.SolarYear = int(y)
	solar.SolarMonth = int(mm)
	solar.SolarDay = int(dd)

	return solar
}
