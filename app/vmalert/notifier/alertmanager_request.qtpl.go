// Code generated by qtc from "alertmanager_request.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line notifier/alertmanager_request.qtpl:1
package notifier

//line notifier/alertmanager_request.qtpl:1
import (
	"strconv"
	"time"
)

//line notifier/alertmanager_request.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line notifier/alertmanager_request.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line notifier/alertmanager_request.qtpl:7
func streamamRequest(qw422016 *qt422016.Writer, alerts []Alert, generatorURL func(string, string) string) {
//line notifier/alertmanager_request.qtpl:7
	qw422016.N().S(`[`)
//line notifier/alertmanager_request.qtpl:9
	for i, alert := range alerts {
//line notifier/alertmanager_request.qtpl:9
		qw422016.N().S(`{"startsAt":`)
//line notifier/alertmanager_request.qtpl:11
		qw422016.N().Q(alert.Start.Format(time.RFC3339Nano))
//line notifier/alertmanager_request.qtpl:11
		qw422016.N().S(`,"generatorURL":`)
//line notifier/alertmanager_request.qtpl:12
		qw422016.N().Q(generatorURL(strconv.FormatUint(alert.GroupID, 10), strconv.FormatUint(alert.ID, 10)))
//line notifier/alertmanager_request.qtpl:12
		qw422016.N().S(`,`)
//line notifier/alertmanager_request.qtpl:13
		if !alert.End.IsZero() {
//line notifier/alertmanager_request.qtpl:13
			qw422016.N().S(`"endsAt":`)
//line notifier/alertmanager_request.qtpl:14
			qw422016.N().Q(alert.End.Format(time.RFC3339Nano))
//line notifier/alertmanager_request.qtpl:14
			qw422016.N().S(`,`)
//line notifier/alertmanager_request.qtpl:15
		}
//line notifier/alertmanager_request.qtpl:15
		qw422016.N().S(`"labels": {"alertname":`)
//line notifier/alertmanager_request.qtpl:17
		qw422016.N().Q(alert.Name)
//line notifier/alertmanager_request.qtpl:18
		for k, v := range alert.Labels {
//line notifier/alertmanager_request.qtpl:18
			qw422016.N().S(`,`)
//line notifier/alertmanager_request.qtpl:19
			qw422016.N().Q(k)
//line notifier/alertmanager_request.qtpl:19
			qw422016.N().S(`:`)
//line notifier/alertmanager_request.qtpl:19
			qw422016.N().Q(v)
//line notifier/alertmanager_request.qtpl:20
		}
//line notifier/alertmanager_request.qtpl:20
		qw422016.N().S(`},"annotations": {`)
//line notifier/alertmanager_request.qtpl:23
		c := len(alert.Annotations)

//line notifier/alertmanager_request.qtpl:24
		for k, v := range alert.Annotations {
//line notifier/alertmanager_request.qtpl:25
			c = c - 1

//line notifier/alertmanager_request.qtpl:26
			qw422016.N().Q(k)
//line notifier/alertmanager_request.qtpl:26
			qw422016.N().S(`:`)
//line notifier/alertmanager_request.qtpl:26
			qw422016.N().Q(v)
//line notifier/alertmanager_request.qtpl:26
			if c > 0 {
//line notifier/alertmanager_request.qtpl:26
				qw422016.N().S(`,`)
//line notifier/alertmanager_request.qtpl:26
			}
//line notifier/alertmanager_request.qtpl:27
		}
//line notifier/alertmanager_request.qtpl:27
		qw422016.N().S(`}}`)
//line notifier/alertmanager_request.qtpl:30
		if i != len(alerts)-1 {
//line notifier/alertmanager_request.qtpl:30
			qw422016.N().S(`,`)
//line notifier/alertmanager_request.qtpl:30
		}
//line notifier/alertmanager_request.qtpl:31
	}
//line notifier/alertmanager_request.qtpl:31
	qw422016.N().S(`]`)
//line notifier/alertmanager_request.qtpl:33
}

//line notifier/alertmanager_request.qtpl:33
func writeamRequest(qq422016 qtio422016.Writer, alerts []Alert, generatorURL func(string, string) string) {
//line notifier/alertmanager_request.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line notifier/alertmanager_request.qtpl:33
	streamamRequest(qw422016, alerts, generatorURL)
//line notifier/alertmanager_request.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line notifier/alertmanager_request.qtpl:33
}

//line notifier/alertmanager_request.qtpl:33
func amRequest(alerts []Alert, generatorURL func(string, string) string) string {
//line notifier/alertmanager_request.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line notifier/alertmanager_request.qtpl:33
	writeamRequest(qb422016, alerts, generatorURL)
//line notifier/alertmanager_request.qtpl:33
	qs422016 := string(qb422016.B)
//line notifier/alertmanager_request.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line notifier/alertmanager_request.qtpl:33
	return qs422016
//line notifier/alertmanager_request.qtpl:33
}
