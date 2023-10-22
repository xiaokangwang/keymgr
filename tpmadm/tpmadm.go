package tpmadm

import (
	"crypto/sha256"
	"encoding/asn1"
	"fmt"
	"github.com/canonical/go-tpm2"
	"github.com/canonical/go-tpm2/linux"
	"github.com/canonical/go-tpm2/mu"
	"github.com/canonical/go-tpm2/templates"
	"golang.org/x/crypto/ssh"
	"math/big"
)

type Adm struct {
	context *tpm2.TPMContext

	primaryKeySeed    string
	primaryKeyUnique  []byte
	primaryKeyContext tpm2.ResourceContext
}

func NewAdm() *Adm {
	adm := &Adm{}
	err := adm.init()
	if err != nil {
		panic(err)
	}
	return adm
}

func (a *Adm) init() error {
	tcti, err := linux.OpenDevice("/dev/tpmrm0")
	if err != nil {
		return err
	}
	context :=
		tpm2.NewTPMContext(tcti)
	a.context = context
	err = a.createPrimaryKey()
	if err != nil {
		return err
	}
	return nil
}

func (a *Adm) GetPrimaryKeyUnique() []byte {
	return a.primaryKeyUnique
}

func (a *Adm) createPrimaryKey() error {
	seed := sha256.Sum256([]byte(a.primaryKeySeed))
	inUnique := make(tpm2.Digest, 32)
	copy(inUnique, seed[:])
	rootTemplate := templates.NewDerivationParentKeyWithDefaults()
	rootTemplate.Unique = &tpm2.PublicIDU{KeyedHash: inUnique}

	context, public, _, _, _, err := a.context.CreatePrimary(a.context.OwnerHandleContext(), nil, rootTemplate, nil, nil, nil)
	if err != nil {
		return err
	}
	a.primaryKeyUnique = public.Unique.KeyedHash
	fmt.Printf("Primary Key Unique: %x\n", public.Unique.KeyedHash)
	a.primaryKeyContext = context
	return err
}

func (a *Adm) SignWithECDSAKey(keySeed string, dataIn []byte) ([]byte, error) {
	return a.signWithECDSAKey(keySeed, dataIn)
}

func (a *Adm) signWithECDSAKey(keySeed string, dataIn []byte) ([]byte, error) {
	_, resourceContext, err := a.createECDSAKey(keySeed)
	if err != nil {
		return nil, err
	}
	defer a.context.FlushContext(resourceContext)
	digest := sha256.Sum256(dataIn)
	signature, err := a.context.Sign(resourceContext, tpm2.Digest(digest[:]),
		&tpm2.SigScheme{Scheme: tpm2.SigSchemeAlgECDSA,
			Details: &tpm2.SigSchemeU{ECDSA: &tpm2.SigSchemeECDSA{HashAlg: tpm2.HashAlgorithmSHA256}}}, nil, nil)
	if err != nil {
		return nil, err
	}
	type asn1Signature struct {
		R, S *big.Int
	}
	var signatureASN1 asn1Signature
	signatureASN1.R = new(big.Int).SetBytes(signature.Signature.ECDSA.SignatureR)
	signatureASN1.S = new(big.Int).SetBytes(signature.Signature.ECDSA.SignatureS)
	resultASN1, err := asn1.Marshal(signatureASN1)
	if err != nil {
		return nil, err
	}
	return resultASN1, err

}

func (a *Adm) CreateECDSAKey(keySeed string) ([]byte, error) {
	pub, resourceContext, err := a.createECDSAKey(keySeed)
	defer a.context.FlushContext(resourceContext)
	if err != nil {
		return nil, err
	}
	pubAsByte, err := mu.MarshalToBytes(pub)
	if err != nil {
		return nil, err
	}
	return pubAsByte, nil
}

func (a *Adm) createECDSAKey(keySeed string) (*tpm2.Public, tpm2.ResourceContext, error) {
	keyTemplate := &tpm2.Public{
		Type:    tpm2.ObjectTypeECC,
		NameAlg: tpm2.HashAlgorithmSHA256,
		Attrs:   tpm2.AttrFixedTPM | tpm2.AttrFixedParent | tpm2.AttrUserWithAuth | tpm2.AttrSign | tpm2.AttrStClear,
		Params: &tpm2.PublicParamsU{ECCDetail: &tpm2.ECCParams{
			Symmetric: tpm2.SymDefObject{
				Algorithm: tpm2.SymObjectAlgorithmNull},
			Scheme:  tpm2.ECCScheme{Scheme: tpm2.ECCSchemeNull},
			CurveID: tpm2.ECCCurveNIST_P256,
			KDF:     tpm2.KDFScheme{Scheme: tpm2.KDFAlgorithmNull, Details: &tpm2.KDFSchemeU{}},
		}},
	}
	digest := sha256.Sum256([]byte(keySeed))
	keyTemplate.Unique = &tpm2.PublicIDU{ECC: &tpm2.ECCPoint{
		X: digest[:16],
		Y: digest[16:],
	}}
	context, _, public, err := a.context.CreateLoaded(a.primaryKeyContext, nil, keyTemplate, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	return public, context, nil
}

func (a *Adm) GetKeyInSSHFormat(public []byte) (string, error) {
	var publicStruct tpm2.Public
	_, err := mu.UnmarshalFromBytes(public, &publicStruct)
	if err != nil {
		return "", err
	}
	return getKeyInSSHFormat(&publicStruct)
}

func getKeyInSSHFormat(public *tpm2.Public) (string, error) {
	key := public.Public()
	sshKey, err := ssh.NewPublicKey(key)
	if err != nil {
		return "", nil
	}
	return string(ssh.MarshalAuthorizedKey(sshKey)), nil
}
