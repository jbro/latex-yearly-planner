package cal

import (
	"math"
	"strconv"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
	"github.com/kudrykv/latex-yearly-planner/app/texx"
)

type Days []*Day
type Day struct {
	Time time.Time
}

func (d Day) Day(today, large interface{}) string {
	if d.Time.IsZero() {
		return ""
	}

	day := strconv.Itoa(d.Time.Day())

	if larg, _ := large.(bool); larg {
		return `\hyperlink{` + d.ref() + `}{\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + day + `\\ \hline\end{tabular}}`
	}

	if td, ok := today.(Day); ok {
		if d.Time.Equal(td.Time) {
			return texx.EmphCell(day)
		}
	}

	return hyper.Link(d.ref(), day)
}

func (d Day) ref() string {
	return d.Time.Format(time.RFC3339)
}

func (d Day) Add(days int) Day {
	return Day{Time: d.Time.AddDate(0, 0, days)}
}

func (d Day) WeekLink() string {
	return hyper.Link(d.ref(), strconv.Itoa(d.Time.Day())+", "+d.Time.Weekday().String())
}

func (d Day) Breadcrumb(prefix string, leaf string) string {
	wpref := ""
	_, wn := d.Time.ISOWeek()
	if wn > 50 && d.Time.Month() == time.January {
		wpref = "fw"
	}

	dayItem := header.NewTextItem(d.Time.Format("Monday, 2")).RefText(d.Time.Format(time.RFC3339))
	items := header.Items{
		header.NewIntItem(d.Time.Year()),
		header.NewTextItem("Q" + strconv.Itoa(int(math.Ceil(float64(d.Time.Month())/3.)))),
		header.NewMonthItem(d.Time.Month()),
		header.NewTextItem(wpref + "Week " + strconv.Itoa(wn)),
	}

	if len(leaf) > 0 {
		items = append(items, dayItem, header.NewTextItem(leaf).RefText(prefix+d.ref()).Ref(true))
	} else {
		items = append(items, dayItem.Ref(true))
	}

	return items.Table(true)
}

func (d Day) LinkLeaf(prefix, leaf string) string {
	return hyper.Link(prefix+d.ref(), leaf)
}

func (d Day) PrevNext(prefix string) header.Items {
	items := header.Items{}

	if d.Time.Month() > time.January || d.Time.Day() > 1 {
		prev := d.Add(-1)
		items = append(items, header.NewTextItem(prev.Time.Format("Mon, 2")).RefText(prefix+prev.ref()))
	}

	if d.Time.Month() < time.December || d.Time.Day() < 31 {
		next := d.Add(1)
		items = append(items, header.NewTextItem(next.Time.Format("Mon, 2")).RefText(prefix+next.ref()))
	}

	return items
}

func (d Day) Hours(bottom, top int) Days {
	moment := time.Date(1, 1, 1, bottom, 0, 0, 0, time.Local)
	list := make(Days, 0, top-bottom+1)

	for i := bottom; i <= top; i++ {
		list = append(list, &Day{moment})
		moment = moment.Add(time.Hour)
	}

	return list
}

func (d Day) FormatHour(ampm interface{}) string {
	if doAmpm, _ := ampm.(bool); doAmpm {
		return d.Time.Format("3 PM")
	}

	return d.Time.Format("15")
}

func (d Day) Quarter() int {
	return int(math.Ceil(float64(d.Time.Month()) / 3.))
}

func (d Day) Month() time.Month {
	return d.Time.Month()
}

func (d Day) HeadingMOS(prefix, leaf string) string {
	day := strconv.Itoa(d.Time.Day())
	if len(leaf) > 0 {
		day = hyper.Link(d.ref(), day)
	}

	return hyper.Target(prefix+d.ref(), "") + `%
\begin{tabular}{@{}l|l}
  \multirow{2}{*}{\resizebox{!}{\myLenHeaderResizeBox}{` + day + `}} & \textbf{` + d.Time.Weekday().String() + `} \\
  & ` + d.Time.Month().String() + `
\end{tabular}`
}
