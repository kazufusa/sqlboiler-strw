// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package database

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Measurement is an object representing the database table.
type Measurement struct {
	ID                        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MeasuredAt                time.Time `boil:"measured_at" json:"measured_at" toml:"measured_at" yaml:"measured_at"`
	Alpha                     float32   `boil:"alpha" json:"alpha" toml:"alpha" yaml:"alpha"`
	Beta                      float32   `boil:"beta" json:"beta" toml:"beta" yaml:"beta"`
	AlphaAlarm                int64     `boil:"alpha_alarm" json:"alpha_alarm" toml:"alpha_alarm" yaml:"alpha_alarm"`
	BetaAlarm                 int64     `boil:"beta_alarm" json:"beta_alarm" toml:"beta_alarm" yaml:"beta_alarm"`
	Code                      int64     `boil:"code" json:"code" toml:"code" yaml:"code"`
	ExternalDoserate          float32   `boil:"external_doserate" json:"external_doserate" toml:"external_doserate" yaml:"external_doserate"`
	FlowRate                  float32   `boil:"flow_rate" json:"flow_rate" toml:"flow_rate" yaml:"flow_rate"`
	IntegratedFlowVolume      float32   `boil:"integrated_flow_volume" json:"integrated_flow_volume" toml:"integrated_flow_volume" yaml:"integrated_flow_volume"`
	AlphaBG                   float32   `boil:"alpha_bg" json:"alpha_bg" toml:"alpha_bg" yaml:"alpha_bg"`
	BetaBG                    float32   `boil:"beta_bg" json:"beta_bg" toml:"beta_bg" yaml:"beta_bg"`
	RadonConc                 float32   `boil:"radon_conc" json:"radon_conc" toml:"radon_conc" yaml:"radon_conc"`
	ThoronConc                float32   `boil:"thoron_conc" json:"thoron_conc" toml:"thoron_conc" yaml:"thoron_conc"`
	RadonCompensation         float32   `boil:"radon_compensation" json:"radon_compensation" toml:"radon_compensation" yaml:"radon_compensation"`
	ConfigurationMode         int64     `boil:"configuration_mode" json:"configuration_mode" toml:"configuration_mode" yaml:"configuration_mode"`
	FilterChanges             int64     `boil:"filter_changes" json:"filter_changes" toml:"filter_changes" yaml:"filter_changes"`
	FilterPrpessure           float32   `boil:"filter_prpessure" json:"filter_prpessure" toml:"filter_prpessure" yaml:"filter_prpessure"`
	AlphaEfficiency           float32   `boil:"alpha_efficiency" json:"alpha_efficiency" toml:"alpha_efficiency" yaml:"alpha_efficiency"`
	AlphaCTCPS                float32   `boil:"alpha_ct_cps" json:"alpha_ct_cps" toml:"alpha_ct_cps" yaml:"alpha_ct_cps"`
	AlphaCTBQPerCubicMeter    float32   `boil:"alpha_ct_bq_per_cubic_meter" json:"alpha_ct_bq_per_cubic_meter" toml:"alpha_ct_bq_per_cubic_meter" yaml:"alpha_ct_bq_per_cubic_meter"`
	AlphaCTAsdCPS             float32   `boil:"alpha_ct_asd_cps" json:"alpha_ct_asd_cps" toml:"alpha_ct_asd_cps" yaml:"alpha_ct_asd_cps"`
	AlphaCTAsdBQPerCubicMeter float32   `boil:"alpha_ct_asd_bq_per_cubic_meter" json:"alpha_ct_asd_bq_per_cubic_meter" toml:"alpha_ct_asd_bq_per_cubic_meter" yaml:"alpha_ct_asd_bq_per_cubic_meter"`
	AlphaLTCPS                float32   `boil:"alpha_lt_cps" json:"alpha_lt_cps" toml:"alpha_lt_cps" yaml:"alpha_lt_cps"`
	AlphaLTBQPerCubicMeter    float32   `boil:"alpha_lt_bq_per_cubic_meter" json:"alpha_lt_bq_per_cubic_meter" toml:"alpha_lt_bq_per_cubic_meter" yaml:"alpha_lt_bq_per_cubic_meter"`
	AlphaLTAsdCPS             float32   `boil:"alpha_lt_asd_cps" json:"alpha_lt_asd_cps" toml:"alpha_lt_asd_cps" yaml:"alpha_lt_asd_cps"`
	AlphaLTAsdBQPerCubicMeter float32   `boil:"alpha_lt_asd_bq_per_cubic_meter" json:"alpha_lt_asd_bq_per_cubic_meter" toml:"alpha_lt_asd_bq_per_cubic_meter" yaml:"alpha_lt_asd_bq_per_cubic_meter"`
	AlphaBGCPS                float32   `boil:"alpha_bg_cps" json:"alpha_bg_cps" toml:"alpha_bg_cps" yaml:"alpha_bg_cps"`
	AlphaBGBQPerCubicMeter    float32   `boil:"alpha_bg_bq_per_cubic_meter" json:"alpha_bg_bq_per_cubic_meter" toml:"alpha_bg_bq_per_cubic_meter" yaml:"alpha_bg_bq_per_cubic_meter"`
	Radon                     float32   `boil:"radon" json:"radon" toml:"radon" yaml:"radon"`
	Thoron                    float32   `boil:"thoron" json:"thoron" toml:"thoron" yaml:"thoron"`
	EqRadonConcN              float32   `boil:"eq_radon_conc_n" json:"eq_radon_conc_n" toml:"eq_radon_conc_n" yaml:"eq_radon_conc_n"`
	BetaEfficiency            float32   `boil:"beta_efficiency" json:"beta_efficiency" toml:"beta_efficiency" yaml:"beta_efficiency"`
	BetaCTCPS                 float32   `boil:"beta_ct_cps" json:"beta_ct_cps" toml:"beta_ct_cps" yaml:"beta_ct_cps"`
	BetaCTBQPerCubicMeter     float32   `boil:"beta_ct_bq_per_cubic_meter" json:"beta_ct_bq_per_cubic_meter" toml:"beta_ct_bq_per_cubic_meter" yaml:"beta_ct_bq_per_cubic_meter"`
	BetaCTAsdCPS              float32   `boil:"beta_ct_asd_cps" json:"beta_ct_asd_cps" toml:"beta_ct_asd_cps" yaml:"beta_ct_asd_cps"`
	BetaCTAsdBQPerCubicMeter  float32   `boil:"beta_ct_asd_bq_per_cubic_meter" json:"beta_ct_asd_bq_per_cubic_meter" toml:"beta_ct_asd_bq_per_cubic_meter" yaml:"beta_ct_asd_bq_per_cubic_meter"`
	BetaBGCPS                 float32   `boil:"beta_bg_cps" json:"beta_bg_cps" toml:"beta_bg_cps" yaml:"beta_bg_cps"`
	BetaBGBQPerCubicMeter     float32   `boil:"beta_bg_bq_per_cubic_meter" json:"beta_bg_bq_per_cubic_meter" toml:"beta_bg_bq_per_cubic_meter" yaml:"beta_bg_bq_per_cubic_meter"`
	GammaComp                 float32   `boil:"gamma_comp" json:"gamma_comp" toml:"gamma_comp" yaml:"gamma_comp"`
	SDGammaComp               float32   `boil:"sd_gamma_comp" json:"sd_gamma_comp" toml:"sd_gamma_comp" yaml:"sd_gamma_comp"`
	FilterChangeInterval      int64     `boil:"filter_change_interval" json:"filter_change_interval" toml:"filter_change_interval" yaml:"filter_change_interval"`

	R *measurementR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L measurementL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MeasurementColumns = struct {
	ID                        string
	MeasuredAt                string
	Alpha                     string
	Beta                      string
	AlphaAlarm                string
	BetaAlarm                 string
	Code                      string
	ExternalDoserate          string
	FlowRate                  string
	IntegratedFlowVolume      string
	AlphaBG                   string
	BetaBG                    string
	RadonConc                 string
	ThoronConc                string
	RadonCompensation         string
	ConfigurationMode         string
	FilterChanges             string
	FilterPrpessure           string
	AlphaEfficiency           string
	AlphaCTCPS                string
	AlphaCTBQPerCubicMeter    string
	AlphaCTAsdCPS             string
	AlphaCTAsdBQPerCubicMeter string
	AlphaLTCPS                string
	AlphaLTBQPerCubicMeter    string
	AlphaLTAsdCPS             string
	AlphaLTAsdBQPerCubicMeter string
	AlphaBGCPS                string
	AlphaBGBQPerCubicMeter    string
	Radon                     string
	Thoron                    string
	EqRadonConcN              string
	BetaEfficiency            string
	BetaCTCPS                 string
	BetaCTBQPerCubicMeter     string
	BetaCTAsdCPS              string
	BetaCTAsdBQPerCubicMeter  string
	BetaBGCPS                 string
	BetaBGBQPerCubicMeter     string
	GammaComp                 string
	SDGammaComp               string
	FilterChangeInterval      string
}{
	ID:                        "id",
	MeasuredAt:                "measured_at",
	Alpha:                     "alpha",
	Beta:                      "beta",
	AlphaAlarm:                "alpha_alarm",
	BetaAlarm:                 "beta_alarm",
	Code:                      "code",
	ExternalDoserate:          "external_doserate",
	FlowRate:                  "flow_rate",
	IntegratedFlowVolume:      "integrated_flow_volume",
	AlphaBG:                   "alpha_bg",
	BetaBG:                    "beta_bg",
	RadonConc:                 "radon_conc",
	ThoronConc:                "thoron_conc",
	RadonCompensation:         "radon_compensation",
	ConfigurationMode:         "configuration_mode",
	FilterChanges:             "filter_changes",
	FilterPrpessure:           "filter_prpessure",
	AlphaEfficiency:           "alpha_efficiency",
	AlphaCTCPS:                "alpha_ct_cps",
	AlphaCTBQPerCubicMeter:    "alpha_ct_bq_per_cubic_meter",
	AlphaCTAsdCPS:             "alpha_ct_asd_cps",
	AlphaCTAsdBQPerCubicMeter: "alpha_ct_asd_bq_per_cubic_meter",
	AlphaLTCPS:                "alpha_lt_cps",
	AlphaLTBQPerCubicMeter:    "alpha_lt_bq_per_cubic_meter",
	AlphaLTAsdCPS:             "alpha_lt_asd_cps",
	AlphaLTAsdBQPerCubicMeter: "alpha_lt_asd_bq_per_cubic_meter",
	AlphaBGCPS:                "alpha_bg_cps",
	AlphaBGBQPerCubicMeter:    "alpha_bg_bq_per_cubic_meter",
	Radon:                     "radon",
	Thoron:                    "thoron",
	EqRadonConcN:              "eq_radon_conc_n",
	BetaEfficiency:            "beta_efficiency",
	BetaCTCPS:                 "beta_ct_cps",
	BetaCTBQPerCubicMeter:     "beta_ct_bq_per_cubic_meter",
	BetaCTAsdCPS:              "beta_ct_asd_cps",
	BetaCTAsdBQPerCubicMeter:  "beta_ct_asd_bq_per_cubic_meter",
	BetaBGCPS:                 "beta_bg_cps",
	BetaBGBQPerCubicMeter:     "beta_bg_bq_per_cubic_meter",
	GammaComp:                 "gamma_comp",
	SDGammaComp:               "sd_gamma_comp",
	FilterChangeInterval:      "filter_change_interval",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperfloat32 struct{ field string }

func (w whereHelperfloat32) EQ(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat32) NEQ(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat32) LT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat32) LTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat32) GT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat32) GTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MeasurementWhere = struct {
	ID                        whereHelperint64
	MeasuredAt                whereHelpertime_Time
	Alpha                     whereHelperfloat32
	Beta                      whereHelperfloat32
	AlphaAlarm                whereHelperint64
	BetaAlarm                 whereHelperint64
	Code                      whereHelperint64
	ExternalDoserate          whereHelperfloat32
	FlowRate                  whereHelperfloat32
	IntegratedFlowVolume      whereHelperfloat32
	AlphaBG                   whereHelperfloat32
	BetaBG                    whereHelperfloat32
	RadonConc                 whereHelperfloat32
	ThoronConc                whereHelperfloat32
	RadonCompensation         whereHelperfloat32
	ConfigurationMode         whereHelperint64
	FilterChanges             whereHelperint64
	FilterPrpessure           whereHelperfloat32
	AlphaEfficiency           whereHelperfloat32
	AlphaCTCPS                whereHelperfloat32
	AlphaCTBQPerCubicMeter    whereHelperfloat32
	AlphaCTAsdCPS             whereHelperfloat32
	AlphaCTAsdBQPerCubicMeter whereHelperfloat32
	AlphaLTCPS                whereHelperfloat32
	AlphaLTBQPerCubicMeter    whereHelperfloat32
	AlphaLTAsdCPS             whereHelperfloat32
	AlphaLTAsdBQPerCubicMeter whereHelperfloat32
	AlphaBGCPS                whereHelperfloat32
	AlphaBGBQPerCubicMeter    whereHelperfloat32
	Radon                     whereHelperfloat32
	Thoron                    whereHelperfloat32
	EqRadonConcN              whereHelperfloat32
	BetaEfficiency            whereHelperfloat32
	BetaCTCPS                 whereHelperfloat32
	BetaCTBQPerCubicMeter     whereHelperfloat32
	BetaCTAsdCPS              whereHelperfloat32
	BetaCTAsdBQPerCubicMeter  whereHelperfloat32
	BetaBGCPS                 whereHelperfloat32
	BetaBGBQPerCubicMeter     whereHelperfloat32
	GammaComp                 whereHelperfloat32
	SDGammaComp               whereHelperfloat32
	FilterChangeInterval      whereHelperint64
}{
	ID:                        whereHelperint64{field: "\"measurements\".\"id\""},
	MeasuredAt:                whereHelpertime_Time{field: "\"measurements\".\"measured_at\""},
	Alpha:                     whereHelperfloat32{field: "\"measurements\".\"alpha\""},
	Beta:                      whereHelperfloat32{field: "\"measurements\".\"beta\""},
	AlphaAlarm:                whereHelperint64{field: "\"measurements\".\"alpha_alarm\""},
	BetaAlarm:                 whereHelperint64{field: "\"measurements\".\"beta_alarm\""},
	Code:                      whereHelperint64{field: "\"measurements\".\"code\""},
	ExternalDoserate:          whereHelperfloat32{field: "\"measurements\".\"external_doserate\""},
	FlowRate:                  whereHelperfloat32{field: "\"measurements\".\"flow_rate\""},
	IntegratedFlowVolume:      whereHelperfloat32{field: "\"measurements\".\"integrated_flow_volume\""},
	AlphaBG:                   whereHelperfloat32{field: "\"measurements\".\"alpha_bg\""},
	BetaBG:                    whereHelperfloat32{field: "\"measurements\".\"beta_bg\""},
	RadonConc:                 whereHelperfloat32{field: "\"measurements\".\"radon_conc\""},
	ThoronConc:                whereHelperfloat32{field: "\"measurements\".\"thoron_conc\""},
	RadonCompensation:         whereHelperfloat32{field: "\"measurements\".\"radon_compensation\""},
	ConfigurationMode:         whereHelperint64{field: "\"measurements\".\"configuration_mode\""},
	FilterChanges:             whereHelperint64{field: "\"measurements\".\"filter_changes\""},
	FilterPrpessure:           whereHelperfloat32{field: "\"measurements\".\"filter_prpessure\""},
	AlphaEfficiency:           whereHelperfloat32{field: "\"measurements\".\"alpha_efficiency\""},
	AlphaCTCPS:                whereHelperfloat32{field: "\"measurements\".\"alpha_ct_cps\""},
	AlphaCTBQPerCubicMeter:    whereHelperfloat32{field: "\"measurements\".\"alpha_ct_bq_per_cubic_meter\""},
	AlphaCTAsdCPS:             whereHelperfloat32{field: "\"measurements\".\"alpha_ct_asd_cps\""},
	AlphaCTAsdBQPerCubicMeter: whereHelperfloat32{field: "\"measurements\".\"alpha_ct_asd_bq_per_cubic_meter\""},
	AlphaLTCPS:                whereHelperfloat32{field: "\"measurements\".\"alpha_lt_cps\""},
	AlphaLTBQPerCubicMeter:    whereHelperfloat32{field: "\"measurements\".\"alpha_lt_bq_per_cubic_meter\""},
	AlphaLTAsdCPS:             whereHelperfloat32{field: "\"measurements\".\"alpha_lt_asd_cps\""},
	AlphaLTAsdBQPerCubicMeter: whereHelperfloat32{field: "\"measurements\".\"alpha_lt_asd_bq_per_cubic_meter\""},
	AlphaBGCPS:                whereHelperfloat32{field: "\"measurements\".\"alpha_bg_cps\""},
	AlphaBGBQPerCubicMeter:    whereHelperfloat32{field: "\"measurements\".\"alpha_bg_bq_per_cubic_meter\""},
	Radon:                     whereHelperfloat32{field: "\"measurements\".\"radon\""},
	Thoron:                    whereHelperfloat32{field: "\"measurements\".\"thoron\""},
	EqRadonConcN:              whereHelperfloat32{field: "\"measurements\".\"eq_radon_conc_n\""},
	BetaEfficiency:            whereHelperfloat32{field: "\"measurements\".\"beta_efficiency\""},
	BetaCTCPS:                 whereHelperfloat32{field: "\"measurements\".\"beta_ct_cps\""},
	BetaCTBQPerCubicMeter:     whereHelperfloat32{field: "\"measurements\".\"beta_ct_bq_per_cubic_meter\""},
	BetaCTAsdCPS:              whereHelperfloat32{field: "\"measurements\".\"beta_ct_asd_cps\""},
	BetaCTAsdBQPerCubicMeter:  whereHelperfloat32{field: "\"measurements\".\"beta_ct_asd_bq_per_cubic_meter\""},
	BetaBGCPS:                 whereHelperfloat32{field: "\"measurements\".\"beta_bg_cps\""},
	BetaBGBQPerCubicMeter:     whereHelperfloat32{field: "\"measurements\".\"beta_bg_bq_per_cubic_meter\""},
	GammaComp:                 whereHelperfloat32{field: "\"measurements\".\"gamma_comp\""},
	SDGammaComp:               whereHelperfloat32{field: "\"measurements\".\"sd_gamma_comp\""},
	FilterChangeInterval:      whereHelperint64{field: "\"measurements\".\"filter_change_interval\""},
}

