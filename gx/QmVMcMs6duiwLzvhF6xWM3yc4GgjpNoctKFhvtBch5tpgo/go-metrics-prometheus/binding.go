package metricsprometheus

import (
	"strings"

	pro "mbfs/go-mbfs/gx/QmTQuFQWHAWy4wMH6ZyPfGiawA5u9T8rs79FENoV8yXaoS/client_golang/prometheus"
	logging "mbfs/go-mbfs/gx/QmcuXC5cxs79ro2cUuHs4HQ2bkDLJUYokwL8aivcX6HW3C/go-log"
	metrics "mbfs/go-mbfs/gx/QmekzFM3hPZjTjUFGTABdQkEnQ3PTiMstY198PwSFr5w1Q/go-metrics-interface"
)

var log logging.EventLogger = logging.Logger("metrics-prometheus")

func Inject() error {
	return metrics.InjectImpl(newCreator)
}

func newCreator(name, helptext string) metrics.Creator {
	return &creator{
		name:     strings.Replace(name, ".", "_", -1),
		helptext: helptext,
	}
}

var _ metrics.Creator = &creator{}

type creator struct {
	name     string
	helptext string
}

func (c *creator) Counter() metrics.Counter {
	res := pro.NewCounter(pro.CounterOpts{
		Name: c.name,
		Help: c.helptext,
	})
	err := pro.Register(res)
	if err != nil {
		if registered, ok := err.(pro.AlreadyRegisteredError); ok {
			if existing, ok := registered.ExistingCollector.(pro.Counter); ok {
				log.Warningf("using existing prometheus collector: %s\n", c.name)
				return existing
			}
		}
		log.Errorf("Registering prometheus collector, name: %s, error: %s\n", c.name, err.Error())
	}
	return res
}
func (c *creator) Gauge() metrics.Gauge {
	res := pro.NewGauge(pro.GaugeOpts{
		Name: c.name,
		Help: c.helptext,
	})
	err := pro.Register(res)
	if err != nil {
		if registered, ok := err.(pro.AlreadyRegisteredError); ok {
			if existing, ok := registered.ExistingCollector.(pro.Gauge); ok {
				log.Warningf("using existing prometheus collector: %s\n", c.name)
				return existing
			}
		}
		log.Errorf("Registering prometheus collector, name: %s, error: %s\n", c.name, err.Error())
	}
	return res
}
func (c *creator) Histogram(buckets []float64) metrics.Histogram {
	res := pro.NewHistogram(pro.HistogramOpts{
		Name:    c.name,
		Help:    c.helptext,
		Buckets: buckets,
	})
	err := pro.Register(res)
	if err != nil {
		if registered, ok := err.(pro.AlreadyRegisteredError); ok {
			if existing, ok := registered.ExistingCollector.(pro.Histogram); ok {
				log.Warningf("using existing prometheus collector: %s\n", c.name)
				return existing
			}
		}
		log.Errorf("Registering prometheus collector, name: %s, error: %s\n", c.name, err.Error())
	}
	return res
}

func (c *creator) Summary(opts metrics.SummaryOpts) metrics.Summary {
	res := pro.NewSummary(pro.SummaryOpts{
		Name: c.name,
		Help: c.helptext,

		Objectives: opts.Objectives,
		MaxAge:     opts.MaxAge,
		AgeBuckets: opts.AgeBuckets,
		BufCap:     opts.BufCap,
	})
	err := pro.Register(res)
	if err != nil {
		if registered, ok := err.(pro.AlreadyRegisteredError); ok {
			if existing, ok := registered.ExistingCollector.(pro.Summary); ok {
				log.Warningf("using existing prometheus collector: %s\n", c.name)
				return existing
			}
		}
		log.Errorf("Registering prometheus collector, name: %s, error: %s\n", c.name, err.Error())
	}
	return res
}
