package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	joonix "github.com/joonix/log"
	log "github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	logJSON    = kingpin.Flag("log-json", "Use structured logging in JSON format").Default("false").Bool()
	logFluentd = kingpin.Flag("log-fluentd", "Use structured logging in GKE Fluentd format").Default("false").Bool()
	logLevel   = kingpin.Flag("log-level", "The level of logging").Default("info").Enum("debug", "info", "warn", "error", "panic", "fatal")
	secrets    = kingpin.Flag("secret", "Secret name (may be repeated)").Short('s').StringMap()
	command    = kingpin.Arg("command", "Command to run").Required().String()
	args       = kingpin.Arg("arg", "Argument").Strings()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.CommandLine.DefaultEnvars()
	kingpin.Parse()

	switch strings.ToLower(*logLevel) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	if *logJSON {
		log.SetFormatter(&log.JSONFormatter{})
	}
	if *logFluentd {
		log.SetFormatter(joonix.NewFormatter())
	}

	log.SetOutput(os.Stderr)
	ctx := context.Background()

	environ := os.Environ()

	for env, name := range *secrets {
		log.Debugf("Setting env %s from secret %s", env, name)

		client, err := secretmanager.NewClient(ctx)
		if err != nil {
			log.Fatal(err)
		}

		req := &secretmanagerpb.AccessSecretVersionRequest{
			Name: name,
		}

		result, err := client.AccessSecretVersion(ctx, req)
		if err != nil {
			log.Fatal(err)
		}

		environ = append(environ, fmt.Sprintf("%s=%s", env, string(result.Payload.Data)))
	}

	log.Debug(environ)

	argv := []string{*command}
	for _, a := range *args {
		argv = append(argv, a)
	}

	err := syscall.Exec(*command, argv, environ)
	if err != nil {
		log.Fatal(err)
	}
}
