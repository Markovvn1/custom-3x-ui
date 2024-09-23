package service

import (
	"bytes"
	"os/exec"
	"strings"

	"x-ui/logger"
	"x-ui/xray"
)

type ServerService struct {
	xrayService    XrayService
	inboundService InboundService
}

func (s *ServerService) RestartXrayService() (string error) {
	s.xrayService.StopXray()
	defer func() {
		err := s.xrayService.RestartXray(true)
		if err != nil {
			logger.Error("start xray failed:", err)
		}
	}()

	return nil
}

func (s *ServerService) GetNewX25519Cert() (interface{}, error) {
	// Run the command
	cmd := exec.Command(xray.GetBinaryPath(), "x25519")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")

	privateKeyLine := strings.Split(lines[0], ":")
	publicKeyLine := strings.Split(lines[1], ":")

	privateKey := strings.TrimSpace(privateKeyLine[1])
	publicKey := strings.TrimSpace(publicKeyLine[1])

	keyPair := map[string]interface{}{
		"privateKey": privateKey,
		"publicKey":  publicKey,
	}

	return keyPair, nil
}
