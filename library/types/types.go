package types

import "time"

type Alertmanager struct {
	Alerts []struct {
		Annotations struct {
			CurrentValue string `json:"current_value"`
		} `json:"annotations"`
		EndsAt time.Time `json:"endsAt"`
		GeneratorURL string `json:"generatorURL"`
		Labels struct {
			AlertName string `json:"alert_name"`
			AlertType string `json:"alert_type"`
			Alertname string `json:"alertname"`
			ClusterName string `json:"cluster_name"`
			Comparison string `json:"comparison"`
			Duration string `json:"duration"`
			Expression string `json:"expression"`
			GroupID string `json:"group_id"`
			Instance string `json:"instance"`
			Prometheus string `json:"prometheus"`
			PrometheusFrom string `json:"prometheus_from"`
			RuleID string `json:"rule_id"`
			Severity string `json:"severity"`
			ThresholdValue string `json:"threshold_value"`
		} `json:"labels"`
		StartsAt time.Time `json:"startsAt"`
		Status string `json:"status"`
	} `json:"alerts"`
	CommonAnnotations struct {
		CurrentValue string `json:"current_value"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		AlertName string `json:"alert_name"`
		AlertType string `json:"alert_type"`
		Alertname string `json:"alertname"`
		ClusterName string `json:"cluster_name"`
		Comparison string `json:"comparison"`
		Duration string `json:"duration"`
		Expression string `json:"expression"`
		GroupID string `json:"group_id"`
		Instance string `json:"instance"`
		Prometheus string `json:"prometheus"`
		PrometheusFrom string `json:"prometheus_from"`
		RuleID string `json:"rule_id"`
		Severity string `json:"severity"`
		ThresholdValue string `json:"threshold_value"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupKey string `json:"groupKey"`
	GroupLabels struct {
		RuleID string `json:"rule_id"`
	} `json:"groupLabels"`
	Receiver string `json:"receiver"`
	Status string `json:"status"`
	Version string `json:"version"`
}