// MeasurementRels is where relationship names are stored.
var MeasurementRels = struct {
}{}

// measurementR is where relationships are stored.
type measurementR struct {
}

// NewStruct creates a new relationship struct
func (*measurementR) NewStruct() *measurementR {
	return &measurementR{}
}

// measurementL is where Load methods for each relationship are stored.
type measurementL struct{}

var (
	measurementAllColumns            = []string{"id", "measured_at", "alpha", "beta", "alpha_alarm", "beta_alarm", "code", "external_doserate", "flow_rate", "integrated_flow_volume", "alpha_bg", "beta_bg", "radon_conc", "thoron_conc", "radon_compensation", "configuration_mode", "filter_changes", "filter_prpessure", "alpha_efficiency", "alpha_ct_cps", "alpha_ct_bq_per_cubic_meter", "alpha_ct_asd_cps", "alpha_ct_asd_bq_per_cubic_meter", "alpha_lt_cps", "alpha_lt_bq_per_cubic_meter", "alpha_lt_asd_cps", "alpha_lt_asd_bq_per_cubic_meter", "alpha_bg_cps", "alpha_bg_bq_per_cubic_meter", "radon", "thoron", "eq_radon_conc_n", "beta_efficiency", "beta_ct_cps", "beta_ct_bq_per_cubic_meter", "beta_ct_asd_cps", "beta_ct_asd_bq_per_cubic_meter", "beta_bg_cps", "beta_bg_bq_per_cubic_meter", "gamma_comp", "sd_gamma_comp", "filter_change_interval"}
	measurementColumnsWithoutDefault = []string{}
	measurementColumnsWithDefault    = []string{"id", "measured_at", "alpha", "beta", "alpha_alarm", "beta_alarm", "code", "external_doserate", "flow_rate", "integrated_flow_volume", "alpha_bg", "beta_bg", "radon_conc", "thoron_conc", "radon_compensation", "configuration_mode", "filter_changes", "filter_prpessure", "alpha_efficiency", "alpha_ct_cps", "alpha_ct_bq_per_cubic_meter", "alpha_ct_asd_cps", "alpha_ct_asd_bq_per_cubic_meter", "alpha_lt_cps", "alpha_lt_bq_per_cubic_meter", "alpha_lt_asd_cps", "alpha_lt_asd_bq_per_cubic_meter", "alpha_bg_cps", "alpha_bg_bq_per_cubic_meter", "radon", "thoron", "eq_radon_conc_n", "beta_efficiency", "beta_ct_cps", "beta_ct_bq_per_cubic_meter", "beta_ct_asd_cps", "beta_ct_asd_bq_per_cubic_meter", "beta_bg_cps", "beta_bg_bq_per_cubic_meter", "gamma_comp", "sd_gamma_comp", "filter_change_interval"}
	measurementPrimaryKeyColumns     = []string{"id"}
)

