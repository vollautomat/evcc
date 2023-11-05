package schedule

import "sort"

type Config []struct {
	Price       float64
	Days, Hours string
}

func FromConfig(cc Config) (Zones, error) {
	var res Zones

	for _, z := range cc {
		days, err := ParseDays(z.Days)
		if err != nil {
			return nil, err
		}

		hours, err := ParseTimeRanges(z.Hours)
		if err != nil && z.Hours != "" {
			return nil, err
		}

		if len(hours) == 0 {
			res = append(res, Zone{
				Price: z.Price,
				Days:  days,
			})
			continue
		}

		for _, h := range hours {
			res = append(res, Zone{
				Price: z.Price,
				Days:  days,
				Hours: h,
			})
		}
	}

	sort.Sort(res)

	return res, nil
}
