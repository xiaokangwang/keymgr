package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/asn1"
	"errors"
	"github.com/xiaokangwang/keymgr/def"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"math/big"
	"net"
	"os"
	"strings"
)

type globalKeyCache struct {
	key string
}

func runSSHAgent(service def.TPMKeyServiceClient) {
	listenOn := mustGetFromEnv("SSH_AUTH_SOCK_LISTEN")
	os.Remove(listenOn)
	listener, err := net.ListenUnix("unix", &net.UnixAddr{Name: listenOn, Net: "unix"})
	if err != nil {
		panic(err)
	}
	keycache := &globalKeyCache{}
	defer listener.Close()
	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			panic(err)
		}
		go handleSSHAgentConnection(&multiProtocolConn{ReadWriteCloser: conn, bufRead: bufio.NewReader(conn), keyCache: keycache}, service, keycache)
	}
}

type multiProtocolConn struct {
	io.ReadWriteCloser
	bufRead       *bufio.Reader
	secondaryRead bool
	keyCache      *globalKeyCache
}

func (m *multiProtocolConn) Read(b []byte) (n int, err error) {
	if !m.secondaryRead {
		m.secondaryRead = true
		magic := "ecdsa-sha2-nistp256"
		data, err := m.bufRead.Peek(len(magic))
		if err != nil {
			return 0, err
		}
		if string(data) == magic {
			line, _, err := m.bufRead.ReadLine()
			if err != nil {
				return 0, err
			}
			m.keyCache.key = string(line)
			m.ReadWriteCloser.Close()
			return 0, io.EOF
		}
	}
	return m.bufRead.Read(b)
}

func handleSSHAgentConnection(conn io.ReadWriteCloser, service def.TPMKeyServiceClient, keyCache *globalKeyCache) {
	defer conn.Close()
	tpmAgent := &sshAgentFromTPMService{service: service, keyCache: keyCache}
	agent.ServeAgent(tpmAgent, conn)
}

type sshAgentFromTPMService struct {
	service  def.TPMKeyServiceClient
	keyCache *globalKeyCache
}

func (s sshAgentFromTPMService) List() ([]*agent.Key, error) {
	publicKey, comment, _, _, err := ssh.ParseAuthorizedKey([]byte(s.keyCache.key))
	if err != nil {
		return nil, err
	}
	return []*agent.Key{
		{
			Format:  publicKey.Type(),
			Blob:    publicKey.Marshal(),
			Comment: comment,
		},
	}, nil
}

func (s sshAgentFromTPMService) Sign(key ssh.PublicKey, data []byte) (*ssh.Signature, error) {
	publicKey, comment, _, _, err := ssh.ParseAuthorizedKey([]byte(s.keyCache.key))
	if err != nil {
		return nil, err
	}
	if bytes.Compare(publicKey.Marshal(), key.Marshal()) != 0 {
		return nil, errors.New("key not found")
	}
	KeySeed := strings.Split(comment, "@")[0] + mustGetFromEnv("KEY_SEED_SUFFIX")
	resp, err := s.service.SignWithECDSAKey(context.Background(), &def.SignWithECDSAKeyReq{KeySeed: KeySeed, Data: data})
	if err != nil {
		return nil, err
	}

	type asn1Signature struct {
		R, S *big.Int
	}
	asn1Sig := new(asn1Signature)
	_, err = asn1.Unmarshal(resp.Signature, asn1Sig)
	if err != nil {
		return nil, err
	}

	return &ssh.Signature{
		Format: publicKey.Type(),
		Blob:   ssh.Marshal(asn1Sig),
	}, nil
}

func (s sshAgentFromTPMService) Add(key agent.AddedKey) error {
	return errors.New("not implemented")
}

func (s sshAgentFromTPMService) Remove(key ssh.PublicKey) error {
	return errors.New("not implemented")
}

func (s sshAgentFromTPMService) RemoveAll() error {
	return errors.New("not implemented")
}

func (s sshAgentFromTPMService) Lock(passphrase []byte) error {
	return errors.New("not implemented")
}

func (s sshAgentFromTPMService) Unlock(passphrase []byte) error {
	return errors.New("not implemented")
}

func (s sshAgentFromTPMService) Signers() ([]ssh.Signer, error) {
	return []ssh.Signer{}, nil
}