type (
	// MeasurementSlice is an alias for a slice of pointers to Measurement.
	// This should generally be used opposed to []Measurement.
	MeasurementSlice []*Measurement

	measurementQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	measurementType                 = reflect.TypeOf(&Measurement{})
	measurementMapping              = queries.MakeStructMapping(measurementType)
	measurementPrimaryKeyMapping, _ = queries.BindMapping(measurementType, measurementMapping, measurementPrimaryKeyColumns)
	measurementInsertCacheMut       sync.RWMutex
	measurementInsertCache          = make(map[string]insertCache)
	measurementUpdateCacheMut       sync.RWMutex
	measurementUpdateCache          = make(map[string]updateCache)
	measurementUpsertCacheMut       sync.RWMutex
	measurementUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single measurement record from the query.
func (q measurementQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Measurement, error) {
	o := &Measurement{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "database: failed to execute a one query for measurements")
	}

	return o, nil
}

// All returns all Measurement records from the query.
func (q measurementQuery) All(ctx context.Context, exec boil.ContextExecutor) (MeasurementSlice, error) {
	var o []*Measurement

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "database: failed to assign all query results to Measurement slice")
	}

	return o, nil
}

// Count returns the count of all Measurement records in the query.
func (q measurementQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "database: failed to count measurements rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q measurementQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "database: failed to check if measurements exists")
	}

	return count > 0, nil
}

