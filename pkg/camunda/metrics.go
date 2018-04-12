package camunda

import (
	"fmt"
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
	"net/http"
	"time"
)

const (
	pathMetrics                 = "metrics"
	pathMetricsSumByNamePattern = pathMetrics + "/%s/sum"

	MetricActivityInstanceStart                  = "activity-instance-start"
	MetricActivityInstanceEnd                    = "activity-instance-end"
	MetricJobAcquisitionAttempt                  = "job-acquisition-attempt"
	MetricJobAcquiredSuccess                     = "job-acquired-success"
	MetricJobAcquiredFailure                     = "job-acquired-failure"
	MetricJobExecutionRejected                   = "job-execution-rejected"
	MetricJobSuccessful                          = "job-successful"
	MetricJobFailed                              = "job-failed"
	MetricJobLockedExclusive                     = "job-locked-exclusive"
	MetricExecutedDecisionElements               = "executed-decision-elements"
	MetricHistoryCleanupRemovedProcessInstances  = "history-cleanup-removed-process-instances"
	MetricHistoryCleanupRemovedCaseInstances     = "history-cleanup-removed-case-instances"
	MetricHistoryCleanupRemovedDecisionInstances = "history-cleanup-removed-decision-instances"
	MetricHistoryCleanupRemovedBatchOperations   = "history-cleanup-removed-batch-operations"
)

var (
	metricsNames = [...]string{
		MetricActivityInstanceStart, MetricActivityInstanceEnd,
		MetricJobAcquisitionAttempt, MetricJobAcquiredSuccess, MetricJobAcquiredFailure,
		MetricJobExecutionRejected, MetricJobSuccessful, MetricJobFailed, MetricJobLockedExclusive,
		MetricExecutedDecisionElements,
		MetricHistoryCleanupRemovedProcessInstances, MetricHistoryCleanupRemovedCaseInstances,
		MetricHistoryCleanupRemovedDecisionInstances, MetricHistoryCleanupRemovedBatchOperations,
	}
)

type Metrics struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Reporter  *string    `json:"reporter,omitempty"`
	Value     *int       `json:"value,omitempty"`
}

func (m Metrics) String() string {
	return util.Stringify(m)
}

type GetMetricsRequestDto struct {
	Id                                string
	Name                              string
	NameLike                          string
	Source                            string
	WithoutSource                     string
	TenantIdIn                        []string
	WithoutTentantId                  bool
	IncludeDeploymentsWithoutTenantId bool
	After                             time.Time
	Before                            time.Time
	SortBy                            string // id / name / deploymentTime / tenantId
	SortOrder                         string // asc / desc
	FirstResult                       int
	MaxResults                        int
}

type MetricsService service

func (mc *MetricsService) GetList(metricName string, hostName string) ([]*Metrics, error) {
	request, err := newJsonRequest(mc.client.config.Endpoint, mc.client.config.UserAgent, http.MethodGet, pathMetrics, nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	if metricName != "" {
		q.Add("name", metricName)
	}
	q.Add("reporter", hostName)
	//	q.Add("startDate", strconv.FormatBool(skipCustomListeners))
	//	q.Add("endDate", strconv.FormatBool(skipIoMappings))
	//	q.Add("firstResult", strconv.FormatBool(skipIoMappings))
	//	q.Add("maxResults", strconv.FormatBool(skipIoMappings))
	//	q.Add("interval", strconv.FormatBool(skipIoMappings))
	request.URL.RawQuery = q.Encode()

	var metrics []*Metrics
	_, err = doRequest(mc.client.ctx, mc.client.httpClient, mc.client.requestInterceptors, request, metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

type metricsResult struct {
	result int `json:"result"`
}

func (mc *MetricsService) GetSum(metricName string, startDate time.Time, endDate time.Time) (int, error) {
	url := fmt.Sprintf(pathMetricsSumByNamePattern, metricName)
	request, err := newJsonRequest(mc.client.config.Endpoint, mc.client.config.UserAgent, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil
	}

	q := request.URL.Query()
	//    q.Add("startDate", strconv.FormatBool(skipCustomListeners))
	//    q.Add("endDate", strconv.FormatBool(skipIoMappings))
	request.URL.RawQuery = q.Encode()

	var metricsResult *metricsResult
	_, err = doRequest(mc.client.ctx, mc.client.httpClient, mc.client.requestInterceptors, request, metricsResult)
	if err != nil {
		return 0, err
	}
	return metricsResult.result, nil
}
