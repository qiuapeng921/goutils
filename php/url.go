package php

import (
	"net/url"
)

// ParseUrl - Parse a URL and return its components @data
func ParseUrl(rawUrl string) (*url.URL, error) {
	return url.Parse(rawUrl)
}

// ParseStr - Parses the string into variables
func ParseStr(query string) (url.Values, error) {
	return url.ParseQuery(query)
}

// RawUrlEncode - URL-encode according to RFC 3986
func RawUrlEncode(s string) string {
	return url.PathEscape(s)
}

// RawUrlDecode - Decodes URL-encoded string
func RawUrlDecode(s string) (string, error) {
	return url.PathUnescape(s)
}

// UrlEncode - URL-encodes string
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

// UrlDecode - Decodes URL-encoded string
func UrlDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// HTTPBuildQuery http_build_query()
func HTTPBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}