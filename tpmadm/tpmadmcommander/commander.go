package tpmadmcommander

import (
	"context"
	"encoding/asn1"
	"errors"
	"github.com/xiaokangwang/keymgr/def"
	"github.com/xiaokangwang/keymgr/tpmadm"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"strings"
)

func NewTpmAdmService() def.TPMKeyServiceServer {

	return &tpmadmService{adm: tpmadm.NewAdm()}
}

type tpmadmService struct {
	adm *tpmadm.Adm
	def.UnsafeTPMKeyServiceServer
}

func (t tpmadmService) SignWithECDSAKey(ctx context.Context, req *def.SignWithECDSAKeyReq) (*def.SignWithECDSAKeyResp, error) {

	err := t.checkCert(ctx, req.KeySeed)
	if err != nil {
		return nil, err
	}

	sig, err := t.adm.SignWithECDSAKey(req.KeySeed, req.Data)
	if err != nil {
		return nil, err
	}
	return &def.SignWithECDSAKeyResp{Signature: sig}, nil
}

func (t tpmadmService) checkCert(ctx context.Context, KeySeed string) error {
	p, ok := peer.FromContext(ctx)
	if ok {
		tlsInfo := p.AuthInfo.(credentials.TLSInfo)
		cert := tlsInfo.State.VerifiedChains[0][0]
		for _, value := range cert.Extensions {
			if value.Id.String() == "2.255.0.1" {
				extensionValue := value.Value
				var usageGranted string
				_, err := asn1.Unmarshal(extensionValue, &usageGranted)
				if err != nil {
					return err
				}
				detailedUsages := strings.Split(usageGranted, ";")
				if len(detailedUsages) == 0 {
					return errors.New("invalid usage granted")
				}
				if detailedUsages[0] != "tpm" {
					return errors.New("invalid usage granted: not tpm")
				}
				if !strings.HasPrefix(KeySeed, detailedUsages[1]) {
					return errors.New("invalid usage granted: not stating with the granted prefix")
				}
			}
		}
	}
	return nil
}

func (t tpmadmService) CreateECDSAKey(ctx context.Context, req *def.CreateECDSAKeyReq) (*def.CreateECDSAKeyResp, error) {
	err := t.checkCert(ctx, req.KeySeed)
	if err != nil {
		return nil, err
	}

	pub, err := t.adm.CreateECDSAKey(req.KeySeed)
	if err != nil {
		return nil, err
	}
	pubAsSsh, err := t.adm.GetKeyInSSHFormat(pub)
	if err != nil {
		return nil, err
	}
	return &def.CreateECDSAKeyResp{PublicBlob: pub, SshPublicKey: pubAsSsh}, nil
}

func (t tpmadmService) GetPrimaryKeyHash(ctx context.Context, req *def.GetPrimaryKeyHashReq) (*def.GetPrimaryKeyHashResp, error) {
	return &def.GetPrimaryKeyHashResp{PrimaryKey: t.adm.GetPrimaryKeyUnique()}, nil
}
