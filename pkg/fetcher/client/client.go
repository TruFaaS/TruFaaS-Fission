package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fission/fission/trufaas"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
	"golang.org/x/net/context/ctxhttp"

	ferror "github.com/fission/fission/pkg/error"
	"github.com/fission/fission/pkg/fetcher"
)

type (
	Client struct {
		logger     *zap.Logger
		url        string
		httpClient *http.Client
	}
)

func MakeClient(logger *zap.Logger, fetcherUrl string) *Client {
	hc := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	return &Client{
		logger:     logger.Named("fetcher_client"),
		url:        strings.TrimSuffix(fetcherUrl, "/"),
		httpClient: hc,
	}
}

func (c *Client) getSpecializeUrl() string {
	return c.url + "/specialize"
}

func (c *Client) getFetchUrl() string {
	return c.url + "/fetch"
}

func (c *Client) getUploadUrl() string {
	return c.url + "/upload"
}

func (c *Client) Specialize(ctx context.Context, req *fetcher.FunctionSpecializeRequest) error {
	_, err := sendRequest(c.logger, ctx, c.httpClient, req, c.getSpecializeUrl())
	return err
}

func (c *Client) Fetch(ctx context.Context, fr *fetcher.FunctionFetchRequest) error {
	_, err := sendRequest(c.logger, ctx, c.httpClient, fr, c.getFetchUrl())
	return err
}

func (c *Client) Upload(ctx context.Context, fr *fetcher.ArchiveUploadRequest) (*fetcher.ArchiveUploadResponse, error) {
	body, err := sendRequest(c.logger, ctx, c.httpClient, fr, c.getUploadUrl())
	if err != nil {
		return nil, err
	}

	uploadResp := fetcher.ArchiveUploadResponse{}
	err = json.Unmarshal(body, &uploadResp)
	if err != nil {
		return nil, err
	}

	return &uploadResp, nil
}

func sendRequest(logger *zap.Logger, ctx context.Context, httpClient *http.Client, req interface{}, url string) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	maxRetries := 20
	var resp *http.Response
	var specializationReq *http.Request

	for i := 0; i < maxRetries; i++ {
		// TruFaaS Modification [Protocol] - Protocol headers added to req sent
		if strings.Contains(url, "/specialize") {
			specializationReq, err = http.NewRequest("POST", url, bytes.NewReader(body))
			if err != nil {
				return nil, err
			}
			specializationReq.Header.Set("Content-Type", "application/json")
			trufaas.AddTrustProtocolHeadersToReq(specializationReq)
			if httpClient == nil {
				httpClient = http.DefaultClient
			}
			resp, err = httpClient.Do(specializationReq.WithContext(ctx))
			// If we got an error, and the context has been canceled, the context's error is probably more useful.
			if err != nil {
				select {
				case <-ctx.Done():
					err = ctx.Err()
				default:
				}
			}
		} else {
			resp, err = ctxhttp.Post(ctx, httpClient, url, "application/json", bytes.NewReader(body))
		}
		// TruFaaS Modification [Protocol] - Protocol headers saved in fetcher service
		trufaas.GetTrustProtocolHeadersFromResp(resp)
		if err == nil {
			if resp.StatusCode == 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					logger.Error("error reading response body", zap.Error(err))
				}
				defer resp.Body.Close()
				return body, err
			}
			err = ferror.MakeErrorFromHTTP(resp)
		}

		// skip retry and return directly due to context deadline exceeded
		if err == context.DeadlineExceeded {
			msg := "error specializing function pod, either increase the specialization timeout for function or check function pod log would help."
			err = errors.Wrap(err, msg)
			logger.Error(msg, zap.Error(err), zap.String("url", url))
			return nil, err
		}

		// TruFaaS Modification - skip retrying if trust verification failed in fetcher /specialize
		if strings.Contains(err.Error(), trufaas.TrustVerificationFailedMsg) {
			return nil, err
		}

		if i < maxRetries-1 {
			time.Sleep(50 * time.Duration(2*i) * time.Millisecond)
			logger.Error("error specializing/fetching/uploading package, retrying", zap.Error(err), zap.String("url", url))
			continue
		}
	}

	return nil, err
}
