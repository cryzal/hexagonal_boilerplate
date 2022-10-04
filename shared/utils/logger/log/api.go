package log

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// Constant variable in API log. 2nd level logging we put into Event.
// Example : log.Level_1().Msg(Level2)

// APILogHandler : handle something who need to do
func APILogHandler(c echo.Context, req, res []byte) {
	c.Response().Header().Set("X-Majoo-ResponseTime", time.Now().Format(time.RFC3339))
	reqTime, err := time.Parse(time.RFC3339, c.Request().Header.Get("X-Majoo-RequestTime"))
	var elapstime time.Duration
	if err == nil {
		elapstime = time.Since(reqTime)
	}

	var handler string
	r := c.Echo().Routes()
	cpath := strings.Replace(c.Path(), "/", "", -1)
	for _, v := range r {
		vpath := strings.Replace(v.Path, "/", "", -1)
		if vpath == cpath && v.Method == c.Request().Method {
			handler = v.Name
			// Handler for wrong route.
			if strings.Contains(handler, "func1") {
				handler = "UndefinedRoute"
			}
			break
		}
	}

	// Get Handler Name
	dir, file := path.Split(handler)
	fileStrings := strings.Split(file, ".")
	packHandler := dir + fileStrings[0]
	funcHandler := strings.Replace(handler, packHandler+".", "", -1)

	logPrint := Info().
		Str("Identifier", "mp-aggregator_http").
		Str("package", packHandler).
		Int64("elapsed_time", elapstime.Nanoseconds()/int64(time.Millisecond)).
		Str("handler", funcHandler).
		Str("ip", c.RealIP()).
		Str("host", c.Request().Host).
		Str("method", c.Request().Method).
		Str("url", c.Request().RequestURI).
		Str("request_time", c.Request().Header.Get("X-Majoo-RequestTime")).
		Int("httpcode", c.Response().Status).
		Str("response_time", c.Response().Header().Get("X-Majoo-ResponseTime"))

	respHeader, _ := json.Marshal(c.Response().Header())
	logPrint.RawJSON("response_header", respHeader)

	reqdata := &bytes.Buffer{}
	if err := json.Compact(reqdata, req); err != nil {
	} else {
		logPrint.RawJSON("request", reqdata.Bytes())
	}

	resdata := &bytes.Buffer{}
	if err := json.Compact(resdata, res); err != nil {
	} else {
		logPrint.RawJSON("response", reqdata.Bytes())
	}

	logPrint.Msg("")
}

// APILogSkipper : rules for APILogHandler
func APILogSkipper(c echo.Context) bool {
	// bool, is this url request include "/api"?
	rules1 := strings.Contains(c.Request().RequestURI, "/api")

	// bool, is this request using method "GET"?
	rules2 := c.Request().Method != "GET"

	// bool, is this url request include "/login"?
	rules3 := strings.Contains(c.Request().RequestURI, "/login")

	// bool, is this url request include "/linkaja"?
	rules4 := strings.Contains(c.Request().RequestURI, "/linkaja")

	rules5 := strings.Contains(c.Request().RequestURI, "/upload")

	if rules1 {
		return false
	}

	if rules2 {
		if !rules3 {
			return false
		}
	}

	if rules4 {
		return false
	}

	if rules5 {
		return false
	}

	return true
}

// InfoLogHandler : handle something who need to do
func InfoLogHandler(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "?"
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	name, _ := os.Hostname()

	Info().
		Str("Identifier", "mp-aggregator_http").
		Str("file", file).
		Str("handler", fnName).
		Str("line", strconv.Itoa(line)).
		Str("host", name).
		Int("port", 8909).
		Msg(msg)
}

// RequestLogHandler : handle something who need to do
func RequestLogHandler(req *http.Request, resp *http.Response, statusCode int, body, reqBody []byte, skipLog bool) {
	// logDebug, _ := strconv.ParseBool(os.Getenv("LOG_DEBUG"))
	// if logDebug && !skipLog {
	Info().
		Str("Identifier", "mp-aggregator_http_out").
		Str("ip", req.RemoteAddr).
		Str("host", req.Host).
		Str("method", req.Method).
		Str("url", req.URL.String()).
		Int("http_code", statusCode).
		RawJSON("response", body).
		RawJSON("request", reqBody).
		Msg("")
	// }
}
