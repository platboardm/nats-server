// Copyright 2012-2024 The NATS Authors
// Licensed under the Apache License, Version 2.0

package server

import (
	"crypto/tls"
	"net"
	"time"
)

const (
	// DefaultHost is the default host to listen on.
	DefaultHost = "0.0.0.0"
	// DefaultPort is the default port for client connections.
	DefaultPort = 4222
	// DefaultMaxConn is the default maximum number of client connections.
	DefaultMaxConn = 64 * 1024
	// DefaultMaxPayload is the default maximum payload size (1MB).
	DefaultMaxPayload = 1024 * 1024
	// DefaultMaxPending is the default maximum pending size (64MB).
	DefaultMaxPending = 64 * 1024 * 1024
	// DefaultPingInterval is the default ping interval.
	DefaultPingInterval = 2 * time.Minute
	// DefaultMaxPingOut is the default maximum number of unanswered pings.
	DefaultMaxPingOut = 2
	// DefaultWriteDeadline is the default write deadline.
	DefaultWriteDeadline = 10 * time.Second
)

// Options block for nats-server.
type Options struct {
	// Host and Port to listen on for client connections.
	Host string `json:"addr,omitempty"`
	Port int    `json:"port,omitempty"`

	// Maximum number of client connections.
	MaxConn int `json:"max_connections,omitempty"`

	// Maximum payload size in bytes.
	MaxPayload int32 `json:"max_payload,omitempty"`

	// Maximum pending bytes per client.
	MaxPending int64 `json:"max_pending,omitempty"`

	// Username and password for client authentication.
	Username string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`

	// Authorization token for client authentication.
	Authorization string `json:"auth_token,omitempty"`

	// TLS configuration.
	TLSConfig *tls.Config `json:"-"`
	TLS       bool        `json:"tls,omitempty"`
	TLSCert   string      `json:"tls_cert,omitempty"`
	TLSKey    string      `json:"tls_key,omitempty"`
	TLSCaCert string      `json:"tls_ca_cert,omitempty"`

	// Cluster configuration.
	Cluster ClusterOpts `json:"cluster,omitempty"`

	// Logging options.
	Debug   bool   `json:"debug,omitempty"`
	Trace   bool   `json:"trace,omitempty"`
	LogFile string `json:"log_file,omitempty"`
	Syslog  bool   `json:"syslog,omitempty"`

	// PID file.
	PidFile string `json:"pid_file,omitempty"`

	// Ping settings.
	PingInterval time.Duration `json:"ping_interval,omitempty"`
	MaxPingsOut  int           `json:"ping_max,omitempty"`

	// Write deadline for clients.
	WriteDeadline time.Duration `json:"write_deadline,omitempty"`

	// NoSig disables signal handling.
	NoSig bool `json:"-"`
}

// ClusterOpts are options for clustering.
type ClusterOpts struct {
	Host        string      `json:"addr,omitempty"`
	Port        int         `json:"cluster_port,omitempty"`
	Username    string      `json:"-"`
	Password    string      `json:"-"`
	AuthTimeout float64     `json:"auth_timeout,omitempty"`
	TLSConfig   *tls.Config `json:"-"`
	TLS         bool        `json:"tls,omitempty"`
	ListenStr   string      `json:"-"`
	Advertise   string      `json:"advertise,omitempty"`
	NoAdvertise bool        `json:"no_advertise,omitempty"`
	Routes      []*net.URL  `json:"-"`
}

// DefaultOptions returns an Options with default values set.
func DefaultOptions() *Options {
	return &Options{
		Host:          DefaultHost,
		Port:          DefaultPort,
		MaxConn:       DefaultMaxConn,
		MaxPayload:    DefaultMaxPayload,
		MaxPending:    DefaultMaxPending,
		PingInterval:  DefaultPingInterval,
		MaxPingsOut:   DefaultMaxPingOut,
		WriteDeadline: DefaultWriteDeadline,
	}
}
