// Copyright (c) 2024 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"strconv"
	"time"

	"github.com/alibaba/opentelemetry-go-auto-instrumentation/test/verifier"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

var port int

func requestHttpServer() {
	client := &fasthttp.Client{}

	reqURL := "http://127.0.0.1:" + strconv.Itoa(port) + "/fiber"

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(reqURL)
	req.Header.SetMethod(fasthttp.MethodGet)

	if err := client.Do(req, resp); err != nil {
		panic(err)
	}
}

func setupFiberServer() {
	app := fiber.New()
	app.Get("/fiber", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.Status(fiber.StatusOK).SendString("Hello, World 👋!")
	})

	var err error
	port, err = verifier.GetFreePort()
	if err != nil {
		panic(err)
	}

	if err := app.Listen("127.0.0.1:" + strconv.Itoa(port)); err != nil {
		panic(err)
	}
}

func main() {
	go setupFiberServer()
	time.Sleep(2 * time.Second)
	requestHttpServer()
	verifier.WaitAndAssertMetrics(map[string]func(metricdata.ResourceMetrics){
		"http.server.request.duration": func(mrs metricdata.ResourceMetrics) {
			if len(mrs.ScopeMetrics) <= 0 {
				panic("No http.server.request.duration metrics received!")
			}
			point := mrs.ScopeMetrics[0].Metrics[0].Data.(metricdata.Histogram[float64])
			if point.DataPoints[0].Count <= 0 {
				panic("http.server.request.duration metrics count is not positive, actually " + strconv.Itoa(int(point.DataPoints[0].Count)))
			}
			verifier.VerifyHttpServerMetricsAttributes(point.DataPoints[0].Attributes.ToSlice(), "GET", "/fiber", "", "http", "", "http", fiber.StatusOK)
		},
		"http.client.request.duration": func(mrs metricdata.ResourceMetrics) {
			if len(mrs.ScopeMetrics) <= 0 {
				panic("No http.client.request.duration metrics received!")
			}
			point := mrs.ScopeMetrics[0].Metrics[0].Data.(metricdata.Histogram[float64])
			if point.DataPoints[0].Count <= 0 {
				panic("http.client.request.duration metrics count is not positive, actually " + strconv.Itoa(int(point.DataPoints[0].Count)))
			}
			verifier.VerifyHttpClientMetricsAttributes(point.DataPoints[0].Attributes.ToSlice(), "GET", "127.0.0.1:"+strconv.Itoa(port), "", "http", "", port, fiber.StatusOK)
		},
	})
}
