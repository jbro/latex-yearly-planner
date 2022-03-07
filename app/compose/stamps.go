package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
	"github.com/kudrykv/latex-yearly-planner/app/tex"
)

func Stamps(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 1)

	modules = append(modules, page.Module{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: map[string]interface{}{
			"Breadcrumb":   "Stamps",
			"HeadingMOS":   tex.Hypertarget("Stamps", "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Stamps`),
			"SideQuarters": year.SideQuarters(0),
			"SideMonths":   year.SideMonths(0),
			"Extra":        make(header.Items, 0),
			"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, true, nil, 1),
		},
	})

	modules = append(modules, page.Module{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: map[string]interface{}{
			"Breadcrumb":   "Scratch",
			"HeadingMOS":   tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Scratch`),
			"SideQuarters": year.SideQuarters(0),
			"SideMonths":   year.SideMonths(0),
			"Extra":        make(header.Items, 0),
			"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, false, nil, 1),
		},
	})

	return modules, nil
}