// Measurements retrieves all the records using an executor.
func Measurements(mods ...qm.QueryMod) measurementQuery {
	mods = append(mods, qm.From("\"measurements\""))
	return measurementQuery{NewQuery(mods...)}
}

// FindMeasurement retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMeasurement(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Measurement, error) {
	measurementObj := &Measurement{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"measurements\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, measurementObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "database: unable to select from measurements")
	}

	return measurementObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Measurement) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("database: no measurements provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(measurementColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	measurementInsertCacheMut.RLock()
	cache, cached := measurementInsertCache[key]
	measurementInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			measurementAllColumns,
			measurementColumnsWithDefault,
			measurementColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(measurementType, measurementMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(measurementType, measurementMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"measurements\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"measurements\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"measurements\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, measurementPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "database: unable to insert into measurements")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == measurementMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "database: unable to populate default values for measurements")
	}

CacheNoHooks:
	if !cached {
		measurementInsertCacheMut.Lock()
		measurementInsertCache[key] = cache
		measurementInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Measurement.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Measurement) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	measurementUpdateCacheMut.RLock()
	cache, cached := measurementUpdateCache[key]
	measurementUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			measurementAllColumns,
			measurementPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return errors.New("database: unable to update measurements, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"measurements\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, measurementPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(measurementType, measurementMapping, append(wl, measurementPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "database: unable to update measurements row")
	}

	if !cached {
		measurementUpdateCacheMut.Lock()
		measurementUpdateCache[key] = cache
		measurementUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q measurementQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "database: unable to update all for measurements")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MeasurementSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("database: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), measurementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"measurements\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, measurementPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "database: unable to update all in measurement slice")
	}

	return nil
}

// Delete deletes a single Measurement record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Measurement) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("database: no Measurement provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), measurementPrimaryKeyMapping)
	sql := "DELETE FROM \"measurements\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "database: unable to delete from measurements")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q measurementQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("database: no measurementQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "database: unable to delete all from measurements")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MeasurementSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), measurementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"measurements\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, measurementPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "database: unable to delete all from measurement slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Measurement) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMeasurement(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MeasurementSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MeasurementSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), measurementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"measurements\".* FROM \"measurements\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, measurementPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "database: unable to reload all in MeasurementSlice")
	}

	*o = slice

	return nil
}

// MeasurementExists checks if the Measurement row exists.
func MeasurementExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"measurements\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "database: unable to check if measurements exists")
	}

	return exists, nil
}
